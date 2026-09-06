package app

type TwitterCard struct {
	Card string

	Site string

	Creator string

	Title string

	Description string

	Image string

	ImageAlt string
}

func (c TwitterCard) toMap() map[string]string { _ = "STUB: not implemented"; return nil }
