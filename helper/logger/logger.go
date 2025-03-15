package logger

import (
	"bankapp2/helper/logger/prettylog"
	"log/slog"
	"os"
)

func NewLogger() *slog.Logger {
	opts := prettylog.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)
	log := slog.New(handler)
	slog.SetDefault(log)

	return log
}
