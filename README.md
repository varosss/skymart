# 🛒 SkyMart

SkyMart — это микросервисная backend-платформа e-commerce, реализованная на Go с применением принципов **Clean Architecture**, **Domain-Driven Design (DDD)** и **Event-Driven Architecture**.

Проект демонстрирует архитектурный подход к построению распределённых систем, изоляцию доменной логики и взаимодействие сервисов через gRPC и Kafka.

---

## 🏗 Архитектура

Проект построен на следующих принципах:

- Clean Architecture
- Domain-Driven Design (DDD)
- Hexagonal Architecture (Ports & Adapters)
- Event-Driven Architecture
- Микросервисный подход

Каждый сервис:
- изолирован
- имеет собственную доменную модель
- взаимодействует с другими сервисами через gRPC (синхронно)
- публикует доменные события через Kafka (асинхронно)

---

## 🧩 Микросервисы

| Сервис  | Назначение |
|----------|------------|
| User     | Регистрация пользователей |
| Auth     | JWT-аутентификация, refresh tokens, роли |
| Seller   | Управление продавцами |
| Buyer    | Управление покупателями |
| Product  | CRUD товаров + публикация |
| Order    | Создание заказов |
| Billing  | Генерация и управление инвойсами |
| Payment  | Обработка платежей |

---

## ⚙️ Технологии

### Язык
- Go

### Коммуникация
- gRPC
- Protocol Buffers
- Kafka

### Хранение данных
- PostgreSQL
- GORM

### Безопасность
- JWT (RS256)
- Refresh Token rotation
- bcrypt

### Инфраструктура
- Docker
- Nginx
- Миграции базы данных

---

## 📂 Структура проекта
internal/
├── domain
│ ├── entity
│ ├── valueobject
│ ├── event
│ └── repository interfaces
│
├── application
│ ├── usecase
│ └── ports
│
└── infrastructure
├── http / grpc handlers
├── kafka adapters
├── gorm repositories
└── security implementations


---

## 🧠 Архитектурные решения

### Domain Layer
- Entities
- Value Objects
- Domain Events
- Repository interfaces
- Domain Errors

Домен полностью изолирован от инфраструктуры.

### Application Layer
- Use Cases
- Оркестрация бизнес-логики
- Работа через абстракции (Ports)

### Infrastructure Layer
- Реализации репозиториев (GORM)
- gRPC / HTTP контроллеры
- Kafka producer / consumer
- JWT signer / verifier
- Password hashing

---

## 🔄 Event-Driven Workflow

Пример бизнес-процесса:

1. Order публикует `OrderCreated`
2. Billing создаёт Invoice
3. Payment создаёт Payment
4. При успешной оплате публикуется `PaymentSucceeded`
5. Billing обновляет статус Invoice

События отделены от инфраструктурного представления через мапперы.

---

## 🔐 Аутентификация

- Access Token (JWT, RS256)
- Refresh Token
- Ролевая модель (admin, seller, buyer)
- Middleware для авторизации

---

## 🧪 Тестирование

- Unit-тесты для UseCase
- Использование fakes вместо моков
- Изоляция бизнес-логики от инфраструктуры

---

## 🚀 Цель проекта

- Продемонстрировать понимание DDD
- Показать умение проектировать микросервисную архитектуру
- Реализовать event-driven взаимодействие
- Строго разделить бизнес-логику и инфраструктуру
- Продемонстрировать production-ready подход к проектированию

---

## 📌 Возможные улучшения

- CI/CD pipeline
- Docker Compose для локального запуска всей системы
- Observability (Prometheus, Grafana)
- Centralized logging
- Rate limiting
- Circuit breaker
- Saga orchestration

---

## 👤 Автор

Varos Simonyan  
Backend Developer (Go)
