package adapter

import (
	"log/syslog"
)

type Syslog struct {
	Syslog *syslog.Writer
}

func (logger *Syslog) IsAsync() bool {
	return false
}

func (logger *Syslog) Info(message string) error {
	logger.Syslog.Info(message)

	return nil
}

func (logger *Syslog) Debug(message string) error {
	logger.Syslog.Debug(message)

	return nil
}

func (logger *Syslog) Warning(message string) error {
	logger.Syslog.Warning(message)

	return nil
}

func (logger *Syslog) Critical(message string) error {
	logger.Syslog.Crit(message)

	return nil
}
