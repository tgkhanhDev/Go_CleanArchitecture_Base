# Stage 1: Build binary
FROM golang:1.24-alpine AS builder
# Build-time argument
ARG APP_NAME=go-ftp-server_latest
ENV APP_NAME=${APP_NAME}
WORKDIR /app
RUN apk add --no-cache git
# Chỉ copy go.mod và go.sum trước để cache dependency
COPY go.mod go.sum ./
RUN go mod download
# Copy toàn bộ source
COPY . .
# Build binary tĩnh, strip debug info
RUN go build -ldflags="-s -w" -o ${APP_NAME} ./cmd/

# Stage 2: Runtime
FROM alpine:latest
# Cài CA certificates
RUN apk --no-cache add ca-certificates
ARG APP_NAME=go-ftp-server_latest
ENV APP_NAME=${APP_NAME}
WORKDIR /root/
COPY --from=builder /app/${APP_NAME} .
COPY --from=builder /app/*.yaml .
COPY --from=builder /app/assets ./assets
COPY --from=builder /app/docs ./docs
COPY --from=builder /app/templates ./templates

EXPOSE 8080
CMD sh -c "./${APP_NAME}"