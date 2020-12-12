package main

import (
	"time"

	"github.com/keaising/auto-mouse-keyboard/device"
)

func main() {
	//tap("m", "cmd")
	time.Sleep(2 * time.Second)
	device.Click("right", false)
	time.Sleep(2 * time.Second)
	device.Click("left", false)
	device.Input("fasjdfaks")
	device.Tap("m", "cmd")
}
