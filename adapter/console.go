package adapter

import (
	"fmt"
)

type Console struct{}

func (logger *Console) IsAsync() bool {
	return false
}

func (logger *Console) Info(message string) error {
	return logger.log(message)
}

func (logger *Console) Debug(message string) error {
	return logger.log(message)
}

func (logger *Console) Warning(message string) error {
	return logger.log(message)
}

func (logger *Console) Critical(message string) error {
	return logger.log(message)
}

func (logger *Console) log(message string) error {
	fmt.Println(message)

	return nil
}