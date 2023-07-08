# Build Stage
FROM golang:1.20-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

# Run Stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

# Final Stage
EXPOSE 8080
CMD ["/app/main"]