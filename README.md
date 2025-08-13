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
curl -X POST -H "Content-Type: application/json" -d '{"title":"Test task 2", "status":"completed"}' http://localhost:8080/tasks
```
- Получение задач/задачи по айди или по статусу (в примере – completed)
```bash
curl http://localhost:8080/tasks
curl http://localhost:8080/tasks/1
curl http://localhost:8080/tasks\?status\=completed
```
