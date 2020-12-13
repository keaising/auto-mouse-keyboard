package command

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/go-vgo/robotgo"
	"github.com/keaising/auto-mouse-keyboard/model"
)

const errTemplate = "ERROR! reason: %s line: %d source: %s"

var (
	ErrNoEqual            = errors.New("no equal")
	ErrNoConfig           = errors.New("no config")
	ErrNoCommaConfig      = errors.New("no comma config")
	ErrNotInt             = errors.New("not int")
	ErrEmptyKey           = errors.New("empty key")
	ErrNotSupportKey      = errors.New("not support key")
	ErrNotSupportOperator = errors.New("not support operator")
)

func ParseCommands(sources []string) ([]*model.Command, error) {
	var (
		lastErr error
		result  []*model.Command
	)
	for i, source := range sources {
		c, err := ParseCommand(i+1, source)
		if err != nil {
			lastErr = err
		}
		if c != nil {
			result = append(result, c)
		}
	}
	return result, lastErr
}

func ParseCommand(lineNumber int, source string) (*model.Command, error) {
	source = strings.TrimSpace(source)
	if source == "" || strings.HasPrefix(source, "#") || strings.HasPrefix(source, "SHIM") || strings.HasPrefix(source, "SCALE") {
		return nil, nil
	}
	switch source[0] {
	case 'M':
		return parseCommandMove(lineNumber, source)
	case 'C':
		return parseCommandClick(lineNumber, source)
	case 'I':
		return parseCommandInput(lineNumber, source)
	case 'T':
		return parseCommandTap(lineNumber, source)
	case 'S':
		return parseCommandSleep(lineNumber, source)
	default:
		log.Printf(errTemplate, "未知的操作类型", lineNumber, source)
		return nil, ErrNotSupportOperator
	}
}

func parseCommandMove(lineNumber int, source string) (*model.Command, error) {
	if len(source) < 3 || source[1] != '=' {
		log.Printf(errTemplate, "缺少等号", lineNumber, source)
		return nil, ErrNoEqual
	}
	var rawConfig = source[2:]
	if len(rawConfig) == 0 {
		log.Printf(errTemplate, "缺少配置项", lineNumber, source)
		return nil, ErrNoConfig
	}
	var raws = strings.Split(rawConfig, ",")
	if len(raws) != 2 {
		log.Printf(errTemplate, "缺少逗号分割的配置", lineNumber, source)
		return nil, ErrNoCommaConfig
	}
	x, err := strconv.Atoi(raws[0])
	if err != nil {
		log.Printf(errTemplate, "不是整数", lineNumber, source)
		return nil, ErrNotInt
	}
	y, err := strconv.Atoi(raws[1])
	if err != nil {
		log.Printf(errTemplate, "不是整数", lineNumber, source)
		return nil, ErrNotInt
	}
	return &model.Command{
		Type: model.CommandTypeMove,
		Line: lineNumber,
		Args: model.MoveArgs{
			X: x,
			Y: y,
		},
	}, nil
}

func parseCommandClick(lineNumber int, source string) (*model.Command, error) {
	if len(source) == 1 {
		return &model.Command{
			Type: model.CommandTypeClick,
			Line: lineNumber,
			Args: model.ClickArgs{
				Type:   model.ClickTypeLeft,
				Double: false,
			},
		}, nil
	}
	if len(source) < 3 {
		log.Printf(errTemplate, "缺少配置项", lineNumber, source)
		return nil, ErrNoConfig
	}
	var (
		raws     = strings.Split(source[2:], ",")
		isDouble = false
		btn      = model.ClickTypeLeft
	)
	if len(raws) >= 2 {
		if strings.EqualFold(raws[1], "double") {
			isDouble = true
		}
	}
	if strings.EqualFold(raws[0], model.ClickTypeRight) {
		btn = model.ClickTypeRight
	}
	return &model.Command{
		Type: model.CommandTypeClick,
		Args: model.ClickArgs{
			Type:   btn,
			Double: isDouble,
		},
	}, nil
}

func parseCommandInput(lineNumber int, source string) (*model.Command, error) {
	if len(source) < 3 || source[1] != '=' {
		log.Printf(errTemplate, "缺少等号", lineNumber, source)
		return nil, ErrNoEqual
	}
	return &model.Command{
		Type: model.CommandTypeInput,
		Line: lineNumber,
		Args: model.InputArgs{
			Content: source[2:],
		},
	}, nil
}

func parseCommandTap(lineNumber int, source string) (*model.Command, error) {
	if len(source) < 3 || source[1] != '=' {
		log.Printf(errTemplate, "缺少等号", lineNumber, source)
		return nil, ErrNoEqual
	}
	var rawConfig = strings.TrimSpace(source[2:])
	if strings.Contains(rawConfig, ",") {
		var keys = strings.Split(rawConfig, ",")
		for _, key := range keys {
			key = strings.TrimSpace(key)
			if key == "" {
				log.Printf(errTemplate, "有空按键", lineNumber, source)
				return nil, ErrEmptyKey
			}
			if _, ok := robotgo.Keycode[key]; !ok {
				log.Printf(errTemplate, fmt.Sprintf("不支持该按键：%s", key), lineNumber, source)
				return nil, ErrNotSupportKey
			}
		}
		return &model.Command{
			Type: model.CommandTypeTap,
			Line: lineNumber,
			Args: model.TapArgs{
				CombineKeys: keys,
				RepeatKeys:  nil,
			},
		}, nil
	}
	if strings.Contains(rawConfig, "*") {
		var config = strings.Split(rawConfig, "*")
		var key string
		if len(config) > 2 {
			key = "*"
		} else {
			key = strings.TrimSpace(config[0])
		}
		if _, ok := robotgo.Keycode[key]; !ok {
			log.Printf(errTemplate, fmt.Sprintf("不支持该按键：%s", key), lineNumber, source)
			return nil, ErrNotSupportKey
		}
		count, err := strconv.Atoi(config[len(config)-1])
		if err != nil {
			log.Printf(errTemplate, "不是整数", lineNumber, source)
			return nil, ErrNotInt
		}
		var repeatKeys []string
		for i := 0; i < count; i++ {
			repeatKeys = append(repeatKeys, key)
		}
		return &model.Command{
			Type: model.CommandTypeTap,
			Line: lineNumber,
			Args: model.TapArgs{
				CombineKeys: nil,
				RepeatKeys:  repeatKeys,
			},
		}, nil
	}
	if _, ok := robotgo.Keycode[rawConfig]; !ok {
		log.Printf(errTemplate, fmt.Sprintf("不支持该按键：%s", rawConfig), lineNumber, source)
		return nil, ErrNotSupportKey
	}
	return &model.Command{
		Type: model.CommandTypeTap,
		Line: lineNumber,
		Args: model.TapArgs{
			CombineKeys: []string{rawConfig},
			RepeatKeys:  nil,
		},
	}, nil
}

func parseCommandSleep(lineNumber int, source string) (*model.Command, error) {
	if len(source) < 3 {
		log.Printf(errTemplate, "缺少等号", lineNumber, source)
		return nil, ErrNoEqual
	}
	duration, err := strconv.Atoi(source[2:])
	if err != nil {
		log.Printf(errTemplate, "duration value not int", lineNumber, source)
		return nil, ErrNotInt
	}
	return &model.Command{
		Type: model.CommandTypeSleep,
		Line: lineNumber,
		Args: model.Sleep{
			Duration: duration,
		},
	}, nil
}
