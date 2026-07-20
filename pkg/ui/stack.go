package ui

import "github.com/maxence-charriere/go-app/v11/pkg/app"

type IStack interface {
	app.UI

	ID(v string) IStack

	Class(v string) IStack

	Style(k, v string) IStack

	Left() IStack

	Center() IStack

	Right() IStack

	Top() IStack

	Middle() IStack

	Bottom() IStack

	Stretch() IStack

	Content(elems ...app.UI) IStack
}

func Stack() IStack { _ = "STUB: not implemented"; return *new(IStack) }

type stack struct {
	app.Compo

	Iid              string
	Iclass           string
	IhorizontalAlign string
	IverticalAlign   string
	Istyles          []style
	Icontent         []app.UI
}

func (s *stack) ID(v string) IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Class(v string) IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Style(k, v string) IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Left() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Center() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Right() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Top() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Middle() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Bottom() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Stretch() IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) Content(elems ...app.UI) IStack { _ = "STUB: not implemented"; return *new(IStack) }

func (s *stack) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (s *stack) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
