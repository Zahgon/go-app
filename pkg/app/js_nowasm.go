//go:build !wasm
// +build !wasm

package app

import (
	"net/url"
)

type value struct{}

func (v value) Bool() bool { _ = "STUB: not implemented"; return false }

func (v value) Call(m string, args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Delete(p string) { _ = "STUB: not implemented"; return }

func (v value) Equal(w Value) bool { _ = "STUB: not implemented"; return false }

func (v value) Float() float64 { _ = "STUB: not implemented"; return 0 }

func (v value) Get(p string) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Index(i int) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) InstanceOf(t Value) bool { _ = "STUB: not implemented"; return false }

func (v value) Int() int { _ = "STUB: not implemented"; return 0 }

func (v value) Invoke(args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) IsNaN() bool { _ = "STUB: not implemented"; return false }

func (v value) IsNull() bool { _ = "STUB: not implemented"; return false }

func (v value) IsUndefined() bool { _ = "STUB: not implemented"; return false }

func (v value) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Length() int { _ = "STUB: not implemented"; return 0 }

func (v value) New(args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Set(p string, x any) { _ = "STUB: not implemented"; return }

func (v value) SetIndex(i int, x any) { _ = "STUB: not implemented"; return }

func (v value) String() string { _ = "STUB: not implemented"; return "" }

func (v value) Truthy() bool { _ = "STUB: not implemented"; return false }

func (v value) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (v value) Then(f func(Value)) { _ = "STUB: not implemented"; return }

func (v value) getAttr(k string) string { _ = "STUB: not implemented"; return "" }

func (v value) setAttr(k, val string) { _ = "STUB: not implemented"; return }

func (v value) delAttr(k string) { _ = "STUB: not implemented"; return }

func (v value) firstChild() Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) appendChild(c Wrapper) { _ = "STUB: not implemented"; return }

func (v value) replaceChild(new, old Wrapper) { _ = "STUB: not implemented"; return }

func (v value) removeChild(c Wrapper) { _ = "STUB: not implemented"; return }

func (v value) firstElementChild() Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) addEventListener(event string, fn Func, options map[string]any) {
	_ = "STUB: not implemented"
	return
}

func (v value) removeEventListener(event string, fn Func) { _ = "STUB: not implemented"; return }

func (v value) setNodeValue(val string) { _ = "STUB: not implemented"; return }

func (v value) setInnerHTML(val string) { _ = "STUB: not implemented"; return }

func (v value) setInnerText(val string) { _ = "STUB: not implemented"; return }

func null() Value { _ = "STUB: not implemented"; return *new(Value) }

func undefined() Value { _ = "STUB: not implemented"; return *new(Value) }

func valueOf(x any) Value { _ = "STUB: not implemented"; return *new(Value) }

type function struct {
	value
}

func (f function) Release() { _ = "STUB: not implemented"; return }

func funcOf(fn func(this Value, args []Value) any) Func {
	_ = "STUB: not implemented"
	return *new(Func)
}

type browserWindow struct {
	value
}

func newBrowserWindow() *browserWindow { _ = "STUB: not implemented"; return nil }

func (w browserWindow) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (w browserWindow) Size() (width, height int) { _ = "STUB: not implemented"; return 0, 0 }

func (w browserWindow) CursorPosition() (x, y int) { _ = "STUB: not implemented"; return 0, 0 }

func (w browserWindow) setCursorPosition(x, y int) { _ = "STUB: not implemented"; return }

func (w *browserWindow) GetElementByID(id string) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (w *browserWindow) ScrollToID(id string) { _ = "STUB: not implemented"; return }

func (w *browserWindow) setBody(body UI) { _ = "STUB: not implemented"; return }

func (w *browserWindow) createElement(tag, xmlns string) (Value, error) {
	_ = "STUB: not implemented"
	return *new(Value), nil
}

func (w *browserWindow) createTextNode(v string) Value {
	_ = "STUB: not implemented"
	return *new(Value)
}

func (w *browserWindow) addHistory(u *url.URL) { _ = "STUB: not implemented"; return }

func (w *browserWindow) replaceHistory(u *url.URL) { _ = "STUB: not implemented"; return }

func copyBytesToGo(dst []byte, src Value) int { _ = "STUB: not implemented"; return 0 }

func copyBytesToJS(dst Value, src []byte) int { _ = "STUB: not implemented"; return 0 }
