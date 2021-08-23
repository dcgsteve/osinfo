package sensors

import (
	"fmt"
	"log/syslog"
	"strconv"
	"strings"
)

func SendToSyslog(result KPIResults, tempResult *HardwareSensors) error {
	var logFormatter []string
	syslogger, err := syslog.Dial("", "", syslog.LOG_CRIT, "os_stat")
	if err != nil {
		return err
	}

	if result.CPUPercUsed >= CPULIMIT || result.MEMPercUsed >= MEMLIMIT || len(result.MountPercUsed) > 0 {
		logFormatter = append(logFormatter, "MachineID: ", result.MachineID)
	}

	if result.CPUPercUsed >= CPULIMIT {
		logFormatter = append(logFormatter, " CPU Usage: ", strconv.Itoa(int(result.CPUPercUsed)))
	}

	if result.MEMPercUsed >= MEMLIMIT {
		logFormatter = append(logFormatter, " Memory: ", strconv.Itoa(int(result.MEMPercUsed)))
	}

	for _, diskInfo := range result.MountPercUsed {
		logFormatter = append(logFormatter, " Mount Point: ", diskInfo.Name, " Mount Point Percentage used: ", strconv.Itoa(diskInfo.PercUsed))
	}
	if len(tempResult.Chips) > 0 {
		for chip := range tempResult.Chips {
			for k, v := range tempResult.Chips[chip] {
				if strings.Contains(k, "Core") {
					//currentTemp, _, criticalTreshold := TemperatureToFloat(v)
					currentTemp, _, _ := TemperatureToFloat(v)
					if currentTemp >= 10 {
						temp := fmt.Sprintf("%v", currentTemp)
						logFormatter = append(logFormatter, " ", k, ": ", temp)
					}
				}
			}
		}
	}

	if len(logFormatter) > 0 {
		if _, err = fmt.Fprintf(syslogger, strings.Join(logFormatter, "")); err != nil {
			return err
		}
	}

	return nil
}
