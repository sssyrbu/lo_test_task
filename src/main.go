package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	repo := NewTaskRepository()
	logger := NewAsyncLogger()
	service := NewTaskService(repo, logger)
	handler := NewTaskHandler(service)

	go logger.Start()

	http.HandleFunc("GET /tasks", handler.GetTasks)
	http.HandleFunc("POST /tasks", handler.CreateTask)
	http.HandleFunc("GET /tasks/{id}", handler.GetTaskByID)

	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("running on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log("startup", 0, "error: "+err.Error())
			log.Fatal(err)
		}
	}()

	<-quit
	logger.Log("shutdown", 0, "shutdown")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logger.Stop()
	if err := server.Shutdown(ctx); err != nil {
		logger.Log("forced shutdown", 0, "error: "+err.Error())
	}
	log.Println("Server stopped gracefully")
}
