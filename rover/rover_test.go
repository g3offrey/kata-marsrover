package rover

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRover_MoveForward(t *testing.T) {
	mars := NewMars(5, 5)

	tests := []struct {
		name             string
		numberOfMoves    int
		expectedPosition Position
	}{
		{
			name:             "Initial position at origin facing North",
			numberOfMoves:    0,
			expectedPosition: Position{0, 0, North},
		},
		{
			name:             "Move forward 1 step facing North",
			numberOfMoves:    1,
			expectedPosition: Position{0, 1, North},
		},
		{
			name:             "Move forward 3 steps facing North",
			numberOfMoves:    3,
			expectedPosition: Position{0, 3, North},
		},
		{
			name:             "Mars is round, so I should be able to come back to origin too far",
			numberOfMoves:    mars.Width + 1,
			expectedPosition: Position{0, 0, North},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := New(mars)

			for i := 0; i < tt.numberOfMoves; i++ {
				rover.MoveForward()
			}

			assert.Equal(t, tt.expectedPosition, rover.Position())
		})
	}
}

func TestRover_TurnRight(t *testing.T) {
	mars := NewMars(5, 5)

	tests := []struct {
		name             string
		numberOfTurns    int
		expectedPosition Position
	}{
		{
			name:             "Single turn right from North to East",
			numberOfTurns:    1,
			expectedPosition: Position{0, 0, East},
		},
		{
			name:             "Two turns right from North to South",
			numberOfTurns:    2,
			expectedPosition: Position{0, 0, South},
		},
		{
			name:             "Three turns right from North to West",
			numberOfTurns:    3,
			expectedPosition: Position{0, 0, West},
		},
		{
			name:             "Four turns right from North back to North",
			numberOfTurns:    4,
			expectedPosition: Position{0, 0, North},
		},
		{
			name:             "Five turns right from North to East",
			numberOfTurns:    5,
			expectedPosition: Position{0, 0, East},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := New(mars)

			for i := 0; i < tt.numberOfTurns; i++ {
				rover.TurnRight()
			}

			assert.Equal(t, tt.expectedPosition, rover.Position())
		})
	}
}

func TestRover_TurnLeft(t *testing.T) {
	mars := NewMars(5, 5)

	tests := []struct {
		name             string
		numberOfTurns    int
		expectedPosition Position
	}{
		{
			name:             "Single turn left from North to West",
			numberOfTurns:    1,
			expectedPosition: Position{0, 0, West},
		},
		{
			name:             "Two turns left from North to South",
			numberOfTurns:    2,
			expectedPosition: Position{0, 0, South},
		},
		{
			name:             "Three turns left from North to East",
			numberOfTurns:    3,
			expectedPosition: Position{0, 0, East},
		},
		{
			name:             "Four turns left from North back to North",
			numberOfTurns:    4,
			expectedPosition: Position{0, 0, North},
		},
		{
			name:             "Five turns left from North to West",
			numberOfTurns:    5,
			expectedPosition: Position{0, 0, West},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rover := New(mars)

			for i := 0; i < tt.numberOfTurns; i++ {
				rover.TurnLeft()
			}

			assert.Equal(t, tt.expectedPosition, rover.Position())
		})
	}
}

func TestRover_MovingRight(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.TurnRight()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{2, 0, East}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingLeft(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.TurnLeft()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{4, 0, West}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingBottom(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.TurnLeft()
	rover.TurnLeft()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{0, 4, South}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingTopRight(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.MoveForward()
	rover.MoveForward()
	rover.TurnRight()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{2, 2, East}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingBottomRight(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.TurnRight()
	rover.MoveForward()
	rover.MoveForward()
	rover.TurnRight()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{2, 4, South}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingBottomLeft(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.TurnLeft()
	rover.MoveForward()
	rover.MoveForward()
	rover.TurnLeft()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{4, 4, South}
	assert.Equal(t, expectedPosition, rover.Position())
}

func TestRover_MovingTopLeft(t *testing.T) {
	mars := NewMars(5, 5)
	rover := New(mars)

	rover.MoveForward()
	rover.MoveForward()
	rover.TurnLeft()
	rover.MoveForward()
	rover.MoveForward()

	expectedPosition := Position{4, 2, West}
	assert.Equal(t, expectedPosition, rover.Position())
}
