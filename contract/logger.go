package contract

// ILogger is the logger interface
type ILogger interface {
	Fatal(string, ...interface{})
	Debug(string, ...interface{})
	Info(string, ...interface{})
	Error(string, ...interface{})
}