# Stage 1: Build
FROM node:20-alpine AS builder
WORKDIR /app
COPY . .
RUN npm install && npm run build

# Stage 2: Serve with dev server (for hot reload)
FROM node:20-alpine
WORKDIR /app
COPY --from=builder /app /app
EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host"]
