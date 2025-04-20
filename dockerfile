# Builder stage
FROM golang:1.23 AS builder

WORKDIR /build
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Copy built binary and static files to a new image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/app .
COPY --from=builder /build/public ./public

EXPOSE 3000
ENTRYPOINT ["./app"]
