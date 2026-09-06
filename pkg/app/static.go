//go:build !wasm
// +build !wasm

package app

import (
	"os"
)

func GenerateStaticWebsite(dir string, h *Handler, pages ...string) error {
	_ = "STUB: not implemented"
	return nil
}

func createStaticDir(dir, path string) error { _ = "STUB: not implemented"; return nil }

func createStaticFile(dir, path string) (*os.File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func createStaticPage(path string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
