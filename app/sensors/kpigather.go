package sensors

import (
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

// Tunable percentage limits
const CPULIMIT = 85
const MEMLIMIT = 85
const DISKLIMIT = 85

// Bitmask settings
const ALERT_CPU = 1
const ALERT_MEM = 2
const ALERT_MOUNTPOINT = 4

type KPI struct {
	host  *host.InfoStat
	load  *load.AvgStat
	mem   *mem.VirtualMemoryStat
	parts []disk.PartitionStat
}

type KPIResults struct {
	MachineID     string
	alertBitmask  uint8
	CPUPercUsed   uint8
	MEMPercUsed   uint8
	MountPercUsed []MountPoint
}

type MountPoint struct {
	Name     string
	PercUsed int
}

func NewKPIGather() (*KPI, error) {
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

	return &KPI{
		host:  host,
		load:  load,
		mem:   mem,
		parts: parts,
	}, nil
}

func (sm *KPI) LoadKPI() KPIResults {

	// Build up response as we go along
	var response KPIResults

	// Default to all ok
	response.alertBitmask = 0

	// Set machineID for reference
	response.MachineID = sm.host.HostID

	// CPU checks - NB. order is important, don't change :)
	if (sm.load.Load5 * 100) > CPULIMIT {
		response.CPUPercUsed = uint8(sm.load.Load5 * 100)
		response.alertBitmask = response.alertBitmask | ALERT_CPU
	}
	if (sm.load.Load15 * 100) > CPULIMIT {
		response.CPUPercUsed = uint8(sm.load.Load15 * 100)
		response.alertBitmask = response.alertBitmask | ALERT_CPU
	}

	// Memory checks
	if sm.mem.UsedPercent > MEMLIMIT {
		response.MEMPercUsed = uint8(sm.mem.UsedPercent)
		response.alertBitmask = response.alertBitmask | ALERT_MEM
	}

	// Partition checks
	for _, part := range sm.parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			continue
		}
		if usage.UsedPercent > DISKLIMIT {
			response.MountPercUsed = append(response.MountPercUsed, MountPoint{part.Mountpoint, int(usage.UsedPercent)})
			response.alertBitmask = response.alertBitmask | ALERT_MOUNTPOINT
		}
	}

	return response

}
