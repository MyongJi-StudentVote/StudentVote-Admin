include .envrc
.PHONY: all templ_generate tailwindcss build status up down
url="postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable"

all: tailwindcss templ_generate build

templ_generate:
	@templ generate

tailwindcss:
	@bun tailwindcss -i ./static/input.css -o ./static/output.css

build:
	@go build -o ./tmp/main ./cmd/web

status:
	@goose postgres ${url} status

up:
	@goose postgres ${url} up

down:
	@goose postgres ${url} down
