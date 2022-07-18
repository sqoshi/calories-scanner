FROM golang:1.18.3-alpine3.16


ENV GIN_MODE=release

COPY --chown=nobody:nobody  . /app

WORKDIR /app

RUN go mod download && \
    go build -o app-bin

USER nobody

ENTRYPOINT ["./app-bin"]