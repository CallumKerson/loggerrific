package loggerrific

type Logger interface {
	WithField(key string, value any) Entry
	WithFields(fields map[string]any) Entry
	WithError(err error) Entry

	SetLevelDebug()
	SetLevelInfo()
	SetLevelWarn()
	SetLevelError()

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)

	Debugln(args ...any)
	Infoln(args ...any)
	Warnln(args ...any)
	Errorln(args ...any)

	IsDebugEnabled() bool
	IsInfoEnabled() bool
	IsWarnEnabled() bool
	IsErrorEnabled() bool
}
