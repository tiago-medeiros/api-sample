FROM golang:1.24.4-alpine3.22 as builder
WORKDIR /app
COPY src/. . 
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8000
CMD ["./main"] 
