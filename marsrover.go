package main

import (
	"github.com/g3offrey/mars-rover/instruction"
	"github.com/g3offrey/mars-rover/rover"
)

type MarsRover struct{}

func (mr *MarsRover) Run(commands string, width, height int) rover.Position {
	mars := rover.NewMars(width, height)
	rover := rover.New(mars)
	instructionRunner := instruction.NewRunner(rover)

	cmds := instruction.CommandsFromString(commands)
	instructionRunner.Execute(cmds)

	return rover.Position()
}
