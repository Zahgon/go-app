package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type seoPage struct {
	app.Compo
}

func newSEOPage() *seoPage { _ = "STUB: not implemented"; return nil }

func (p *seoPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *seoPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *seoPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
