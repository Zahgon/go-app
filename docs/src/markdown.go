package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

const (
	getMarkdown = "/markdown/get"
)

func handleGetMarkdown(ctx app.Context, a app.Action) { _ = "STUB: not implemented"; return }

func markdownState(src string) string { _ = "STUB: not implemented"; return "" }

type markdownContent struct {
	Status status
	Err    error
	Data   string
}
