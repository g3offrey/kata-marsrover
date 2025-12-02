package main

import (
	"flag"
	"fmt"
)

func main() {
	marsWidth := flag.Int("width", 5, "-width=5")
	marsHeight := flag.Int("height", 5, "-height=5")
	commandsString := flag.String("commands", "", `-commands="MMRMM"`)
	flag.Parse()

	marsRover := &MarsRover{}
	position := marsRover.Run(*commandsString, *marsWidth, *marsHeight)

	fmt.Println("Rover final position:", position.X, position.Y, string(position.Direction))
}
