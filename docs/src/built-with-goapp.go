package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type builtWithGoapp struct {
	app.Compo

	Iid    string
	Iclass string
}

func newBuiltWithGoapp() *builtWithGoapp { _ = "STUB: not implemented"; return nil }

func (b *builtWithGoapp) ID(v string) *builtWithGoapp { _ = "STUB: not implemented"; return nil }

func (b *builtWithGoapp) Class(v string) *builtWithGoapp { _ = "STUB: not implemented"; return nil }

func (b *builtWithGoapp) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

type builtWithGoappItem struct {
	app.Compo

	Iclass       string
	Iimage       string
	Iname        string
	Idescription string
	Ihref        string
}

func newBuiltWithGoappItem() *builtWithGoappItem { _ = "STUB: not implemented"; return nil }

func (i *builtWithGoappItem) Class(v string) *builtWithGoappItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *builtWithGoappItem) Image(v string) *builtWithGoappItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *builtWithGoappItem) Name(v string) *builtWithGoappItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *builtWithGoappItem) Description(v string) *builtWithGoappItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *builtWithGoappItem) Href(v string) *builtWithGoappItem {
	_ = "STUB: not implemented"
	return nil
}

func (i *builtWithGoappItem) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
