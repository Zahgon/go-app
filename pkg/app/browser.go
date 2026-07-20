package app

import (
	"time"
)

type browser struct {
	AppUpdatable bool

	anchorClick      Func
	popState         Func
	navigationFromJS Func
	appUpdate        Func
	appInstallChange Func
	appResize        Func
	resizeTimer      *time.Timer
}

func (b *browser) HandleEvents(ctx Context, notifyComponentEvent func(any)) {
	_ = "STUB: not implemented"
	return
}

func (b *browser) handleAnchorClick(ctx Context) { _ = "STUB: not implemented"; return }

func (b *browser) handlePopState(ctx Context) { _ = "STUB: not implemented"; return }

func (b *browser) handleNavigationFromJS(ctx Context) { _ = "STUB: not implemented"; return }

func (b *browser) handleAppUpdate(ctx Context, notifyComponentEvent func(any)) {
	_ = "STUB: not implemented"
	return
}

func (b *browser) handleAppInstallChange(ctx Context, notifyComponentEvent func(any)) {
	_ = "STUB: not implemented"
	return
}

func (b *browser) handleAppResize(ctx Context, notifyComponentEvent func(any)) {
	_ = "STUB: not implemented"
	return
}
