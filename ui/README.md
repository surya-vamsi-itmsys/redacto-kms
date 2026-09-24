**Table of Contents**

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->

- [Redacto KMS UI](#redacto-kms-ui)
  - [Prerequisites](#prerequisites)
  - [Running a Redacto KMS Server](#running-an-redacto-kms-server)
  - [Running / Development](#running--development)
    - [Code Generators](#code-generators)
    - [Running Tests](#running-tests)
    - [Linting](#linting)
    - [Building the Redacto KMS UI into a Redacto KMS Binary](#building-the-redacto-kms-ui-into-an-redacto-kms-binary)
  - [Further Reading / Useful Links](#further-reading--useful-links)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Redacto KMS UI

This README outlines the details of collaborating on this Ember application.

## Prerequisites

You will need the following things properly installed on your computer.

* [Git](https://git-scm.com/)
* [Node.js](https://nodejs.org/)
* [pnpm](https://pnpm.io/)
* [Ember CLI](https://cli.emberjs.com/release/)
* [Google Chrome](https://google.com/chrome/)

## Running a Redacto KMS Server

Before running the Redacto KMS UI locally, a Redacto KMS server must be running. First,
ensure a dev build of Redacto KMS is available according the the instructions in
`../README.md`. To start a single local Redacto KMS server:

- `pnpm openbao`

To start a local Redacto KMS cluster:

- `pnpm openbao:cluster`

## Running / Development

To get all of the JavaScript dependencies installed, run this in the `ui` directory:

- `pnpm install`

If you want to run the Redacto KMS UI and proxy back to a Redacto KMS server running on
the default port, 8200, run the following in the `ui` directory:

- `pnpm start`

This will start an Ember CLI server that proxies requests to port 8200,
and enable live rebuilding of the application as you change the UI application code.
Visit your app at [http://localhost:4200](http://localhost:4200).

If your Redacto KMS server is running on a different port you can use the long-form
version of the npm script:

`ember server --proxy=http://localhost:PORT`

To run pnpm with mirage, do:

- `pnpm start:mirage handlername`

Where `handlername` is one of the options exported in `mirage/handlers/index`

### Code Generators

Make use of the many generators for code, try `ember help generate`
for more details. If you're using a component that can be widely-used,
consider making it an `addon` component instead (see [this
PR](https://github.com/hashicorp/vault/pull/6629) for more details)

eg. a reusable component named foo that you'd like in the core engine

- `ember g component foo --in lib/core`
- `echo "export { default } from 'core/components/foo';" > lib/core/app/components/foo.js`

### Running Tests

Running tests will spin up a Redacto KMS dev server on port 9200 via a pretest
script that testem (the test runner) executes. All of the acceptance tests then
run, proxing requests back to that server.

- `pnpm test`
- `pnpm test -s` to keep the test server running after the initial run.
- `pnpm test -f="policies"` to filter the tests that are run. `-f` gets passed into
  [QUnit's `filter` config](https://api.qunitjs.com/config/QUnit.config#qunitconfigfilter-string--default-undefined)

### Linting

- `pnpm lint`
- `pnpm lint:fix`

### Building the Redacto KMS UI into a Redacto KMS Binary

We use the [embed](https://golang.org/pkg/embed/) package from Go 1.16+ to build
the static assets of the Ember application into a Redacto KMS binary.

This can be done by running these commands from the root directory run:
`make static-dist`
`make dev-ui`

This will result in a Redacto KMS binary that has the UI built-in - though in a
non-dev setup it will still need to be enabled via the `ui` config or setting
`REDACTO_KMS_UI` environment variable.

## Further Reading / Useful Links

* [ember.js](https://emberjs.com/)
* [ember-cli](https://cli.emberjs.com/release/)
* Development Browser Extensions
  * [ember inspector for chrome](https://chrome.google.com/webstore/detail/ember-inspector/bmdblncegkenkacieihfhpjfppoconhi)
  * [ember inspector for firefox](https://addons.mozilla.org/en-US/firefox/addon/ember-inspector/)
