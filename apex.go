package apex // import "go.unistack.org/micro-logger-apex/v3"

import (
	"context"
	j "encoding/json"
	"fmt"

	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"go.unistack.org/micro/v3/logger"
)

type apex struct {
	apexLog.Interface
	opts logger.Options
}

// Fields set fields to always be logged
func (l *apex) Fields(fields ...interface{}) logger.Logger {
	data := make(apexLog.Fields, len(fields))
	for i := 0; i < len(fields); i += 2 {
		data[fields[i].(string)] = fields[i+1]
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

	switch li := l.Interface.(type) {
	case *apexLog.Logger:
		switch h := li.Handler.(type) {
		case *text.Handler:
			h.Writer = l.opts.Out
		case *json.Handler:
			h.Encoder = j.NewEncoder(l.opts.Out)
		case *cli.Handler:
			h.Writer = l.opts.Out
		default:
			li.Handler = json.New(l.opts.Out)
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

func (l *apex) Clone(opts ...logger.Option) logger.Logger {
	nl := &apex{
		Interface: l.Interface,
		opts:      l.opts,
	}

	for _, o := range opts {
		o(&nl.opts)
	}

	if nl.opts.Context != nil {
		if al, ok := nl.opts.Context.Value(loggerKey{}).(apexLog.Interface); ok {
			nl.Interface = al
			return nil
		}

		if h, ok := nl.opts.Context.Value(handlerKey{}).(apexLog.Handler); ok {
			apexLog.SetHandler(h)
		}

		if lvl, ok := nl.opts.Context.Value(levelKey{}).(logger.Level); ok {
			nl.setLevel(lvl)
		}
	}

	return nl
}

func (l *apex) Level(lvl logger.Level) {
	l.setLevel(lvl)
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
