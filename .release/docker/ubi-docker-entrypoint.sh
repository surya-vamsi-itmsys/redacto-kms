#!/bin/sh
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

set -e

# Prevent core dumps
ulimit -c 0

# Allow setting REDACTO_KMS_REDIRECT_ADDR and REDACTO_KMS_CLUSTER_ADDR using an interface
# name instead of an IP address. The interface name is specified using
# REDACTO_KMS_REDIRECT_INTERFACE and REDACTO_KMS_CLUSTER_INTERFACE environment variables. If
# REDACTO_KMS_*_ADDR is also set, the resulting URI will combine the protocol and port
# number with the IP of the named interface.
get_addr () {
    local if_name=$1
    local uri_template=$2
    ip addr show dev $if_name | awk -v uri=$uri_template '/\s*inet\s/ { \
      ip=gensub(/(.+)\/.+/, "\\1", "g", $2); \
      print gensub(/^(.+:\/\/).+(:.+)$/, "\\1" ip "\\2", "g", uri); \
      exit}'
}

if [ -n "$REDACTO_KMS_REDIRECT_INTERFACE" ]; then
    export REDACTO_KMS_REDIRECT_ADDR=$(get_addr $REDACTO_KMS_REDIRECT_INTERFACE ${REDACTO_KMS_REDIRECT_ADDR:-"http://0.0.0.0:8200"})
    echo "Using $REDACTO_KMS_REDIRECT_INTERFACE for REDACTO_KMS_REDIRECT_ADDR: $REDACTO_KMS_REDIRECT_ADDR"
fi
if [ -n "$REDACTO_KMS_CLUSTER_INTERFACE" ]; then
    export REDACTO_KMS_CLUSTER_ADDR=$(get_addr $REDACTO_KMS_CLUSTER_INTERFACE ${REDACTO_KMS_CLUSTER_ADDR:-"https://0.0.0.0:8201"})
    echo "Using $REDACTO_KMS_CLUSTER_INTERFACE for REDACTO_KMS_CLUSTER_ADDR: $REDACTO_KMS_CLUSTER_ADDR"
fi

# REDACTO_KMS_CONFIG_DIR isn't exposed as a volume but you can compose additional
# config files in there if you use this image as a base, or use
# REDACTO_KMS_LOCAL_CONFIG below.
REDACTO_KMS_CONFIG_DIR=/redacto-kms/config

# You can also set the REDACTO_KMS_LOCAL_CONFIG environment variable to pass some
# Redacto KMS configuration JSON without having to bind any volumes.
if [ -n "$REDACTO_KMS_LOCAL_CONFIG" ]; then
    echo "$REDACTO_KMS_LOCAL_CONFIG" > "$REDACTO_KMS_CONFIG_DIR/local.json"
fi

# Due to OpenShift environment compatibility, we have to allow group write
# access to the Redacto KMS configuration. This requires us to disable the stricter
# file permissions checks (introduced upstream in Vault v1.11.0).
export REDACTO_KMS_DISABLE_FILE_PERMISSIONS_CHECK=true

# If the user is trying to run Redacto KMS directly with some arguments, then
# pass them to Redacto KMS.
if [ "${1:0:1}" = '-' ]; then
    set -- redacto-kms "$@"
fi

# Look for Redacto KMS subcommands.
if [ "$1" = 'server' ]; then
    shift
    set -- redacto-kms server \
        -config="$REDACTO_KMS_CONFIG_DIR" \
        -dev-root-token-id="$REDACTO_KMS_DEV_ROOT_TOKEN_ID" \
        -dev-listen-address="${REDACTO_KMS_DEV_LISTEN_ADDRESS:-"0.0.0.0:8200"}" \
        "$@"
elif [ "$1" = 'version' ]; then
    # This needs a special case because there's no help output.
    set -- redacto-kms "$@"
elif redacto-kms --help "$1" 2>&1 | grep -q "redacto-kms $1"; then
    # We can't use the return code to check for the existence of a subcommand, so
    # we have to use grep to look for a pattern in the help output.
    set -- redacto-kms "$@"
fi

# If we are running Redacto KMS, make sure it executes as the proper user.
if [ "$1" = 'redacto-kms' ]; then
    if [ -z "$SKIP_CHOWN" ]; then
        # If the config dir is bind mounted then chown it
        if [ "$(stat -c %u /redacto-kms/config)" != "$(id -u redacto-kms)" ]; then
            chown -R redacto-kms:redacto-kms /redacto-kms/config || echo "Could not chown /redacto-kms/config (may not have appropriate permissions)"
        fi

        # If the logs dir is bind mounted then chown it
        if [ "$(stat -c %u /redacto-kms/logs)" != "$(id -u redacto-kms)" ]; then
            chown -R redacto-kms:redacto-kms /redacto-kms/logs
        fi

        # If the file dir is bind mounted then chown it
        if [ "$(stat -c %u /redacto-kms/file)" != "$(id -u redacto-kms)" ]; then
            chown -R redacto-kms:redacto-kms /redacto-kms/file
        fi
    fi

fi

# This script is first executed as root, however, it is
# later rerun as the Redacto KMS user, which no longer needs to
# chown directories (previously done on the first run).
if [[ "$(id -u)" == '0' ]]
then
    export SKIP_CHOWN="true"
    exec su redacto-kms -p "$0" -- "$@"
else
    exec "$@"
fi
