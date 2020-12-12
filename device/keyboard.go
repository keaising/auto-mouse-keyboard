package device

import "github.com/go-vgo/robotgo"

func Input(str string) {
	robotgo.TypeStr(str)
}

func Tap(keys ...string) {
	robotgo.KeyTap(keys[0], keys[1:])
}
