package tlogger

import "testing"

type TEntry struct {
	t *testing.T
}

func (e *TEntry) Debugf(format string, args ...any) {
	e.t.Logf(levelDebug+": "+format, args...)
}

func (e *TEntry) Infof(format string, args ...any) {
	e.t.Logf(levelInfo+": "+format, args...)
}

func (e *TEntry) Warnf(format string, args ...any) {
	e.t.Logf(levelWarn+": "+format, args...)
}

func (e *TEntry) Errorf(format string, args ...any) {
	e.t.Logf(levelError+": "+format, args...)
}

func (e *TEntry) Debugln(args ...any) {
	e.t.Log(append([]any{levelDebug}, args...)...)
}

func (e *TEntry) Infoln(args ...any) {
	e.t.Log(append([]any{levelInfo}, args...)...)
}

func (e *TEntry) Warnln(args ...any) {
	e.t.Log(append([]any{levelWarn}, args...)...)
}

func (e *TEntry) Errorln(args ...any) {
	e.t.Log(append([]any{levelError}, args...)...)
}
