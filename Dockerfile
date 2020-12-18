FROM golang:1.11.13
WORKDIR /go/src/github.com/irenicaa/go-dice-generator
COPY . .
RUN go install -v ./...
CMD ["go-dice-generator"]
