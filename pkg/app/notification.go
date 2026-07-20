package app

type Notification struct {
	Title string `json:"title"`

	Path string `json:"path"`

	Lang string `json:"lang,omitempty"`

	Badge string `json:"badge,omitempty"`

	Body string `json:"body,omitempty"`

	Tag string `json:"tag,omitempty"`

	Icon string `json:"icon,omitempty"`

	Image string `json:"image,omitempty"`

	Data map[string]any `json:"data"`

	Renotify bool `json:"renotify,omitempty"`

	RequireInteraction bool `json:"requireInteraction,omitempty"`

	Silent bool `json:"silent,omitempty"`

	Vibrate []int `json:"vibrate,omitempty"`

	Actions []NotificationAction `json:"actions,omitempty"`
}

type NotificationAction struct {
	Action string `json:"action"`

	Title string `json:"title"`

	Icon string `json:"icon,omitempty"`

	Path string `json:"path"`
}

type NotificationSubscription struct {
	Endpoint string `json:"endpoint"`

	Keys struct {
		Auth string `json:"auth"`

		P256dh string `json:"p256dh"`
	} `json:"keys"`
}

type NotificationPermission string

const (
	NotificationDefault NotificationPermission = "default"

	NotificationGranted NotificationPermission = "granted"

	NotificationDenied NotificationPermission = "denied"

	NotificationNotSupported NotificationPermission = "unsupported"
)

type NotificationService struct{}

func (s NotificationService) Permission() NotificationPermission {
	_ = "STUB: not implemented"
	return *new(NotificationPermission)
}

func (s NotificationService) RequestPermission() NotificationPermission {
	_ = "STUB: not implemented"
	return *new(NotificationPermission)
}

func (s NotificationService) New(n Notification) { _ = "STUB: not implemented"; return }

func (s NotificationService) Subscribe(vapIDPublicKey string) (NotificationSubscription, error) {
	_ = "STUB: not implemented"
	return *new(NotificationSubscription), nil
}
