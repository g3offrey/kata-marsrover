package rover

type Mars struct {
	Width  int
	Height int
}

func NewMars(width, height int) *Mars {
	return &Mars{Width: width, Height: height}
}
