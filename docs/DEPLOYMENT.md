# Deployment

Local: `make verify && make up` (Docker Compose with Postgres, Mongo, Redis, NATS, Nginx).
Production: build the multi-stage Docker image, deploy the `runtime` stage, configure env vars from a secret manager, run migrations, then start `/server`.
