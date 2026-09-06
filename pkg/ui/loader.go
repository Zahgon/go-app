package ui

import (
	"time"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

const (
	defaultLoaderErrorIcon = `<svg style="width:%vpx;height:%vpx" viewBox="0 0 24 24">
    	<path fill="currentColor" d="M22 14H21C21 10.13 17.87 7 14 7H13V5.73C13.6 5.39 14 4.74 14 4C14 2.9 13.11 2 12 2S10 2.9 10 4C10 4.74 10.4 5.39 11 5.73V7H10C6.13 7 3 10.13 3 14H2C1.45 14 1 14.45 1 15V18C1 18.55 1.45 19 2 19H3V20C3 21.11 3.9 22 5 22H19C20.11 22 21 21.11 21 20V19H22C22.55 19 23 18.55 23 18V15C23 14.45 22.55 14 22 14M9.86 16.68L8.68 17.86L7.5 16.68L6.32 17.86L5.14 16.68L6.32 15.5L5.14 14.32L6.32 13.14L7.5 14.32L8.68 13.14L9.86 14.32L8.68 15.5L9.86 16.68M18.86 16.68L17.68 17.86L16.5 16.68L15.32 17.86L14.14 16.68L15.32 15.5L14.14 14.32L15.32 13.14L16.5 14.32L17.68 13.14L18.86 14.32L17.68 15.5L18.86 16.68Z" />
	</svg>`
)

type ILoader interface {
	app.UI

	ID(v string) ILoader

	Class(v string) ILoader

	Style(k, v string) ILoader

	Loading(v bool) ILoader

	Size(px int) ILoader

	Color(v string) ILoader

	Speed(v time.Duration) ILoader

	Spacing(px int) ILoader

	Label(v string) ILoader

	Err(err error) ILoader

	ErrIcon(v string) ILoader
}

func Loader() ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

type loader struct {
	app.Compo

	Iid      string
	Iclass   string
	Istyles  []style
	Iloading bool
	Isize    int
	Icolor   string
	Ispacing int
	Ilabel   string
	Ispeed   time.Duration
	Ierr     error
	IerrIcon string
}

func (l *loader) ID(v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Class(v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Style(k, v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Loading(v bool) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Size(px int) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Color(v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Speed(v time.Duration) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Spacing(px int) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Label(v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Err(err error) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) ErrIcon(v string) ILoader { _ = "STUB: not implemented"; return *new(ILoader) }

func (l *loader) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }
