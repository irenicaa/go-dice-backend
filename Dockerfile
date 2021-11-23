FROM golang:1.15.10 AS builder
WORKDIR /go/src/github.com/irenicaa/go-dice-backend
COPY . .
RUN CGO_ENABLED=0 go install -v ./...

FROM scratch
COPY --from=builder /go/bin/go-dice-backend /usr/local/bin/go-dice-backend
CMD ["/usr/local/bin/go-dice-backend"]
