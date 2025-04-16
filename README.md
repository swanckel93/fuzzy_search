# 🔍 Fuzzy Search App

A full-stack application for uploading files, performing fuzzy search on their contents, and expanding sentence context. Built with:

- **Frontend**: React (or similar)
- **Backend**: Go
- **Fuzzy Search**: Custom algorithm with caching
- **API Docs**: Swagger (OpenAPI 3.0)

---

## 🚀 Features

- 📄 Upload text files for indexing
- 🔍 Perform fuzzy searches across uploaded documents
- 🧠 Expand sentence context
- ⚡ In-memory LRU caching for optimized search performance
- 🧾 Swagger-powered API docs

---

## 🚀 Running the App with Docker Compose

Start both the backend and frontend with:

```bash
docker-compose up --build
```
Then open your browser:

🌐 Frontend: http://localhost:5173
Note: Frontend runs in dev mode

⚙️ API: http://localhost:8080

📚 Swagger UI: Visit the Swagger JSON below in Swagger Editor

## 🧾 API Documentation (Swagger)
This project uses Swagger (OpenAPI 3.0) to document the backend.

To view it:
👉 [**Open in Swagger Editor**](https://editor.swagger.io/?url=https://raw.githubusercontent.com/swanckel93/fuzzy_search/main/backend/docs/swagger.json)

Make sure swagger.json is committed and up to date in the backend/docs directory.