# 🧠 Task Manager

Task Manager — это микросервис для постановки и обработки задач в очереди. Он реализован по принципам **чистой архитектуры** и поддерживает:

- REST API (через Gin)
- gRPC интерфейс
- PostgreSQL хранилище
- Очередь задач с асинхронной обработкой
- Docker-сборку и контейнеризацию

---

## 📦 Стек технологий

- **Golang** - 1.24
- **Gin** — REST API
- **gRPC** — сервисный интерфейс
- **PostgreSQL** — хранилище задач
- **Docker / Docker Compose**
- **protobuf** — сериализация gRPC

---

## 🚀 Быстрый старт

### 1. Клонирование репозитория

```bash
git clone https://github.com/vegitobluefan/task-manager.git
cd task-manager
```

### 2. Запуск через Docker Compose

```bash
docker-compose up --build
```

Ожидается, что PostgreSQL и API поднимутся автоматически. REST API будет доступно на:  
`http://localhost:8080`

---

## 🧪 Примеры запросов

### ➕ Создать задачу

```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"type": "sleep", "payload": "5"}'
```

### 🔍 Получить задачу по ID

```bash
curl http://localhost:8080/tasks/<id>
```

### 📋 Получить список всех задач

```bash
curl http://localhost:8080/tasks
```

---

## 📚 Структура проекта
```
.
├── cmd/                 # Точка входа (main.go)
├── api/                 # HTTP обработчики
├── grpc/                # gRPC сервер
├── dispatcher/          # Очередь задач
├── domain/              # Бизнес-логика (сущности + интерфейсы)
├── repository/          # Реализация интерфейса доступа к данным
├── usecase/             # Реализация бизнес-логики
├── proto/               # .proto файлы и gRPC код
├── docker-compose.yml
└── README.md
```

---

## ⚙️ Команды для отладки

### Пересобрать проект:
```bash
docker-compose up --build
```
### Подключиться к БД:
```bash
docker exec -it task_pg psql -U postgres -d taskdb
```

---

# Автор: [Аринов Данияр](https://github.com/vegitobluefan)