package noop_test

import (
	"testing"

	"github.com/CallumKerson/loggerrific"
	"github.com/CallumKerson/loggerrific/noop"
)

func TestNoOpLogger(t *testing.T) {
	logger := noop.New()

	// Test interface compliance
	var _ loggerrific.Logger = logger

	if logger == nil {
		t.Fatal("New() should return a non-nil logger")
	}
}

func TestNoOpLoggerMethods(t *testing.T) {
	logger := noop.New()

	// All methods should be no-ops and not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NoOpLogger methods should not panic: %v", r)
		}
	}()

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
	entry = logger.WithError(nil)
	if entry == nil {
		t.Error("WithError should return a non-nil Entry")
	}

	// Test level setting methods
	logger.SetLevelDebug()
	logger.SetLevelInfo()
	logger.SetLevelWarn()
	logger.SetLevelError()

	// Test logging methods
	logger.Debugf("debug: %s", "test")
	logger.Infof("info: %s", "test")
	logger.Warnf("warn: %s", "test")
	logger.Errorf("error: %s", "test")

	logger.Debugln("debug", "test")
	logger.Infoln("info", "test")
	logger.Warnln("warn", "test")
	logger.Errorln("error", "test")
}

func TestNoOpLoggerLevelChecking(t *testing.T) {
	logger := noop.New()

	// NoOp logger should always return false for level checking
	if logger.IsDebugEnabled() {
		t.Error("NoOp logger should return false for IsDebugEnabled()")
	}
	if logger.IsInfoEnabled() {
		t.Error("NoOp logger should return false for IsInfoEnabled()")
	}
	if logger.IsWarnEnabled() {
		t.Error("NoOp logger should return false for IsWarnEnabled()")
	}
	if logger.IsErrorEnabled() {
		t.Error("NoOp logger should return false for IsErrorEnabled()")
	}
}

func TestNoOpEntry(t *testing.T) {
	logger := noop.New()
	entry := logger.WithField("test", "value")

	// Test interface compliance
	var _ = entry

	// All entry methods should be no-ops and not panic
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("NoOpEntry methods should not panic: %v", r)
		}
	}()

	// Test formatted logging methods
	entry.Debugf("debug: %s", "test")
	entry.Infof("info: %s", "test")
	entry.Warnf("warn: %s", "test")
	entry.Errorf("error: %s", "test")

	// Test line logging methods
	entry.Debugln("debug", "test")
	entry.Infoln("info", "test")
	entry.Warnln("warn", "test")
	entry.Errorln("error", "test")
}

func BenchmarkNoOpLogger(b *testing.B) {
	logger := noop.New()

	b.Run("Debugf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			logger.Debugf("debug message: %s", "test")
		}
	})

	b.Run("WithField", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			entry := logger.WithField("key", "value")
			entry.Infof("message: %s", "test")
		}
	})

	b.Run("WithFields", func(b *testing.B) {
		fields := map[string]interface{}{
			"field1": "value1",
			"field2": 42,
		}
		for i := 0; i < b.N; i++ {
			entry := logger.WithFields(fields)
			entry.Infof("message: %s", "test")
		}
	})
}
