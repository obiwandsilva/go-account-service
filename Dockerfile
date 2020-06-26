FROM golang:1.14.3-alpine3.11

WORKDIR /opt/go-account-service

COPY . .

RUN ["go", "build", "cmd/account-service/main.go"]

EXPOSE 80

CMD [ "./main", "80" ]
