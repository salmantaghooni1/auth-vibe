FROM golang:1.24-bullseye as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o server ./cmd/main.go


FROM debian:bullseye
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive \
    rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml
COPY --from=builder /app/docs /app/docs

ENV APP_ENV docker
ENV PORT ${Port}

CMD [ "/app/server" ]
