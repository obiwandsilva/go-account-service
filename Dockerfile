FROM golang:1.14.3-alpine3.11 AS builder
WORKDIR /tmp/go-account-service
COPY . .
RUN ["go", "build", "cmd/account-service/main.go"]

FROM alpine:3.11
WORKDIR /opt/account-service
COPY --from=builder /tmp/go-account-service/main .
EXPOSE 3000
CMD [ "./main", "3000" ]
