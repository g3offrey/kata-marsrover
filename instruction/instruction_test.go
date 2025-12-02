package instruction

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandsFromString(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedOutput []Command
	}{
		{
			name:           "Empty string",
			input:          "",
			expectedOutput: []Command{},
		},
		{
			name:           "Single Move Forward command",
			input:          "M",
			expectedOutput: []Command{MoveForward},
		},
		{
			name:           "Multiple commands",
			input:          "MRL",
			expectedOutput: []Command{MoveForward, TurnRight, TurnLeft},
		},
		{
			name:           "Invalid characters are ignored",
			input:          "MXRLYZ",
			expectedOutput: []Command{MoveForward, TurnRight, TurnLeft},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CommandsFromString(tt.input)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

type FakeMover struct {
	MoveForwardCalled int
	TurnRightCalled   int
	TurnLeftCalled    int
}

func (f *FakeMover) MoveForward() {
	f.MoveForwardCalled++
}

func (f *FakeMover) TurnRight() {
	f.TurnRightCalled++
}

func (f *FakeMover) TurnLeft() {
	f.TurnLeftCalled++
}

func TestInstructionRunner_Execute(t *testing.T) {
	tests := []struct {
		name              string
		commands          []Command
		expectedMoveForwd int
		expectedTurnRight int
		expectedTurnLeft  int
	}{
		{
			name:              "No commands",
			commands:          []Command{},
			expectedMoveForwd: 0,
			expectedTurnRight: 0,
			expectedTurnLeft:  0,
		},
		{
			name:              "Move Forward command",
			commands:          []Command{MoveForward, MoveForward},
			expectedMoveForwd: 2,
			expectedTurnRight: 0,
			expectedTurnLeft:  0,
		},
		{
			name:              "Turn Right command",
			commands:          []Command{TurnRight, TurnRight},
			expectedMoveForwd: 0,
			expectedTurnRight: 2,
			expectedTurnLeft:  0,
		},
		{
			name:              "Turn Left command",
			commands:          []Command{TurnLeft, TurnLeft},
			expectedMoveForwd: 0,
			expectedTurnRight: 0,
			expectedTurnLeft:  2,
		},
		{
			name:              "Mixed commands",
			commands:          []Command{MoveForward, TurnRight, MoveForward, TurnLeft, MoveForward},
			expectedMoveForwd: 3,
			expectedTurnRight: 1,
			expectedTurnLeft:  1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fakeMover := &FakeMover{}
			instructions := InstructionRunner{mover: fakeMover}

			instructions.Execute(tt.commands)

			assert.Equal(t, tt.expectedMoveForwd, fakeMover.MoveForwardCalled)
			assert.Equal(t, tt.expectedTurnRight, fakeMover.TurnRightCalled)
			assert.Equal(t, tt.expectedTurnLeft, fakeMover.TurnLeftCalled)
		})
	}
}
