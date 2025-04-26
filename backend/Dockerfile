FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN go mod tidy

COPY . .

RUN go install github.com/air-verse/air@latest

COPY .air.toml ./

EXPOSE 8080

CMD ["air"]