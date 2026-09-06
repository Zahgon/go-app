package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type referencePage struct {
	app.Compo
}

func newReferencePage() *referencePage { _ = "STUB: not implemented"; return nil }

func (p *referencePage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *referencePage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *referencePage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
