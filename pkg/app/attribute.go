package app

type attributes map[string]string

func (a attributes) Set(name string, value any) { _ = "STUB: not implemented"; return }

type attributeURLResolver func(string) string

func toAttributeValue(v any) string { _ = "STUB: not implemented"; return "" }

func resolveAttributeURLValue(name, value string, resolve attributeURLResolver) string {
	_ = "STUB: not implemented"
	return ""
}

func setJSAttribute(jsElement Value, name, value string) { _ = "STUB: not implemented"; return }

func deleteJSAttribute(jsElement Value, name string) { _ = "STUB: not implemented"; return }
