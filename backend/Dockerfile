FROM golang:1.23.5 AS builder

WORKDIR /backend_app_build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o backend_microservice ./cmd/main/

FROM alpine:latest

WORKDIR /backend_app

COPY --from=builder /backend_app_build/backend_microservice .

CMD ["./backend_microservice"]