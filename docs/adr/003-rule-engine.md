# ADR 003 — Rule Engine

## Status
Accepted

## Context
Rules map events to gamification actions (award points, grant badge, etc.) using a safe expression language.

## Decision
- Use `antonmedv/expr` for condition evaluation.
- Per-rule evaluation timeout: 100 ms.
- Max complexity: 1000 AST nodes.
- Pre-compile expressions; cache compiled rules; async evaluation via NATS JetStream.
- Actions: award_points, deduct_points, grant_badge, unlock_achievement, update_streak, increment_quest_progress, notify.

## Consequences
expr is sandboxed but not Turing-complete in practice; complex logic may need native Go handlers.
