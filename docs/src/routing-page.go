package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type routingPage struct {
	app.Compo
}

func newRoutingPage() *routingPage { _ = "STUB: not implemented"; return nil }

func (p *routingPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *routingPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *routingPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
