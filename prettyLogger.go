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

func SetUpLogrusAndSlog(logLevel string) {
	logrus.SetFormatter(&ColoredFormatter{TimestampFormat: time.RFC3339})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(ConvertLogLevel(logLevel))

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
	if strings.Contains(absPath, ":") {

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
	} else {
		return ""
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
				gotit := false
				if len(lines) >= 2 {
					for _, line := range lines[1:] {
						if strings.Contains(line, file) {
							stack = stack + fmt.Sprintf("\t\t\t\t    → %s  %s\n", getRelativePath(line), message)
							gotit = true
							break
						}
					}
				}
				if !gotit {
					stack = stack + fmt.Sprintf("\t\t\t\t    → %s  %s\n", getRelativePath(file), message)
				}
			}
		} else {
			message := lines[0]
			if len(lines) < 3 {
				if len(lines) == 2 {
					file := lines[1]
					stack = stack + fmt.Sprintf("\t\t\t\t    → %s  %s\n", getRelativePath(file), message)
				} else if len(lines) == 1 {
					file := lines[0]
					stack = stack + fmt.Sprintf("\t\t\t\t    → %s  %s\n", getRelativePath(file), message)
				} else {
					stack = stack + fmt.Sprintf("\t\t\t\t    → %s\n", message)
				}
			} else {
				file := lines[2]
				stack = stack + fmt.Sprintf("\t\t\t\t    → %s  %s\n", getRelativePath(file), message)
			}
		}

	}
	return stack
}
func WrapError(err error, format string, a ...any) error {
	if err == nil {
		return withStackSkip(fmt.Errorf(format, a...), 2)
	}
	msg := fmt.Sprintf(format, a...)
	file := getLocation(2)
	return errs.Wrap(err, fmt.Sprintf("\n---\nfile=\"%s\" msg=\"%s\"", file, msg))
}

type stackTracer interface {
	StackTrace() errs.StackTrace
}

func LogNewError(format string, a ...any) {
	slog.Error("", withStackSkip(fmt.Errorf(format, a...), 2))
}

func NewError(format string, a ...any) error {
	err := fmt.Errorf(format, a...)
	return withStackSkip(err, 1)
}

func withStackSkip(err error, skip int) error {
	if err == nil {
		return nil
	}
	pcs := make([]uintptr, 32)
	n := runtime.Callers(skip+2, pcs)
	pcs = pcs[:n]
	return &withStack{
		error: err,
		stack: pcs,
	}
}

type withStack struct {
	error
	stack []uintptr
}

func (w *withStack) StackTrace() errs.StackTrace {
	frames := make([]errs.Frame, len(w.stack))
	for i := range frames {
		frames[i] = errs.Frame(w.stack[i])
	}
	return frames
}

func (w *withStack) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		if s.Flag('+') {
			fmt.Fprintf(s, "%+v", w.error)
			for _, pc := range w.stack {
				f := errs.Frame(pc)
				fmt.Fprintf(s, "\n%+v", f)
			}
			return
		}
		fallthrough
	case 's':
		fmt.Fprint(s, w.error.Error())
	case 'q':
		fmt.Fprintf(s, "%q", w.error.Error())
	}
}

func wrapError(err error, msg string) error {
	file := getLocation(5)
	return errs.Wrap(err, fmt.Sprintf("\n---\nfile=\"%s\" msg=\"%s\"", file, msg))
}
