package cli

import (
	"io"

	"github.com/maxence-charriere/go-app/v11/pkg/errors"
)

var (
	errNoRootCmd = errors.New("no root command")
)

type command struct {
	help    string
	name    string
	options interface{}
}

func (c *command) Help(h string) Command { _ = "STUB: not implemented"; return *new(Command) }

func (c *command) Options(o interface{}) Command { _ = "STUB: not implemented"; return *new(Command) }

type commandManager struct {
	out      io.Writer
	commands map[string]Command
}

func (m *commandManager) register(cmd ...string) Command {
	_ = "STUB: not implemented"
	return *new(Command)
}

func (m *commandManager) parse(args ...string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func commandString(cmd ...string) string { _ = "STUB: not implemented"; return "" }

func splitCommand(args []string) (cmd, opts []string) { _ = "STUB: not implemented"; return nil, nil }

func commandEndIndex(args []string) int { _ = "STUB: not implemented"; return 0 }

type writerNoop struct{}

func (w writerNoop) Write([]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
