# lo_test_task
## Билд
```bash
docker build -t test_task_sssyrbu .
```
## Деплой
API работает на http://localhost:8080
```bash
docker run -p 8080:8080 test_task_sssyrbu
```
## Функционал
- Создание задач (обязательный параметр – title, опционально можно прописать status – pending, in_progress, completed (default: pending)
```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"Test task 1"}' http://localhost:8080/tasks
```
```bash
curl -X POST -H "Content-Type: application/json" -d '{"title":"Test task 2", "status":"completed"}' http://localhost:8080/tasks
```
- Получение задач/задачи по айди или по статусу (в примере – completed)
```bash
curl http://localhost:8080/tasks
```
```bash
curl http://localhost:8080/tasks/1
```
```bash
curl http://localhost:8080/tasks\?status\=completed
```
- Также на стороне сервера пишутся логи, например:
```bash
2025/08/13 18:04:15 running on :8080
2025/08/13 18:04:25 [2025-08-13 18:04:25] VALIDATION TaskID:  0 Invalid status - defaulting to 'pending'
2025/08/13 18:04:25 [2025-08-13 18:04:25] CREATE  TaskID:  1 Task created: Test task 1
2025/08/13 18:04:36 [2025-08-13 18:04:36] shutdown TaskID:  0 shutdown
2025/08/13 18:04:36 Server stopped gracefully
```
