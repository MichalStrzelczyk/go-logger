# Logger

This library can be used for simple logging of messages to several sources.

## 1. Adapters

Three logging adapters are available to use:
- syslog
- console
- file

If you want to create your own adapter, you can do it by implementing `AdapterInterface`
```
type LoggerInterface interface {
    IsAsync() bool
    Debug(message string) error
    Info(message string) error
    Warning(message string) error
    Critical(message string) error
}
```  

## 2. Formatters

All log messages can be decorated by formatters. For this moment there are two default formatters available:
- name
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
// [Name formatter] [Time formatter] [significance formatter] error message
[MyProject] [2020-10-18 12:45:14.106169] [critical] Jakiś tekst 

```

**CAUTION** The order of the formatters does matter

## 3. Usage

```
package main

import (
	myLogger "github.com/MichalStrzelczyk/go-logger"
)

func main() {
	myLogger:= myLogger.CreateLoggerManager()
	myLogger.Critical("Error message ")
}
```