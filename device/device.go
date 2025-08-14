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
		return fmt.Errorf("device %s is not safe!", d.name)
	}
	d.status = StatusOn
	return nil
}
func (d *Device) SetSafeStatus(isSafe bool) {
	d.isSafe = isSafe
}
func (d *Device) String() string {
	status := "OFF"
	if d.status == StatusOn {
		status = "ON"
	}
	safe := "safe"
	if !d.isSafe {
		safe = "unsafe"
	}

	return fmt.Sprintf("device %s is %s and %s, power: %dW", d.name, status, safe, d.power)

}
func (s Status) String() string {
	switch s {
	case StatusOn:
		return "ON"
	case StatusOff:
		return "OFF"
	}
	return "UNKNOWN"
}
