package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type markdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Imd    string
}

func newMarkdownDoc() *markdownDoc { _ = "STUB: not implemented"; return nil }

func (d *markdownDoc) ID(v string) *markdownDoc { _ = "STUB: not implemented"; return nil }

func (d *markdownDoc) Class(v string) *markdownDoc { _ = "STUB: not implemented"; return nil }

func (d *markdownDoc) MD(v string) *markdownDoc { _ = "STUB: not implemented"; return nil }

func (d *markdownDoc) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *markdownDoc) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *markdownDoc) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (d *markdownDoc) highlightCode(ctx app.Context) { _ = "STUB: not implemented"; return }

func parseMarkdown(md []byte) []byte { _ = "STUB: not implemented"; return nil }

type remoteMarkdownDoc struct {
	app.Compo

	Iid    string
	Iclass string
	Isrc   string

	md markdownContent
}

func newRemoteMarkdownDoc() *remoteMarkdownDoc { _ = "STUB: not implemented"; return nil }

func (d *remoteMarkdownDoc) ID(v string) *remoteMarkdownDoc { _ = "STUB: not implemented"; return nil }

func (d *remoteMarkdownDoc) Class(v string) *remoteMarkdownDoc {
	_ = "STUB: not implemented"
	return nil
}

func (d *remoteMarkdownDoc) Src(v string) *remoteMarkdownDoc { _ = "STUB: not implemented"; return nil }

func (d *remoteMarkdownDoc) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *remoteMarkdownDoc) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *remoteMarkdownDoc) load(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *remoteMarkdownDoc) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
