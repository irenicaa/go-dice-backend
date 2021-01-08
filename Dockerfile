FROM golang:1.11.13 AS builder
WORKDIR /go/src/github.com/irenicaa/go-dice-generator
COPY . .
RUN CGO_ENABLED=0 go install -v ./...

FROM scratch
COPY --from=builder /go/bin/go-dice-generator /usr/local/bin/go-dice-generator
CMD ["/usr/local/bin/go-dice-generator"]
