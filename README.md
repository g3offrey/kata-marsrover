# Mars Rover (Go kata)

A small implementation of the classic "Mars Rover" kata written in Go.

This repo contains an implementation of a rover that can move around a rectangular map of Mars and respond to simple commands:

- `M` - Move forward one grid square in the direction the rover is currently facing
- `R` - Turn right
- `L` - Turn left

## Key features

- Start position: (0, 0) facing North
- Inclusive map coordinates: `width` and `height` are _the maximum coordinates_; the grid is 0..width and 0..height
- Global wrap-around: moving across the edge wraps to the other side (like a torus/planets are round)

## Project structure

- `main.go` — Command-line runner for quick experiments
- `marsrover.go` — Entry point to run a sequence of commands on the rover
- `rover/` — Contains the `Rover` implementation and the `Mars` map
  - `rover.go` — The rover implementation
  - `mars.go` — Mars size and helper
  - `rover_test.go` — Unit tests for `Rover` behavior
- `instruction/` — Command parsing and runner
  - `instruction.go` — Command types and Runner (executes commands over a Mover interface)
  - `instruction_test.go` — Unit tests for instructions
- `go.mod` / `go.sum` — Go module definition and deps

## Requirements

- Go >= 1.25 (see `go.mod`)
- `github.com/stretchr/testify` used in tests

## Build & Run

Build the binary:

```bash
go build -o mars-rover ./
```

Run from the source (quick):

```bash
go run main.go -commands="MMRMM" -width=5 -height=5
```

Flags (defaults):

- `-commands` — string of commands (e.g., `"MMRMM"`)
- `-width` — maximum X coordinate (default `5`, inclusive)
- `-height` — maximum Y coordinate (default `5`, inclusive)

## Example

1. Move forward twice:

```bash
go run main.go -commands="MM"
# Output:
# Rover final position: 0 2 N
```

2. Move forward, turn right, move forward:

```bash
go run main.go -commands="MRM"
# Output:
# Rover final position: 1 1 E
```

## API / Library usage

You can use the `MarsRover` type programmatically (see `marsrover.go`):

```go
marsRover := &MarsRover{}
pos := marsRover.Run("MRM", 5, 5) // pos: rover.Position
fmt.Println(pos.X, pos.Y, string(pos.Direction))
```

## Testing

Run all tests:

```bash
go test ./...
```

Run a single package tests:

```bash
go test ./rover
go test ./instruction
```

## Design notes

- `rover.Rover` contains the movement rules and wraps around past the edges of a `Mars`.
- `instruction` package defines a `Mover` interface used by `InstructionRunner`, which allows the runner to be tested with a fake mover.
- Coordinates are inclusive — `width` or `height` acts as the maximum coordinate. This means when `width` is `5`, `x` can go from 0 to 5 inclusive.

## Contributing

Contributions are welcome. Suggested improvements:

- Add more property-based tests (random command sequences)
- Add multiple rovers/obstacles functionality for more complicated rules
- Add integration tests or CLI options

Running tests locally and opening a PR with short explanation is the usual workflow.

## License

This project is licensed under the MIT License — see the `LICENSE` file for details.

## Acknowledgments

This is a small kata used for TDD and katas practice.

# kata-marsrover
