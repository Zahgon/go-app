package app

type Composer interface {
	UI

	Render() UI

	setRef(Composer) Composer
	depth() uint
	setDepth(uint) Composer
	parent() UI
	root() UI
	setRoot(UI) Composer
}

type Initializer interface {
	OnInit()
}

type PreRenderer interface {
	OnPreRender(Context)
}

type Mounter interface {
	OnMount(Context)
}

type Dismounter interface {
	OnDismount()
}

type Navigator interface {
	OnNav(Context)
}

type DismountEnforcer interface {
	CompoID() string
}

type Updater interface {
	OnUpdate(Context)
}

type AppUpdater interface {
	OnAppUpdate(Context)
}

type AppInstaller interface {
	OnAppInstallChange(Context)
}

type Resizer interface {
	OnResize(Context)
}

type Compo struct {
	treeDepth     uint
	ref           Composer
	parentElement UI
	rootElement   UI
}

func (c *Compo) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (c *Compo) Mounted() bool { _ = "STUB: not implemented"; return false }

func (c *Compo) Render() UI { _ = "STUB: not implemented"; return *new(UI) }

func (c *Compo) ValueTo(v any) EventHandler { _ = "STUB: not implemented"; return *new(EventHandler) }

func (c *Compo) setRef(v Composer) Composer { _ = "STUB: not implemented"; return *new(Composer) }

func (c *Compo) depth() uint { _ = "STUB: not implemented"; return 0 }

func (c *Compo) setDepth(v uint) Composer { _ = "STUB: not implemented"; return *new(Composer) }

func (c *Compo) parent() UI { _ = "STUB: not implemented"; return *new(UI) }

func (c *Compo) setParent(p UI) UI { _ = "STUB: not implemented"; return *new(UI) }

func (c *Compo) root() UI { _ = "STUB: not implemented"; return *new(UI) }

func (c *Compo) setRoot(v UI) Composer { _ = "STUB: not implemented"; return *new(Composer) }
