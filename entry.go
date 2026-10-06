package loggerrific

type Entry interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)

	Debugln(args ...any)
	Infoln(args ...any)
	Warnln(args ...any)
	Errorln(args ...any)
}
