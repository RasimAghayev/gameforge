FROM golang:1.27-alpine AS base
RUN apk add --no-cache git ca-certificates tzdata && update-ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

FROM base AS deps
COPY . .
RUN go mod verify

FROM base AS build
COPY --from=deps /app /app
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /server ./cmd/server
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /worker ./cmd/worker
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /scheduler ./cmd/scheduler

FROM base AS test
COPY --from=deps /app /app
RUN go test ./... -race -cover

FROM gcr.io/distroless/static-debian12:nonroot AS runtime
COPY --from=build /server /server
COPY --from=build /worker /worker
COPY --from=build /scheduler /scheduler
EXPOSE 8080
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 CMD ["/server", "healthz"] || exit 1
ENTRYPOINT ["/server"]
