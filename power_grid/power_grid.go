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
func (pg *PowerGrid) AutoDisable() error {
	if len(pg.devices) == 0 {
		return fmt.Errorf("no devices in grid")
	}
	for pg.TotalConsumption() > pg.maxPower {
		maxDevice := pg.findMaxPowerDevice()
		if maxDevice == nil {
			return fmt.Errorf("no device to turn off")
		}
		maxDevice.TurnOff()
	}
	return nil
}

func (pg *PowerGrid) findMaxPowerDevice() *device.Device {
	var maxDevice *device.Device
	for _, d := range pg.devices {
		if d.IsOn() && (maxDevice == nil || d.Power() > maxDevice.Power()) {
			maxDevice = d
		}
	}
	return maxDevice

}
