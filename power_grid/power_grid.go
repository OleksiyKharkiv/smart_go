package power_grid

import "smart_go/device"

type PowerGrid struct {
	devices [] *device.Device
	maxPower int
}

func NewPowerGrid(macPower int) *PowerGrid{
	return nil
}

func (pg *PowerGrid) AddDevice (device *device.Device){
	pg.devices = append(pg.devices, device)
}

func