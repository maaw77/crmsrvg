# Сервис учета горюче-смазочных материалов
Микросервис, предоставляющий данные учета горюче-смазочных материалов в JSON-формате.
## Запуск
1. Установить docker 
2. В папке с проектом создать файл `.env`, содержащий параметры соединения с сервером PostgreSQL.<br>
Пример:
```text
POSTGRES_DB=postgres
POSTGRES_USER=postgres
POSTGRES_PASSWORD=************
HOST_DB=db
PORT_DB=5432
```
3. Обновить ./config/config.yaml, задав свою конфигурацию серверов.
 Пример:
```yaml
server:
  Addr: 0.0.0.0:8080
  WriteTimeout: 25 # Seconds
  ReadTimeout: 25 # Seconds
  IdleTimeout: 70 # Seconds
  ShutdownTimeout: 25 # Seconds
db:
  DB: postgres
  Ueser: postgres
  Password: crmpassword
  Host: db
  Port: 5432
  PoolMaxConns: 7
```
4. В папке с проектом выполнить командy
```commandline
 docker compose up -d
```
## Описание API
`http://localhost:8080/api/v1/swagger/index.html`