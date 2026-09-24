#!/usr/bin/env bash
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


set -e

# Generate an OpenAPI document for all backends.
#
# Assumptions:
#
#   1. Redacto KMS has been checked out at an appropriate version and built
#   2. redacto-kms executable is in your path
#   3. Redacto KMS isn't already running
#   4. jq is installed

cd "$(dirname "${BASH_SOURCE[0]}")"

echo "Starting Redacto KMS..."
if pgrep -x "bao" > /dev/null
then
    echo "Redacto KMS is already running. Aborting."
    exit 1
fi

redacto-kms server -dev -dev-root-token-id=root &
REDACTO_KMS_PID=$!

# Allow time for Redacto KMS to start its HTTP listener
sleep 1

defer_stop_bao() {
    echo "Stopping Redacto KMS..."
    kill $REDACTO_KMS_PID
    # Allow time for Redacto KMS to print final logging and exit,
    # before this script ends, and the shell prints its next prompt
    sleep 1
}

trap defer_stop_bao INT TERM EXIT

export REDACTO_KMS_ADDR=http://127.0.0.1:8200

echo "Unmounting the default kv-v2 secrets engine ..."

# Unmount the default kv-v2 engine so that we can remount it at 'kv_v2/' later.
# The mount path will be reflected in the resultant OpenAPI document.
redacto-kms secrets disable "secret/"

echo "Mounting all builtin plugins ..."

# Enable auth plugins
redacto-kms auth enable "approle"
redacto-kms auth enable "cert"
redacto-kms auth enable "jwt"
redacto-kms auth enable "kubernetes"
redacto-kms auth enable "userpass"

# Enable secrets plugins
redacto-kms secrets enable "database"
redacto-kms secrets enable "kubernetes"
redacto-kms secrets enable -path="kv-v1/" -version=1 "kv"
redacto-kms secrets enable -path="kv-v2/" -version=2 "kv"
redacto-kms secrets enable "pki"
redacto-kms secrets enable "rabbitmq"
redacto-kms secrets enable "ssh"
redacto-kms secrets enable "totp"
redacto-kms secrets enable "transit"

# Output OpenAPI, optionally formatted
if [ "$1" == "-p" ]; then
    curl --header 'X-Vault-Token: root' \
         --data '{"generic_mount_paths": true}' \
            'http://127.0.0.1:8200/v1/sys/internal/specs/openapi' | jq > openapi.json
else
    curl --header 'X-Vault-Token: root' \
         --data '{"generic_mount_paths": true}' \
            'http://127.0.0.1:8200/v1/sys/internal/specs/openapi' > openapi.json
fi

echo
echo "openapi.json generated"
echo
