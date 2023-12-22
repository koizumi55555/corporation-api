FROM golang:1.20.7 as modules
RUN echo test
RUN echo test
COPY go.mod go.sum /modules/
WORKDIR /modules
RUN go mod download

FROM golang:1.20.7 as builder
COPY --from=modules /go/pkg /go/pkg
COPY . /app
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /bin/corporation-api ./cmd/app/main.go
EXPOSE 8080
ENTRYPOINT ["corporation-api"]