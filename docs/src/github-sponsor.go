package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type githubSponsor struct {
	app.Compo

	Iclass string
}

func newGithubSponsor() *githubSponsor { _ = "STUB: not implemented"; return nil }

func (s *githubSponsor) Class(v string) *githubSponsor { _ = "STUB: not implemented"; return nil }

func (s *githubSponsor) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
