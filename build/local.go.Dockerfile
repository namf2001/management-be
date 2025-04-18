FROM golang:1.24.2-alpine

RUN apk add --no-cache git make bash curl

WORKDIR /api

# Copy go.mod, go.sum and vendor directory
COPY go.mod go.sum ./
COPY vendor/ ./vendor/

# Install air for live reloading
RUN go install github.com/cosmtrek/air@v1.40.4

# Install mockery for generating mocks
RUN go install github.com/vektra/mockery/v2@v2.20.0

# Install migrate for database migrations
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

# Install entimport for generating models
RUN go install ariga.io/entimport/cmd/entimport@latest

# Install specific version of Ent that matches go.mod
RUN go install entgo.io/ent/cmd/ent@latest

# Set environment variables
ENV CGO_ENABLED=0
ENV GO111MODULE=on

# Default command
CMD ["air", "-c", ".air.toml"]
