package cli

import (
	"context"
	"os"
)

var (
	defaultManager = commandManager{out: os.Stderr}
	currentUsage   func()
	exitOnError    = true
	programArgs    = os.Args[1:]
)

type Command interface {
	Help(string) Command

	Options(interface{}) Command
}

func Register(cmd ...string) Command { _ = "STUB: not implemented"; return *new(Command) }

func Load() (cmd string) { _ = "STUB: not implemented"; return "" }

func Usage() { _ = "STUB: not implemented"; return }

func Error(err error) { _ = "STUB: not implemented"; return }

func ContextWithSignals(parent context.Context, sig ...os.Signal) (ctx context.Context, cancel func()) {
	_ = "STUB: not implemented"
	return *new(context.Context), nil
}
