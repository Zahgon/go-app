package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type testingPage struct {
	app.Compo
}

func newTestingPage() *testingPage { _ = "STUB: not implemented"; return nil }

func (p *testingPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *testingPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *testingPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
