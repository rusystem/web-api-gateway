# Этап 1: Сборка приложения
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum отдельно для кэширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем все остальные файлы проекта
COPY . .

# Сборка приложения
RUN go build -o web-api-gateway ./cmd/main.go

# Этап 2: Минимальный образ
FROM debian:bullseye-slim

# Устанавливаем необходимые зависимости (если нужны, например, psql)
RUN apt-get update && apt-get install -y libpq-dev && rm -rf /var/lib/apt/lists/*

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем собранное приложение из предыдущего этапа
COPY --from=builder /app/web-api-gateway /app/web-api-gateway

# Настраиваем права доступа
RUN chmod +x web-api-gateway

# Запуск приложения
CMD ["./web-api-gateway"]