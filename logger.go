package loggerrific

type Logger interface {
	WithField(key string, value interface{}) Entry
	WithFields(fields map[string]interface{}) Entry
	WithError(err error) Entry

	SetLevelDebug()
	SetLevelInfo()
	SetLevelWarn()
	SetLevelError()

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})

	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
}
