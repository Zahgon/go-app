package analytics

type Backend interface {
	Identify(userID string, traits map[string]interface{})

	Track(event string, properties map[string]interface{})

	Page(name string, properties map[string]interface{})
}

func Identify(userID string, traits map[string]interface{}) { _ = "STUB: not implemented"; return }

func Track(event string, properties map[string]interface{}) { _ = "STUB: not implemented"; return }

func Page(name string, properties map[string]interface{}) { _ = "STUB: not implemented"; return }

func Add(b Backend) { _ = "STUB: not implemented"; return }

var (
	backends []Backend
)

func sanitizeValue(v interface{}) interface{} { _ = "STUB: not implemented"; return nil }
