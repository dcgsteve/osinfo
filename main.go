package main

import (
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// Tunable limits
const CPULIMIT float64 = 0.8
const MEMLIMIT float64 = 0.9

type StatMan struct {
	host *host.InfoStat
	load *load.AvgStat
	mem  *mem.VirtualMemoryStat
	// disk *disk.PartitionStat
}

type KPIResults struct {
	alertStatus bool
	CPUPerc     int
	MEMPerc     int
}

func NewStatMan() (*StatMan, error) {
	// Host info
	host, err := host.Info()
	if err != nil {
		return nil, err
	}

	// CPU load (average over 1 minute, 5 minutes, 15 minutes)
	load, err := load.Avg()
	if err != nil {
		return nil, err
	}

	// Memory info
	mem, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	// Partition info
	// disk, err := disk.Partitions(false)
	// if err != nil {
	// 	return nil, err
	// }

	return &StatMan{
		host: host,
		load: load,
		mem:  mem,
	}, nil
}

func (sm *StatMan) LoadKPI() KPIResults {
	var response KPIResults
	response.alertStatus = false

	// CPU checks - NB. order is important, don't change :)
	if sm.load.Load5 > CPULIMIT {
		response.CPUPerc = int(sm.load.Load5 * 100)
		response.alertStatus = true
	}
	if sm.load.Load15 > CPULIMIT {
		response.CPUPerc = int(sm.load.Load15 * 100)
		response.alertStatus = true
	}

	// Memory checks
	if sm.mem.UsedPercent > (MEMLIMIT * 100) {
		response.MEMPerc = int(sm.mem.UsedPercent)
		response.alertStatus = true
	}

	return response

}

// func getDiskInfo() {
// 	parts, err := disk.Partitions(true)
// 	if err != nil {
// 		fmt.Printf("get Partitions failed, err:%v\n", err)
// 		return
// 	}

// 	fmt.Printf("DISK INFO\n")
// 	for _, part := range parts {
// 		fmt.Printf("part:%v\n", part.String())
// 	}
// }
