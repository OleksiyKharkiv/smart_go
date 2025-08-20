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
		}
	}
	return total
}
func (pg *PowerGrid) AutoDisable() (*device.Device, error) {
	if len(pg.devices) == 0 {
		return nil, fmt.Errorf("no devices in grid")
	}
	for pg.TotalConsumption() > pg.maxPower {
		maxDevice := pg.findMaxPowerDevice()
		if maxDevice == nil {
			return nil, fmt.Errorf("no device to turn off")
		}
		maxDevice.TurnOff()
		return maxDevice, nil
	}
	return nil, nil
}

func (pg *PowerGrid) findMaxPowerDevice() *device.Device {
	var MaxDevice *device.Device
	for _, d := range pg.devices {
		if d.IsOn() && (MaxDevice == nil || d.Power() > MaxDevice.Power()) {
			MaxDevice = d
		}
	}
	return MaxDevice

}
func (pg *PowerGrid) MaxPower() int {
	return pg.maxPower
}
