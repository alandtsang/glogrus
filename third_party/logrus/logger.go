package logrus

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	// The logs are `io.Copy`'d to this in a mutex. It's common to set this to a
	// file, or leave it default which is `os.Stderr`. You can also set this to
	// something more adventorous, such as logging to Kafka.
	Out io.Writer
	// Hooks for the logger instance. These allow firing events based on logging
	// levels and log entries. For example, to send errors to an error tracking
	// service, log to StatsD or dump the core on fatal errors.
	Hooks LevelHooks
	// All log entries pass through the formatter before logged to Out. The
	// included formatters are `TextFormatter` and `JSONFormatter` for which
	// TextFormatter is the default. In development (when a TTY is attached) it
	// logs with colors, but to a file it wouldn't. You can easily implement your
	// own that implements the `Formatter` interface, see the `README` or included
	// formatters for examples.
	Formatter Formatter
	// The logging level the logger should log at. This is typically (and defaults
	// to) `logrus.Info`, which allows Info(), Warn(), Error() and Fatal() to be
	// logged. `logrus.Debug` is useful in
	Level Level
	// Used to sync writing to the log.
	mu sync.Mutex
}

// Creates a new logger. Configuration should be set by changing `Formatter`,
// `Out` and `Hooks` directly on the default logger instance. You can also just
// instantiate your own:
//
//    var log = &Logger{
//      Out: os.Stderr,
//      Formatter: new(JSONFormatter),
//      Hooks: make(LevelHooks),
//      Level: logrus.DebugLevel,
//    }
//
// It's recommended to make this a global instance called `log`.
func New() *Logger {
	return &Logger{
		Out:       os.Stderr,
		Formatter: new(TextFormatter),
		Hooks:     make(LevelHooks),
		Level:     InfoLevel,
	}
}

// Adds a field to the log entry, note that you it doesn't log until you call
// Debug, Print, Info, Warn, Fatal or Panic. It only creates a log entry.
// If you want multiple fields, use `WithFields`.
func (logger *Logger) WithField(key string, value interface{}) *Entry {
	f, _ := Trace(3)
	f[key] = value
	return NewEntry(logger).WithFields(f)
}

// Adds a struct of fields to the log entry. All it does is call `WithField` for
// each `Field`.
func (logger *Logger) WithFields(fields Fields) *Entry {
	f, _ := Trace(3)
	for k, v := range f {
		fields[k] = v
	}
	return NewEntry(logger).WithFields(fields)
}

// Add an error as single field to the log entry.  All it does is call
// `WithError` for the given `error`.
func (logger *Logger) WithError(err error) *Entry {
	return NewEntry(logger).WithError(err)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	if logger.Level >= DebugLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Debugf(format, args...)
	}
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	if logger.Level >= InfoLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Infof(format, args...)
	}
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	fields, _ := Trace(3)
	NewEntry(logger).WithFields(fields).Printf(format, args...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warnf(format, args...)
	}
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warnf(format, args...)
	}
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	if logger.Level >= ErrorLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Errorf(format, args...)
	}
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	if logger.Level >= FatalLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Fatalf(format, args...)
	}
	os.Exit(1)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	if logger.Level >= PanicLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Panicf(format, args...)
	}
}

func (logger *Logger) Debug(args ...interface{}) {
	if logger.Level >= DebugLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Debug(args...)
	}
}

func (logger *Logger) Info(args ...interface{}) {
	if logger.Level >= InfoLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Info(args...)
	}
}

func (logger *Logger) Print(args ...interface{}) {
	fields, _ := Trace(3)
	NewEntry(logger).WithFields(fields).Info(args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warn(args...)
	}
}

func (logger *Logger) Warning(args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warn(args...)
	}
}

func (logger *Logger) Error(args ...interface{}) {
	if logger.Level >= ErrorLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Error(args...)
	}
}

func (logger *Logger) Fatal(args ...interface{}) {
	if logger.Level >= FatalLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Fatal(args...)
	}
	os.Exit(1)
}

func (logger *Logger) Panic(args ...interface{}) {
	if logger.Level >= PanicLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Panic(args...)
	}
}

func (logger *Logger) Debugln(args ...interface{}) {
	if logger.Level >= DebugLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Debugln(args...)
	}
}

func (logger *Logger) Infoln(args ...interface{}) {
	if logger.Level >= InfoLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Infoln(args...)
	}
}

func (logger *Logger) Println(args ...interface{}) {
	fields, _ := Trace(3)
	NewEntry(logger).WithFields(fields).Println(args...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warnln(args...)
	}
}

func (logger *Logger) Warningln(args ...interface{}) {
	if logger.Level >= WarnLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Warnln(args...)
	}
}

func (logger *Logger) Errorln(args ...interface{}) {
	if logger.Level >= ErrorLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Errorln(args...)
	}
}

func (logger *Logger) Fatalln(args ...interface{}) {
	if logger.Level >= FatalLevel {
		fields, _ := Trace(3)
		NewEntry(logger).WithFields(fields).Fatalln(args...)
	}
	os.Exit(1)
}

func (logger *Logger) Panicln(args ...interface{}) {
	if logger.Level >= PanicLevel {
		NewEntry(logger).Panicln(args...)
	}
}
