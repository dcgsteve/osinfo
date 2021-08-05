package main

import (
	"fmt"
	"testing"
)

func TestNewStatMan(t *testing.T) {
	sm, err := NewStatMan()
	if err != nil {
		t.FailNow()
	}
	var results = sm.LoadKPI()
	if results.alertStatus {
		fmt.Printf("Oh no - something needs to be looked at!\nCPU percentage is %v, Memory percentage is %v", results.CPUPercUsed, results.MEMPercUsed)
	} else {
		fmt.Println("All ok - phew!")
	}
}
