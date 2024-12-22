# Use multi-stage build to optimize image size

# Stage 1: Build TailwindCSS assets
FROM oven/bun:latest AS tailwind-builder

WORKDIR /app

# Install dependencies
COPY package.json ./
RUN bun install

# Copy TailwindCSS configuration and source files
COPY tailwind.config.js ./
COPY static ./static

# Build TailwindCSS output
RUN bunx tailwindcss -i ./static/input.css -o ./static/output.css --minify

# Stage 2: Build Go application
FROM golang:1.22 AS go-builder

WORKDIR /app

# Install Templ and Go dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . ./

# Build Go binary
RUN CGO_ENABLED=0 go build -o app ./cmd/web

# Stage 3: Final production image
FROM gcr.io/distroless/base-debian11

WORKDIR /app

# Copy Go binary
COPY --from=go-builder /app/app ./

# Copy TailwindCSS assets
COPY --from=tailwind-builder /app/static ./static

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./app"]
