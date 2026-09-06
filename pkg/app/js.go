package app

import (
	"net/url"
)

type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string { _ = "STUB: not implemented"; return "" }

type Wrapper interface {
	JSValue() Value
}

type Value interface {
	Bool() bool

	Call(m string, args ...any) Value

	Delete(p string)

	Equal(w Value) bool

	Float() float64

	Get(p string) Value

	Index(i int) Value

	InstanceOf(t Value) bool

	Int() int

	Invoke(args ...any) Value

	IsNaN() bool

	IsNull() bool

	IsUndefined() bool

	JSValue() Value

	Length() int

	New(args ...any) Value

	Set(p string, x any)

	SetIndex(i int, x any)

	String() string

	Truthy() bool

	Type() Type

	Then(f func(Value))

	getAttr(k string) string
	setAttr(k, v string)
	delAttr(k string)
	firstChild() Value
	appendChild(c Wrapper)
	replaceChild(new, old Wrapper)
	removeChild(c Wrapper)
	firstElementChild() Value
	addEventListener(event string, fn Func, options map[string]any)
	removeEventListener(event string, fn Func)
	setNodeValue(v string)
	setInnerHTML(v string)
	setInnerText(v string)
}

func Null() Value { _ = "STUB: not implemented"; return *new(Value) }

func Undefined() Value { _ = "STUB: not implemented"; return *new(Value) }

func ValueOf(x any) Value { _ = "STUB: not implemented"; return *new(Value) }

type Func interface {
	Value

	Release()
}

func FuncOf(fn func(this Value, args []Value) any) Func {
	_ = "STUB: not implemented"
	return *new(Func)
}

type BrowserWindow interface {
	Value

	URL() *url.URL

	Size() (w, h int)

	CursorPosition() (x, y int)

	setCursorPosition(x, y int)

	GetElementByID(id string) Value

	ScrollToID(id string)

	setBody(body UI)
	createElement(tag, xmlns string) (Value, error)
	createTextNode(v string) Value
	addHistory(u *url.URL)
	replaceHistory(u *url.URL)
}

func CopyBytesToGo(dst []byte, src Value) int { _ = "STUB: not implemented"; return 0 }

func CopyBytesToJS(dst Value, src []byte) int { _ = "STUB: not implemented"; return 0 }
