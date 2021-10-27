package apex

import (
	apexLog "github.com/apex/log"
	"github.com/apex/log/handlers/cli"
	"github.com/apex/log/handlers/json"
	"github.com/apex/log/handlers/text"
	"go.unistack.org/micro/v3/logger"
)

type (
	handlerKey struct{}
	levelKey   struct{}
)

type loggerKey struct{}

// WithLogger sets the logger
func WithLogger(l apexLog.Interface) logger.Option {
	return logger.SetOption(loggerKey{}, l)
}

// WithLevel allows to set the level for Log Output
func WithLevel(level logger.Level) logger.Option {
	return logger.SetOption(levelKey{}, level)
}

// WithHandler allows to set a customHandler for Log Output
func WithHandler(handler apexLog.Handler) logger.Option {
	return logger.SetOption(handlerKey{}, handler)
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
