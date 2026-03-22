# Production build stage
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app main.go

# Final runtime stage
FROM alpine:3.18
WORKDIR /app
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
COPY .env .

ENV PORT=8000
EXPOSE 8000

CMD ["./app"]
