include .env

local:
	go run ./cmd/web \
		-addr=":8080" \
		-dsn="postgres://postgres:${DB_PASSWORD}@localhost:5432/geomaps?sslmode=disable" \
		-mapsURL="https://maps.googleapis.com/maps/api" \
		-mapsKey=${MAPS_API_KEY}