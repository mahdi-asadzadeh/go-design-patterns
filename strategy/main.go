package main

import "fmt"

// Abstraction
type LoggerStrategy interface {
	Write(string) error
}

// Strategy 1
type FileLogger struct{}

func (s *FileLogger) Write(data string) error {
	fmt.Println("Writed to file...")
	return nil
}

// Strategy 2
type DbLogger struct{}

func (s *DbLogger) Write(data string) error {
	fmt.Println("Writed to db..")
	return nil
}

type Logger struct {
	logger LoggerStrategy
}

func (l *Logger) Log(data string) error {
	l.logger.Write(data)
	return nil
}

// Client
func main() {
	lo := Logger{logger: &FileLogger{}}
	lo.Log("Test file log")

	lo = Logger{logger: &DbLogger{}}
	lo.Log("Test db log")
}
