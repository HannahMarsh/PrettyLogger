package pkg

import (
	"github.com/HannahMarsh/PrettyLogger"
	"log/slog"
)

func f2() error {
	err1 := f1()
	slog.Info("f2", "err1", err1)
	return PrettyLogger.WrapError(err1, "f2 message")
}
