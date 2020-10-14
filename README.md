# Logger

This library is used for simple logging of messages to several sources.

## 1. Adapters

Three logging adapters are available to use:
- syslog
- console
- file

If you want to create your own adapter, you can do it by implementing `AdapterInterface`
```
type LoggerInterface interface {
	IsAsync() bool
	Info(message string) error
	Critical(message string) error
}
```  

## 2. Formatters

All log messages can be decorated by formatters. For this moment there are two default formatters available:
- time
- significance 

If you want to create your own adapter, you can do it by implementing `FormatterInterface`

```
type FormatterInterface interface {
	Format(message string,logLevel string) string
}
```

An example log message with formatters:
```
// [Time formatter]          // [significance formatter]
[2020-10-14 06:53:33.267086] [critical] Log message
```

**CAUTION** The order of the formatters does matter

## 3. Usage

```
package main

import (
	myLogger "github.com/digitalfarm/logger"
)

func main() {
	myLogger:= myLogger.CreateLoggerManager()
	myLogger.Critical("Error message ")
}
```