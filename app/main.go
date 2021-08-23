package main

import (
	"github.com/dcgsteve/osinfo/app/sensors"
	"log"
)

func main() {
	sm, err := sensors.NewKPIGather()
	if err != nil {
		log.Fatal(err)
	}
	sensorTemperature, err := sensors.NewSensorSystem()
	if err != nil {
		log.Fatal(err)
	}

	results := sm.LoadKPI()

	if err = sensors.SendToSyslog(results, sensorTemperature); err != nil {
		log.Printf("error occured suring logging to os syslog: %v", err)
	}
}
