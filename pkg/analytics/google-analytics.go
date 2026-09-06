package analytics

func GoogleAnalyticsHeader(propertyID string) string { _ = "STUB: not implemented"; return "" }

func NewGoogleAnalytics() Backend { _ = "STUB: not implemented"; return *new(Backend) }

type googleAnalytics struct {
}

func (a googleAnalytics) Identify(userID string, traits map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a googleAnalytics) Track(event string, properties map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a googleAnalytics) Page(name string, properties map[string]interface{}) {
	_ = "STUB: not implemented"
	return
}

func (a googleAnalytics) gtag(args ...interface{}) { _ = "STUB: not implemented"; return }
