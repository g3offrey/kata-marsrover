package instruction

type Command string

const (
	MoveForward Command = "M"
	TurnRight   Command = "R"
	TurnLeft    Command = "L"
)

func CommandsFromString(s string) []Command {
	commands := []Command{}
	for _, char := range s {
		charStr := string(char)
		if charStr == "M" || charStr == "R" || charStr == "L" {
			commands = append(commands, Command(charStr))
		}
	}

	return commands
}

type Mover interface {
	MoveForward()
	TurnRight()
	TurnLeft()
}

type InstructionRunner struct {
	mover Mover
}

func (i *InstructionRunner) Execute(commands []Command) {
	for _, command := range commands {
		switch command {
		case MoveForward:
			i.mover.MoveForward()
		case TurnRight:
			i.mover.TurnRight()
		case TurnLeft:
			i.mover.TurnLeft()
		}
	}
}

func NewRunner(mover Mover) *InstructionRunner {
	return &InstructionRunner{mover: mover}
}
