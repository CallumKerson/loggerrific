package tlogger_test

import (
	"fmt"
	"testing"

	"github.com/CallumKerson/loggerrific"
	"github.com/CallumKerson/loggerrific/tlogger"
)

var errTest = fmt.Errorf("test error")

func TestTLogger(t *testing.T) {
	logger := tlogger.NewTLogger(t)

	// Test interface compliance
	var _ loggerrific.Logger = logger

	if logger == nil {
		t.Fatal("NewTLogger() should return a non-nil logger")
	}
}

func TestTLoggerMethods(t *testing.T) {
	// Use the actual testing.T instance
	logger := tlogger.NewTLogger(t)

	// Test WithField
	entry := logger.WithField("key", "value")
	if entry == nil {
		t.Error("WithField should return a non-nil Entry")
	}

	// Test WithFields
	fields := map[string]interface{}{
		"field1": "value1",
		"field2": 42,
		"field3": true,
	}
	entry = logger.WithFields(fields)
	if entry == nil {
		t.Error("WithFields should return a non-nil Entry")
	}

	// Test WithError
	entry = logger.WithError(errTest)
	if entry == nil {
		t.Error("WithError should return a non-nil Entry")
	}

	// Test level setting methods (should not panic)
	logger.SetLevelDebug()
	logger.SetLevelInfo()
	logger.SetLevelWarn()
	logger.SetLevelError()
}

func TestTLoggerFormatted(t *testing.T) {
	logger := tlogger.NewTLogger(t)

	// Test formatted logging methods don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TLogger formatted methods should not panic: %v", r)
		}
	}()

	logger.Debugf("debug message: %s", "test")
	logger.Infof("info message: %s", "test")
	logger.Warnf("warn message: %s", "test")
	logger.Errorf("error message: %s", "test")
}

func TestTLoggerLine(t *testing.T) {
	logger := tlogger.NewTLogger(t)

	// Test line logging methods don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TLogger line methods should not panic: %v", r)
		}
	}()

	logger.Debugln("debug", "message")
	logger.Infoln("info", "message")
	logger.Warnln("warn", "message")
	logger.Errorln("error", "message")
}

func TestTLoggerLevelChecking(t *testing.T) {
	logger := tlogger.NewTLogger(t)

	// TLogger should always return true for level checking
	if !logger.IsDebugEnabled() {
		t.Error("TLogger should return true for IsDebugEnabled()")
	}
	if !logger.IsInfoEnabled() {
		t.Error("TLogger should return true for IsInfoEnabled()")
	}
	if !logger.IsWarnEnabled() {
		t.Error("TLogger should return true for IsWarnEnabled()")
	}
	if !logger.IsErrorEnabled() {
		t.Error("TLogger should return true for IsErrorEnabled()")
	}
}

func TestTEntry(t *testing.T) {
	logger := tlogger.NewTLogger(t)
	entry := logger.WithField("test", "value")

	// Test interface compliance
	var _ = entry

	// Test formatted logging methods don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TEntry formatted methods should not panic: %v", r)
		}
	}()

	entry.Debugf("debug message: %s", "test")
	entry.Infof("info message: %s", "test")
	entry.Warnf("warn message: %s", "test")
	entry.Errorf("error message: %s", "test")
}

func TestTEntryLine(t *testing.T) {
	logger := tlogger.NewTLogger(t)
	entry := logger.WithField("test", "value")

	// Test line logging methods don't panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TEntry line methods should not panic: %v", r)
		}
	}()

	entry.Debugln("debug", "message")
	entry.Infoln("info", "message")
	entry.Warnln("warn", "message")
	entry.Errorln("error", "message")
}
