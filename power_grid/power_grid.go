package power_grid

import (
	"fmt"
	"smart_go/device"
)

type PowerGrid struct {
	devices  []*device.Device
	maxPower int
}

func NewPowerGrid(maxPower int) *PowerGrid {
	return &PowerGrid{
		devices:  []*device.Device{},
		maxPower: maxPower,
	}
}

func (pg *PowerGrid) AddDevice(d *device.Device) {
	pg.devices = append(pg.devices, d)
	fmt.Printf("Device %s succefuly added \n", d.Name())
}

func (pg *PowerGrid) TotalConsumption() int {
	total := 0
	for _, d := range pg.devices {
		if d.IsOn() {
			total += d.Power()
			fmt.Printf("Device %s is on\n added to grid with power  %d \n", d.Name(), d.Power())
		}
	}
	return total
}
