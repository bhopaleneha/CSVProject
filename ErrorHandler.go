package main

import (
	"github.com/sirupsen/logrus"
)

// Event stores messages to log later
type Event struct {
	Useid   string
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.JSONFormatter{}

	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgValueMessage = Event{"userId", "Invalid value for user with name: %s: %v"}
)

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentId string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentId)
}
