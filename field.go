package logger

import (
	"go.uber.org/zap"
)

type Field struct {
	Key   string
	Value string
}

type Fields []Field

func (f Fields) zap() []zap.Field {
	var fields = make([]zap.Field, len(f))

	for index, field := range f {
		fields[index] = zap.String(field.Key, field.Value)
	}

	return fields
}
