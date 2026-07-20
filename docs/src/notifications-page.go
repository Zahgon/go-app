package main

import (
	"github.com/maxence-charriere/go-app/v11/pkg/app"
)

type notificationsPage struct {
	app.Compo

	notificationPermission app.NotificationPermission
}

func newNotificationsPage() *notificationsPage { _ = "STUB: not implemented"; return nil }

func (p *notificationsPage) OnMount(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *notificationsPage) OnNav(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *notificationsPage) initPage(ctx app.Context) { _ = "STUB: not implemented"; return }

func (p *notificationsPage) Render() app.UI { _ = "STUB: not implemented"; return *new(app.UI) }

func (p *notificationsPage) enableNotifications(ctx app.Context, e app.Event) {
	_ = "STUB: not implemented"
	return
}

func (p *notificationsPage) testNotification(ctx app.Context, e app.Event) {
	_ = "STUB: not implemented"
	return
}

func (p *notificationsPage) registerSubscription(ctx app.Context) {
	_ = "STUB: not implemented"
	return
}
