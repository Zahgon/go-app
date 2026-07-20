//go:build wasm
// +build wasm

package app

func GenerateStaticWebsite(dir string, h *Handler, pages ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func wasmExecJS() string { _ = "STUB: not implemented"; return "" }
