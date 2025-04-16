# 🔍 Fuzzy Search App

A full-stack application for uploading files, performing fuzzy search on their contents, and expanding sentence context. Built with:

- **Frontend**: React, Typescript, TailwindCSS
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

## 🔧 Areas for Improvement

### Backend
- **Storage Abstraction**: Introduce a `StorageInterface` to decouple the storage implementation, enabling future migration to relational databases or other storage systems.
- **Search Configuration**: Add support for configurable fuzzy search parameters (e.g., Levenshtein threshold, case sensitivity).
- **Search History**: Implement a search history feature (e.g. linked-list-based) after persistence is in place.
- **Format Support**: Extend support for additional document formats (e.g., PDF, DOCX).
- **Preprocessing**: Add document preprocessing pipelines to improve input quality (e.g., cleaning, normalization).

### Frontend
- **Layout Refinement**: Improve the header layout to better balance space between actions like uploading and searching.
- **Result Display**: Make `SearchResultCard` content scrollable. Prevent the expand button from shifting UI elements.
- **Navigation**: Add a navigation bar inspired by the OpenAI UI to switch between document-specific search sessions.
- **Context Navigation**: Add a button to jump to the document position corresponding to a search result.

### CI/CD and DevOps
- **HTTPS Support**: Set up SSL using Certbot for secure connections and automatic certificate renewal.
- **Testing**: Expand test coverage, including unit tests for core components like the LRU cache.
- **Automation**: Add GitHub Actions workflows to run tests and checks on pull requests.

