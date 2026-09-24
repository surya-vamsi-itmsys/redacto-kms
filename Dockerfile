# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

# This is a helper stage that ensures the binary layer is always the same, no
# matter which base image it is copied into:
#
# 1. Always use /usr/bin/redacto-kms, not /bin/redacto-kms etc.
# 2. Apply the same file permissions across the /usr and /usr/bin directories.
#    Specifically, UBI is missing an u+w bit on /usr/bin that Alpine and
#    Distroless have.
#
# Together with SOURCE_DATE_EPOCH and rewrite-timestamp, this results in an
# identical binary layer digest across all distributions below, i.e., a given
# release binary is only ever pushed to a registry once, even if there is more
# than one container image flavor packaging it.
FROM scratch AS bin
ARG TARGETARCH
COPY --chmod=555 bin/${TARGETARCH}/redacto-kms /usr/bin/redacto-kms

# This is the redacto-kms container image.
FROM alpine:3.24.2@sha256:294b683cb724975bec92580e1e685676bd4b50bda910ddb8c51d4cabeaec77e6 AS default

COPY LICENSE /licenses/mozilla.txt

# Create a non-root user to run the software.
RUN addgroup redacto-kms && adduser -S -G redacto-kms redacto-kms

RUN apk add --no-cache ca-certificates libcap su-exec dumb-init tzdata gcompat

# Copy the binary stage.
COPY --from=bin . /

# /redacto-kms/logs is made available to use as a location to store audit logs, if
# desired; /redacto-kms/file is made available to use as a location with the file
# storage backend, if desired; the server will be started with /redacto-kms/config
# as the configuration directory so you can add additional config files in that
# location.
RUN mkdir -p /redacto-kms/logs && \
    mkdir -p /redacto-kms/file && \
    mkdir -p /redacto-kms/config && \
    chown -R redacto-kms:redacto-kms /redacto-kms

# 8200/tcp is the primary interface that applications use to interact with
# Redacto KMS.
EXPOSE 8200

# Use the Redacto KMS user as the default user for starting this container.
USER redacto-kms

# The entry point script uses dumb-init as the top-level process to reap any
# zombie processes created by Redacto KMS sub-processes.
COPY .release/docker/docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]

# By default you'll get a single-node development server that stores everything
# in RAM and bootstraps itself. Don't use this configuration for production.
CMD ["server", "-dev", "-dev-no-store-token"]


# This is the redacto-kms-ubi container image.
FROM registry.access.redhat.com/ubi10-minimal:10.2@sha256:e3a5632d7ae8a97e06f634522d06187f12793e90ac0d7b51bc671c83a96d8eda AS ubi

COPY LICENSE /licenses/mozilla.txt

# Overwrite Red Hat-specific labels present on the UBI base image.
LABEL io.k8s.description="Redacto KMS is a key management and secrets management service" \
      io.k8s.display-name="Redacto KMS" \
      io.openshift.expose-services="8200/tcp:https"

# Set up ca-certificates & base tooling.
RUN microdnf install -y ca-certificates gnupg openssl libcap tzdata procps shadow-utils util-linux

# Create a non-root user to run the software.
RUN groupadd --gid 1000 redacto-kms && \
    adduser --uid 100 --system -g redacto-kms redacto-kms && \
    usermod -a -G root redacto-kms

# Copy the binary stage.
COPY --from=bin . /

# /redacto-kms/logs is made available to use as a location to store audit logs, if
# desired; /redacto-kms/file is made available to use as a location with the file
# storage backend, if desired; the server will be started with /redacto-kms/config
# as the configuration directory so you can add additional config files in that
# location.
ENV HOME=/home/redacto-kms
RUN mkdir -p /redacto-kms/logs && \
    mkdir -p /redacto-kms/file && \
    mkdir -p /redacto-kms/config && \
    mkdir -p $HOME && \
    chown -R redacto-kms /redacto-kms && chown -R redacto-kms $HOME && \
    chgrp -R 0 $HOME && chmod -R g+rwX $HOME && \
    chgrp -R 0 /redacto-kms && chmod -R g+rwX /redacto-kms

# 8200/tcp is the primary interface that applications use to interact with
# Redacto KMS.
EXPOSE 8200

# Use the Redacto KMS user as the default user for starting this container.
USER redacto-kms

# The entry point script uses dumb-init as the top-level process to reap any
# zombie processes created by Redacto KMS sub-processes.
COPY .release/docker/ubi-docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["docker-entrypoint.sh"]

# By default you'll get a single-node development server that stores everything
# in RAM and bootstraps itself. Don't use this configuration for production.
CMD ["server", "-dev", "-dev-no-store-token"]


# This is the redacto-kms-distroless container image.
FROM gcr.io/distroless/static:nonroot@sha256:e2e927ec666bae08560abb3c55d0659eceabb657f56b6782ab500a9fc7f555e3 AS distroless

COPY LICENSE /licenses/mozilla.txt

# Copy the binary stage.
COPY --from=bin . /

# 8200/tcp is the primary interface that applications use to interact with
# Redacto KMS.
EXPOSE 8200

# By default you'll get a single-node development server that stores everything
# in RAM and bootstraps itself. Don't use this configuration for production.
ENTRYPOINT ["/usr/bin/redacto-kms"]
CMD ["server", "-dev", "-dev-no-store-token"]
