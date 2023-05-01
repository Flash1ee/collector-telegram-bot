package internal

import "github.com/google/uuid"

type Logger interface {
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
}

type UUID = uuid.UUID
