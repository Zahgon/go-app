package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
	"golang.org/x/net/html"
)

const (
	getReference   = "/reference/get"
	referenceState = "/reference"
)

func handleGetReference(ctx app.Context, a app.Action) { _ = "STUB: not implemented"; return }

type htmlContent struct {
	Status  status
	Err     error
	Index   string
	Content string
}

func getHTML(n *html.Node, class string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func findHTMLNode(n *html.Node, sel string) (*html.Node, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func normalizeHTMLNode(n *html.Node) { _ = "STUB: not implemented"; return }

func refLinkID(v string) string { _ = "STUB: not implemented"; return "" }
