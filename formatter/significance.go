package formatter

type Significance struct { }

func (formatter *Significance) Format(message string, logLevel string) string {
	return "[" + logLevel + "] " + message
}
