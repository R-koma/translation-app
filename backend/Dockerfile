FROM golang:1.22-alpine

WORKDIR /app

# ENV GO111MODULE=on

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .
RUN go build -o  main

EXPOSE 8080

CMD [ "./main" ]