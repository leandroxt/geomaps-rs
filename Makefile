include .env

local:
	go run ./cmd/web \
		-addr=":8080" \
		-dsn="postgres://postgres:${DB_PASSWORD}@localhost:5432/geomaps?sslmode=disable"