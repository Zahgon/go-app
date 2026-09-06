package errors

import (
	"sync/atomic"
)

type encoderFunc func(any) ([]byte, error)

var encoder atomic.Value

func init() {
	SetInlineEncoder()
}

func SetEncoder(fn func(any) ([]byte, error)) { _ = "STUB: not implemented"; return }

func SetInlineEncoder() { _ = "STUB: not implemented"; return }

func SetIndentEncoder() { _ = "STUB: not implemented"; return }

func getEncoder() encoderFunc { _ = "STUB: not implemented"; return *new(encoderFunc) }
