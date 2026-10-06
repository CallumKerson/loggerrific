package tlogger

import "testing"

type TEntry struct {
	t *testing.T
}

func (e *TEntry) Debugf(format string, args ...interface{}) {
	e.t.Logf(levelDebug+": "+format, args...)
}

func (e *TEntry) Infof(format string, args ...interface{}) {
	e.t.Logf(levelInfo+": "+format, args...)
}

func (e *TEntry) Warnf(format string, args ...interface{}) {
	e.t.Logf(levelWarn+": "+format, args...)
}

func (e *TEntry) Errorf(format string, args ...interface{}) {
	e.t.Logf(levelError+": "+format, args...)
}

func (e *TEntry) Debugln(args ...interface{}) {
	e.t.Log(append([]interface{}{levelDebug}, args...)...)
}

func (e *TEntry) Infoln(args ...interface{}) {
	e.t.Log(append([]interface{}{levelInfo}, args...)...)
}

func (e *TEntry) Warnln(args ...interface{}) {
	e.t.Log(append([]interface{}{levelWarn}, args...)...)
}

func (e *TEntry) Errorln(args ...interface{}) {
	e.t.Log(append([]interface{}{levelError}, args...)...)
}
