package logger

import (
	"github.com/digitalfarm/logger/adapter"
	"github.com/digitalfarm/logger/formatter"
	"log/syslog"
)

const INFO = "info"
const DEBUG = "debug"
const WARNING = "warning"
const CRITICAL = "critical"

type LoggerInterface interface {
	IsAsync() bool
	Debug(message string) error
	Info(message string) error
	Warning(message string) error
	Critical(message string) error
}

type FormatterInterface interface {
	Format(message string,logLevel string) string
}

type LoggerManager struct {
	Adapters   []LoggerInterface
	Formatters []FormatterInterface
}

func (loggerManager *LoggerManager) AddAdapter(adapter LoggerInterface) {
	loggerManager.Adapters = append(loggerManager.Adapters, adapter)
}

func (loggerManager *LoggerManager) AddFormatter(formatter FormatterInterface) {
	loggerManager.Formatters = append(loggerManager.Formatters, formatter)
}

func (loggerManager *LoggerManager) Format(message string, logLevel string) string {
	for _, formatter := range loggerManager.Formatters {
		message = formatter.Format(message, logLevel)
	}

	return message
}

func (loggerManager *LoggerManager) Info(message string) error {
	message = loggerManager.Format(message, INFO)
	for _, logAdapter := range loggerManager.Adapters {
		if logAdapter.IsAsync() == true {
			go logAdapter.Info(message)
		} else {
			err := logAdapter.Info(message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (loggerManager *LoggerManager) Debug(message string) error {
	message = loggerManager.Format(message, DEBUG)
	for _, logAdapter := range loggerManager.Adapters {
		if logAdapter.IsAsync() == true {
			go logAdapter.Debug(message)
		} else {
			err := logAdapter.Debug(message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (loggerManager *LoggerManager) Warning(message string) error {
	message = loggerManager.Format(message, WARNING)
	for _, logAdapter := range loggerManager.Adapters {
		if logAdapter.IsAsync() == true {
			go logAdapter.Warning(message)
		} else {
			err := logAdapter.Warning(message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (loggerManager *LoggerManager) Critical(message string) error {
	message = loggerManager.Format(message, CRITICAL)
	for _, logAdapter := range loggerManager.Adapters {
		if logAdapter.IsAsync() == true {
			go logAdapter.Critical(message)
		} else {
			err := logAdapter.Critical(message)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func CreateLoggerManager() *LoggerManager {
	syslogWriter, _ := syslog.Dial("", "", syslog.LOG_DEBUG, "my-log")
	syslogLogger := adapter.SyslogLogger{syslogWriter}
	consoleLogger := adapter.ConsoleLogger{}
	fileLogger := adapter.FileLogger{"/tmp/aaa.log"}

	timeFormatter := formatter.Time{}
	significanceFormatter := formatter.Significance{}

	logManager := LoggerManager{}
	logManager.AddAdapter(&consoleLogger)
	logManager.AddAdapter(&syslogLogger)
	logManager.AddAdapter(&fileLogger)
	logManager.AddFormatter(&significanceFormatter)
	logManager.AddFormatter(&timeFormatter)

	return &logManager
}
