# Ping Results

## Функционал
- Пинг разных портов контейнеров через каждый заданный интервал времени с помощью Docker API.
- Получение и отправка результатов пинга через Backend API.
- Сохранение и обновление результатов пинга в базе данных PostgreSQL.
- React-приложение для отображения данных.

## Запуск проекта

### Требования

- Docker
- Docker Compose

### Шаги для запуска

1. Создайте файл `.env` в директории, где находится `docker-compose.yml`, и добавьте необходимые переменные окружения. Пример содержимого для `.env` можно найти в файле `.env.example`.

2. Для запуска проекта используйте следующую команду:
   ```bash
   docker-compose --env-file .env up --build
   
