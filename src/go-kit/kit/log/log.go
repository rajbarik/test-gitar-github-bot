package log

import (
	"fmt"
	"io"
	"log"
	"os"
)

// Logger interface for go-kit/kit/log
type Logger interface {
	Log(keyvals ...interface{}) error
}

// JSONLogger implements Logger interface
type JSONLogger struct {
	logger *log.Logger
}

// NewJSONLogger creates a new JSON logger
func NewJSONLogger(w io.Writer) Logger {
	if w == nil {
		w = os.Stdout
	}
	return &JSONLogger{
		logger: log.New(w, "", log.LstdFlags),
	}
}

// Log logs key-value pairs in JSON-like format
func (l *JSONLogger) Log(keyvals ...interface{}) error {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "MISSING_VALUE")
	}

	var pairs []string
	for i := 0; i < len(keyvals); i += 2 {
		key := fmt.Sprintf("%v", keyvals[i])
		value := fmt.Sprintf("%v", keyvals[i+1])
		pairs = append(pairs, fmt.Sprintf(`"%s":"%s"`, key, value))
	}

	l.logger.Printf("{%s}", fmt.Sprintf("%s", pairs))
	return nil
}

// LogfmtLogger implements Logger interface for logfmt format
type LogfmtLogger struct {
	logger *log.Logger
}

// NewLogfmtLogger creates a new logfmt logger
func NewLogfmtLogger(w io.Writer) Logger {
	if w == nil {
		w = os.Stderr
	}
	return &LogfmtLogger{
		logger: log.New(w, "", log.LstdFlags),
	}
}

// Log logs key-value pairs in logfmt format
func (l *LogfmtLogger) Log(keyvals ...interface{}) error {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "MISSING_VALUE")
	}

	var pairs []string
	for i := 0; i < len(keyvals); i += 2 {
		key := fmt.Sprintf("%v", keyvals[i])
		value := fmt.Sprintf("%v", keyvals[i+1])
		pairs = append(pairs, fmt.Sprintf("%s=%s", key, value))
	}

	l.logger.Printf("%s", fmt.Sprintf("%s", pairs))
	return nil
}
