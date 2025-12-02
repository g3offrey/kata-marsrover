package main

import (
	"github.com/g3offrey/mars-rover/instruction"
	"github.com/g3offrey/mars-rover/rover"
)

type MarsRover struct {
	mars *rover.Mars
}

func NewMarsRover(width int, height int) *MarsRover {
	return &MarsRover{mars: rover.NewMars(width, height)}
}

func (mr *MarsRover) Run(commands string) rover.Position {
	rover := rover.New(mr.mars)
	instructionRunner := instruction.NewRunner(rover)

	cmds := instruction.CommandsFromString(commands)
	instructionRunner.Execute(cmds)

	return rover.Position()
}
