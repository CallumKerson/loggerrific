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

func (l *TLogger) WithField(_ string, _ any) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) WithFields(_ map[string]any) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) WithError(_ error) loggerrific.Entry {
	return &TEntry{t: l.T}
}

func (l *TLogger) SetLevelDebug() {}

func (l *TLogger) SetLevelInfo() {}

func (l *TLogger) SetLevelWarn() {}

func (l *TLogger) SetLevelError() {}

func (l *TLogger) Debugf(format string, args ...any) {
	l.T.Logf(levelDebug+": "+format, args...)
}

func (l *TLogger) Infof(format string, args ...any) {
	l.T.Logf(levelInfo+": "+format, args...)
}

func (l *TLogger) Warnf(format string, args ...any) {
	l.T.Logf(levelWarn+": "+format, args...)
}

func (l *TLogger) Errorf(format string, args ...any) {
	l.T.Logf(levelError+": "+format, args...)
}

func (l *TLogger) Debugln(args ...any) {
	l.T.Log(append([]any{levelDebug}, args...)...)
}

func (l *TLogger) Infoln(args ...any) {
	l.T.Log(append([]any{levelInfo}, args...)...)
}

func (l *TLogger) Warnln(args ...any) {
	l.T.Log(append([]any{levelWarn}, args...)...)
}

func (l *TLogger) Errorln(args ...any) {
	l.T.Log(append([]any{levelError}, args...)...)
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
