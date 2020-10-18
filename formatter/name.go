package formatter

type Name struct {
	ProjectName string
}

func (formatter *Name) Format(message string, logLevel string) string {
	return "[" + formatter.ProjectName + "] " + message
}
