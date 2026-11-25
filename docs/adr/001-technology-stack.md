# ADR 001 — Technology Stack

## Status
Accepted

## Context
The GameForge gamification server needs to serve multiple tenants with low-latency event ingestion, rule evaluation, leaderboards, and real-time updates.

## Decision
- Go 1.22 — concurrency, performance, single binary.
- Chi v5 — lightweight HTTP router.
- gRPC + protobuf — internal service communication.
- gorilla/websocket — real-time leaderboard/player updates.
- PostgreSQL 16 + pgx/v5 — config/rules, ACID, JSONB.
- MongoDB 7 + mongo-driver — high-throughput event storage.
- Redis 7 + go-redis/v9 — leaderboards via sorted sets, cache.
- NATS JetStream + nats.go — event bus, at-least-once delivery.
- expr — safe rule expression evaluation.
- golang-jwt/v5 + API keys + mTLS — auth.
- casbin — tenant-scoped RBAC.
- OpenTelemetry + Prometheus + Loki — observability.
- React + TS + Tailwind + shadcn/ui — admin UI.

## Consequences
Single Go binary for server/worker/scheduler simplifies deployment. Multi-store architecture increases operational complexity but matches access patterns.

## ADRs referenced
ADR 002 (multi-tenancy), ADR 003 (rule engine), ADR 004 (leaderboard architecture).
