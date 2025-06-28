package tlogger

import "testing"

type TEntry struct {
	t *testing.T
}

func (e *TEntry) Debugf(format string, args ...interface{}) {
	e.t.Logf("debug: "+format, args...)
}

func (e *TEntry) Infof(format string, args ...interface{}) {
	e.t.Logf("info: "+format, args...)
}

func (e *TEntry) Warnf(format string, args ...interface{}) {
	e.t.Logf("warn: "+format, args...)
}

func (e *TEntry) Errorf(format string, args ...interface{}) {
	e.t.Logf("error: "+format, args...)
}

func (e *TEntry) Debugln(args ...interface{}) {
	e.t.Log(append([]interface{}{"debug"}, args...)...)
}

func (e *TEntry) Infoln(args ...interface{}) {
	e.t.Log(append([]interface{}{"info"}, args...)...)
}

func (e *TEntry) Warnln(args ...interface{}) {
	e.t.Log(append([]interface{}{"warn"}, args...)...)
}

func (e *TEntry) Errorln(args ...interface{}) {
	e.t.Log(append([]interface{}{"error"}, args...)...)
}
