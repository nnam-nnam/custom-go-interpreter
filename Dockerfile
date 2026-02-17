FROM golang:1.21 AS builder
WORKDIR /app/src/monkey
# Copy everything
COPY ./src/monkey .
# Disable CGO for a static binary
RUN CGO_ENABLED=0 go build -o monkey

FROM alpine:latest
WORKDIR /app
# Copy the binary from the builder stage
COPY --from=builder /app/src/monkey/monkey /usr/local/bin/
CMD ["monkey"]