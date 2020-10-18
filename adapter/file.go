package adapter

import (
	"os"
)

type File struct {
	FilePath string
}

func (logger *File) IsAsync() bool {
	return false
}

func (logger *File) Info(message string) error {
	return logger.log(message)
}

func (logger *File) Debug(message string) error {
	return logger.log(message)
}

func (logger *File) Warning(message string) error {
	return logger.log(message)
}

func (logger *File) Critical(message string) error {
	return logger.log(message)
}

func (logger *File) log(message string) error {
	f, err := os.OpenFile(logger.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	f.WriteString(message + "\r\n")
	f.Close()

	return nil
}
