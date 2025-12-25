package log

import (
	"context"
	"io"
	"log"
	"log/slog"

	"github.com/radovskyb/not-a-rule-engine/services"
)

type Logger interface {
	Log(ctx context.Context, value string, level slog.Level)
}

type logger struct {
	l *slog.Logger
}

func New(w io.Writer) *logger {
	h := slog.NewTextHandler(w, nil)
	return &logger{
		l: slog.New(h),
	}
}

func (l *logger) Funcs() (int, map[string]services.FncParam) {
	log.Println("retrieving allowed funcs + params for log service")
	return services.LogServiceID, nil
}

func (l *logger) Call(ctx context.Context, params any) (any, error) {
	log.Println("calling log service")

	return nil, nil
}

func (l *logger) Log(ctx context.Context, value string, level slog.Level) {
	l.l.Log(ctx, level, value)
}
