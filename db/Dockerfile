FROM golang:1.20.6-alpine3.18 as builder

ARG SQL_MIGRATE_VERSION

RUN \
 CGO_ENABLED=0 \
 GOOS=linux \
 GOARCH=amd64 \
 go install "github.com/rubenv/sql-migrate/...1.5.2"

FROM scratch
COPY ./dbconfig.yml /
COPY ./migrations /migrations/
ENTRYPOINT [ "/sql-migrate" ]