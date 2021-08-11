// Gulu - Golang common utilities for everyone.
// Copyright (c) 2019-present, b3log.org
//
// Gulu is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package gulu

import (
	"fmt"
	"io"
	stdlog "log"
	"os"
	"strings"
)

type guluLog struct{}

// Log utilities.
var Log = guluLog{}

// Logging level.
const (
	Off = iota
	Trace
	Debug
	Info
	Warn
	Error
	Fatal
)

// all loggers.
var loggers []*Logger

// the global default logging level, it will be used for creating logger.
var logLevel = Debug

// Logger represents a simple logger with level.
// The underlying logger is the standard Go logging "log".
type Logger struct {
	level  int
	logger *stdlog.Logger
}

// NewLogger creates a logger.
func (*guluLog) NewLogger(out io.Writer) *Logger {
	ret := &Logger{level: logLevel, logger: stdlog.New(out, "", stdlog.Ldate|stdlog.Ltime|stdlog.Lshortfile)}

	loggers = append(loggers, ret)

	return ret
}

// SetLevel sets the logging level of all loggers.
func (*guluLog) SetLevel(level string) {
	logLevel = getLevel(level)

	for _, l := range loggers {
		l.SetLevel(level)
	}
}

// getLevel gets logging level int value corresponding to the specified level.
func getLevel(level string) int {
	level = strings.ToLower(level)

	switch level {
	case "off":
		return Off
	case "trace":
		return Trace
	case "debug":
		return Debug
	case "info":
		return Info
	case "warn":
		return Warn
	case "error":
		return Error
	case "fatal":
		return Fatal
	default:
		return Info
	}
}

// SetLevel sets the logging level of a logger.
func (l *Logger) SetLevel(level string) {
	l.level = getLevel(level)
}

// IsTraceEnabled determines whether the trace level is enabled.
func (l *Logger) IsTraceEnabled() bool {
	return l.level <= Trace
}

// IsDebugEnabled determines whether the debug level is enabled.
func (l *Logger) IsDebugEnabled() bool {
	return l.level <= Debug
}

// IsWarnEnabled determines whether the debug level is enabled.
func (l *Logger) IsWarnEnabled() bool {
	return l.level <= Warn
}

// Trace prints trace level message.
func (l *Logger) Trace(v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Tracef prints trace level message with format.
func (l *Logger) Tracef(format string, v ...interface{}) {
	if Trace < l.level {
		return
	}

	l.logger.SetPrefix("T ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Debug prints debug level message.
func (l *Logger) Debug(v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Debugf prints debug level message with format.
func (l *Logger) Debugf(format string, v ...interface{}) {
	if Debug < l.level {
		return
	}

	l.logger.SetPrefix("D ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Info prints info level message.
func (l *Logger) Info(v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Infof prints info level message with format.
func (l *Logger) Infof(format string, v ...interface{}) {
	if Info < l.level {
		return
	}

	l.logger.SetPrefix("I ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Warn prints warning level message.
func (l *Logger) Warn(v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Warnf prints warning level message with format.
func (l *Logger) Warnf(format string, v ...interface{}) {
	if Warn < l.level {
		return
	}

	l.logger.SetPrefix("W ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Error prints error level message.
func (l *Logger) Error(v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprintln(v...))
}

// Errorf prints error level message with format.
func (l *Logger) Errorf(format string, v ...interface{}) {
	if Error < l.level {
		return
	}

	l.logger.SetPrefix("E ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
}

// Fatal prints fatal level message and exit process with code 1.
func (l *Logger) Fatal(v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Fatalf prints fatal level message with format and exit process with code 1.
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if Fatal < l.level {
		return
	}

	l.logger.SetPrefix("F ")
	l.logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}
