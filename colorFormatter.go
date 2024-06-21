package PrettyLogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type ColoredFormatter struct {
	TimestampFormat string
}

func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := getColorByLevel(entry.Level)
	timestamp := entry.Time.Format(f.TimestampFormat)

	message := fmt.Sprintf("%s %-18s %s\n",
		timestamp,
		fmt.Sprintf("[%s]", levelColor(strings.ToUpper(entry.Level.String()))),
		entry.Message)

	if entry.Level == logrus.ErrorLevel {
		if stack, ok := entry.Data["stack"]; ok {
			message += fmt.Sprintf("%s\n", stack)
		}
	}

	return []byte(message), nil
}

func getColorByLevel(level logrus.Level) func(string) string {
	switch level {
	case logrus.DebugLevel:
		return colorCyan
	case logrus.InfoLevel:
		return colorGreen
	case logrus.WarnLevel:
		return colorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return colorRed
	default:
		return colorWhite
	}
}

func colorCyan(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s)
}

func colorGreen(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func colorYellow(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s)
}

func colorRed(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func colorWhite(s string) string {
	return fmt.Sprintf("\033[37m%s\033[0m", s)
}
