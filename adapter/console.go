package adapter

import (
	"fmt"
)

type ConsoleLogger struct{}

func (logger *ConsoleLogger) IsAsync() bool {
	return false
}

func (logger *ConsoleLogger) Info(message string) error {
	return logger.log(message)
}

func (logger *ConsoleLogger) Critical(message string) error {
	return logger.log(message)
}

func (logger *ConsoleLogger) log(message string) error {
	fmt.Println(message)

	return nil
}