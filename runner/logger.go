package runner

import (
	"fmt"
	logPkg "log"

	colorable "github.com/mattn/go-colorable"
)

type logFunc func(string, ...interface{})

var logger = logPkg.New(colorable.NewColorableStderr(), "", 0)

func newLogFunc(prefix string) func(string, ...interface{}) {
	color, clear := "", ""
	if settings["colors"] == "1" {
		color = fmt.Sprintf("\033[%sm", logColor(prefix))
		clear = fmt.Sprintf("\033[%sm", colors["reset"])
	}
	prefix = fmt.Sprintf("%-11s", prefix)

	return func(format string, v ...interface{}) {
		format = fmt.Sprintf("%s %s |%s %s", color, prefix, clear, format)
		logger.Printf(format, v...)
	}
}

func fatal(err error) {
	logger.Fatal(err)
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	appLog(string(p))

	return len(p), nil
}
