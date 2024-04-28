package logger

type Level int8

const (
	Debug Level = iota - 1
	Info
	Warn
	Error
)
