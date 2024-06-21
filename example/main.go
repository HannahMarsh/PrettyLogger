package main

import (
	"errors"
	"github.com/HannahMarsh/PrettyLogger"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"os"
	"time"
)

func main() {

	logrus.SetFormatter(&PrettyLogger.ColoredFormatter{TimestampFormat: time.RFC3339})
	logrus.SetOutput(os.Stdout)

	// integrate Logrus with the slog logger
	slog.New(PrettyLogger.NewLogrusHandler(logrus.StandardLogger()))

	err := f2()
	if err != nil {

		PrettyLogger.PrintError(err, "main message")
	}

}

func f1() error {
	return errors.New("f1 error")
}

func f2() error {
	err := f1()
	return PrettyLogger.WrapError(err, "f2 message")
}
