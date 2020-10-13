package adapter

import (
	"log/syslog"
)

type SyslogLogger struct {
	Syslog *syslog.Writer
}

func (logger *SyslogLogger) IsAsync() bool {
	return false
}

func (logger *SyslogLogger) Info(message string) error {
	logger.Syslog.Info(message)

	return nil
}

func (logger *SyslogLogger) Critical(message string) error {
	logger.Syslog.Crit(message)

	return nil
}
