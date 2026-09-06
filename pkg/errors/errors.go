package errors

func init() {
	SetInlineEncoder()
}

func Unwrap(err error) error { _ = "STUB: not implemented"; return nil }

func Is(err, target error) bool { _ = "STUB: not implemented"; return false }

func As(err error, target any) bool { _ = "STUB: not implemented"; return false }

func Type(err error) string { _ = "STUB: not implemented"; return "" }

func HasType(err error, v string) bool { _ = "STUB: not implemented"; return false }

func UIError(err error) string { _ = "STUB: not implemented"; return "" }

func Tag(err error, k string) any { _ = "STUB: not implemented"; return *new(any) }

type Error struct {
	Line        string
	Message     string
	DefinedType string
	UIMessage   string
	Tags        map[string]any
	WrappedErr  error
}

func New(msg string) Error { _ = "STUB: not implemented"; return *new(Error) }

func Newf(msgFormat string, v ...any) Error { _ = "STUB: not implemented"; return *new(Error) }

func makeError(v string) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) WithType(v string) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) Type() string { _ = "STUB: not implemented"; return "" }

func (e Error) WithTag(k string, v any) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) Tag(k string) any { _ = "STUB: not implemented"; return *new(any) }

func (e Error) WithUIError(msg string) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) UIError() string { _ = "STUB: not implemented"; return "" }

func (e Error) Wrap(err error) Error { _ = "STUB: not implemented"; return *new(Error) }

func (e Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e Error) Error() string { _ = "STUB: not implemented"; return "" }

func (e Error) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e Error) Is(err error) bool { _ = "STUB: not implemented"; return false }

func isSameErr(a, b error) bool { _ = "STUB: not implemented"; return false }

type jsonError struct {
	Line        string         `json:"line,omitempty"`
	Message     string         `json:"message"`
	UIMessage   string         `json:"ui,omitempty"`
	DefinedType string         `json:"type,omitempty"`
	Tags        map[string]any `json:"tags,omitempty"`
	WrappedErr  any            `json:"wrap,omitempty"`
}

func makeJSONError(err Error) jsonError { _ = "STUB: not implemented"; return *new(jsonError) }

func makeJSONWrappedErr(err error) any { _ = "STUB: not implemented"; return *new(any) }

func makeJSONTags(tags map[string]any) map[string]any { _ = "STUB: not implemented"; return nil }
