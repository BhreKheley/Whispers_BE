# syntax=docker/dockerfile:1

FROM golang:1.23.2-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Add wait-for.sh to wait for PostgreSQL
COPY wait-for.sh /wait-for.sh
RUN chmod +x /wait-for.sh

RUN go build -o main .

EXPOSE 8080

CMD ["/wait-for.sh", "./main"]
