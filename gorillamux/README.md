# go-boilerplates

## Prerequisites

- Install the latest version of [Go](https://go.dev/dl).
- Set `GOPATH` environment variable.
- Add `$(go env GOPATH)/bin` to `PATH`.
- Install [`mkcert`](https://github.com/FiloSottile/mkcert).

## Getting started

```bash
mkcert -install && mkcert localhost

git submodule --remote --recursive

go get

make server
```
