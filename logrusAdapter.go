package PrettyLogger

// refs:
// https://josephwoodward.co.uk/2022/11/slog-structured-logging-proposal
// https://thedevelopercafe.com/articles/logging-in-go-with-slog-a7bb489755c2

import (
	"context"
	"fmt"
	"log/slog"
	"sort"
	"strings"

	"github.com/sirupsen/logrus"
)

type LogrusHandler struct {
	logger *logrus.Logger
}

func NewLogrusHandler(logger *logrus.Logger) *LogrusHandler {
	return &LogrusHandler{
		logger: logger,
	}
}

func ConvertLogLevel(level string) logrus.Level {
	var l logrus.Level

	switch strings.ToLower(level) {
	case "Error":
		l = logrus.ErrorLevel
	case "warm":
		l = logrus.WarnLevel
	case "info":
		l = logrus.InfoLevel
	case "debug":
		l = logrus.DebugLevel
	default:
		l = logrus.InfoLevel
	}

	return l
}

//func (h *LogrusHandler) Enabled(_ slog.Level) bool {
//	// support all logging levels
//	return true
//}

func (h *LogrusHandler) Enabled(_ context.Context, _ slog.Level) bool {
	// support all logging levels
	return true
}

func (h *LogrusHandler) Handle(ctx context.Context, rec slog.Record) error {
	fields := make(map[string]interface{}, rec.NumAttrs())

	rec.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()
		return true // continue iteration
	})

	entry := h.logger.WithFields(fields)

	printMsg := ""

	if rec.Level != slog.LevelError {
		fnName := Italic(ColorPurple(GetFuncNameWithSkip(4)))
		loc := getLocation(4)
		if len(fields) > 0 {
			//str := ""
			// Step 1: Extract the keys from the map
			keys := make([]string, 0, len(fields))
			for key, _ := range fields {
				keys = append(keys, key)
			}
			// Step 2: Sort the keys
			sort.Strings(keys)

			vars := make([]string, len(keys))

			for i, k := range keys {
				paramStr := interfaceToString(fields[k])

				vars[i] = fmt.Sprintf("%s=%s", Italic(ColorCyan(k)), Bold(paramStr))
				//str = str + fmt.Sprintf("%s=%v, ", k, fields[k])
			}
			str := strings.Join(vars, ", ")
			printMsg = fmt.Sprintf("%s (%s) → %s | %s", loc, fnName, ColorBrightWhite(rec.Message), str)
		} else {
			printMsg = fmt.Sprintf("%s (%s) → %s", loc, fnName, ColorBrightWhite(rec.Message))
		}
	}

	switch rec.Level {
	case slog.LevelDebug:
		entry.Debug(printMsg)
	case slog.LevelInfo:
		entry.Info(printMsg)
	case slog.LevelWarn:
		entry.Warn(printMsg)
	case slog.LevelError:
		//stack := ""
		b := false

		for _, v := range fields {
			//if k == "err" {
			if e, ok := v.(error); ok {
				ee := wrapError(e, rec.Message)
				str := fmt.Sprintf("%+v", ee)
				ew := parseWrappedError(str)
				ew = strings.TrimSpace(ew)
				i := strings.Index(ew, "\n")
				fmt.Println()
				entry.Error(ColorRed(ew[:i]) + "\n" + ew[i+1:] + "\n")
				b = true
			}
			//}
		}
		if !b {
			entry.Error(rec.Message)
		}
	}

	return nil
}

func (h *LogrusHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// not implemented for brevity
	return h
}

func (h *LogrusHandler) WithGroup(name string) slog.Handler {
	// not implemented for brevity
	return h
}
