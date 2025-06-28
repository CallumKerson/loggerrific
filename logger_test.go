package loggerrific_test

import (
	"fmt"
	"testing"

	"github.com/CallumKerson/loggerrific"
	"github.com/CallumKerson/loggerrific/noop"
	"github.com/CallumKerson/loggerrific/tlogger"
)

var errTest = fmt.Errorf("test error")

// TestLoggerInterface ensures all implementations satisfy the Logger interface
func TestLoggerInterface(t *testing.T) {
	tests := []struct {
		name   string
		logger loggerrific.Logger
	}{
		{"NoOpLogger", noop.New()},
		{"TLogger", tlogger.NewTLogger(t)},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Test interface compliance
			var _ = testCase.logger

			// Test WithField returns Entry
			entry := testCase.logger.WithField("key", "value")
			if entry == nil {
				t.Error("WithField should return a non-nil Entry")
			}

			// Test WithFields returns Entry
			fields := map[string]interface{}{
				"field1": "value1",
				"field2": 42,
			}
			entry = testCase.logger.WithFields(fields)
			if entry == nil {
				t.Error("WithFields should return a non-nil Entry")
			}

			// Test WithError returns Entry
			entry = testCase.logger.WithError(errTest)
			if entry == nil {
				t.Error("WithError should return a non-nil Entry")
			}

			// Test all logging methods don't panic
			testLoggingMethods(t, testCase.logger)

			// Test level setting methods don't panic
			testLevelMethods(t, testCase.logger)

			// Test level checking methods
			testLevelChecking(t, testCase.logger)
		})
	}
}

func testLoggingMethods(t *testing.T, logger loggerrific.Logger) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Logging methods should not panic: %v", r)
		}
	}()

	// Test formatted logging methods
	logger.Debugf("debug message: %s", "test")
	logger.Infof("info message: %s", "test")
	logger.Warnf("warn message: %s", "test")
	logger.Errorf("error message: %s", "test")

	// Test line logging methods
	logger.Debugln("debug", "message")
	logger.Infoln("info", "message")
	logger.Warnln("warn", "message")
	logger.Errorln("error", "message")
}

func testLevelMethods(t *testing.T, logger loggerrific.Logger) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Level setting methods should not panic: %v", r)
		}
	}()

	logger.SetLevelDebug()
	logger.SetLevelInfo()
	logger.SetLevelWarn()
	logger.SetLevelError()
}

func testLevelChecking(t *testing.T, logger loggerrific.Logger) {
	t.Helper()

	// These methods should return boolean values without panicking
	_ = logger.IsDebugEnabled()
	_ = logger.IsInfoEnabled()
	_ = logger.IsWarnEnabled()
	_ = logger.IsErrorEnabled()
}

// TestEntryInterface ensures all Entry implementations work correctly
func TestEntryInterface(t *testing.T) {
	tests := []struct {
		name   string
		logger loggerrific.Logger
	}{
		{"NoOpLogger", noop.New()},
		{"TLogger", tlogger.NewTLogger(t)},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			entry := testCase.logger.WithField("test", "value")

			// Test interface compliance
			var _ = entry

			// Test all entry logging methods don't panic
			testEntryLoggingMethods(t, entry)
		})
	}
}

func testEntryLoggingMethods(t *testing.T, entry loggerrific.Entry) {
	t.Helper()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Entry logging methods should not panic: %v", r)
		}
	}()

	// Test formatted logging methods
	entry.Debugf("debug message: %s", "test")
	entry.Infof("info message: %s", "test")
	entry.Warnf("warn message: %s", "test")
	entry.Errorf("error message: %s", "test")

	// Test line logging methods
	entry.Debugln("debug", "message")
	entry.Infoln("info", "message")
	entry.Warnln("warn", "message")
	entry.Errorln("error", "message")
}
