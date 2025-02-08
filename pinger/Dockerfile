FROM golang:1.23.5 as builder

WORKDIR /pinger_app_build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o pinger_microservice ./cmd/main/

FROM alpine:latest

WORKDIR /pinger_app

COPY --from=builder /pinger_app_build/pinger_microservice .

CMD ["./pinger_microservice"]