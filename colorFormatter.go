package PrettyLogger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
)

type ColoredFormatter struct {
	TimestampFormat string
	LogLevel        string
}

func (f *ColoredFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	levelColor := f.getColorByLevel(entry.Level)
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

func (f *ColoredFormatter) getColorByLevel(level logrus.Level) func(string) string {
	switch level {
	case logrus.DebugLevel:
		return f.ColorCyan
	case logrus.InfoLevel:
		return f.ColorGreen
	case logrus.WarnLevel:
		return f.ColorYellow
	case logrus.ErrorLevel, logrus.FatalLevel, logrus.PanicLevel:
		return f.ColorRed
	default:
		return f.ColorGrey
	}
}

func (f *ColoredFormatter) ColorRed(s string) string {
	if f.LogLevel == "debug" {
		return fmt.Sprintf("\033[31m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightRed(s string) string {
	if f.LogLevel == "debug" {
		return fmt.Sprintf("\033[91m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorYellow(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[33m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightYellow(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[93m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorGreen(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[32m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightGreen(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[92m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorCyan(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[36m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightCyan(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[96m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBlue(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[34m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightBlue(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[94m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorPurple(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[35m%s\033[0m", s) // Pink/magenta color
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorPink(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[95m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBrightWhite(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[97m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorGrey(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[37m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) ColorBlack(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[90m%s\033[0m", s)
	} else {
		return s
	}
}

func (f *ColoredFormatter) Bold(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[1m%s\033[0m", s) // Bold text
	} else {
		return s
	}
}

func (f *ColoredFormatter) Italic(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[3m%s\033[0m", s) // Italic text
	} else {
		return s
	}
}

func (f *ColoredFormatter) Underline(s string) string {
	if f.LogLevel == "debug" {

		return fmt.Sprintf("\033[4m%s\033[0m", s) // Underlined text
	} else {
		return s
	}
}
