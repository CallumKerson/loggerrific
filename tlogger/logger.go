package tlogger

import (
	"testing"

	"github.com/CallumKerson/loggerrific"
)

type TLogger struct {
	T *testing.T
}

func NewTLogger(t *testing.T) *TLogger {
	return &TLogger{T: t}
}

func (l *TLogger) WithField(_ string, _ interface{}) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) WithFields(_ map[string]interface{}) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) WithError(_ error) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) SetLevelDebug() {}

func (l *TLogger) SetLevelInfo() {}

func (l *TLogger) SetLevelWarn() {}

func (l *TLogger) SetLevelError() {}

func (l *TLogger) Debugf(format string, args ...interface{}) {
	l.T.Logf("debug: "+format, args...)
}

func (l *TLogger) Infof(format string, args ...interface{}) {
	l.T.Logf("info: "+format, args...)
}

func (l *TLogger) Warnf(format string, args ...interface{}) {
	l.T.Logf("warn: "+format, args...)
}

func (l *TLogger) Errorf(format string, args ...interface{}) {
	l.T.Logf("error: "+format, args...)
}

func (l *TLogger) Debugln(args ...interface{}) {
	l.T.Log(append([]interface{}{"debug"}, args...)...)
}

func (l *TLogger) Infoln(args ...interface{}) {
	l.T.Log(append([]interface{}{"info"}, args...)...)
}

func (l *TLogger) Warnln(args ...interface{}) {
	l.T.Log(append([]interface{}{"warn"}, args...)...)
}

func (l *TLogger) Errorln(args ...interface{}) {
	l.T.Log(append([]interface{}{"error"}, args...)...)
}

func (l *TLogger) IsDebugEnabled() bool {
	return true
}

func (l *TLogger) IsInfoEnabled() bool {
	return true
}

func (l *TLogger) IsWarnEnabled() bool {
	return true
}

func (l *TLogger) IsErrorEnabled() bool {
	return true
}
