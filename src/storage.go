package main

import (
	"sync"
	"time"
)

var taskCounter uint64 = 0

type Task struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	Status    string    `json:"status"` // pending || in_progress || completed
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskRepository struct {
	sync.RWMutex
	tasks map[uint64]Task
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		tasks: make(map[uint64]Task),
	}
}

func (r *TaskRepository) Save(task Task) {
	r.Lock()
	defer r.Unlock()
	r.tasks[task.ID] = task
}

func (r *TaskRepository) GetByID(id uint64) (Task, bool) {
	r.RLock()
	defer r.RUnlock()
	task, exists := r.tasks[id]
	return task, exists
}

func (r *TaskRepository) GetAll() []Task {
	r.RLock()
	defer r.RUnlock()
	tasks := make([]Task, 0, len(r.tasks))
	for _, task := range r.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

func (s *TaskService) NewTask(title string, optionalStatus ...string) Task {
	status := "pending"
	if len(optionalStatus) > 0 {
		proposedStatus := optionalStatus[0]
		switch proposedStatus {
		case "pending", "in_progress", "completed":
			status = proposedStatus
		default:
			s.logger.Log("VALIDATION", taskCounter+1, "Invalid status - defaulting to 'pending'")
		}
	}

	taskCounter += 1
	return Task{
		ID:        taskCounter,
		Title:     title,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
