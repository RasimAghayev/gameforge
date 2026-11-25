# ADR 002 — Multi-tenancy

## Status
Accepted

## Context
GameForge serves multiple games/apps/tenants from a single deployment. Tenant data must be isolated.

## Decision
- All queries scoped by `tenant_id`.
- PostgreSQL: Row-Level Security (RLS) enabled on tenant-scoped tables.
- MongoDB: `tenant_id` in all documents + index.
- Redis: `tenant_id` in all keys (`leaderboard:{tenant_id}:{leaderboard_id}`).
- API keys and JWT scoped per tenant.

## Consequences
RLS adds operational overhead; integration tests must verify cross-tenant isolation.
