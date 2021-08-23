package sensors

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type HardwareSensors struct {
	Content string                 `json:"-"`
	Chips   map[string]ChipEntries `json:"chips"`
}

type ChipEntries map[string]string
type ChipData map[string]interface{}

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

func construct(in []byte) *ChipData {
	var out ChipData
	fmt.Println(string(in))
	if err := json.Unmarshal(in, &out); err != nil {
		fmt.Println(err)
	}
	return &out
}

func NewSensorSystem() (*HardwareSensors, error) {
	executionResult, err := exec.Command("sensors").Output()
	if err != nil {
		return &HardwareSensors{}, errors.New("lmsensors package missing, please install")
	}

	return construction(string(executionResult)), nil
}

func NewJsonSensorSystem() (*ChipData, error) {
	executionResult, err := exec.Command("sensors", "-j").Output()
	if err != nil {
		return &ChipData{}, errors.New("lmsensors package missing, please install")
	}

	return construct(executionResult), nil
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

func TemperatureToFloat(value string) (currentTemp, highWarning, criticalTreshold float64) {
	strTemp := strings.TrimLeft(value, "+")
	splitted := strings.SplitAfter(strTemp, "°C")
	currentTemp, _ = strconv.ParseFloat(strings.TrimRight(splitted[0], "°C"), 64)
	temp1 := strings.TrimLeft(splitted[1], "(high = +")
	temp2 := strings.TrimLeft(splitted[2], ", crit = +")
	highWarning, _ = strconv.ParseFloat(strings.TrimRight(temp1, "°C"), 64)
	criticalTreshold, _ = strconv.ParseFloat(strings.TrimRight(temp2, "°C"), 64)
	return currentTemp, highWarning, criticalTreshold
}
