package app

import (
	"net/url"
	"syscall/js"
)

type value struct {
	jsValue
}

func (v value) Call(m string, args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Delete(p string) { _ = "STUB: not implemented"; return }

func (v value) Equal(w Value) bool { _ = "STUB: not implemented"; return false }

func (v value) Get(p string) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Set(p string, x any) { _ = "STUB: not implemented"; return }

func (v value) Index(i int) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) InstanceOf(t Value) bool { _ = "STUB: not implemented"; return false }

func (v value) Invoke(args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) JSValue() Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) New(args ...any) Value { _ = "STUB: not implemented"; return *new(Value) }

func (v value) Type() Type { _ = "STUB: not implemented"; return *new(Type) }

func (v value) Release() { _ = "STUB: not implemented"; return }

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

func valueOf(v any) Value { _ = "STUB: not implemented"; return *new(Value) }

func funcOf(function func(this Value, args []Value) any) Func {
	_ = "STUB: not implemented"
	return *new(Func)
}

type jsValue interface {
	Bool() bool
	Call(string, ...any) js.Value
	Delete(string)
	Equal(js.Value) bool
	Float() float64
	Get(string) js.Value
	Index(int) js.Value
	InstanceOf(js.Value) bool
	Int() int
	Invoke(...any) js.Value
	IsNaN() bool
	IsNull() bool
	IsUndefined() bool
	Length() int
	New(...any) js.Value
	Set(string, any)
	SetIndex(int, any)
	String() string
	Truthy() bool
	Type() js.Type
}

type jsFunc interface {
	jsValue
	Release()
}

type jsError interface {
	jsValue
	Error() string
}

func syscallJSArgs(v []any) []any { _ = "STUB: not implemented"; return nil }

func syscalJSValueOf(v any) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }

type browserWindow struct {
	value

	body    UI
	cursorX int
	cursorY int
}

func newBrowserWindow() *browserWindow { _ = "STUB: not implemented"; return nil }

func (w *browserWindow) URL() *url.URL { _ = "STUB: not implemented"; return nil }

func (w *browserWindow) Size() (width int, height int) { _ = "STUB: not implemented"; return 0, 0 }

func (w *browserWindow) CursorPosition() (x, y int) { _ = "STUB: not implemented"; return 0, 0 }

func (w *browserWindow) setCursorPosition(x, y int) { _ = "STUB: not implemented"; return }

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
