package device

import "fmt"

type Status int

const (
	StatusOff Status = iota
	StatusOn
)

type Device struct {
	name   string
	power  int
	isSafe bool
	status Status
}

func NewDevice(name string, power int) *Device {
	return &Device{
		name:   name,
		power:  power,
		status: StatusOff,
		isSafe: true,
	}
}
func (d *Device) TurnOn() error {
	if !d.isSafe {
		return fmt.Errorf("device %s is n ot safe!", d.name)
	}
	d.status = StatusOn
	return nil
}
