package main

import (
	"github.com/HannahMarsh/PrettyLogger"
	"github.com/HannahMarsh/PrettyLogger/example/pkg"
	"golang.org/x/exp/slog"
)

func main() {

	PrettyLogger.SetUpLogrusAndSlog()

	err2 := pkg.F3()

	slog.Error("main error", err2)
	slog.Info("main info", "info", "info")
	slog.Warn("main warn", "warn", "warn")
	slog.Debug("main debug", "debug", "debug")

}
