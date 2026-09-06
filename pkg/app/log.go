package app

import (
	"runtime"
)

var (
	DefaultLogger func(format string, v ...any)

	defaultColor string
	errorColor   string
	infoColor    string
)

func init() {
	goarch := runtime.GOARCH
	if goarch == "wasm" {
		DefaultLogger = clientLog
		return
	}

	if goarch != "window" {
		defaultColor = "\033[00m"
		errorColor = "\033[91m"
		infoColor = "\033[94m"
	}
	DefaultLogger = serverLog
}

func Log(v ...any) { _ = "STUB: not implemented"; return }

func Logf(format string, v ...any) { _ = "STUB: not implemented"; return }

func serverLog(format string, v ...any) { _ = "STUB: not implemented"; return }

func clientLog(format string, v ...any) { _ = "STUB: not implemented"; return }
