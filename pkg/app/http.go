package app

import (
	"net/http"
	"sync"
)

const (
	defaultThemeColor = "#2d2c2c"
)

type Handler struct {
	Name string

	ShortName string

	Icon Icon

	BackgroundColor string

	ThemeColor string

	LoadingLabel string

	Lang string

	Libraries []Library

	Title string

	Description string

	Domain string

	Author string

	Keywords []string

	Image string

	Styles []string

	Fonts []string

	Scripts []string

	CacheableResources []string

	RawHeaders []string

	HTML func() HTMLHtml

	Body func() HTMLBody

	Env map[string]string

	InternalURLs []string

	Preconnect []string

	ProxyResources []ProxyResource

	Resources ResourceResolver

	StartURL string

	Version string

	WasmContentLength string

	WasmContentLengthHeader string

	ServiceWorkerTemplate string

	once                 sync.Once
	etag                 string
	libraries            map[string][]byte
	proxyResources       map[string]ProxyResource
	cachedProxyResources *memoryCache
	cachedPWAResources   *memoryCache
}

func (h *Handler) init() {
	h.initVersion()
	h.initStaticResources()
	h.initLibraries()
	h.initLinks()
	h.initServiceWorker()
	h.initIcon()
	h.initPWA()
	h.initPageContent()
	h.initPWAResources()
	h.initProxyResources()
}

func (h *Handler) initVersion() { _ = "STUB: not implemented"; return }

func (h *Handler) initStaticResources() { _ = "STUB: not implemented"; return }

func (h *Handler) initLibraries() { _ = "STUB: not implemented"; return }

func (h *Handler) initLinks() { _ = "STUB: not implemented"; return }

func (h *Handler) initServiceWorker() { _ = "STUB: not implemented"; return }

func (h *Handler) initIcon() { _ = "STUB: not implemented"; return }

func (h *Handler) initPWA() { _ = "STUB: not implemented"; return }

func (h *Handler) initPageContent() { _ = "STUB: not implemented"; return }

func (h *Handler) initPWAResources() { _ = "STUB: not implemented"; return }

func (h *Handler) makeAppJS() []byte { _ = "STUB: not implemented"; return nil }

func (h *Handler) makeAppWorkerJS() []byte { _ = "STUB: not implemented"; return nil }

func (h *Handler) makeManifestJSON() []byte { _ = "STUB: not implemented"; return nil }

func (h *Handler) initProxyResources() { _ = "STUB: not implemented"; return }

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) serveCachedItem(w http.ResponseWriter, i cacheItem) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) serveProxyResource(resource ProxyResource, w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) servePage(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *Handler) serveLibrary(w http.ResponseWriter, r *http.Request, library []byte) {
	_ = "STUB: not implemented"
	return
}

type Icon struct {
	Default string

	Large string

	SVG string

	Maskable string
}

func isRemoteLocation(path string) bool { _ = "STUB: not implemented"; return false }

func isStaticResourcePath(path string) bool { _ = "STUB: not implemented"; return false }

type httpResource struct {
	URL         string
	LoadingMode string
	CrossOrigin string
}

func (r httpResource) toLink() HTMLLink { _ = "STUB: not implemented"; return *new(HTMLLink) }

func (r httpResource) toScript() HTMLScript { _ = "STUB: not implemented"; return *new(HTMLScript) }

func parseHTTPResource(v string) httpResource { _ = "STUB: not implemented"; return *new(httpResource) }
