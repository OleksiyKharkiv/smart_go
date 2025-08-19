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
		devices:  make([]*device.Device, 0, maxPower),
		maxPower: maxPower,
	}
}

func (pg PowerGrid) AddDevice(d *device.Device) {
	pg.devices = append(pg.devices, d)
	fmt.Printf("Device %s succefuly added \n", d.GetName())
}

func (pg PowerGrid) TotalConsumption() int {
	total := 0
	for _, d := range pg.devices {
		if d.StatusIsOn() {
			total += d.GetPower()
		}
	}
	return total
}
