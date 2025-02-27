package logger

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/jwalton/gchalk"
)

var opts = colorfulHandlerOptions{
	SlogOpts: slog.HandlerOptions{
		Level: slog.LevelDebug,
	},
}
var handler = newColorfulHandler(os.Stdout, opts)
var logger = slog.New(handler)

type colorfulHandlerOptions struct {
	SlogOpts slog.HandlerOptions
}

type colorfulHandler struct {
	slog.Handler
	l *log.Logger
}

func (h *colorfulHandler) Handle(ctx context.Context, r slog.Record) error {
	level := r.Level.String()
	level = strings.Join([]string{"[", level, "]"}, "")
	switch r.Level {
	case slog.LevelDebug:
		level = gchalk.Magenta(level)
	case slog.LevelInfo:
		level = gchalk.Cyan(level)
	case slog.LevelWarn:
		level = gchalk.Yellow(level)
	case slog.LevelError:
		level = gchalk.Red(level)
	}

	fields := make(map[string]any, r.NumAttrs())
	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true
	})

	b, err := json.MarshalIndent(fields, "", "  ")
	if err != nil {
		return err
	}

	timeStr := gchalk.Grey(r.Time.Format("2006-01-02 15:04:05.999"))

	jsonStr := string(b)
	if len(jsonStr) == 2 {
		jsonStr = ""
	}

	h.l.Printf("%s %s %s %s", timeStr, level, r.Message, jsonStr)

	return nil
}

func newColorfulHandler(out io.Writer, opts colorfulHandlerOptions) *colorfulHandler {
	h := &colorfulHandler{Handler: slog.NewJSONHandler(out, &opts.SlogOpts),
		l: log.New(out, "", 0),
	}
	return h
}

func Debug(msg string, args ...any) {
	logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	logger.Error(msg, args...)
}
