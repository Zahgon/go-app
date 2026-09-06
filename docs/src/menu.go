package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type menu struct {
	app.Compo

	Iclass string

	appInstallable bool
}

func newMenu() *menu { _ = "STUB: not implemented"; return nil }

func (m *menu) Class(v string) *menu { _ = "STUB: not implemented"; return nil }

func (m *menu) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (m *menu) OnAppInstallChange(ctx app.Context) { _ = "STUB: not implemented"; return }

func (m *menu) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (m *menu) installApp(ctx app.Context, e app.Event) { _ = "STUB: not implemented"; return }
