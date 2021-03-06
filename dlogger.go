package dlogger

import (
	"fmt"
	"io"
	"os"
)

// DebugLogger is a logger for debug.
type DebugLogger struct {
	w     io.Writer
	debug bool
	tag   string
}

// New creates a new Debug Logger.
func New(w io.Writer) *DebugLogger {
	dlogger := &DebugLogger{w: w, tag: "[DEBUG] "}

	if len(os.Getenv("DEBUG")) != 0 {
		dlogger.debug = true
	}

	return dlogger
}

// Printf print log with format if DEBUG env specified.
func (dlogger *DebugLogger) Printf(format string, a ...interface{}) (n int, err error) {
	if dlogger.debug {
		return fmt.Fprintf(dlogger.w, dlogger.tag+format, a...)
	}

	return 0, nil
}

// Print print log if DEBUG env specified.
func (dlogger *DebugLogger) Print(a ...interface{}) (n int, err error) {
	if dlogger.debug {
		a = append([]interface{}{dlogger.tag}, a...)
		return fmt.Fprint(dlogger.w, a...)
	}

	return 0, nil
}
