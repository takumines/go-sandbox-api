# syntax=docker/dockerfile:1
FROM golang:1.20-bullseye as local

WORKDIR /app
COPY ./app .

RUN go get -u github.com/cosmtrek/air && \
    go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]