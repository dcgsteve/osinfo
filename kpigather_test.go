package main

import (
	"fmt"
	"testing"
)

func TestKPIGather(t *testing.T) {
	sm, err := NewKPIGather()
	if err != nil {
		t.FailNow()
	}

	var results = sm.LoadKPI()

	fmt.Printf("Checking machine with ID %v\n", results.MachineID)

	if results.alertBitmask&1 != 0 {
		fmt.Printf("    CPU alert! %v percent used\n", results.CPUPercUsed)
	}

	if results.alertBitmask&2 != 0 {
		fmt.Printf("    Memory alert! %v percent used\n", results.MEMPercUsed)
	}

	if results.alertBitmask&4 != 0 {
		fmt.Println("    Disk space alert!")
		for _, part := range results.MountPercUsed {
			fmt.Printf("     - Mount point %v %v percent used\n", part.Name, part.PercUsed)
		}
	}

	fmt.Println("Check complete")
}
