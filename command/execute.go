package command

import (
	"fmt"
	"log"
	"time"

	"github.com/keaising/auto-mouse-keyboard/device"
	"github.com/keaising/auto-mouse-keyboard/model"
)

func ExecuteCommand(cmd *model.Command, com *model.Common) error {
	log.Println(cmd.Type, cmd.Args)
	switch cmd.Type {
	case model.CommandTypeMove:
		{
			args, ok := cmd.Args.(model.MoveArgs)
			if !ok {
				log.Println("wrong move args", cmd)
				return fmt.Errorf("wrong move args %v", cmd)
			}
			device.Move(args.X, args.Y, com.Scale)
		}

	case model.CommandTypeClick:
		{
			args, ok := cmd.Args.(model.ClickArgs)
			if !ok {
				log.Println("wrong click args", cmd)
				return fmt.Errorf("wrong click args %v", cmd)
			}
			device.Click(args.Type, args.Double)
		}
	case model.CommandTypeInput:
		{
			args, ok := cmd.Args.(model.InputArgs)
			if !ok {
				log.Println("wrong input args", cmd)
				return fmt.Errorf("wrong input args %v", cmd)
			}
			device.Input(args.Content)
		}
	case model.CommandTypeTap:
		{
			args, ok := cmd.Args.(model.TapArgs)
			if !ok {
				log.Println("wrong tap args", cmd)
				return fmt.Errorf("wrong tap args %v", cmd)
			}
			if len(args.CombineKeys) != 0 {
				device.Tap(args.CombineKeys...)
			}
			if len(args.RepeatKeys) != 0 {
				for _, key := range args.RepeatKeys {
					device.Tap(key)
				}
			}
		}
	case model.CommandTypeSleep:
		{
			args, ok := cmd.Args.(model.SleepArgs)
			if !ok {
				log.Println("wrong sleep args", cmd)
				return fmt.Errorf("wrong sleep args %v", cmd)
			}
			time.Sleep(time.Duration(args.Duration) * time.Second)
		}
	}
	return nil
}
