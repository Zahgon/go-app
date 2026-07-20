package app

import (
	"bytes"
	"context"
	"net/url"
	"sync"
)

type engineX struct {
	ctx context.Context

	localStorage   BrowserStorage
	sessionStorage BrowserStorage
	browser        browser

	routes         *router
	internalURLs   []string
	resolveURL     func(string) string
	originPage     *requestPage
	lastVisitedURL *url.URL

	nodes   nodeManager
	updates updateManager
	body    HTMLBody

	dispatches chan func()
	defers     chan func()
	goroutines sync.WaitGroup

	asynchronousActionHandlers map[string]ActionHandler
	actions                    actionManager
	states                     stateManager
}

func newEngine(ctx context.Context, routes *router, resolveURL func(string) string, originPage *requestPage, actionHandlers map[string]ActionHandler) *engineX {
	_ = "STUB: not implemented"
	return nil
}

func (e *engineX) baseContext() Context { _ = "STUB: not implemented"; return *new(Context) }

func (e *engineX) Navigate(destination *url.URL, updateHistory bool) {
	_ = "STUB: not implemented"
	return
}

func (e *engineX) initBrowser() { _ = "STUB: not implemented"; return }

func (e *engineX) notifyComponentEvent(event any) { _ = "STUB: not implemented"; return }

func (e *engineX) externalNavigation(v *url.URL) bool { _ = "STUB: not implemented"; return false }

func (e *engineX) mailTo(v *url.URL) bool { _ = "STUB: not implemented"; return false }

func (e *engineX) internalURL(v *url.URL) bool { _ = "STUB: not implemented"; return false }

func (e *engineX) page() Page { _ = "STUB: not implemented"; return *new(Page) }

func (e *engineX) Load(v Composer) error { _ = "STUB: not implemented"; return nil }

func (e *engineX) Start(framerate int) { _ = "STUB: not implemented"; return }

func (e *engineX) processFrame() { _ = "STUB: not implemented"; return }

func (e *engineX) executeDefers() { _ = "STUB: not implemented"; return }

func (e *engineX) ConsumeNext() { _ = "STUB: not implemented"; return }

func (e *engineX) ConsumeAll() { _ = "STUB: not implemented"; return }

func (e *engineX) Encode(w *bytes.Buffer, document HTMLHtml) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *engineX) dispatch(v func()) { _ = "STUB: not implemented"; return }

func (e *engineX) defere(v func()) { _ = "STUB: not implemented"; return }

func (e *engineX) async(v func()) { _ = "STUB: not implemented"; return }
