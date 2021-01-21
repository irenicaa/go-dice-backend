# go-dice-generator

[![GoDoc](https://godoc.org/github.com/irenicaa/go-dice-generator?status.svg)](https://godoc.org/github.com/irenicaa/go-dice-generator)
[![Go Report Card](https://goreportcard.com/badge/github.com/irenicaa/go-dice-generator)](https://goreportcard.com/report/github.com/irenicaa/go-dice-generator)
[![Build Status](https://travis-ci.org/irenicaa/go-dice-generator.svg?branch=master)](https://travis-ci.org/irenicaa/go-dice-generator)
[![codecov](https://codecov.io/gh/irenicaa/go-dice-generator/branch/master/graph/badge.svg)](https://codecov.io/gh/irenicaa/go-dice-generator)

The web service that implements dice rolling.

## Installation

```
$ go get github.com/irenicaa/go-dice-generator/...
```

## Usage

```
$ go-dice-generator -h | -help | --help
```

Options:

- `-h`, `-help`, `--help` &mdash; show the help message and exit.

Environment variables:

- `PORT` &mdash; server port (default: `8080`).

## Docs

- [swagger.yaml](docs/swagger.yaml) &mdash; Swagger definition of the server API
- [postman_collection.json](docs/postman_collection.json) &mdash; Postman collection of the server API

## Output Example

```
2021/01/22 00:53:32.261794 GET /stats 28.706µs
2021/01/22 00:53:32.262589 GET /dice?tries=37&faces=17 57.924µs
2021/01/22 00:53:32.263880 GET /dice?tries=55&faces=88 55.419µs
2021/01/22 00:53:32.265180 GET /dice?tries=19&faces=58 12.94µs
2021/01/22 00:53:32.266425 GET /dice?tries=100&faces=84 56.536µs
2021/01/22 00:53:32.267702 GET /dice?tries=28&faces=84 12.995µs
2021/01/22 00:53:32.269032 GET /dice?tries=42&faces=39 21.617µs
2021/01/22 00:53:32.270590 GET /dice?tries=24&faces=53 62.206µs
2021/01/22 00:53:32.271839 GET /dice?tries=43&faces=65 44.863µs
2021/01/22 00:53:32.273147 GET /dice?tries=23&faces=6 16.589µs
2021/01/22 00:53:32.274443 GET /dice?tries=43&faces=57 13.36µs
2021/01/22 00:53:32.275757 GET /stats 66.444µs
```

## License

The MIT License (MIT)

Copyright &copy; 2020-2021 irenica
