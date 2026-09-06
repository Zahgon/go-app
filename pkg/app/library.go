package app

type Library interface {
	Styles() (path, styles string)
}
