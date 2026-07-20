package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type migratePage struct {
	app.Compo
}

func newMigratePage() *migratePage { _ = "STUB: not implemented"; return nil }

func (p *migratePage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *migratePage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *migratePage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
