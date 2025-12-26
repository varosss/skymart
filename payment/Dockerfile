# Stage 1: Build
FROM golang:1.25-alpine AS builder

WORKDIR /srv

COPY go.mod go.sum ./
RUN go mod download

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

COPY . ./app
WORKDIR /srv/app
RUN go build -o main .

COPY ./migrations /srv/migrations

# Stage 2: Run
FROM alpine:latest
WORKDIR /srv

COPY --from=builder /srv/app/main .  
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /srv/migrations /srv/migrations

ENV PATH="/usr/local/bin:${PATH}"

EXPOSE 80

CMD ["./main"]
