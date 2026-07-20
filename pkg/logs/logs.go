package logs

var (
	Encoder func(any) ([]byte, error)
)

func init() {
	SetInlineEncoder()
}

func SetInlineEncoder() { _ = "STUB: not implemented"; return }

func SetIndentEncoder() { _ = "STUB: not implemented"; return }

type Entry struct {
	Line    string         `json:"line,omitempty"`
	Message string         `json:"message"`
	Tags    map[string]any `json:"tags,omitempty"`
}

func New(v string) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func Newf(msgFormat string, v ...any) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func makeEntry(v string) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (e Entry) WithTag(k string, v any) Entry { _ = "STUB: not implemented"; return *new(Entry) }

func (e Entry) String() string { _ = "STUB: not implemented"; return "" }
