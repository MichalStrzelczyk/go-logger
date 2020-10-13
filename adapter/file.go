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
	f, err := os.OpenFile(logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.WriteString(message + "\r\n")
	f.Close()

	return nil
}

func (logger *FileLogger) Critical(message string) error {
	f, err := os.OpenFile(logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.WriteString(message + "\r\n")
	f.Close()

	return nil
}
