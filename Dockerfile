FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go test ./... && CGO_ENABLED=0 go build -o /out/server ./cmd/server && CGO_ENABLED=0 go build -o /out/migrate ./cmd/migrate
FROM alpine:3.21
WORKDIR /app
RUN adduser -D -u 10001 app && mkdir -p /app/data && chown -R app:app /app
COPY --from=build /out/server /app/server
COPY --from=build /out/migrate /app/migrate
USER app
ENV APP_DB_PATH=/app/data/service.sqlite3
EXPOSE 8080
ENTRYPOINT ["/app/server"]
