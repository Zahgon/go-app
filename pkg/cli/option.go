package cli

import (
	"flag"
	"reflect"
)

type optionParser struct {
	flags   *flag.FlagSet
	options []option
}

func (p *optionParser) parse(v interface{}) ([]option, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *optionParser) parseStruct(prefix string, v reflect.Value) {
	_ = "STUB: not implemented"
	return
}

type option struct {
	name   string
	help   string
	envKey string
	value  reflect.Value
}

func (o option) IsBoolFlag() bool { _ = "STUB: not implemented"; return false }

func (o option) String() string { _ = "STUB: not implemented"; return "" }

func (o option) Set(s string) error { _ = "STUB: not implemented"; return nil }

func setDuration(v reflect.Value, s string) error { _ = "STUB: not implemented"; return nil }

func normalizeOptionName(name string, sep string) string { _ = "STUB: not implemented"; return "" }

func normalizeCLIOptionName(name string) string { _ = "STUB: not implemented"; return "" }

func normalizeEnvOptionName(name string) string { _ = "STUB: not implemented"; return "" }

func isUpperCase(b byte) bool { _ = "STUB: not implemented"; return false }
