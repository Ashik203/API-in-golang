package logger

import "log/slog"

type LogKey string

var Error = slog.Error

const (
	ExtraKey LogKey = "extra"
)

func Extra(value any) slog.Attr {
	return slog.String(string(ExtraKey), ConvertToJson(value))
}
