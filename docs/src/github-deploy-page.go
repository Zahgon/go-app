package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type githubDeployPage struct {
	app.Compo
}

func newGithubDeployPage() *githubDeployPage { _ = "STUB: not implemented"; return nil }

func (p *githubDeployPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *githubDeployPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *githubDeployPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
