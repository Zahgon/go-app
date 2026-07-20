package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type ILink interface {
	app.UI

	ID(v string) ILink

	Class(v string) ILink

	Style(k, v string) ILink

	Icon(v string) ILink

	IconSize(px int) ILink

	IconSpace(px int) ILink

	Padding(v int) ILink

	Label(v string) ILink

	Help(v string) ILink

	Href(v string) ILink

	OnClick(v app.EventHandler) ILink
}

func Link() ILink { _ = "STUB: not implemented"; return *new(ILink) }

type link struct {
	app.Compo

	Iid        string
	Iclass     string
	Istyles    []style
	Iicon      string
	IiconSize  int
	IiconSpace int
	Ipadding   int
	Ilabel     string
	Ihelp      string
	Ihref      string
	IonClick   app.EventHandler
}

func (l *link) ID(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Class(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Style(k, v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Icon(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) IconSize(px int) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) IconSpace(px int) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Padding(px int) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Label(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Help(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Href(v string) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) OnClick(v app.EventHandler) ILink { _ = "STUB: not implemented"; return *new(ILink) }

func (l *link) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (l *link) onClick(ctx app.Context, e app.Event) { _ = "STUB: not implemented"; return }
