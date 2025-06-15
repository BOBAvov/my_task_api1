# Task API (English)

## Getting Started

### Windows 11
1. Download the project from Github
2. In the console, navigate to the project root and run: `./cmd/app/app.exe`

### Any Operating System
1. Install Go compiler
2. Download the project from Github
3. In the console, navigate to the project root and run: `go run cmd/app/main.go`

## Configuration
The server port can be changed in `config.yaml` by modifying `localhost:8080` to your desired port (e.g., `localhost:8081`)

## API Endpoints

Postman is recommended for sending POST and GET requests.

### Available Endpoints

1. **Health Check**
   - GET `http://127.0.0.1:8080`
   - Response: `{"status": "connect"}` if server is running

2. **Add Task**
   - POST `http://127.0.0.1:8080/tasks/add`
   - Body: `{"name": "task_name"}`
   - Responses:
     - Success: `{"status": "add task"}`
     - Duplicate: `{"status": "task already exists"}`

3. **Get Task Details**
   - GET `http://127.0.0.1:8080/tasks/{task_name}`
   - Responses:
     - Not Found: `{"status": "task not found"}`
     - Success Example:
     ```json
     {
         "name": "to eat",
         "status": "completed",  // completed | in progress
         "lead_time_min": 5,     // 3-5 minutes
         "end_time": "2025-06-14T15:28:50.1430089+03:00"
     }
     ```

4. **List All Tasks**
   - GET `http://127.0.0.1:8080/tasks`
   - Returns list of all active tasks

5. **Delete Task**
   - POST `http://127.0.0.1:8080/tasks/delete`
   - Body: `{"name": "task_name"}`
   - Responses:
     - Success: `{"status": "delete task"}`
     - Not Found: `{"status": "task not found"}`

## Shutdown
Press `Ctrl + C` in the console running the server

---

# Task API (Русский)

## Запуск

### Windows 11
1. Скачайте проект с Github
2. В консоли в корне проекта выполните: `./cmd/app/app.exe`

### Любая операционная система
1. Установите компилятор Go
2. Скачайте проект с Github
3. В консоли в корне проекта выполните: `go run cmd/app/main.go`

## Конфигурация
Порт сервера можно изменить в файле `config.yaml`, заменив `localhost:8080` на нужный порт (например, `localhost:8081`)

## API Endpoints

Рекомендуется использовать Postman для отправки POST и GET запросов.

### Доступные эндпоинты

1. **Проверка работоспособности**
   - GET `http://127.0.0.1:8080`
   - Ответ: `{"status": "connect"}` если сервер запущен

2. **Добавление задачи**
   - POST `http://127.0.0.1:8080/tasks/add`
   - Тело запроса: `{"name": "название_задачи"}`
   - Ответы:
     - Успех: `{"status": "add task"}`
     - Дубликат: `{"status": "task already exists"}`

3. **Получение информации о задаче**
   - GET `http://127.0.0.1:8080/tasks/{название_задачи}`
   - Ответы:
     - Не найдено: `{"status": "task not found"}`
     - Пример успешного ответа:
     ```json
     {
         "name": "to eat",
         "status": "completed",  // completed - выполнена | in progress - в процессе
         "lead_time_min": 5,     // от 3 до 5 минут
         "end_time": "2025-06-14T15:28:50.1430089+03:00"
     }
     ```

4. **Список всех задач**
   - GET `http://127.0.0.1:8080/tasks`
   - Возвращает список всех активных задач

5. **Удаление задачи**
   - POST `http://127.0.0.1:8080/tasks/delete`
   - Тело запроса: `{"name": "название_задачи"}`
   - Ответы:
     - Успех: `{"status": "delete task"}`
     - Не найдено: `{"status": "task not found"}`

## Завершение работы
Нажмите `Ctrl + C` в консоли, где запущен сервер
