// Package noop provides a no operation (noop) implementation of the
// loggerrific interface that can be used as a default implementation if no
// other implementation is provided.
package noop

import (
	"github.com/CallumKerson/loggerrific"
)

type NoOpLogger struct{}

func New() *NoOpLogger {
	noOpLogger := &NoOpLogger{}
	return noOpLogger
}

func (l *NoOpLogger) WithField(key string, value any) loggerrific.Entry {
	return &NoOpEntry{}
}

func (l *NoOpLogger) WithFields(fields map[string]any) loggerrific.Entry {
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

func (l *NoOpLogger) Debugf(format string, args ...any) {}
func (l *NoOpLogger) Infof(format string, args ...any)  {}
func (l *NoOpLogger) Warnf(format string, args ...any)  {}
func (l *NoOpLogger) Errorf(format string, args ...any) {}

func (l *NoOpLogger) Debugln(args ...any) {}
func (l *NoOpLogger) Infoln(args ...any)  {}
func (l *NoOpLogger) Warnln(args ...any)  {}
func (l *NoOpLogger) Errorln(args ...any) {}

type NoOpEntry struct{}

func (e *NoOpEntry) Debugf(format string, args ...any) {}
func (e *NoOpEntry) Infof(format string, args ...any)  {}
func (e *NoOpEntry) Warnf(format string, args ...any)  {}
func (e *NoOpEntry) Errorf(format string, args ...any) {}

func (e *NoOpEntry) Debugln(args ...any) {}
func (e *NoOpEntry) Infoln(args ...any)  {}
func (e *NoOpEntry) Warnln(args ...any)  {}
func (e *NoOpEntry) Errorln(args ...any) {}
