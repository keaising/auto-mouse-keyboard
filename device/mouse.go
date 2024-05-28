package device

import (
	"github.com/go-vgo/robotgo"
	"github.com/keaising/auto-mouse-keyboard/model"
)

// default mouse speed
const mouseMoveArg = 0.5

func Move(x, y int, scale float64) {
	x = int(float64(x) / scale)
	y = int(float64(y) / scale)
	robotgo.MoveSmooth(x, y, mouseMoveArg, mouseMoveArg)
}

func Click(clickType string, isDouble bool) {
	if clickType != model.ClickTypeLeft && clickType != model.ClickTypeRight {
		clickType = "left"
	}
	robotgo.Click(clickType, isDouble)
}
