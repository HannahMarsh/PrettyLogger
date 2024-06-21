package PrettyLogger

// refs:
// https://josephwoodward.co.uk/2022/11/slog-structured-logging-proposal
// https://thedevelopercafe.com/articles/logging-in-go-with-slog-a7bb489755c2

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"

	"github.com/go-errors/errors"
)

// ErrorsStackHook is a Logrus hook that adds stack traces to errors.
type ErrorsStackHook struct{}

func (hook *ErrorsStackHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (hook *ErrorsStackHook) Fire(entry *logrus.Entry) error {
	if err, ok := entry.Data[logrus.ErrorKey].(error); ok {
		stackErr := errors.Wrap(err, 2)
		entry.Data["stack"] = fmt.Sprintf("%+v", stackErr)
	}
	return nil
}

// logErrorWithStack logs an error with a stack trace.
func LogErrorWithStack(message string, err error) {

	// notice that we're using 1, so it will actually log the where
	// the error happened, 0 = this function, we don't want that.
	if pc, filename, line, ok := runtime.Caller(1); !ok {
		fmt.Println("failed to get caller")
		return
	} else if workingDir, e1 := os.Getwd(); e1 != nil {
		fmt.Println(e1)
		return
	} else if relativePath, e2 := filepath.Rel(workingDir, filename); e2 != nil {
		fmt.Println(e2)
		return
	} else {
		//s := debug.Stack()
		//fmt.Printf("%v\n", s)
		s := fmt.Sprintf("%+v", err)
		fmt.Printf(s)
		fmt.Printf("\n%s:%d\n %v\n", relativePath, line, err)

		stackErr := errors.Wrap(err, 1)

		logrus.WithField("stack", fmt.Sprintf("%+v", stackErr)).Error(fmt.Sprintf("[error] in %s %s:%d  %v", runtime.FuncForPC(pc).Name(), filename, line, err))
	}
}

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
	case "error":
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

func (h *LogrusHandler) Enabled(_ slog.Level) bool {
	// support all logging levels
	return true
}

func (h *LogrusHandler) Handle(rec slog.Record) error {
	fields := make(map[string]interface{}, rec.NumAttrs())

	rec.Attrs(func(a slog.Attr) {
		fields[a.Key] = a.Value.Any()
	})

	entry := h.logger.WithFields(fields)

	switch rec.Level {
	case slog.DebugLevel:
		entry.Debug(rec.Message)
	case slog.InfoLevel.Level():
		entry.Info(rec.Message)
	case slog.WarnLevel:
		entry.Warn(rec.Message)
	case slog.ErrorLevel:
		entry.Error(rec.Message)
	}

	// Handle error stack trace
	if err, ok := fields[logrus.ErrorKey].(error); ok {
		stackErr := errors.Wrap(err, 2)
		entry = entry.WithField("stack", stackErr.ErrorStack())
		entry.Error(rec.Message)
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
