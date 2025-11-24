# ---------------------------
# Stage 1: Build
# ---------------------------
FROM golang:1.25-alpine AS builder

ARG SERVICE
ENV SERVICE=${SERVICE}

WORKDIR /srv

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /srv/services/${SERVICE}/cmd
RUN go build -o /srv/main .

RUN mkdir -p /srv/migrations && cp -r /srv/services/${SERVICE}/migrations/* /srv/migrations/

# ---------------------------
# Stage 2: Run
# ---------------------------
FROM alpine:latest

WORKDIR /srv

COPY --from=builder /srv/main ./main
COPY --from=builder /go/bin/migrate /usr/local/bin/migrate
COPY --from=builder /srv/migrations /srv/migrations

ENV PATH="/usr/local/bin:${PATH}"

EXPOSE 80

CMD ["./main"]
