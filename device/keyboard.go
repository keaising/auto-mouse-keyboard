package device

import (
	"log"

	"github.com/go-vgo/robotgo"
)

func Input(str string) {
	robotgo.TypeStr(str)
}

func Tap(keys ...string) {
	if len(keys) == 0 {
		return
	}
	if len(keys) == 1 {
		result := robotgo.KeyTap(keys[0])
		if len(result) != 0 {
			log.Println("tap result:", result)
		}
		return
	}
	result := robotgo.KeyTap(keys[0], keys[1:])
	if len(result) != 0 {
		log.Println("tap result:", result)
	}
}
