FROM golang:1.11 as builder
WORKDIR /go/src/github.com/anpryl/paymentsvc
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o paymentsvc ./cmd/paymentsvc/main.go

FROM alpine:3.6
WORKDIR /
RUN apk update && apk add ca-certificates
COPY --from=builder /go/src/github.com/anpryl/paymentsvc/paymentsvc .
COPY migrations /go/src/github.com/anpryl/paymentsvc/migrations
EXPOSE 80
CMD ["/paymentsvc"]
