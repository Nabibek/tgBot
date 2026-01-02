package logger

import (
	"fmt"
	"log"
	"os"
	"sync"
)

type Logger struct {
	fileLogger *log.Logger
	console    bool
	mu         sync.Mutex
}

func New(logFile string, consoleOutput bool) (*Logger, error) {
	file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open log file: %w", err)
	}

	return &Logger{
		fileLogger: log.New(file, "", log.LstdFlags),
		console:    consoleOutput,
	}, nil
}

func (l *Logger) Info(format string, args ...interface{}) {
	l.log("INFO", format, args...)
}

func (l *Logger) Error(format string, args ...interface{}) {
	l.log("ERROR", format, args...)
}

func (l *Logger) Warn(format string, args ...interface{}) {
	l.log("WARN", format, args...)
}

func (l *Logger) Debug(format string, args ...interface{}) {
	l.log("DEBUG", format, args...)
}

func (l *Logger) log(level, format string, args ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	message := fmt.Sprintf("[%s] %s", level, fmt.Sprintf(format, args...))

	l.fileLogger.Println(message)

	if l.console {
		fmt.Println(message)
	}
}

func (l *Logger) Close() {
	// Файл закрывается автоматически при выходе из программы
}
