package power_grid

import (
	"fmt"
	"smart_go/device"
)

type PowerGrid struct {
	devices [] *device.Device
	maxPower int
}

func NewPowerGrid(maxPower int) *PowerGrid{
	return &PowerGrid{
		devices: make([] *device.Device, maxPower),
		maxPower: maxPower,
	}
}

func (pg *PowerGrid) AddDevice (d *device.Device){
	pg.devices = append(pg.devices, d)
	fmt.Printf("Device %s succefuly added \n", d.GetName())
}

func (pg PowerGrid) TotalConsumption int{


}