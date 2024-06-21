package PrettyLogger

import (
	"fmt"
	errs "github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/exp/slog"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

func InitDefault() {
	logrus.SetFormatter(&ColoredFormatter{TimestampFormat: time.RFC3339})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)

	// integrate Logrus with the slog logger
	slog.SetDefault(slog.New(NewLogrusHandler(logrus.StandardLogger())))

}

// logErrorWithStack logs an Error with a stack trace.
func getLocation(skip int) (file string) {
	if _, filename, line, ok := runtime.Caller(skip); !ok {
		fmt.Println("failed to get caller")
		return
	} else if workingDir, e1 := os.Getwd(); e1 != nil {
		fmt.Println(e1)
		return
	} else if relativePath, e2 := filepath.Rel(workingDir, filename); e2 != nil {
		fmt.Println(e2)
		return
	} else {
		return fmt.Sprintf("%s:%d", relativePath, line)
	}
}

func getRelativePath(absPath string) string {
	absPath = strings.TrimSpace(absPath)
	i := strings.LastIndex(absPath, ":")
	path := absPath[:i]
	line := absPath[i+1:]
	if workingDir, e1 := os.Getwd(); e1 != nil {
		return absPath
	} else if relativePath, e2 := filepath.Rel(workingDir, path); e2 != nil {
		return absPath
	} else {
		return relativePath + ":" + line
	}
}

func parseWrappedError(str string) string {
	stack := ""
	spl := strings.Split(str, "\n---\n")
	for _, s := range spl {
		lines := strings.Split(strings.TrimSpace(s), "\n")
		msg := lines[0]
		if strings.Contains(msg, "file=\"") {
			file := strings.Split(" "+msg, "file=\"")[1]
			if strings.Contains(file, "\" msg=\"") {
				message := strings.TrimSuffix(strings.Split(" "+file, "\" msg=\"")[1], "\"")
				file = strings.TrimSuffix(strings.Split(file, "\" msg=\"")[0], "\"")
				for _, line := range lines[1:] {
					if strings.Contains(line, file) {
						stack = stack + fmt.Sprintf("\t\t\t\t    - %s  %s\n", getRelativePath(line), message)
						break
					}
				}
			}
		} else {
			message := lines[0]
			file := lines[2]
			stack = stack + fmt.Sprintf("\t\t\t\t    - %s  %s\n", getRelativePath(file), message)
		}

	}
	return stack
}
func WrapError(err error, msg string) error {
	file := getLocation(2)
	return errs.Wrap(err, fmt.Sprintf("\n---\nfile=\"%s\" msg=\"%s\"", file, msg))
}

func wrapError(err error, msg string) error {
	file := getLocation(5)
	return errs.Wrap(err, fmt.Sprintf("\n---\nfile=\"%s\" msg=\"%s\"", file, msg))
}
