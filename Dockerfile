# Используем официальный образ Go
FROM golang:1.23-alpine

# Устанавливаем зависимости
RUN apk add --no-cache git

# Устанавливаем рабочую директорию
WORKDIR .

# Копируем файлы
COPY . .

# Загружаем зависимости
RUN go mod download

# Сборка приложения
RUN go build -o main ./main.go

# Назначаем права на выполнение
RUN chmod +x main

# Запуск приложения
CMD ["./main"]
