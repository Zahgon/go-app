package ui

var (
	BlockPadding = 30

	BlockMobilePadding = 18

	BlockContentWidth = 580

	BaseHPadding = 36

	BaseMobileHPadding = 12

	BaseAdHPadding = BaseHPadding / 2

	BaseVPadding = 12

	DefaultIconSize = 24

	DefaultIconSpace = 6

	DefaultFlowItemWidth = 372
)

const (
	defaultHeaderHeight = 90
)

func pxToString(px int) string { _ = "STUB: not implemented"; return "" }

type alignment int

const (
	stretch alignment = iota
	top
	right
	bottom
	left
	middle
)

type style struct {
	key   string
	value string
}
