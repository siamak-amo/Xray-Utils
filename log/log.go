// SPDX-License-Identifier: GPL-3.0-or-later
package log

import "os"

const (
	Verbose int = iota
	Info
	Warning
	Error
	None // no log at all
)

var (
	LogLevel int
)

func Must(level int) bool {
	return (Verbose == LogLevel) ||
		(None != LogLevel  &&  level >= LogLevel);
}

func Verbosef(format string, args ...interface{}) {
	if (Verbose == LogLevel) {
		MiniFlogf (os.Stderr, format, args...)
	}
}

func Infof(format string, args ...interface{}) {
	if Must (Info) {
		Flogf (os.Stderr, "Info", format, args...)
	}
}

func Warnf(format string, args ...interface{}) {
	if Must (Warning) {
		Flogf (os.Stderr, "Warning", format, args...)
	}
}

func Errorf(format string, args ...interface{}) {
	if Must (Error) {
		Flogf (os.Stderr, "Error", format, args...)
	}
}
