# syntax=docker/dockerfile:1.4
FROM --platform=$BUILDPLATFORM golang:1.21-alpine3.18

ENV PROJECT_DIR=/gobeh \
    GO111MODULE=on \
    CGO_ENABLED=0

WORKDIR /gobeh
RUN mkdir "bin"
COPY . .
RUN go get github.com/githubnemo/CompileDaemon
RUN go install github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon -build="go build -o /gobeh/bin/gobeh" -command="/gobeh/bin/gobeh"