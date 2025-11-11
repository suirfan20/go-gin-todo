# go-gin-todo

Small demo service built with **Go + Gin** to showcase **CI/CD + DevOps**:
- Lint & Test via **GitHub Actions**
- Containerized via **Docker** (multi-stage, distroless)
- One-click deploy to **Render (Free)** via **Deploy Hook**

## Endpoints
- `GET /healthz` → `{ "status": "ok" }`
- `GET /v1/todos`
- `POST /v1/todos` → `{ "title": "..." }`
- `DELETE /v1/todos/:id`

## Local Dev
```bash
make init
make run
# or docker:
docker compose up --build
