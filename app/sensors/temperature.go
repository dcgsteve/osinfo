package sensors

import (
	"encoding/json"
	"errors"
	"log"
	"os/exec"
	"strings"
)

type HardwareSensors struct {
	Content string                 `json:"-"`
	Chips   map[string]ChipEntries `json:"chips"`
}

type ChipEntries map[string]string

func construction(content string) *HardwareSensors {
	var chip string
	sensor := &HardwareSensors{}
	sensor.Content = content
	sensor.Chips = map[string]ChipEntries{}
	lines := strings.Split(sensor.Content, "\n")
	for _, line := range lines {
		if len(line) > 0 {
			if !strings.Contains(line, ":") {
				chip = line
				sensor.Chips[chip] = ChipEntries{}
			} else if len(chip) > 0 {
				parttition := strings.Split(line, ":")
				entry := parttition[0]
				val := strings.TrimRight(strings.TrimLeft(parttition[1], " "), " ")
				sensor.Chips[chip][entry] = val
			}
		}
	}
	return sensor
}

func NewSensorSystem() (*HardwareSensors, error) {
	executionResult, err := exec.Command("sensors").Output()
	if err != nil {
		return &HardwareSensors{}, errors.New("lmsensors package missing, please install")
	}

	return construction(string(executionResult)), nil
}

func (sensor *HardwareSensors) JsonFomatter() string {
	outputResult, err := json.Marshal(sensor)
	if err != nil {
		log.Println("json formatter error: ", err)
	}
	return string(outputResult)
}

func (sensor *HardwareSensors) StringFormatter() string {
	return sensor.JsonFomatter()
}
