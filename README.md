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

## Run on Windows

### Option 1: Run with Go

1. Install Go 1.23 or newer from https://go.dev/dl/.
2. Download this repository with **Code → Download ZIP**, then extract it.
3. Open PowerShell in the extracted `openping` folder.
4. Download dependencies:

```powershell
go mod tidy
```

5. Start OpenPing:

```powershell
go run ./cmd/openping
```

6. Open **http://localhost:8080** in your browser.

You should see the OpenPing dashboard. Keep the PowerShell window open while the server is running.

> If local compilation reports that `gcc` is missing, use the Docker method below or install a C compiler/toolchain for Windows.

### Option 2: Run with Docker

1. Install Docker Desktop from https://www.docker.com/products/docker-desktop/.
2. Open PowerShell in the OpenPing folder.
3. Run:

```powershell
docker compose up -d --build
```

4. Open **http://localhost:8080**.

To stop it:

```powershell
docker compose down
```

## Quick start on macOS / Linux

Requirements: Go 1.23+

```bash
go mod tidy
go run ./cmd/openping
```

Open http://localhost:8080

## Configuration

Edit `config.yaml`. Target `type` can be `http` or `tcp`.

Example:

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
