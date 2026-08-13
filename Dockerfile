FROM golang:1.23-alpine AS build
RUN apk add --no-cache gcc musl-dev
WORKDIR /src
COPY . .
RUN go mod download && CGO_ENABLED=1 go build -trimpath -ldflags="-s -w" -o /openping ./cmd/openping
FROM alpine:3.21
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build /openping /app/openping
COPY config.yaml /app/config.yaml
COPY web /app/web
RUN mkdir -p /app/data
EXPOSE 8080
VOLUME ["/app/data"]
ENTRYPOINT ["/app/openping"]
