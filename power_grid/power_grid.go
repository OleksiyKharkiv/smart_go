package power_grid

import (
	"fmt"
	"runtime"
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
func (pg *PowerGrid) AutoDisable() {
	var maxDevice *device.Device
	for _, d := range pg.devices {
		if maxDevice == nil || d.Power() > maxDevice.Power() {
			if d.IsOn() {
				maxDevice = d
			}
		}
	}
	if maxDevice == nil {
		runtime.Breakpoint()
	}
	maxDevice.TurnOff()
	fmt.Printf("=================================\n "+
		"Power grid is overloaded! \n Device %s is off\n================================= \n", maxDevice.Name())
	fmt.Printf("Actual total consumer of grid is : %d \n", pg.TotalConsumption())
}
