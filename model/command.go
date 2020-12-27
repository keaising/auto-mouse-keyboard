package model

type Command struct {
	Type CommandType
	Line int
	Args interface{}
}

type CommandType string

const (
	CommandTypeMove             CommandType = "move"
	CommandTypeClick            CommandType = "click"
	CommandTypeInput            CommandType = "input"
	CommandTypeTap              CommandType = "tap"
	CommandTypeSleep            CommandType = "sleep"
	CommandTypeMillisecondSleep CommandType = "millisecond_sleep"
	CommandTypeLoop             CommandType = "loop"
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

type SleepArgs struct {
	Duration int
}

type LoopArgs struct {
	// 循环次数
	Times int
	// 每次循环的全部命令
	Commands []*Command
}
