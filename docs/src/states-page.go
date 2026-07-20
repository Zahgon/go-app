package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type statesPage struct {
	app.Compo
}

func newStatesPage() *statesPage { _ = "STUB: not implemented"; return nil }

func (p *statesPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *statesPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *statesPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
