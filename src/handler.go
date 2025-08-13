package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type TaskHandler struct {
	service *TaskService
}

func NewTaskHandler(service *TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (h *TaskHandler) GetTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	tasks := h.service.GetAll(status)
	jsonResponse(w, http.StatusOK, tasks)
}

func (h *TaskHandler) GetTaskByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/tasks/")

	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid task id format"})
		return
	}

	task, exists := h.service.GetByID(id)
	if !exists {
		jsonResponse(w, http.StatusNotFound, map[string]string{"error": "Task not found"})
		return
	}
	jsonResponse(w, http.StatusOK, task)
}

func (h *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title  string `json:"title"`
		Status string `json:"status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
		return
	}

	task := h.service.Create(req.Title, req.Status)
	jsonResponse(w, http.StatusCreated, task)
}

func jsonResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}
