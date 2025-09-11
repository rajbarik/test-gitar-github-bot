package ctxlog

import (
	"context"
	"log"
)

// Info logs an info message with context
func Info(ctx context.Context, msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

// Debug logs a debug message with context
func Debug(ctx context.Context, msg string, args ...interface{}) {
	log.Printf("[DEBUG] "+msg, args...)
}

// Error logs an error message with context
func Error(ctx context.Context, msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}

// Warn logs a warning message with context
func Warn(ctx context.Context, msg string, args ...interface{}) {
	log.Printf("[WARN] "+msg, args...)
}
