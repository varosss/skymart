# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /srv
RUN mkdir ./app

WORKDIR /srv/app

ENV SERVICE_DIR=./app/auth-service

COPY $SERVICE_DIR/go.mod $SERVICE_DIR/go.sum ./

RUN go mod download
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY $SERVICE_DIR ./
RUN go build -o main .

# Stage 2: Run
FROM alpine:latest
WORKDIR /srv

COPY --from=builder /srv/app/main .

ENV PATH="/go/bin:${PATH}"

EXPOSE 8002
CMD ["./main"]
