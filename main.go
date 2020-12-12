package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/keaising/auto-mouse-keyboard/command"
	"github.com/keaising/auto-mouse-keyboard/model"
)

//tap("m", "cmd")
//time.Sleep(2 * time.Second)
//device.Click("right", false)
//time.Sleep(2 * time.Second)
//device.Click("left", false)
//device.Input("fasjdfaks")
//device.Tap("m", "cmd")

func main() {
	file, err := os.Open("amk.conf")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sources []string
	for scanner.Scan() {
		sources = append(sources, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sources, common, err := getCommon(sources)
	if err != nil {
		log.Println("Get common value error, please check and retry")
		return
	}
	commands, err := command.ParseCommands(sources)
	if err != nil {
		log.Println("ERROR! correct amk.conf and run again")
		return
	}
	for _, cmd := range commands {
		if err = command.ExecuteCommand(cmd); err != nil {
			log.Println("Execute command error!!! Please check and retry")
			break
		}
		time.Sleep(time.Duration(common.Shim) * time.Millisecond)
	}
}

func getCommon(sources []string) ([]string, *model.Common, error) {
	var (
		filtered []string
		common   model.Common
	)

	for _, source := range sources {
		source = strings.TrimSpace(source)
		if source == "" || strings.HasPrefix(source, "#") {
			continue
		}
		if strings.HasPrefix(source, "SHIM") {
			if len(source) < 3 {
				log.Println("no shim value")
				return nil, nil, fmt.Errorf("no shim value")
			}
			shim, err := strconv.Atoi(source[5:])
			if err != nil {
				log.Println("shim value not int", err)
				return nil, nil, fmt.Errorf("shim value not int %v", err)
			}
			common.Shim = shim
			continue
		}
		filtered = append(filtered, source)
	}
	return filtered, &common, nil
}
