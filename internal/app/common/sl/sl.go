package sl

import (
	"log/slog"
)


func Error(msg string) slog.Attr {
	return slog.Attr{
		Key: "error",
		Value: slog.StringValue(msg),
	}
}