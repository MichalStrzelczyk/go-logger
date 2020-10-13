package formatter

import (
	"time"
)

type Time struct {}

func (formatter *Time) Format(message string, logLevel string) string {
	time := time.Now().UTC().Format("2006-01-02 15:04:05.000000")

	return "[" + time  + "] " + message
}
