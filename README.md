# GameForge

Go-based gamification server platform. Central gamification engine for multiple games/apps/tenants — points/XP, badges, leaderboards, achievements, quests, levels, streaks, rewards — with REST, gRPC, WebSocket, and SDKs. Multi-tenant with PostgreSQL (config/rules), MongoDB (events/history), Redis (leaderboards/cache), and NATS JetStream (event bus).

## Tech stack

| Layer | Choice |
|-------|--------|
| Language | Go 1.22+ |
| HTTP router | Chi (chi/v5) |
| gRPC | google.golang.org/grpc + protobuf |
| WebSocket | gorilla/websocket |
| Config DB | PostgreSQL 16 (pgx/v5) |
| Events DB | MongoDB 7 (mongo-driver) |
| Cache / leaderboards | Redis 7 (go-redis/v9) |
| Message queue | NATS JetStream (nats.go) |
| Rule engine | expr (antonmedv/expr) |
| Auth | golang-jwt/v5 + API keys + mTLS |
| RBAC | casbin |
| Observability | OpenTelemetry + Prometheus + Loki |

## Repository structure

See `docs/ARCHITECTURE.md` and `docs/REPOSITORY_AUDIT.md`.

## Quick start (Docker-first)

No host runtime/toolchain is used. Everything builds and runs through Docker Compose.

```bash
make verify   # build image + run test stage
make up       # start Postgres, Mongo, Redis, NATS, Nginx, then the server
```

See `docs/DEPLOYMENT.md` and `docs/OPERATIONS.md` for details.

## Community standards

This repo follows the Contributor Covenant (see `CODE_OF_CONDUCT.md`). Bug reports, feature requests, and pull requests use the templates in `.github/`. License: MIT (`LICENSE`).

A short Azerbaijani summary is acceptable; primary language is English.
