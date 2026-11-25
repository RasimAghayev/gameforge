# Architecture

GameForge is a modular monolith gamification engine. Unified server exposes REST, gRPC, and WebSocket. A background worker consumes NATS JetStream events and runs the rule engine; a scheduler handles recurring quests, leaderboard resets, and cleanup.

Domains: tenant, player, points, badge, achievement, leaderboard, quest, level, streak, reward, rule, event.
Storage: PostgreSQL (config/rules), MongoDB (events/history), Redis (leaderboards/cache).
Event pipeline: External Event → Webhook/API → NATS JetStream → Event Processor → Rule Engine → Gamification Actions → Leaderboard (Redis) → WebSocket / Webhook.
