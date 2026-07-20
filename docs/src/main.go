package main

import (
	"context"
	"net/http"
	"os"
	"sync"
	"syscall"

	"github.com/SherClockHolmes/webpush-go"
	"github.com/maxence-charriere/go-app/v11/pkg/analytics"
	"github.com/maxence-charriere/go-app/v11/pkg/app"
	"github.com/maxence-charriere/go-app/v11/pkg/cli"
	"github.com/maxence-charriere/go-app/v11/pkg/ui"
)

const (
	defaultTitle       = "A Go package for building Progressive Web Apps"
	defaultDescription = "A package for building progressive web apps (PWA) with the Go programming language (Golang) and WebAssembly (Wasm). It uses a declarative syntax that allows creating and dealing with HTML elements only by using Go, and without writing any HTML markup."
	backgroundColor    = "#2e343a"

	buyMeACoffeeURL     = "https://www.buymeacoffee.com/maxence"
	openCollectiveURL   = "https://opencollective.com/go-app"
	githubURL           = "https://github.com/maxence-charriere/go-app"
	githubSponsorURL    = "https://github.com/sponsors/maxence-charriere"
	twitterURL          = "https://twitter.com/jonhymaxoo"
	coinbaseBusinessURL = "https://commerce.coinbase.com/checkout/851320a4-35b5-41f1-897b-74dd5ee207ae"
)

type localOptions struct {
	Port            int    `cli:"p"                 env:"GOAPP_DOCS_PORT"   help:"The port used by the server that serves the PWA."`
	VAPIDPrivateKey string `cli:"vapid-private-key" env:"VAPID_PRIVATE_KEY" help:"The VAP id private key to sign push notifications."`
	VAPIDPublicKey  string `cli:"vapid-public-key"  env:"VAPID_PUBLIC_KEY"  help:"The VAP id public key to verify push notifications."`
}

type githubOptions struct {
	Output string `cli:"o" env:"-" help:"The directory where static resources are saved."`
}

func main() {
	ui.BaseHPadding = 42
	ui.BlockPadding = 18
	analytics.Add(analytics.NewGoogleAnalytics())

	app.Route("/", app.NewZeroComponentFactory(newHomePage()))
	app.Route("/getting-started", app.NewZeroComponentFactory(newGettingStartedPage()))
	app.Route("/architecture", app.NewZeroComponentFactory(newArchitecturePage()))
	app.Route("/reference", app.NewZeroComponentFactory(newReferencePage()))

	app.Route("/components", app.NewZeroComponentFactory(newComponentsPage()))
	app.Route("/declarative-syntax", app.NewZeroComponentFactory(newDeclarativeSyntaxPage()))
	app.Route("/routing", app.NewZeroComponentFactory(newRoutingPage()))
	app.Route("/static-resources", app.NewZeroComponentFactory(newStaticResourcePage()))
	app.Route("/js", app.NewZeroComponentFactory(newJSPage()))
	app.Route("/concurrency", app.NewZeroComponentFactory(newConcurrencyPage()))
	app.Route("/seo", app.NewZeroComponentFactory(newSEOPage()))
	app.Route("/lifecycle", app.NewZeroComponentFactory(newLifecyclePage()))
	app.Route("/install", app.NewZeroComponentFactory(newInstallPage()))
	app.Route("/testing", app.NewZeroComponentFactory(newTestingPage()))
	app.Route("/actions", app.NewZeroComponentFactory(newActionPage()))
	app.Route("/states", app.NewZeroComponentFactory(newStatesPage()))
	app.Route("/notifications", app.NewZeroComponentFactory(newNotificationsPage()))

	app.Route("/migrate", app.NewZeroComponentFactory(newMigratePage()))
	app.Route("/github-deploy", app.NewZeroComponentFactory(newGithubDeployPage()))

	app.Route("/privacy-policy", app.NewZeroComponentFactory(newPrivacyPolicyPage()))

	app.Handle(installApp, handleAppInstall)
	app.Handle(updateApp, handleAppUpdate)
	app.Handle(getMarkdown, handleGetMarkdown)
	app.Handle(getReference, handleGetReference)

	app.RunWhenOnBrowser()

	ctx, cancel := cli.ContextWithSignals(context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()
	defer exit()

	localOpts := localOptions{Port: 7777}
	cli.Register("local").
		Help(`Launches a server that serves the documentation app in a local environment.`).
		Options(&localOpts)

	githubOpts := githubOptions{}
	cli.Register("github").
		Help(`Generates the required resources to run the documentation app on GitHub Pages.`).
		Options(&githubOpts)

	h := app.Handler{
		Name:        "Documentation for go-app",
		Title:       defaultTitle,
		Description: defaultDescription,
		Author:      "Maxence Charriere",
		Image:       "https://go-app.dev/web/images/go-app.png",
		Keywords: []string{
			"go-app",
			"go",
			"golang",
			"app",
			"pwa",
			"progressive web app",
			"webassembly",
			"web assembly",
			"webapp",
			"web",
			"gui",
			"ui",
			"user interface",
			"graphical user interface",
			"frontend",
			"opensource",
			"open source",
			"github",
		},
		BackgroundColor: backgroundColor,
		ThemeColor:      backgroundColor,
		LoadingLabel:    "go-app documentation {progress}%",
		Styles: []string{

			"/web/css/prism.css",
			"/web/css/docs.css",
		},
		Scripts: []string{
			"/web/js/prism.js defer",
			"https://pagead2.googlesyndication.com/pagead/js/adsbygoogle.js?client=ca-pub-1013306768105236 async crossorigin=anonymous",
		},
		RawHeaders: []string{
			analytics.GoogleAnalyticsHeader("G-SW4FQEM9VM"),
		},
		CacheableResources: []string{
			"/web/documents/what-is-go-app.md",
			"/web/documents/updates.md",
			"/web/documents/home.md",
			"/web/documents/home-next.md",
		},
	}

	switch cli.Load() {
	case "local":
		runLocal(ctx, &h, localOpts)

	case "github":
		generateGitHubPages(ctx, &h, githubOpts)
	}
}

func runLocal(ctx context.Context, h *app.Handler, opts localOptions) {
	_ = "STUB: not implemented"
	return
}

func generateGitHubPages(ctx context.Context, h *app.Handler, opts githubOptions) {
	_ = "STUB: not implemented"
	return
}

func exit() { _ = "STUB: not implemented"; return }

type notificationHandler struct {
	VAPIDPrivateKey string
	VAPIDPublicKey  string

	mutex         sync.Mutex
	subscriptions map[string]webpush.Subscription
}

func (h *notificationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *notificationHandler) handleRegistrations(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

func (h *notificationHandler) handleTests(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
