package ui

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type IFlyer interface {
	app.UI

	ID(v string) IFlyer

	Class(v string) IFlyer

	HeaderHeight(px int) IFlyer

	PremiumHeight(px int) IFlyer

	FooterHeight(px int) IFlyer

	Banner(v ...app.UI) IFlyer

	Premium(v ...app.UI) IFlyer

	Bonus(v ...app.UI) IFlyer
}

func Flyer() IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

type flyer struct {
	app.Compo

	Iid            string
	Iclass         string
	IheaderHeight  int
	IpremiumHeight int
	IfooterHeight  int
	Ibanner        []app.UI
	Ipremium       []app.UI
	Ibonus         []app.UI

	hpadding      int
	vpadding      int
	bannerHeight  int
	premiumHeight int
	bonusHeight   int
	layoutID      string
}

func (f *flyer) ID(v string) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) Class(v string) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) HeaderHeight(px int) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) PremiumHeight(px int) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) FooterHeight(px int) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) Banner(v ...app.UI) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) Premium(v ...app.UI) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) Bonus(v ...app.UI) IFlyer { _ = "STUB: not implemented"; return *new(IFlyer) }

func (f *flyer) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flyer) OnResize(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flyer) OnUpdate(ctx app.Context) { _ = "STUB: not implemented"; return }

func (f *flyer) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (f *flyer) resize(ctx app.Context) { _ = "STUB: not implemented"; return }
