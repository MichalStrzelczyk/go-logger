package adapter

import (
	"os"
)

type FileLogger struct {
	FilePath string
}

func (logger *FileLogger) IsAsync() bool {
	return false
}

func (logger *FileLogger) Info(message string) error {
	return logger.log(message)
}

func (logger *FileLogger) Debug(message string) error {
	return logger.log(message)
}

func (logger *FileLogger) Warning(message string) error {
	return logger.log(message)
}

func (logger *FileLogger) Critical(message string) error {
	return logger.log(message)
}

func (logger *FileLogger) log(message string) error {
	f, err := os.OpenFile(logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.WriteString(message + "\r\n")
	f.Close()

	return nil
}
