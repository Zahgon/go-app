package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type declarativeSyntaxPage struct {
	app.Compo
}

func newDeclarativeSyntaxPage() *declarativeSyntaxPage { _ = "STUB: not implemented"; return nil }

func (p *declarativeSyntaxPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *declarativeSyntaxPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *declarativeSyntaxPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
