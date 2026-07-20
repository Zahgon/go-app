package app

import (
	"context"
	"net/url"
	"time"
)

type Context struct {
	context.Context

	page                  func() Page
	appUpdatable          bool
	resolveURL            func(string) string
	navigate              func(*url.URL, bool)
	localStorage          BrowserStorage
	sessionStorage        BrowserStorage
	dispatch              func(func())
	defere                func(func())
	async                 func(func())
	addComponentUpdate    func(Composer, int)
	removeComponentUpdate func(Composer)
	handleAction          func(string, UI, bool, ActionHandler)
	postAction            func(Context, Action)
	observeState          func(Context, string, any) Observer
	getState              func(Context, string, any)
	setState              func(Context, string, any) State
	delState              func(Context, string)

	sourceElement        UI
	notifyComponentEvent func(Context, UI, any)
}

func (ctx Context) Src() UI { _ = "STUB: not implemented"; return *new(UI) }

func (ctx Context) JSSrc() Value { _ = "STUB: not implemented"; return *new(Value) }

func (ctx Context) AppUpdateAvailable() bool { _ = "STUB: not implemented"; return false }

func (ctx Context) IsAppInstallable() bool { _ = "STUB: not implemented"; return false }

func (ctx Context) IsAppleBrowser() bool { _ = "STUB: not implemented"; return false }

func (ctx Context) ShowAppInstallPrompt() { _ = "STUB: not implemented"; return }

func (ctx Context) DeviceID() string { _ = "STUB: not implemented"; return "" }

func (ctx Context) Page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (ctx Context) Reload() { _ = "STUB: not implemented"; return }

func (ctx Context) Navigate(rawURL string) { _ = "STUB: not implemented"; return }

func (ctx Context) NavigateTo(u *url.URL) { _ = "STUB: not implemented"; return }

func (ctx Context) ResolveStaticResource(v string) string { _ = "STUB: not implemented"; return "" }

func (ctx Context) ScrollTo(id string) { _ = "STUB: not implemented"; return }

func (ctx Context) LocalStorage() BrowserStorage {
	_ = "STUB: not implemented"
	return *new(BrowserStorage)
}

func (ctx Context) SessionStorage() BrowserStorage {
	_ = "STUB: not implemented"
	return *new(BrowserStorage)
}

func (ctx Context) Encrypt(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ctx Context) Decrypt(crypted []byte, v any) error { _ = "STUB: not implemented"; return nil }

func (ctx Context) cryptoKey() string { _ = "STUB: not implemented"; return "" }

func (ctx Context) Notifications() NotificationService {
	_ = "STUB: not implemented"
	return *new(NotificationService)
}

func (ctx Context) Dispatch(v func(Context)) { _ = "STUB: not implemented"; return }

func (ctx Context) Defer(v func(Context)) { _ = "STUB: not implemented"; return }

func (ctx Context) Async(v func()) { _ = "STUB: not implemented"; return }

func (ctx Context) After(d time.Duration, f func(Context)) { _ = "STUB: not implemented"; return }

func (ctx Context) PreventUpdate() { _ = "STUB: not implemented"; return }

func (ctx Context) Update() { _ = "STUB: not implemented"; return }

func (ctx Context) Handle(action string, h ActionHandler) { _ = "STUB: not implemented"; return }

func (ctx Context) NewAction(action string, tags ...Tagger) { _ = "STUB: not implemented"; return }

func (ctx Context) NewActionWithValue(action string, v any, tags ...Tagger) {
	_ = "STUB: not implemented"
	return
}

func (ctx Context) ObserveState(state string, recv any) Observer {
	_ = "STUB: not implemented"
	return *new(Observer)
}

func (ctx Context) GetState(state string, recv any) { _ = "STUB: not implemented"; return }

func (ctx Context) SetState(state string, v any) State {
	_ = "STUB: not implemented"
	return *new(State)
}

func (ctx Context) DelState(state string) { _ = "STUB: not implemented"; return }

func (ctx Context) ResizeContent() { _ = "STUB: not implemented"; return }
