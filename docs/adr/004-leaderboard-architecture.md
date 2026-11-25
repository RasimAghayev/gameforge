# ADR 004 — Leaderboard Architecture

## Status
Accepted

## Context
Leaderboards must support top-N, player rank, and real-time updates with sub-ms reads.

## Decision
- PostgreSQL: leaderboard catalog (metadata).
- Redis sorted sets: key `leaderboard:{tenant_id}:{leaderboard_id}`, score = metric value, member = player_id.
- Operations: ZADD, ZREVRANGE, ZREVRANK, ZSCORE, DEL.
- Reset: DEL (or rename to archive) via scheduler for daily/weekly/monthly.
- Real-time: score update → NATS subject `leaderboard.{tenant}.{lb}.updated` → WebSocket push.
- Debounce: 100 ms batch window; push only if rank changed.

## Consequences
Redis is the source of truth for scores; PostgreSQL catalog is eventually consistent. Periodic resync jobs recommended.
