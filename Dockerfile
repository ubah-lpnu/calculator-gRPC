FROM golang:1.22.2-alpine3.19 as builder

WORKDIR /app

COPY . .
RUN go mod download && go mod verify && go build -o server ./cmd/server/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server /app/server
RUN chmod +x /app/server
EXPOSE 50051
CMD ["./server"]


