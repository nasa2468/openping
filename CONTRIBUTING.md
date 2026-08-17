# Contributing to OpenPing

Thanks for helping improve OpenPing.

## Before you start

- Search existing issues before opening a new one.
- For larger changes, open an issue first so the design can be discussed.
- Keep pull requests focused on one change.

## Development

Requirements:

- Go 1.23+

Run the test suite:

```bash
go test ./...
```

Run the race detector:

```bash
go test -race ./...
```

Check formatting:

```bash
gofmt -w .
```

Run static checks:

```bash
go vet ./...
```

Build OpenPing:

```bash
go build ./cmd/openping
```

## Pull requests

Please include:

- a concise description of the change;
- tests for new or changed behavior where practical;
- documentation updates when user-facing behavior changes;
- any compatibility or migration notes.

Keep commits and pull requests focused. Avoid unrelated formatting or refactoring in feature and bug-fix changes.

## Security

Do not disclose security vulnerabilities in public issues. See `SECURITY.md` for the preferred reporting process.
