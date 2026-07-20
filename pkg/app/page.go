package app

import (
	"net/url"
)

type Page interface {
	Title() string

	SetTitle(v string)

	SetTitlef(format string, v ...any)

	Lang() string

	SetLang(v string)

	Description() string

	SetDescription(v string)

	SetDescriptionf(format string, v ...any)

	Author() string

	SetAuthor(v string)

	SetAuthorf(format string, v ...any)

	Keywords() string

	SetKeywords(v ...string)

	Preloads() []Preload

	SetPreloads(v ...Preload)

	SetLoadingLabel(v string)

	SetLoadingLabelf(format string, v ...any)

	Image() string

	SetImage(v string)

	SetImagef(format string, v ...any)

	URL() *url.URL

	ReplaceURL(v *url.URL)

	Size() (w int, h int)

	SetTwitterCard(v TwitterCard)

	SetCanonicalLink(v string)

	SetCanonicalLinkf(format string, v ...any)
}

type requestPage struct {
	url        *url.URL
	resolveURL func(string) string

	title          string
	lang           string
	description    string
	author         string
	keywords       string
	preloads       []Preload
	loadingLabel   string
	image          string
	canonicalLink  string
	width          int
	height         int
	twitterCardMap map[string]string
}

func makeRequestPage(origin *url.URL, resolveURL func(string) string) requestPage {
	_ = "STUB: not implemented"
	return *new(requestPage)
}

func (p *requestPage) Title() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetTitle(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetTitlef(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p *requestPage) Lang() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetLang(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) Description() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetDescription(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetDescriptionf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p *requestPage) Author() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetAuthor(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetAuthorf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p *requestPage) Keywords() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetKeywords(v ...string) { _ = "STUB: not implemented"; return }

func (p *requestPage) Preloads() []Preload { _ = "STUB: not implemented"; return nil }

func (p *requestPage) SetPreloads(v ...Preload) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetLoadingLabel(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetLoadingLabelf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p *requestPage) Image() string { _ = "STUB: not implemented"; return "" }

func (p *requestPage) SetImage(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetImagef(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p *requestPage) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (p *requestPage) ReplaceURL(v *url.URL) { _ = "STUB: not implemented"; return }

func (p *requestPage) Size() (width int, height int) { _ = "STUB: not implemented"; return 0, 0 }

func (p *requestPage) SetTwitterCard(v TwitterCard) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetCanonicalLink(v string) { _ = "STUB: not implemented"; return }

func (p *requestPage) SetCanonicalLinkf(format string, v ...any) { _ = "STUB: not implemented"; return }

type browserPage struct {
	resolveURL func(string) string
}

func makeBrowserPage(resolveURL func(string) string) browserPage {
	_ = "STUB: not implemented"
	return *new(browserPage)
}

func (p browserPage) Title() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetTitle(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetTitlef(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) Lang() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetLang(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) Description() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetDescription(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetDescriptionf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) Author() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetAuthor(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetAuthorf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) Keywords() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetKeywords(v ...string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetLoadingLabel(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetLoadingLabelf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) Preloads() []Preload { _ = "STUB: not implemented"; return nil }

func (p browserPage) SetPreloads(v ...Preload) { _ = "STUB: not implemented"; return }

func (p browserPage) Image() string { _ = "STUB: not implemented"; return "" }

func (p browserPage) SetImage(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetImagef(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (p browserPage) ReplaceURL(v *url.URL) { _ = "STUB: not implemented"; return }

func (p browserPage) Size() (width int, height int) { _ = "STUB: not implemented"; return 0, 0 }

func (p browserPage) SetTwitterCard(v TwitterCard) { _ = "STUB: not implemented"; return }

func (p browserPage) SetCanonicalLink(v string) { _ = "STUB: not implemented"; return }

func (p browserPage) SetCanonicalLinkf(format string, v ...any) { _ = "STUB: not implemented"; return }

func (p browserPage) metaByName(v string) Value { _ = "STUB: not implemented"; return *new(Value) }

func (p browserPage) metaByProperty(v string) Value { _ = "STUB: not implemented"; return *new(Value) }

type Preload struct {
	Type          string
	As            string
	Href          string
	FetchPriority string
}
