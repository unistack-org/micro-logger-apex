package apex

import (
	"context"
	"fmt"

	apexLog "github.com/apex/log"
	"github.com/unistack-org/micro/v3/logger"
)

type apex struct {
	apexLog.Interface
	opts logger.Options
}

// Fields set fields to always be logged
func (l *apex) Fields(fields map[string]interface{}) logger.Logger {
	data := make(apexLog.Fields, len(fields))
	for k, v := range fields {
		data[k] = v
	}
	return newLogger(l.WithFields(data), l.opts)
}

// Init initializes options
func (l *apex) Init(opts ...logger.Option) error {
	for _, o := range opts {
		o(&l.opts)
	}

	if l.opts.Context != nil {
		if al, ok := l.opts.Context.Value(loggerKey{}).(apexLog.Interface); ok {
			l.Interface = al
			return nil
		}

		if h, ok := l.opts.Context.Value(handlerKey{}).(apexLog.Handler); ok {
			apexLog.SetHandler(h)
		}

		if lvl, ok := l.opts.Context.Value(levelKey{}).(logger.Level); ok {
			l.setLevel(lvl)
		}
	}

	return nil
}

func (l *apex) Options() logger.Options {
	return l.opts
}

func (l *apex) setLevel(level logger.Level) {
	apexLog.SetLevel(convertToApexLevel(level))
}

func (l *apex) Debug(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.DebugLevel, args...)
}

func (l *apex) Debugf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.DebugLevel, format, args)
}

func (l *apex) Error(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.ErrorLevel, args...)
}

func (l *apex) Errorf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.ErrorLevel, format, args)
}

func (l *apex) Info(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.InfoLevel, args...)
}

func (l *apex) Infof(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.InfoLevel, format, args)
}

func (l *apex) Fatal(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.FatalLevel, args...)
}

func (l *apex) Fatalf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.FatalLevel, format, args)
}

func (l *apex) Trace(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.TraceLevel, args...)
}

func (l *apex) Tracef(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.TraceLevel, format, args)
}

func (l *apex) Warn(ctx context.Context, args ...interface{}) {
	l.Log(ctx, logger.WarnLevel, args...)
}

func (l *apex) Warnf(ctx context.Context, format string, args ...interface{}) {
	l.Logf(ctx, logger.WarnLevel, format, args)
}

func (l *apex) V(level logger.Level) bool {
	return l.opts.Level >= level
}

// Log insets a log entry.  Arguments are handled in the manner of fmt.Printf
func (l *apex) Log(ctx context.Context, level logger.Level, v ...interface{}) {
	if !l.V(level) {
		return
	}

	apexlevel := convertToApexLevel(level)
	switch apexlevel {
	case apexLog.FatalLevel:
		l.Interface.Fatal(fmt.Sprint(v...))
	case apexLog.ErrorLevel:
		l.Interface.Error(fmt.Sprint(v...))
	case apexLog.WarnLevel:
		l.Interface.Warn(fmt.Sprint(v...))
	case apexLog.DebugLevel:
		l.Interface.Debug(fmt.Sprint(v...))
	default:
		l.Interface.Info(fmt.Sprint(v...))
	}
}

// Logf insets a log entry.  Arguments are handled in the manner of fmt.Printf
func (l *apex) Logf(ctx context.Context, level logger.Level, format string, v ...interface{}) {
	if !l.V(level) {
		return
	}

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

func newLogger(logInstance apexLog.Interface, opts logger.Options) logger.Logger {
	return &apex{
		Interface: logInstance,
		opts:      opts,
	}
}

// New returns a new ApexLogger instance
func NewLogger(opts ...logger.Option) logger.Logger {
	options := logger.NewOptions(opts...)
	l := newLogger(apexLog.Log, options)
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
