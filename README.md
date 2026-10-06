# loggerrific

[![Go Report Card](https://goreportcard.com/badge/github.com/CallumKerson/loggerrific?style=flat-square)](https://goreportcard.com/report/github.com/CallumKerson/loggerrific)
[![Go Reference](https://pkg.go.dev/badge/github.com/CallumKerson/loggerrific.svg)](https://pkg.go.dev/github.com/CallumKerson/loggerrific)
[![Release](https://img.shields.io/github/release/CallumKerson/loggerrific.svg?style=flat-square)](https://github.com/CallumKerson/loggerrific/releases/latest)

`loggerrific` is a logging interface for abstracting the choice of logging framework for Go applications.
It provides a unified API that allows you to switch between different logging implementations without changing your application code.

## Features

- **Framework Agnostic**: Works with any logging framework
- **Structured Logging**: Support for fields and key-value pairs
- **Level Control**: Dynamic log level management
- **Zero Dependencies**: Pure Go with no external dependencies
- **Testing Support**: Built-in testing logger for unit tests
- **Performance**: Includes no-op implementation for high-performance scenarios

## Quick Start

```go
package main

import (
    "github.com/CallumKerson/loggerrific"
    "github.com/CallumKerson/loggerrific/noop"
)

func main() {
    // Use the no-op logger for high performance
    logger := noop.New()

    // Or use your preferred logging framework implementation
    // logger := mylogger.New()

    // Log with structured data
    logger.WithField("user_id", 12345).
           WithField("action", "login").
           Infof("User logged in successfully")

    // Set log levels dynamically
    logger.SetLevelDebug()

    // Check if logging is enabled before expensive operations
    if logger.IsDebugEnabled() {
        logger.Debugf("Debug info: %+v", expensiveDebugData())
    }
}
```

## Built-in Implementations

### No-Op Logger

Perfect for production environments where logging overhead needs to be minimized:

```go
import "github.com/CallumKerson/loggerrific/noop"

logger := noop.New()
logger.Infof("This won't output anything") // Zero overhead
```

### Testing Logger

Integrates with Go's testing framework for unit tests:

```go
import (
    "testing"
    "github.com/CallumKerson/loggerrific/tlogger"
)

func TestMyFunction(t *testing.T) {
    logger := tlogger.NewTLogger(t)

    // Logs will appear in test output
    myFunction(logger)
}
```

## Interface

The core `Logger` interface provides:

```go
type Logger interface {
    // Structured logging
    WithField(key string, value any) Entry
    WithFields(fields map[string]any) Entry
    WithError(err error) Entry

    // Level management
    SetLevelDebug()
    SetLevelInfo()
    SetLevelWarn()
    SetLevelError()

    // Formatted logging
    Debugf(format string, args ...any)
    Infof(format string, args ...any)
    Warnf(format string, args ...any)
    Errorf(format string, args ...any)

    // Line logging
    Debugln(args ...any)
    Infoln(args ...any)
    Warnln(args ...any)
    Errorln(args ...any)

    // Level checking
    IsDebugEnabled() bool
    IsInfoEnabled() bool
    IsWarnEnabled() bool
    IsErrorEnabled() bool
}
```

## Creating Custom Implementations

An example implementation for [Logrus](https://github.com/sirupsen/logrus):

```go
package logrus

import (
	"github.com/sirupsen/logrus"

	"github.com/CallumKerson/loggerrific"
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

func (l *Logger) WithField(key string, value any) loggerrific.Entry {
	return l.Logger.WithField(key, value)
}

func (l *Logger) WithFields(fields map[string]any) loggerrific.Entry {
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

## Installation

```bash
go get github.com/CallumKerson/loggerrific
```

## Development

This project uses [mise](https://mise.jdx.dev/) for tools and tasks, and [hk](https://hk.jdx.dev/) for linting:

```bash
# Install the pinned tools
mise install

# Run tests
mise run test

# Fix, then check, everything
mise run fix-all
mise run check-all

# See all available tasks
mise tasks
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run `mise run fix-all`, `mise run check-all` and `mise run test`
5. Submit a pull request with a [conventional commit](https://www.conventionalcommits.org/) title

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
