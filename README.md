# Redacto KMS

<p align="center">
  <img width="360" alt="Redacto KMS" src="./branding/redacto-brandlogo.png">
</p>

**Redacto KMS is a software solution to manage, store, and distribute sensitive
data including secrets, certificates, and keys.**

A modern system requires access to a multitude of secrets: database credentials,
API keys for external services, credentials for service-oriented architecture
communication, etc. Understanding who is accessing what secrets is already very
difficult and platform-specific. Adding on key rolling, secure storage, and
detailed audit logs is almost impossible without a custom solution. This is
where Redacto KMS steps in.

The key features of Redacto KMS are:

* **Secure Secret Storage**: Arbitrary key/value secrets can be stored in
  Redacto KMS. Redacto KMS encrypts these secrets prior to writing them to
  persistent storage, so gaining access to the raw storage isn't enough to
  access your secrets. Redacto KMS can write to disk,
  [PostgreSQL](https://www.postgresql.org/), and more.

* **Dynamic Secrets**: Redacto KMS can generate secrets on-demand for some
  systems, such as AWS or SQL databases. For example, when an application needs
  to access an S3 bucket, it asks Redacto KMS for credentials, and Redacto KMS
  will generate an AWS keypair with valid permissions on demand. After creating
  these dynamic secrets, Redacto KMS will also automatically revoke them after
  the lease is up.

* **Data Encryption**: Redacto KMS can encrypt and decrypt data without storing
  it. This allows security teams to define encryption parameters and developers
  to store encrypted data in a location such as a SQL database without having to
  design their own encryption methods.

* **Leasing and Renewal**: All secrets in Redacto KMS have a _lease_ associated
  with them. At the end of the lease, Redacto KMS will automatically revoke that
  secret. Clients are able to renew leases via built-in renew APIs.

* **Revocation**: Redacto KMS has built-in support for secret revocation.
  Redacto KMS can revoke not only single secrets, but a tree of secrets, for
  example, all secrets read by a specific user, or all secrets of a particular
  type. Revocation assists in key rolling as well as locking down systems in the
  case of an intrusion.

## Documentation

Product documentation lives in [`website/content/docs`](./website/content/docs).

## Using Redacto KMS

The command-line interface and server are a single binary, `redacto-kms`:

```sh
$ redacto-kms server -dev
$ export REDACTO_KMS_ADDR='http://127.0.0.1:8200'
$ redacto-kms status
$ redacto-kms version
```

### Environment variables

All public environment variables use the `REDACTO_KMS_` prefix, for example
`REDACTO_KMS_ADDR`, `REDACTO_KMS_TOKEN`, `REDACTO_KMS_TOKEN_PATH`,
`REDACTO_KMS_CACERT`, `REDACTO_KMS_CAPATH`, `REDACTO_KMS_CLIENT_CERT`,
`REDACTO_KMS_CLIENT_KEY`, `REDACTO_KMS_CLIENT_TIMEOUT`, `REDACTO_KMS_NAMESPACE`,
`REDACTO_KMS_FORMAT`, `REDACTO_KMS_API_ADDR`, `REDACTO_KMS_CLUSTER_ADDR`,
`REDACTO_KMS_LOG_LEVEL` and `REDACTO_KMS_UI`. Legacy `BAO_*` and `VAULT_*`
variables are **not** read.

### Default locations

| Purpose | Location |
| --- | --- |
| CLI configuration file | `~/.redacto-kms` |
| CLI token file | `~/.redacto-kms-token` |
| Package configuration | `/etc/redacto-kms/redacto-kms.hcl`, `/etc/redacto-kms/redacto-kms.env` |
| Package data / TLS | `/opt/redacto-kms/data`, `/opt/redacto-kms/tls` |
| systemd unit | `redacto-kms.service` |
| Container config / logs / file storage | `/redacto-kms/config`, `/redacto-kms/logs`, `/redacto-kms/file` |

### Container image

The [`Dockerfile`](./Dockerfile) builds `redacto-kms` images (Alpine, UBI and
distroless targets) from a pre-built binary at `bin/<arch>/redacto-kms`:

```sh
$ docker build --target default -t redacto-kms:dev .
$ docker run --rm -p 8200:8200 -e REDACTO_KMS_DEV_ROOT_TOKEN_ID=root redacto-kms:dev
```

## Developing Redacto KMS

You'll first need [Go](https://www.golang.org) installed on your machine. The
Go toolchain version used in CI and releases is pinned at
[`.go-version`](./.go-version).

To build a `redacto-kms` binary:

```sh
$ mkdir -p bin
$ go build -o bin/redacto-kms .
```

To run the Redacto KMS server in development mode:

```sh
$ go run . server -dev # Or `./bin/redacto-kms server -dev` if you've built the binary already.
```

To test a package:

```sh
$ go test ./some/package
```

Some additional notes on development:

- There is also a [`Makefile`](./Makefile) available for advanced build
  configurations and maintenance tasks (`make dev`, `make bin`).
- The web UI lives in [`ui`](./ui) and the documentation site in
  [`website`](./website). Development instructions are available at
  [`ui/README.md`](./ui/README.md) and [`website/README.md`](./website/README.md).
- Brand assets are stored in [`branding`](./branding) (source logo) and
  [`ui/public`](./ui/public) (derived UI assets). UI colours are defined once in
  `ui/app/styles/utils/_redacto_brand.scss` and exposed as design tokens in
  `ui/app/styles/redacto-tokens.scss`.

### Go module paths

Go module and package paths (for example `github.com/openbao/openbao/v2`,
`internal/vault`) are intentionally unchanged to keep this repository
synchronisable with its upstream source. They are not user-facing.

## License

License: *

Source files retain their original license notices (see [`LICENSE`](./LICENSE)
and the per-file headers).
