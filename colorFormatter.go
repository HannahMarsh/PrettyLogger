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
		return ColorCyan
	case logrus.InfoLevel:
		return ColorGreen
	case logrus.WarnLevel:
		return ColorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return ColorRed
	default:
		return ColorGrey
	}
}

func ColorRed(s string) string {
	return fmt.Sprintf("\033[31m%s\033[0m", s)
}

func ColorBrightRed(s string) string {
	return fmt.Sprintf("\033[91m%s\033[0m", s)
}

func ColorYellow(s string) string {
	return fmt.Sprintf("\033[33m%s\033[0m", s)
}

func ColorBrightYellow(s string) string {
	return fmt.Sprintf("\033[93m%s\033[0m", s)
}

func ColorGreen(s string) string {
	return fmt.Sprintf("\033[32m%s\033[0m", s)
}

func ColorBrightGreen(s string) string {
	return fmt.Sprintf("\033[92m%s\033[0m", s)
}

func ColorCyan(s string) string {
	return fmt.Sprintf("\033[36m%s\033[0m", s)
}

func ColorBrightCyan(s string) string {
	return fmt.Sprintf("\033[96m%s\033[0m", s)
}

func ColorBlue(s string) string {
	return fmt.Sprintf("\033[34m%s\033[0m", s)
}

func ColorBrightBlue(s string) string {
	return fmt.Sprintf("\033[94m%s\033[0m", s)
}

func ColorPurple(s string) string {
	return fmt.Sprintf("\033[35m%s\033[0m", s) // Pink/magenta color
}

func ColorPink(s string) string {
	return fmt.Sprintf("\033[95m%s\033[0m", s)
}

func ColorBrightWhite(s string) string {
	return fmt.Sprintf("\033[97m%s\033[0m", s)
}

func ColorGrey(s string) string {
	return fmt.Sprintf("\033[37m%s\033[0m", s)
}

func ColorBlack(s string) string {
	return fmt.Sprintf("\033[90m%s\033[0m", s)
}

func Bold(s string) string {
	return fmt.Sprintf("\033[1m%s\033[0m", s) // Bold text
}

func Italic(s string) string {
	return fmt.Sprintf("\033[3m%s\033[0m", s) // Italic text
}

func Underline(s string) string {
	return fmt.Sprintf("\033[4m%s\033[0m", s) // Underlined text
}
