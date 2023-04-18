// package noopt provides an no operation (noopt) implementation of the l
// loggerrific interface that can be used as a default implementation if no
// other implementation is provided.
package noopt

import (
	"github.com/CallumKerson/loggerrific"
)

type NoOptLogger struct {
}

func New() *NoOptLogger {
	noOptLogger := &NoOptLogger{}
	return noOptLogger
}

func (l *NoOptLogger) WithField(key string, value interface{}) loggerrific.Entry {
	return &NoOptEntry{}
}

func (l *NoOptLogger) WithFields(fields map[string]interface{}) loggerrific.Entry {
	return &NoOptEntry{}
}

func (l *NoOptLogger) WithError(err error) loggerrific.Entry {
	return &NoOptEntry{}
}

func (l *NoOptLogger) SetLevelDebug() {}

func (l *NoOptLogger) SetLevelInfo() {}

func (l *NoOptLogger) SetLevelWarn() {}

func (l *NoOptLogger) SetLevelError() {}

func (l *NoOptLogger) IsDebugEnabled() bool {
	return false
}

func (l *NoOptLogger) IsInfoEnabled() bool {
	return false
}

func (l *NoOptLogger) IsWarnEnabled() bool {
	return false
}

func (l *NoOptLogger) IsErrorEnabled() bool {
	return false
}

func (e *NoOptLogger) Debugf(format string, args ...interface{}) {}
func (e *NoOptLogger) Infof(format string, args ...interface{})  {}
func (e *NoOptLogger) Warnf(format string, args ...interface{})  {}
func (e *NoOptLogger) Errorf(format string, args ...interface{}) {}

func (e *NoOptLogger) Debugln(args ...interface{}) {}
func (e *NoOptLogger) Infoln(args ...interface{})  {}
func (e *NoOptLogger) Warnln(args ...interface{})  {}
func (e *NoOptLogger) Errorln(args ...interface{}) {}

type NoOptEntry struct {
}

func (e *NoOptEntry) Debugf(format string, args ...interface{}) {}
func (e *NoOptEntry) Infof(format string, args ...interface{})  {}
func (e *NoOptEntry) Warnf(format string, args ...interface{})  {}
func (e *NoOptEntry) Errorf(format string, args ...interface{}) {}

func (e *NoOptEntry) Debugln(args ...interface{}) {}
func (e *NoOptEntry) Infoln(args ...interface{})  {}
func (e *NoOptEntry) Warnln(args ...interface{})  {}
func (e *NoOptEntry) Errorln(args ...interface{}) {}
