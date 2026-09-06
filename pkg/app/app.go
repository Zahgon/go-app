//go:generate go run gen/html.go
//go:generate go run gen/scripts.go
//go:generate go fmt

package app

import (
	"runtime"
)

const (
	IsClient = runtime.GOARCH == "wasm" && runtime.GOOS == "js"

	IsServer = runtime.GOARCH != "wasm" || runtime.GOOS != "js"
)

var (
	routes = makeRouter()
	window = newBrowserWindow()
)

func Getenv(k string) string { _ = "STUB: not implemented"; return "" }

func KeepBodyClean() (close func()) { _ = "STUB: not implemented"; return nil }

func Window() BrowserWindow { _ = "STUB: not implemented"; return *new(BrowserWindow) }

func RunWhenOnBrowser() { _ = "STUB: not implemented"; return }

func displayLoadError(err any) { _ = "STUB: not implemented"; return }

func Route(path string, newComponent func() Composer) { _ = "STUB: not implemented"; return }

func RouteWithRegexp(pattern string, newComponent func() Composer) {
	_ = "STUB: not implemented"
	return
}

func NewZeroComponentFactory(c Composer) func() Composer { _ = "STUB: not implemented"; return nil }

func TryUpdate() { _ = "STUB: not implemented"; return }
