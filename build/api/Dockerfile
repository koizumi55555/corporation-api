FROM golang:1.20.7 as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/app

FROM alpine
COPY --from=builder /build/api /opt/app/
ENTRYPOINT ["/opt/app/api"]