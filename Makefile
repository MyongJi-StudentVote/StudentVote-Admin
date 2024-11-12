.PHONY: all templ_generate tailwindcss build

all: templ_generate tailwindcss build

templ_generate:
	@templ generate

tailwindcss:
	@bun tailwindcss -i ./static/input.css -o ./static/output.css

build:
	@go build -o ./tmp/main ./cmd/web
