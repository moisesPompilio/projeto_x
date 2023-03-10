FROM golang:1.19 as builder
WORKDIR /app
COPY . ./
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o src/cmd/server/main src/cmd/server/main.go

ENTRYPOINT [ "/app/src/cmd/server/main" ]