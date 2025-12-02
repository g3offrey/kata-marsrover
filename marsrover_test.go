package main

import (
	"testing"

	"github.com/g3offrey/mars-rover/rover"
	"github.com/stretchr/testify/assert"
)

func TestMarsRover_Run(t *testing.T) {
	tests := []struct {
		name             string
		commands         string
		expectedPosition rover.Position
	}{
		{
			name:     "Move forward twice",
			commands: "MM",
			expectedPosition: rover.Position{
				X:         0,
				Y:         2,
				Direction: rover.North,
			},
		},
		{
			name:     "Move forward, turn right, move forward",
			commands: "MRM",
			expectedPosition: rover.Position{
				X:         1,
				Y:         1,
				Direction: rover.East,
			},
		},
		{
			name:     "Full rotation and move",
			commands: "RRRRM",
			expectedPosition: rover.Position{
				X:         0,
				Y:         1,
				Direction: rover.North,
			},
		},
		{
			name:     "Complex movement",
			commands: "MMLMRM",
			expectedPosition: rover.Position{
				X:         5,
				Y:         3,
				Direction: rover.North,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			marsRover := &MarsRover{}

			position := marsRover.Run(tt.commands, 5, 5)

			assert.Equal(t, tt.expectedPosition, position)
		})
	}
}
