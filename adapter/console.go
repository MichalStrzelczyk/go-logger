package adapter

import (
	"fmt"
)

type ConsoleLogger struct{}

func (logger *ConsoleLogger) IsAsync() bool {
	return false
}

func (logger *ConsoleLogger) Info(message string) error {
	fmt.Println(message)

	return nil
}

func (logger *ConsoleLogger) Critical(message string) error {
	fmt.Println(message)

	return nil
}
