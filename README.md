# Ston Tactics backend

### Запуск сервера:
1. Переименуйте [.env.example](.env.example) -> .env
2. Настройте авторизацию: Заполните в .env все необходимые данные (VK_CLIENT_SECRET, VK_CLIENT_ID, GOOGLE_CLIENT_SECRET, GOOGLE_CLIENT_ID).
3. Установите и запустите [docker](https://docs.docker.com/engine/install/)
4. Запустите сервер:
```bash
docker compose up --build
```
## Документация:
> Swagger: http://localhost:8080/docs/index.html
