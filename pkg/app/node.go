package app

import (
	"bytes"
	"io"
	"reflect"
)

type UI interface {
	JSValue() Value

	Mounted() bool

	parent() UI
	setParent(UI) UI
}

func FilterUIElems(v ...UI) []UI { _ = "STUB: not implemented"; return nil }

func HTMLString(ui UI) string { _ = "STUB: not implemented"; return "" }

func PrintHTML(w io.Writer, ui UI) { _ = "STUB: not implemented"; return }

type nav struct{}
type appUpdate struct{}
type appInstallChange struct{}
type resize struct{}

type nodeManager struct {
}

func (m nodeManager) Mount(ctx Context, depth uint, v UI) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) mountText(v *text) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) mountHTML(ctx Context, depth uint, v HTML) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) mountHTMLAttributes(ctx Context, v HTML) { _ = "STUB: not implemented"; return }

func (m nodeManager) mountHTMLEventHandlers(ctx Context, v HTML) { _ = "STUB: not implemented"; return }

func (m nodeManager) mountHTMLEventHandler(ctx Context, v HTML, handler eventHandler) eventHandler {
	_ = "STUB: not implemented"
	return *new(eventHandler)
}

func (m nodeManager) mountComponent(ctx Context, depth uint, v Composer) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) renderComponent(v Composer) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) mountRawHTML(depth uint, v *raw) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) Dismount(v UI) { _ = "STUB: not implemented"; return }

func (m nodeManager) dismountHTML(v HTML) { _ = "STUB: not implemented"; return }

func (m nodeManager) dismountHTMLEventHandler(handler eventHandler) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) dismountComponent(v Composer) { _ = "STUB: not implemented"; return }

func (m nodeManager) dismountRawHTML(v *raw) { _ = "STUB: not implemented"; return }

func (m nodeManager) CanUpdate(v, new UI) bool { _ = "STUB: not implemented"; return false }

func (m nodeManager) Update(ctx Context, v, new UI) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) updateText(v, new *text) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) updateHTML(ctx Context, v, new HTML) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) updateHTMLAttributes(ctx Context, v HTML, newAttrs attributes) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) updateHTMLEventHandlers(ctx Context, v HTML, newEvents eventHandlers) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) updateComponent(ctx Context, v, new Composer) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) UpdateComponentRoot(ctx Context, v Composer) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) updateRawHTML(ctx Context, v, new *raw) (UI, error) {
	_ = "STUB: not implemented"
	return *new(UI), nil
}

func (m nodeManager) context(ctx Context, v UI) Context {
	_ = "STUB: not implemented"
	return *new(Context)
}

func (m nodeManager) NotifyComponentEvent(ctx Context, root UI, event any) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) Encode(ctx Context, w *bytes.Buffer, v UI) { _ = "STUB: not implemented"; return }

func (m nodeManager) encode(ctx Context, w *bytes.Buffer, depth int, v UI) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) encodeText(w *bytes.Buffer, depth int, v *text) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) encodeIndent(w *bytes.Buffer, depth int) { _ = "STUB: not implemented"; return }

func (m nodeManager) encodeHTML(ctx Context, w *bytes.Buffer, depth int, v HTML) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) encodeHTMLAttribute(ctx Context, w *bytes.Buffer, name, value string) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) encodeComponent(ctx Context, w *bytes.Buffer, depth int, v Composer) {
	_ = "STUB: not implemented"
	return
}

func (m nodeManager) encodeRawHTML(w *bytes.Buffer, depth int, v *raw) {
	_ = "STUB: not implemented"
	return
}

func canUpdateValue(v, new reflect.Value) bool { _ = "STUB: not implemented"; return false }

func component(v UI) (Composer, bool) { _ = "STUB: not implemented"; return *new(Composer), false }
