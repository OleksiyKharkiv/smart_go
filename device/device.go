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

func (d *Device) TurnOff() {
	d.status = StatusOff
}

func (d *Device) SetSafeStatus(isSafe bool) {
	d.isSafe = isSafe
}

func (d *Device) Name() string {
	return d.name
}

func (d *Device) Power() int {
	return d.power
}

func (d *Device) IsSafe() bool {
	return d.isSafe
}

func (d *Device) IsOn() bool {
	return d.status == StatusOn
}

// String interface Stringer
func (d *Device) String() string {
	return fmt.Sprintf("device %s is %s and %s, power: %dW",
		d.name,
		d.status.String(),
		d.safetyString(),
		d.power)
}

func (d *Device) safetyString() string {
	if d.isSafe {
		return "safe"
	}
	return "unsafe"
}

// String для Status
func (s Status) String() string {
	switch s {
	case StatusOn:
		return "ON"
	case StatusOff:
		return "OFF"
	default:
		return "UNKNOWN"
	}
}
