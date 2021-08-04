/*

Checks whether average CPU load over last 5 minutes > 80%
Checks if any partitions are > 60% used
Checks if any partitions are > 80% used

Checks if available memory is < 20%  **WE NEED TO LOOK AT THIS OVER x MINS NOT JUST ONE OFF - HOW? - DO WE KEEP LAST x RUNS LOCALLY?**

*/

package main

import (
	"fmt"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
)

func main() {
	getHostInfo()
	getCpuInfo()
	getCpuLoad()
	getMemInfo()
	getDiskInfo()
}

func getHostInfo() {
	hInfo, _ := host.Info()
	fmt.Printf("HOST INFO\n%v\n\n", hInfo)
}

func getCpuLoad() {
	info, _ := load.Avg()
	fmt.Printf("CPU LOAD\n%v\n\n", info)
}

func getMemInfo() {
	memInfo, _ := mem.VirtualMemory()
	fmt.Printf("MEMORY INFO\n%v\n\n", memInfo)
}

func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get Partitions failed, err:%v\n", err)
		return
	}

	fmt.Printf("DISK INFO\n")
	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
	}
}

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}
	fmt.Printf("CORE COUNT %v\n", len(cpuInfos))
	fmt.Printf("CPU INFO FROM CORE 1\n%v\n", cpuInfos[0])
}
