# BACKSTACK

[Go][1] stack to start on.

[![E2E](https://github.com/jerome-bienaime/backstack/actions/workflows/e2e.yml/badge.svg)](https://github.com/jerome-bienaime/backstack/actions/workflows/e2e.yml)
[![Go](https://github.com/jerome-bienaime/backstack/actions/workflows/go.yml/badge.svg)](https://github.com/jerome-bienaime/backstack/actions/workflows/go.yml)

## Requirements

+ [Go][1] 1.22.1
+ [SQLite][3] 3.45.1

### Docker

```bash
docker build -t web .
docker-compose up # server at :8080
```
  
## Install

```bash
make install
```

### .env

```bash
cp .example.env .dev.env
```

## Usage

```bash
ENV=DEV go run --tags="fts5" main.go . # start server on :8080
```

### Enable SSL

Enable SSL by setting ``HTTP_SSL`` to ``1`` in .*.env

Example with ``.dev.env``

```bash
# .dev.env
HTTP_SSL=1
```

You will need to create a [certificate](https://www.filecloud.com/blog/2022/09/create-an-ssl-certificate-in-5-easy-steps/)

#### Create certificate locally (with openssl)

```bash
# interactive
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 365

# non-interactive and 10 years expiration
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -sha256 -days 3650 -nodes -subj "/C=XX/ST=StateName/L=CityName/O=CompanyName/OU=CompanySectionName/CN=CommonNameOrHostname"
```

__Known issues__:

+ [tls: failed to parse private key](https://github.com/stellar/go/issues/64#issuecomment-319402456)
Then copy ``key.unencrypted.pem`` to ``key.pem``

```bash
cp key.unencrypted.pem key.pem
```

## Build

```bash
ENV=DEV go build --tags="fts5" . # build ./web executable
./web # launch server at :8080
```

## Reset database

### Test

```bash
make re-test # delete test database, & reseed
```

### Dev

```bash
make re-dev # reset dev database, & reseed
```

## Tests

### Unit && Integration

Github actions uses [go.yml file](./.github/workflows/go.yml) to execute
unit and integration tests.

```bash
make # make clean && make install
ENV=TEST go test -tags fts5 -timeout 30s -v ./...
```

### E2E

[Cypress](https://www.cypress.io/) is used to do E2E tests.

You will need [NodeJS](https://nodejs.org) and a package manager:
[pnpm](https://pnpm.io/) or [npm](https://www.npmjs.com/) or [yarn](https://yarnpkg.com/)
installed on your computer.
Github actions uses [e2e.yml file](./.github/workflows/e2e.yml) to execute E2E tests.
The tests are located under [cypress directory](./cypress/).

#### Cypress (application)

In a terminal run the server:

```bash
ENV=TEST go run --tags="fts5" main.go .
```

In another terminal run the application

```bash
pnpm cypress open 
# or
npx cypress open
# or
yarn cypress open
```

#### Cypress (cli)

The CLI version does not need you to launch server, however you will need to build the go binary.
[See build section](#build)

Once you have build go binary, launch in terminal

```bash
pnpm run ci 
# or
npm run ci
# or
yarn ci
```

[LICENSE](./LICENSE)

[1]: https://go.dev/
[3]: https://www.sqlite.org
