package apex

import (
	apexLog "github.com/apex/log"
	"github.com/unistack-org/micro/v3/logger"
)

type apex struct {
	apexLog.Interface
	opts Options
}

// Fields set fields to always be logged
func (l *apex) Fields(fields map[string]interface{}) logger.Logger {
	data := make(apexLog.Fields, len(fields))
	for k, v := range fields {
		data[k] = v
	}
	return newLogger(l.WithFields(data))
}

// Init initializes options
func (l *apex) Init(opts ...logger.Option) error {
	options := &Options{}
	for _, o := range opts {
		o(&options.Options)
	}

	if options.Context != nil {
		if al, ok := options.Context.Value(loggerKey{}).(apexLog.Interface); ok {
			l.Interface = al
			return nil
		}

		if h, ok := options.Context.Value(handlerKey{}).(apexLog.Handler); ok {
			apexLog.SetHandler(h)
		}

		if lvl, ok := options.Context.Value(levelKey{}).(logger.Level); ok {
			l.setLogLevel(lvl)
		}
	}

	return nil
}

func (l *apex) Options() logger.Options {
	// FIXME: How to return full opts?
	return l.opts.Options
}

func (l *apex) setLogLevel(level logger.Level) {
	apexLog.SetLevel(convertToApexLevel(level))
}

func (l *apex) Debug(args ...interface{}) {
	l.logf(logger.DebugLevel, "%s", args)
}

func (l *apex) Debugf(format string, args ...interface{}) {
	l.logf(logger.DebugLevel, format, args)
}

func (l *apex) Error(args ...interface{}) {
	l.logf(logger.ErrorLevel, "%s", args)
}

func (l *apex) Errorf(format string, args ...interface{}) {
	l.logf(logger.ErrorLevel, format, args)
}

func (l *apex) Info(args ...interface{}) {
	l.logf(logger.InfoLevel, "%s", args)
}

func (l *apex) Infof(format string, args ...interface{}) {
	l.logf(logger.InfoLevel, format, args)
}

func (l *apex) Fatal(args ...interface{}) {
	l.logf(logger.FatalLevel, "%s", args)
}

func (l *apex) Fatalf(format string, args ...interface{}) {
	l.logf(logger.FatalLevel, format, args)
}
func (l *apex) Trace(args ...interface{}) {
	l.logf(logger.TraceLevel, "%s", args)
}

func (l *apex) Tracef(format string, args ...interface{}) {
	l.logf(logger.TraceLevel, format, args)
}
func (l *apex) Warn(args ...interface{}) {
	l.logf(logger.WarnLevel, "%s", args)
}

func (l *apex) Warnf(format string, args ...interface{}) {
	l.logf(logger.WarnLevel, format, args)
}

func (l *apex) V(level logger.Level) bool {
	return l.opts.Level >= level
}

// Logf insets a log entry.  Arguments are handled in the manner of
// fmt.Printf.
func (l *apex) logf(level logger.Level, format string, v ...interface{}) {
	apexlevel := convertToApexLevel(level)
	switch apexlevel {
	case apexLog.FatalLevel:
		l.Interface.Fatalf(format, v...)
	case apexLog.ErrorLevel:
		l.Interface.Errorf(format, v...)
	case apexLog.WarnLevel:
		l.Interface.Warnf(format, v...)
	case apexLog.DebugLevel:
		l.Interface.Debugf(format, v...)
	default:
		l.Interface.Infof(format, v...)
	}
}

// String returns the name of logger
func (l *apex) String() string {
	return "apex"
}

func newLogger(logInstance apexLog.Interface) logger.Logger {
	return &apex{
		Interface: logInstance,
		opts: Options{
			logger.Options{
				Level: logger.InfoLevel,
			},
		},
	}
}

// New returns a new ApexLogger instance
func NewLogger(opts ...logger.Option) logger.Logger {
	l := newLogger(apexLog.Log)
	_ = l.Init(opts...)
	return l
}

func convertToApexLevel(level logger.Level) apexLog.Level {
	switch level {
	case logger.DebugLevel:
		return apexLog.DebugLevel
	case logger.InfoLevel:
		return apexLog.InfoLevel
	case logger.WarnLevel:
		return apexLog.WarnLevel
	case logger.ErrorLevel:
		return apexLog.ErrorLevel
	case logger.FatalLevel:
		return apexLog.FatalLevel
	case logger.TraceLevel:
		return apexLog.DebugLevel
	default:
		return apexLog.InfoLevel
	}
}
