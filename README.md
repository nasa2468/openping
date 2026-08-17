# OpenPing

Open-source uptime and network monitoring for developers, homelabs, and small teams.

OpenPing periodically checks HTTP/HTTPS endpoints and TCP ports, stores historical results in SQLite, and provides a web dashboard with uptime, latency, incidents, and recent checks.

> **Status:** Early-stage open-source project. Contributions and feedback are welcome.

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
- Docker and Docker Compose support
- GitHub Actions CI
- MIT License

## Requirements

- Go 1.23+ for source builds
- Docker for containerized deployment

## Quick start

### Go

```bash
go mod tidy
go run ./cmd/openping
```

Open **http://localhost:8080**.

### Docker Compose

```bash
docker compose up -d --build
```

Open **http://localhost:8080**.

Stop the service with:

```bash
docker compose down
```

## Configuration

Edit `config.yaml`. Target `type` can be `http` or `tcp`.

```yaml
server:
  address: ":8080"

database: "data/openping.db"

targets:
  - name: OpenAI
    type: http
    address: https://openai.com
    interval_seconds: 30
    timeout_seconds: 10

  - name: Cloudflare DNS
    type: tcp
    address: 1.1.1.1:53
    interval_seconds: 30
    timeout_seconds: 5
```

After changing the configuration, restart OpenPing.

## API

| Endpoint | Purpose |
| --- | --- |
| `GET /api/targets` | List configured targets |
| `GET /api/recent?limit=100` | Recent check results |
| `GET /api/incidents` | Monitoring incidents |
| `GET /api/export.csv` | Export monitoring history |
| `GET /metrics` | Prometheus metrics |
| `GET /healthz` | Health check |

## Development

Run the test suite:

```bash
go test ./...
```

Run tests with the race detector:

```bash
go test -race ./...
```

Run static analysis:

```bash
go vet ./...
```

Build the application:

```bash
go build ./cmd/openping
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for contribution guidelines.

## Roadmap

### Monitoring

- [ ] ICMP checks
- [ ] TLS certificate expiry monitoring
- [ ] Maintenance windows

### Notifications

- [ ] Webhook notifications
- [ ] Telegram / Discord notifications
- [ ] Email notifications

### Platform

- [ ] Public status pages
- [ ] Improved historical statistics
- [ ] Multi-node monitoring
- [ ] Automated cross-platform releases

## Security

Please see [SECURITY.md](SECURITY.md) for vulnerability reporting guidance.

## License

MIT
