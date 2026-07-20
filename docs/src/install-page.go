package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type installPage struct {
	app.Compo
}

func newInstallPage() *installPage { _ = "STUB: not implemented"; return nil }

func (p *installPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *installPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *installPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
