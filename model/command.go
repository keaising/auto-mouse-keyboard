package model

type Command struct {
	Type CommandType
	Line int
	Args interface{}
}

type CommandType string

const (
	CommandTypeMove  CommandType = "move"
	CommandTypeClick CommandType = "click"
	CommandTypeInput CommandType = "input"
	CommandTypeTap   CommandType = "tap"
	CommandTypeSleep CommandType = "sleep"
)

type MoveArgs struct {
	X int
	Y int
}

const (
	ClickTypeLeft  = "left"
	ClickTypeRight = "right"
)

type ClickArgs struct {
	Type   string
	Double bool
}

type InputArgs struct {
	Content string
}

type TapArgs struct {
	CombineKeys []string
	RepeatKeys  []string
}

type Sleep struct {
	Duration int
}
