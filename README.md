# Secgram

### Work in progress

Little messanger with E2E encryption written in Golang.

Used:

- Golang with default web server
- PostgreSQL for saving data (sqlx)
- WebSocket for real-time messaging
- (RSA for encryption (TODO))

Building backend:

```bash
go build -o secgram ./cmd/server/main.go
# or with make (see Makefile)
make build
# and run (default port :8080)
./secgram
```

Env:

```env
# for generate jwt tokens
JWT_KEY=pomodoro
DATABASE_URL=postgres://postgres:postgres@localhost:5432/secgram?sslmode=disable
```
