# go-dice-backend

[![GoDoc](https://godoc.org/github.com/irenicaa/go-dice-backend?status.svg)](https://godoc.org/github.com/irenicaa/go-dice-backend)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-dice-backend)](https://goreportcard.com/report/github.com/irenicaa/go-dice-backend)
[![Build Status](https://app.travis-ci.com/irenicaa/go-dice-backend.svg?branch=master)](https://app.travis-ci.com/irenicaa/go-dice-backend)
[![codecov](https://codecov.io/gh/irenicaa/go-dice-backend/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-dice-backend)

The web service that implements dice rolling.

## Installation

```
$ go get github.com/irenicaa/go-dice-backend/...
```

## Usage

```
$ go-dice-backend -h | -help | --help
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit.

Environment variables:

- `PORT` &mdash; server port (default: `8080`).

## Testing

Running of the unit tests:

```
$ go test -race -cover ./...
```

Running of the integration tests:

```
$ docker-compose up -d
$ go test -race -cover -tags integration ./...
```

## Docs

- [swagger.yaml](docs/swagger.yaml) &mdash; Swagger definition of the server API
- [postman_collection.json](docs/postman_collection.json) &mdash; Postman collection of the server API

## Output Example

```
2021/01/22 00:53:32.261794 GET /api/v1/stats 28.706µs
2021/01/22 00:53:32.262589 GET /api/v1/dice?tries=37&faces=17 57.924µs
2021/01/22 00:53:32.263880 GET /api/v1/dice?tries=55&faces=88 55.419µs
2021/01/22 00:53:32.265180 GET /api/v1/dice?tries=19&faces=58 12.94µs
2021/01/22 00:53:32.266425 GET /api/v1/dice?tries=100&faces=84 56.536µs
2021/01/22 00:53:32.267702 GET /api/v1/dice?tries=28&faces=84 12.995µs
2021/01/22 00:53:32.269032 GET /api/v1/dice?tries=42&faces=39 21.617µs
2021/01/22 00:53:32.270590 GET /api/v1/dice?tries=24&faces=53 62.206µs
2021/01/22 00:53:32.271839 GET /api/v1/dice?tries=43&faces=65 44.863µs
2021/01/22 00:53:32.273147 GET /api/v1/dice?tries=23&faces=6 16.589µs
2021/01/22 00:53:32.274443 GET /api/v1/dice?tries=43&faces=57 13.36µs
2021/01/22 00:53:32.275757 GET /api/v1/stats 66.444µs
```

## License

The MIT License (MIT)

Copyright &copy; 2020-2021 irenica
