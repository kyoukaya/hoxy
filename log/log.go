package log

import (
	"bufio"
	"fmt"
	"github.com/kyoukaya/hoxy/utils"
	"io"
	stdLog "log"
	"os"
	"runtime"
	"time"

	"github.com/logrusorgru/aurora"
	"github.com/mattn/go-colorable"
)

var (
	fileLogger   *stdLog.Logger
	stdOutLogger *stdLog.Logger
	buffer       *bufio.Writer
	Verbose      bool
)

const (
	defaultFlags  = stdLog.Ltime | stdLog.Lshortfile
	defaultPrefix = "Hoxy "
	// FileOut       = 1 << iota
	// StdOut
)

// InitLogger sets up the default logger.
func InitLogger(stdOut, fileOut bool, filePath string) {
	Verbose = utils.BoolFlags("hoxy-verbose")

	// Support for colored stdout output on windows.
	var output io.Writer
	if runtime.GOOS == "windows" {
		output = colorable.NewColorableStdout()
	} else {
		output = os.Stdout
	}

	if stdOut {
		stdOutLogger = stdLog.New(output, defaultPrefix, defaultFlags)
	}

	if fileOut {
		now := time.Now()
		var dir string
		var fileName string

		if filePath == "" {
			dir = utils.PackageRoot + "/logs/proxy/"
			fileName = now.Format("2006-01-02_15.04.05") + ".log"
		}

		err := os.MkdirAll(dir, 0755)
		utils.Check(err)

		f, err := os.Create(dir + fileName)
		utils.Check(err)

		buffer = bufio.NewWriter(f)

		fileLogger = stdLog.New(buffer, defaultPrefix, defaultFlags)
	}
}

// Flush all buffers associated with the standard logger, if any.
func Flush() {
	if buffer != nil {
		buffer.Flush()
	}
}

// These functions write to the standard loggers.
// A lot of code is duplicated from the stdlib to provide the same interface.

func Output(calldepth int, color func(interface{}) aurora.Value, prefix, str string) {
	calldepth++
	if fileLogger != nil {
		fileLogger.Output(calldepth, prefix+str)
	}

	if stdOutLogger != nil {
		if color != nil {
			prefix = color(prefix).String()
		}
		// Don't print long strings in stdout, truncate them to 120 chars.
		if len(str) > 1000 {
			str = str[0:120] + "..."
		}
		stdOutLogger.Output(calldepth, prefix+str)
	}
}

// Verboseln calls output to print to the standard logger, only if program is launched with
// the verbose flag.
func Verboseln(v ...interface{}) {
	if Verbose {
		Output(2, aurora.Blue, "INFO ", fmt.Sprintln(v...))
	}
}

// Verbosef calls output to print to the standard logger, only if program is launched with
// the verbose flag.
func Verbosef(format string, v ...interface{}) {
	if Verbose {
		Output(2, aurora.Blue, "INFO ", fmt.Sprintf(format, v...))
	}
}

// Infof calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Infof.
func Infof(format string, v ...interface{}) {
	Output(2, aurora.Green, "INFO ", fmt.Sprintf(format, v...))
}

// Infoln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Infoln.
func Infoln(v ...interface{}) {
	Output(2, aurora.Green, "INFO ", fmt.Sprintln(v...))
}

// Warnf calls Output to print a warning to the standard logger.
// Arguments are handled in the manner of fmt.Printf.
func Warnf(format string, v ...interface{}) {
	Output(2, aurora.Red, "WARN ", fmt.Sprintf(format, v...))
}

// Warnln calls Output to print to the standard logger.
// Arguments are handled in the manner of fmt.Infoln.
func Warnln(v ...interface{}) {
	Output(2, aurora.Red, "WARN ", fmt.Sprintln(v...))
}

// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	Output(2, aurora.BrightRed, "FATAL ", fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	Output(2, aurora.BrightRed, "FATAL ", fmt.Sprintf(format, v...))
	os.Exit(1)
}
