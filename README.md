# loggerrific

[![Go Report Card](https://goreportcard.com/badge/github.com/CallumKerrEdwards/loggerrific?style=flat-square)](https://goreportcard.com/report/github.com/CallumKerrEdwards/loggerrific)
[![Go Reference](https://pkg.go.dev/badge/github.com/CallumKerrEdwards/loggerrific.svg)](https://pkg.go.dev/github.com/CallumKerrEdwards/loggerrific)
[![Release](https://img.shields.io/github/release/CallumKerrEdwards/loggerrific.svg?style=flat-square)](https://github.com/CallumKerrEdwards/loggerrific/releases/latest)

`loggerrific` is a logging interface for abstracting the choice of logging framework
for an application written in Go.

## Example

An example implementation for [Logrus](https://github.com/sirupsen/logrus):
```
package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/CallumKerrEdwards/loggerrific"
)

type Logger struct {
	*logrus.Logger
}

func NewLogger() *Logger {
	logrusLogger := &Logger{
		logrus.New(),
	}
	logrusLogger.SetLevel(logrus.InfoLevel)
	return logrusLogger
}

func (l *Logger) WithField(key string, value interface{}) loggerrific.Entry {
	return l.Logger.WithField(key, value)
}

func (l *Logger) WithFields(fields map[string]interface{}) loggerrific.Entry {
	return l.Logger.WithFields(fields)
}

func (l *Logger) WithError(err error) loggerrific.Entry {
	return l.Logger.WithError(err)
}

func (l *Logger) SetLevelDebug() {
	l.SetLevel(logrus.DebugLevel)
}

func (l *Logger) SetLevelInfo() {
	l.SetLevel(logrus.InfoLevel)
}

func (l *Logger) SetLevelWarn() {
	l.SetLevel(logrus.WarnLevel)
}

func (l *Logger) SetLevelError() {
	l.SetLevel(logrus.ErrorLevel)
}

func (l *Logger) IsDebugEnabled() bool {
	return l.IsLevelEnabled(logrus.DebugLevel)
}

func (l *Logger) IsInfoEnabled() bool {
	return l.IsLevelEnabled(logrus.InfoLevel)
}

func (l *Logger) IsWarnEnabled() bool {
	return l.IsLevelEnabled(logrus.WarnLevel)
}

func (l *Logger) IsErrorEnabled() bool {
	return l.IsLevelEnabled(logrus.ErrorLevel)
}

```
