package main

import (
	"log"
	"time"
)

type Logger struct {
	logChan chan LogEntry
}

type LogEntry struct {
	Timestamp time.Time
	Action    string
	TaskID    uint64
	Details   string
}

func NewAsyncLogger() *Logger {
	return &Logger{
		logChan: make(chan LogEntry, 100),
	}
}

func (l *Logger) Start() {
	for entry := range l.logChan {
		log.Printf("[%s] %-7s %-8s %v %s",
			entry.Timestamp.Format("2006-01-02 15:04:05"),
			entry.Action,
			"TaskID:",
			entry.TaskID,
			entry.Details,
		)
	}
}

func (l *Logger) Log(action string, taskID uint64, details string) {
	l.logChan <- LogEntry{
		Timestamp: time.Now(),
		Action:    action,
		TaskID:    taskID,
		Details:   details,
	}
}

func (l *Logger) Stop() {
	close(l.logChan)
}
