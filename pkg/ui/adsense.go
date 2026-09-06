package ui

import (
	"sync"
	"time"

	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IAdsenseDisplay interface {
	app.UI

	ID(v string) IAdsenseDisplay

	Class(v string) IAdsenseDisplay

	Client(v string) IAdsenseDisplay

	Slot(v string) IAdsenseDisplay
}

func AdsenseDisplay() IAdsenseDisplay { _ = "STUB: not implemented"; return *new(IAdsenseDisplay) }

type adsenseDisplay struct {
	app.Compo

	Iid     string
	Iclass  string
	Iclient string
	Islot   string

	id          string
	currentPath string
	width       int
	height      int
}

func (d *adsenseDisplay) ID(v string) IAdsenseDisplay {
	_ = "STUB: not implemented"
	return *new(IAdsenseDisplay)
}

func (d *adsenseDisplay) Class(v string) IAdsenseDisplay {
	_ = "STUB: not implemented"
	return *new(IAdsenseDisplay)
}

func (d *adsenseDisplay) Client(v string) IAdsenseDisplay {
	_ = "STUB: not implemented"
	return *new(IAdsenseDisplay)
}

func (d *adsenseDisplay) Slot(v string) IAdsenseDisplay {
	_ = "STUB: not implemented"
	return *new(IAdsenseDisplay)
}

func (d *adsenseDisplay) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *adsenseDisplay) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *adsenseDisplay) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *adsenseDisplay) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (d *adsenseDisplay) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (d *adsenseDisplay) containerID() string { _ = "STUB: not implemented"; return "" }

func (d *adsenseDisplay) client() string { _ = "STUB: not implemented"; return "" }

func (d *adsenseDisplay) slot() string { _ = "STUB: not implemented"; return "" }

func (d *adsenseDisplay) load(ctx app.Context) { _ = "STUB: not implemented"; return }

type adUnit interface {
	app.UI

	containerID() string
	client() string
	slot() string
}

type adPimp struct {
	mutex    sync.Mutex
	units    map[adUnit]struct{}
	interval time.Duration
}

func newAdPimp() *adPimp { _ = "STUB: not implemented"; return nil }

func (p *adPimp) Push(ctx app.Context, units ...adUnit) { _ = "STUB: not implemented"; return }

func (p *adPimp) push(ctx app.Context, units ...adUnit) { _ = "STUB: not implemented"; return }

func (p *adPimp) addUnit(u adUnit) { _ = "STUB: not implemented"; return }

func (p *adPimp) pushUnit(u adUnit) { _ = "STUB: not implemented"; return }

func (p *adPimp) removeUnit(u adUnit) { _ = "STUB: not implemented"; return }

var (
	ads = newAdPimp()
)
