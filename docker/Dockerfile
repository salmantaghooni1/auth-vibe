FROM golang:1.24.8 AS builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -v -o server ./cmd/main.go


FROM debian:buster-slim
RUN sed -i 's|deb.debian.org|archive.debian.org|g' /etc/apt/sources.list && \
    sed -i 's|security.debian.org|archive.debian.org|g' /etc/apt/sources.list && \
    echo 'Acquire::Check-Valid-Until "false";' > /etc/apt/apt.conf.d/99ignore-valid-until

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml
COPY --from=builder /app/docs /app/docs

ENV APP_ENV docker
ENV PORT ${Port}

CMD ["/app/server"]
