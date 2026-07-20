package cli

import (
	"io"
)

const (
	defaultColor = "\033[0m"
	errorColor   = "\033[91m"
	successColor = "\033[92m"
	accentColor  = "\033[94m"
	focusColor   = "\033[1m"
	subColor     = "\033[2m"
)

func commandUsage(w io.Writer, cmd *command, opts []option) func() {
	_ = "STUB: not implemented"
	return nil
}

func commandUsageIndex(w io.Writer, cmds map[string]Command) func() {
	_ = "STUB: not implemented"
	return nil
}

func printError(w io.Writer, err error) { _ = "STUB: not implemented"; return }

func indent(w io.Writer, level int) int { _ = "STUB: not implemented"; return 0 }

func writeText(w io.Writer, text string, level, maxLen int) { _ = "STUB: not implemented"; return }

type optionFormatInfo struct {
	nameLen int
	typeLen int
}

func optionsInfo(opts []option) optionFormatInfo {
	_ = "STUB: not implemented"
	return *new(optionFormatInfo)
}
