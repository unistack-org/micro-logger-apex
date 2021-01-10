package apex

import (
	"context"

	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"github.com/unistack-org/micro/v3/logger"
)

type handlerKey struct{}
type levelKey struct{}

// Options is used when applying custom options
type Options struct {
	logger.Options
}

type loggerKey struct{}

// WithLogger sets the logger
func WithLogger(l apexLog.Interface) logger.Option {
	return setOption(loggerKey{}, l)
}

// WithLevel allows to set the level for Log Output
func WithLevel(level logger.Level) logger.Option {
	return setOption(levelKey{}, level)
}

// WithHandler allows to set a customHandler for Log Output
func WithHandler(handler apexLog.Handler) logger.Option {
	return setOption(handlerKey{}, handler)
}

// WithTextHandler sets the Text Handler for Log Output
func WithTextHandler() logger.Option {
	return WithHandler(text.Default)
}

// WithJSONHandler sets the JSON Handler for Log Output
func WithJSONHandler() logger.Option {
	return WithHandler(json.Default)
}

// WithCLIHandler sets the CLI Handler for Log Output
func WithCLIHandler() logger.Option {
	return WithHandler(cli.Default)
}

func setOption(k, v interface{}) logger.Option {
	return func(o *logger.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, k, v)
	}
}
