package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type privacyPolicyPage struct {
	app.Compo
}

func newPrivacyPolicyPage() *privacyPolicyPage { _ = "STUB: not implemented"; return nil }

func (p *privacyPolicyPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *privacyPolicyPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *privacyPolicyPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
