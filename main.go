package main

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// Tunable limits
const CPULIMIT float64 = 0.8
const MEMLIMIT float64 = 0.9

type StatMan struct {
	host  *host.InfoStat
	load  *load.AvgStat
	mem   *mem.VirtualMemoryStat
	parts []*disk.PartitionStat
}

type KPIResults struct {
	alertStatus bool
	CPUPercUsed int
	MEMPercUsed int
	Partitions  []PartitionUsed
}

type PartitionUsed struct {
	Name     string
	PercUsed int
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
	parts, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}

	return &StatMan{
		host:  host,
		load:  load,
		mem:   mem,
		parts: parts,
	}, nil
}

func (sm *StatMan) LoadKPI() KPIResults {
	var response KPIResults
	response.alertStatus = false

	// CPU checks - NB. order is important, don't change :)
	if sm.load.Load5 > CPULIMIT {
		response.CPUPercUsed = int(sm.load.Load5 * 100)
		response.alertStatus = true
	}
	if sm.load.Load15 > CPULIMIT {
		response.CPUPercUsed = int(sm.load.Load15 * 100)
		response.alertStatus = true
	}

	// Memory checks
	if sm.mem.UsedPercent > (MEMLIMIT * 100) {
		response.MEMPercUsed = int(sm.mem.UsedPercent)
		response.alertStatus = true
	}

	// Partition checks
	for _, part := range sm.parts {
		fmt.Printf("part:%v\n", part.String())
	}

	return response

}
