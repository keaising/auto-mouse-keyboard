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
		err := robotgo.KeyTap(keys[0])
		if err != nil {
			log.Println("tap result:", err)
		}
		return
	}
	err := robotgo.KeyTap(keys[0], keys[1:])
	if err != nil {
		log.Println("tap result:", err)
	}
}
