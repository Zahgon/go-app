package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type gettingStartedPage struct {
	app.Compo
}

func newGettingStartedPage() *gettingStartedPage { _ = "STUB: not implemented"; return nil }

func (p *gettingStartedPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *gettingStartedPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *gettingStartedPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
