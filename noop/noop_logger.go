// Package noop provides a no operation (noop) implementation of the
// loggerrific interface that can be used as a default implementation if no
// other implementation is provided.
package noop

import (
	"github.com/CallumKerson/loggerrific"
)

type NoOpLogger struct {
}

func New() *NoOpLogger {
	noOpLogger := &NoOpLogger{}
	return noOpLogger
}

func (l *NoOpLogger) WithField(key string, value interface{}) loggerrific.Entry {
	return &NoOpEntry{}
}

func (l *NoOpLogger) WithFields(fields map[string]interface{}) loggerrific.Entry {
	return &NoOpEntry{}
}

func (l *NoOpLogger) WithError(err error) loggerrific.Entry {
	return &NoOpEntry{}
}

func (l *NoOpLogger) SetLevelDebug() {}

func (l *NoOpLogger) SetLevelInfo() {}

func (l *NoOpLogger) SetLevelWarn() {}

func (l *NoOpLogger) SetLevelError() {}

func (l *NoOpLogger) IsDebugEnabled() bool {
	return false
}

func (l *NoOpLogger) IsInfoEnabled() bool {
	return false
}

func (l *NoOpLogger) IsWarnEnabled() bool {
	return false
}

func (l *NoOpLogger) IsErrorEnabled() bool {
	return false
}

func (l *NoOpLogger) Debugf(format string, args ...interface{}) {}
func (l *NoOpLogger) Infof(format string, args ...interface{})  {}
func (l *NoOpLogger) Warnf(format string, args ...interface{})  {}
func (l *NoOpLogger) Errorf(format string, args ...interface{}) {}

func (l *NoOpLogger) Debugln(args ...interface{}) {}
func (l *NoOpLogger) Infoln(args ...interface{})  {}
func (l *NoOpLogger) Warnln(args ...interface{})  {}
func (l *NoOpLogger) Errorln(args ...interface{}) {}

type NoOpEntry struct {
}

func (e *NoOpEntry) Debugf(format string, args ...interface{}) {}
func (e *NoOpEntry) Infof(format string, args ...interface{})  {}
func (e *NoOpEntry) Warnf(format string, args ...interface{})  {}
func (e *NoOpEntry) Errorf(format string, args ...interface{}) {}

func (e *NoOpEntry) Debugln(args ...interface{}) {}
func (e *NoOpEntry) Infoln(args ...interface{})  {}
func (e *NoOpEntry) Warnln(args ...interface{})  {}
func (e *NoOpEntry) Errorln(args ...interface{}) {}
