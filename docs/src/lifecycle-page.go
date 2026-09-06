package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type lifecyclePage struct {
	app.Compo
}

func newLifecyclePage() *lifecyclePage { _ = "STUB: not implemented"; return nil }

func (p *lifecyclePage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *lifecyclePage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *lifecyclePage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
