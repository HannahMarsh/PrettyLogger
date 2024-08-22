package pkg

import (
	"github.com/HannahMarsh/PrettyLogger"
	"log/slog"
)

func F3() error {
	err3 := f2()
	slog.Info("f3", "err3", err3)
	return PrettyLogger.WrapError(err3, "f3 message")
}
