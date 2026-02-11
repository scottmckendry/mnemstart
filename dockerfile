# Builder stage
FROM golang:1.26 AS builder

WORKDIR /build
COPY . .

RUN go install github.com/a-h/templ/cmd/templ@latest
RUN templ generate
RUN CGO_ENABLED=0 GOOS=linux go build -o mnemstart

# Copy built binary and static files to a new image
FROM alpine:latest

RUN adduser -D -u 1000 mnemstart

WORKDIR /app
COPY --from=builder /build/mnemstart .
COPY --from=builder /build/public ./public

RUN chown -R root:root /app && \
    chmod -R 755 /app && \
    mkdir -p /app/data && \
    chown -R mnemstart:mnemstart /app

USER mnemstart
EXPOSE 3000
ENTRYPOINT ["./mnemstart"]
