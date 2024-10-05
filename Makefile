run:
	make build && ./secgram
build:
	go build -o secgram -ldflags="-s -w" ./cmd/api/main.go

migrate:
	DATABASE_URL=postgres://zethange:	tomato@localhost:5432/secgram?sslmode=disable dbmate up