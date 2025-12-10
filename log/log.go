// SPDX-License-Identifier: GPL-3.0-or-later
package log

import "os"

func Warnf(format string, args ...interface{}) {
	Flogf (os.Stderr, "Warning", format, args...)
}

func Errorf(format string, args ...interface{}) {
	Flogf (os.Stderr, "Error", format, args...)
}
