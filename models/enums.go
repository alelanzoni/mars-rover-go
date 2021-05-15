package models

type Direction string

const (
	DirectionNorth Direction = "N"
	DirectionSouth Direction = "S"
	DirectionEast  Direction = "E"
	DirectionWest  Direction = "W"
)

type Command string

const (
	CommandForward  Command = "f"
	CommandBackward Command = "b"
	CommandLeft     Command = "l"
	CommandRight    Command = "r"
)

type MoveType int

const (
	MoveTypeForward  MoveType = 1
	MoveTypeBackward MoveType = -1
)
