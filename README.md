# DragonDB Init & Cleanup Tool

> Утилита для инициализации и удаления таблиц в PostgreSQL, с выводом системной информации и поддержкой логгирования.

## 🧩 Назначение

Этот инструмент:
- Проверяет наличие таблиц в базе данных
- Создаёт отсутствующие таблицы
- После проверки — удаляет их
- Показывает информацию о системе (CPU, память)

Полезно для CI/CD пайплайнов, тестирования схем БД и демонстрационных целей.

---

## 🚀 Установка

```bash
git clone https://github.com/vangdevops/test.git
cd test
go build -o dragon .
```
---
## ⚙️ Переменные окружения

Перед запуском установите параметры подключения:
```bash
Переменная	Назначение	Обязательная

DBUSER	Пользователь PostgreSQL	  ✅
DBPASS	Пароль PostgreSQL	  ✅
DBHOST	Хост (например localhost) ✅
DBNAME	Название базы данных 	  ✅
```

Если переменные не заданы — используются значения по умолчанию из pkg.

---
## 🛠 Использование
```bash
./dragon
```

Программа выведет информацию о CPU, памяти, проверит таблицы, создаст отсутствующие, затем удалит все.

---
## 📌 Пример вывода
```bash
██████╗ ██████╗  █████╗  ██████╗  ██████╗ ███╗   ██╗
INFO CPU: Intel(R) Xeon(R) | Memory: 2048MB
INFO Database version: PostgreSQL 14.0
DEBUG Created table: users
INFO Checking Successfully! in: 137ms
DEBUG Deleted table: users
INFO Deleted Successfully! in: 98ms
```

---

## 📄 Лицензия

MIT © vangdevops
