# OpenPing

Open-source uptime and network monitoring for developers, homelabs, and small teams.

OpenPing periodically checks HTTP/HTTPS endpoints and TCP ports, stores historical results in SQLite, and provides a web dashboard with uptime, latency, incidents, and recent checks.

## Features

- HTTP/HTTPS monitoring
- TCP port monitoring
- Configurable interval and timeout
- Uptime and latency statistics
- SQLite history
- REST API
- Prometheus metrics
- Web dashboard
- CSV export
- Docker support
- GitHub Actions CI
- MIT License

## Quick start

Requirements: Go 1.23+

```bash
go mod tidy
go run ./cmd/openping
```

Open http://localhost:8080

## Docker

```bash
docker compose up -d --build
```

## Configuration

Edit `config.yaml`. Target `type` can be `http` or `tcp`.

## API

- `GET /api/targets`
- `GET /api/recent?limit=100`
- `GET /api/incidents`
- `GET /api/export.csv`
- `GET /metrics`
- `GET /healthz`

## Roadmap

- Notifications
- ICMP checks
- Maintenance windows
- Public status pages
- TLS certificate expiry monitoring
- Multi-node monitoring

## License

MIT
