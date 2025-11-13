# ---- build stage ----
FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -ldflags="-s -w" -o server ./cmd/server

# ---- run stage ----
FROM gcr.io/distroless/base-debian12
WORKDIR /app
ENV PORT=8081
COPY --from=builder /app/server /app/server
EXPOSE 8081
USER nonroot:nonroot
ENTRYPOINT ["/app/server"]
