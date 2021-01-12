package apex

import (
	"context"
	"fmt"

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
			l.setLevel(lvl)
		}
	}

	return nil
}

func (l *apex) Options() logger.Options {
	// FIXME: How to return full opts?
	return l.opts.Options
}

func (l *apex) setLevel(level logger.Level) {
	apexLog.SetLevel(convertToApexLevel(level))
}

func (l *apex) Debug(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.DebugLevel, "%s", args)
}

func (l *apex) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.DebugLevel, format, args)
}

func (l *apex) Error(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.ErrorLevel, "%s", args)
}

func (l *apex) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.ErrorLevel, format, args)
}

func (l *apex) Info(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.InfoLevel, "%s", args)
}

func (l *apex) Infof(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.InfoLevel, format, args)
}

func (l *apex) Fatal(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.FatalLevel, "%s", args)
}

func (l *apex) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.FatalLevel, format, args)
}
func (l *apex) Trace(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.TraceLevel, "%s", args)
}

func (l *apex) Tracef(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.TraceLevel, format, args)
}
func (l *apex) Warn(ctx context.Context, args ...interface{}) {
	l.Logf(ctx, logger.WarnLevel, "%s", args)
}

func (l *apex) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.WarnLevel, format, args)
}

func (l *apex) V(level logger.Level) bool {
	return l.opts.Level >= level
}

// Log insets a log entry.  Arguments are handled in the manner of fmt.Printf
func (l *apex) Log(ctx context.Context, level logger.Level, v ...interface{}) {
	format := ""
	for i := 0; i < len(v); i++ {
		format += " %v"
	}
	apexlevel := convertToApexLevel(level)
	switch apexlevel {
	case apexLog.FatalLevel:
		l.Interface.Fatal(fmt.Sprintf(format, v...))
	case apexLog.ErrorLevel:
		l.Interface.Error(fmt.Sprintf(format, v...))
	case apexLog.WarnLevel:
		l.Interface.Warn(fmt.Sprintf(format, v...))
	case apexLog.DebugLevel:
		l.Interface.Debug(fmt.Sprintf(format, v...))
	default:
		l.Interface.Info(fmt.Sprintf(format, v...))
	}
}

// Logf insets a log entry.  Arguments are handled in the manner of fmt.Printf
func (l *apex) Logf(ctx context.Context, level logger.Level, format string, v ...interface{}) {
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
