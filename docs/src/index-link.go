package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type indexLink struct {
	app.Compo

	Iclass string
	Ititle string
	Ihref  string
}

func newIndexLink() *indexLink { _ = "STUB: not implemented"; return nil }

func (l *indexLink) Class(v string) *indexLink { _ = "STUB: not implemented"; return nil }

func (l *indexLink) Title(v string) *indexLink { _ = "STUB: not implemented"; return nil }

func (l *indexLink) Href(v string) *indexLink { _ = "STUB: not implemented"; return nil }

func (l *indexLink) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (l *indexLink) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func titleToFragment(v string) string { _ = "STUB: not implemented"; return "" }
