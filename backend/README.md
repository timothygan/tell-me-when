# Backend

This directory contains the code for the Tell Me When backend API and worker services.

## Suggested Tech Stack
- Node.js (Express/NestJS) or Python (FastAPI)
- PostgreSQL
- BullMQ (Node.js) or Celery (Python) for scheduling
- LLM API integration (OpenAI, etc.)

## Suggested Structure
- `/src` - Source code
  - `/api` - API endpoints
  - `/models` - Database models/schemas
  - `/services` - Business logic, AI integration, notifications
  - `/workers` - Background jobs/schedulers
  - `/config` - Configuration files
- `/tests` - Unit/integration tests
