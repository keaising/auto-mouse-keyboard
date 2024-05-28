package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/keaising/auto-mouse-keyboard/command"
	"github.com/keaising/auto-mouse-keyboard/model"
)

func main() {
	if len(os.Args) > 1 {
		log.Println("specific file", os.Args[1])
		readFileAndRun(os.Args[1])
		return
	}
	file, err := os.Open(".")
	if err != nil {
		log.Fatalln("fail opening directory")
	}
	list, err := file.Readdirnames(0)
	if err != nil {
		log.Fatalln("fail read directory")
	}
	for _, name := range list {
		if strings.HasSuffix(name, ".conf") {
			readFileAndRun(name)
		}
	}
}

func readFileAndRun(fileName string) {
	file, err := os.Open(fileName)
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

	common, err := getCommon(sources)
	if err != nil {
		log.Println("Get common value error, please check and retry")
		return
	}
	commands, err := command.ParseCommands(sources)
	if err != nil {
		log.Println("ERROR! correct amk.conf and run again")
		return
	}
	_ = command.RunCommand(commands, common)

}

var commonItems = []model.CommonItem{
	{
		Name: "SHIM",
		Type: model.CommonItemTypeInt,
	},
	{
		Name: "SCALE",
		Type: model.CommonItemTypeFloat64,
	},
}

func getCommon(sources []string) (*model.Common, error) {
	var (
		common model.Common
		kv     = make(map[string]interface{})
	)

	for _, source := range sources {
		source = strings.TrimSpace(source)
		if source == "" || strings.HasPrefix(source, "#") {
			continue
		}
		for _, item := range commonItems {
			if strings.HasPrefix(source, item.Name) {
				if len(source) < len(item.Name)+1 {
					log.Println("no value of", item.Name)
					return nil, fmt.Errorf("no value of %s", item.Name)
				}
				v, err := getValueOfConfig(source, item)
				if err != nil {
					return nil, err
				}
				kv[item.Name] = v
			}
		}
		// Convert map to struct
		data, err := json.Marshal(kv)
		if err != nil {
			log.Println("marshal kv failed", err)
			return nil, fmt.Errorf("marshal kv failed %v", err)
		}
		err = json.Unmarshal(data, &common)
		if err != nil {
			log.Println("unmarshal kv failed", err)
			return nil, fmt.Errorf("unmarshal kv failed %v", err)
		}
	}
	return &common, nil
}

// Convert config into config
func getValueOfConfig(source string, item model.CommonItem) (interface{}, error) {
	rawValue := source[(len(item.Name) + 1):]
	switch item.Type {
	case model.CommonItemTypeInt:
		{
			v, err := strconv.Atoi(rawValue)
			if err != nil {
				log.Println(item.Name, "value not int", err)
				return nil, fmt.Errorf("%s value not int %v", item.Name, err)
			}
			return v, nil
		}
	case model.CommonItemTypeFloat64:
		{
			v, err := strconv.ParseFloat(rawValue, 64)
			if err != nil {
				log.Println(item.Name, "value not float64", err)
				return nil, fmt.Errorf("%s value not float64 %v", item.Name, err)
			}
			return v, nil
		}
	case model.CommonItemTypeString:
		fallthrough
	default:
		return rawValue, nil
	}
}
