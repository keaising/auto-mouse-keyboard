package device

import (
	"github.com/go-vgo/robotgo"
	"github.com/keaising/auto-mouse-keyboard/model"
)

const mouseMoveArg = 0.5

func Move(x, y int) {
	robotgo.MoveMouseSmooth(x, y, mouseMoveArg, mouseMoveArg)
}

func Click(clickType string, isDouble bool) {
	if clickType != model.ClickTypeLeft && clickType != model.ClickTypeRight {
		clickType = "left"
	}
	robotgo.Click(clickType, isDouble)
}
