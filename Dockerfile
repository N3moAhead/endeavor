FROM node:20-alpine AS frontend
WORKDIR /app

# Copy package files and install dependencies
COPY package.json package-lock.json ./
RUN npm install

# Copy tailwind source files
COPY web/input.css ./web/input.css
COPY web/templates ./web/templates

# Build and minify CSS
RUN mkdir -p ./web/static/css
RUN npx tailwindcss -i ./web/input.css -o ./web/static/css/style.css --minify

# Stage 2: Build Go Binary
FROM docker.io/library/golang:1.24-alpine AS backend
WORKDIR /app

# Copy go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -ldflags="-s -w" -o /app/endeavor ./cmd/endeavor/main.go

# Stage 3: Final Production Image
FROM alpine:latest
WORKDIR /app

# Copy the binary from the backend stage
COPY --from=backend /app/endeavor .

# Copy the static assets from the frontend stage
COPY --from=frontend /app/web/static ./web/static

# Copy the templates
COPY web/templates ./web/templates

EXPOSE 9090

# Command to run the application
CMD ["./endeavor"]
