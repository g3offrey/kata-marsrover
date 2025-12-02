package rover

type Position struct {
	X         int
	Y         int
	Direction Direction
}

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

type Rover struct {
	mars      *Mars
	x         int
	y         int
	direction Direction
}

func (r *Rover) MoveForward() {
	nextY := r.y
	nextX := r.x

	switch r.direction {
	case North:
		nextY++
	case East:
		nextX++
	case South:
		nextY--
	case West:
		nextX--
	}

	if nextX > r.mars.Width {
		nextX = 0
	} else if nextX < 0 {
		nextX = r.mars.Width
	}

	if nextY > r.mars.Height {
		nextY = 0
	} else if nextY < 0 {
		nextY = r.mars.Height
	}

	r.x = nextX
	r.y = nextY
}

func (r *Rover) TurnRight() {
	switch r.direction {
	case North:
		r.direction = East
	case East:
		r.direction = South
	case South:
		r.direction = West
	case West:
		r.direction = North
	}
}

func (r *Rover) TurnLeft() {
	switch r.direction {
	case North:
		r.direction = West
	case West:
		r.direction = South
	case South:
		r.direction = East
	case East:
		r.direction = North
	}
}

func (r *Rover) Position() Position {
	return Position{
		X:         r.x,
		Y:         r.y,
		Direction: r.direction,
	}
}

func New(mars *Mars) *Rover {
	return &Rover{mars: mars, x: 0, y: 0, direction: North}
}
