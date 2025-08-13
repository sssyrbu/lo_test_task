package main

import "strings"

type TaskService struct {
	repo   *TaskRepository
	logger *Logger
}

func NewTaskService(repo *TaskRepository, logger *Logger) *TaskService {
	return &TaskService{repo: repo, logger: logger}
}
func (s *TaskService) Create(title string, status string) Task {
	task := NewTask(title, status)
	s.repo.Save(task)
	s.logger.Log("CREATE", task.ID, "Task created: "+title)
	return task
}

func (s *TaskService) GetByID(id uint64) (Task, bool) {
	task, exists := s.repo.GetByID(id)
	if exists {
		s.logger.Log("READ", id, "Task retrieved")
	}
	return task, exists
}

func (s *TaskService) GetAll(status string) []Task {
	tasks := s.repo.GetAll()
	if status != "" {
		filtered := make([]Task, 0)
		for _, task := range tasks {
			if strings.EqualFold(task.Status, status) {
				filtered = append(filtered, task)
			}
		}
		return filtered
	}
	return tasks
}
