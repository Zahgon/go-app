package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type actionPage struct {
	app.Compo
}

func newActionPage() *actionPage { _ = "STUB: not implemented"; return nil }

func (p *actionPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *actionPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *actionPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
