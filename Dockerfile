FROM golang:latest AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/test_task ./src

FROM debian:bookworm-slim

WORKDIR /app

RUN apt-get update
COPY --from=builder /app/test_task /app/

EXPOSE 8080

CMD ["/app/test_task"]