package command

import (
	"fmt"
	"log"
	"time"

	"github.com/keaising/auto-mouse-keyboard/device"
	"github.com/keaising/auto-mouse-keyboard/model"
)

func RunCommand(cmds []*model.Command, com *model.Common) error {
	var toBeRunCommands []*model.Command
	for i := 0; i < len(cmds); i++ {
		if cmds[i].Type != model.CommandTypeLoop {
			toBeRunCommands = append(toBeRunCommands, cmds[i])
			continue
		}
		cmdArgs, ok := cmds[i].Args.(model.LoopArgs)
		if !ok {
			log.Println("wrong loop args", cmds[i])
			return fmt.Errorf("wrong loop args %v", cmds[i])
		}
		var loopCommands []*model.Command
		for j := i + 1; j < len(cmds); j++ {
			if cmds[j].Type != model.CommandTypeLoop {
				loopCommands = append(loopCommands, cmds[j])
			} else {
				i = j
				break
			}
		}
		cmdArgs.Commands = loopCommands
		cmds[i].Args = cmdArgs
		toBeRunCommands = append(toBeRunCommands, cmds[i])
	}
	for _, cmd := range toBeRunCommands {
		if err := ExecuteCommand(cmd, com); err != nil {
			log.Println("Execute command error!!! Please check and retry")
			return err
		}
		time.Sleep(time.Duration(com.Shim) * time.Millisecond)
	}
	return nil
}

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
			time.Sleep(time.Duration(args.Duration) * time.Millisecond)
		}
	case model.CommandTypeLoop:
		{
			args, ok := cmd.Args.(model.LoopArgs)
			if !ok {
				log.Println("wrong loop args", cmd)
				return fmt.Errorf("wrong loop args %v", cmd)
			}
			for i := 0; i < args.Times; i++ {
				for j, c := range args.Commands {
					if err := ExecuteCommand(c, com); err != nil {
						log.Println("execute loop command", i, j, c)
						return fmt.Errorf("execute loop command %d %d %v", i, j, c)
					}
					time.Sleep(time.Duration(com.Shim) * time.Millisecond)
				}
			}
		}
	}
	return nil
}
