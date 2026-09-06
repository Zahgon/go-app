package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type referenceContent struct {
	app.Compo

	Iid    string
	Iclass string
	Iindex bool

	content         htmlContent
	currentFragment string
}

func newReferenceContent() *referenceContent { _ = "STUB: not implemented"; return nil }

func (c *referenceContent) ID(v string) *referenceContent { _ = "STUB: not implemented"; return nil }

func (c *referenceContent) Class(v string) *referenceContent { _ = "STUB: not implemented"; return nil }

func (c *referenceContent) Index(v bool) *referenceContent { _ = "STUB: not implemented"; return nil }

func (c *referenceContent) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) load(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (c *referenceContent) handleFragment(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) unfocusCurrentIndex(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) focusCurrentIndex(ctx app.Context) { _ = "STUB: not implemented"; return }

func (c *referenceContent) scrollTo(ctx app.Context) { _ = "STUB: not implemented"; return }
