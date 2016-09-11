package gohab

import (
	"errors"
	"github.com/sajal/milightgo"
	"strconv"
)

var (
	errColorUnsupported = errors.New("milight module currently only supports white light")
)

//Milight represents a single light bulb zone managed using wifi controller box
type Milight struct {
	addr       string //ip address/port of controller
	zone       int    //Zone .. 1 - 4
	brighness  int    //Current brightness - using state machine 1..10
	warmth     int    //Current warmth - using state machine 1..10
	on         bool   //Is this light on? - using state machine
	controller *milight.Controller
}

func (m *Milight) On() error {
	err := m.controller.ZoneOn(m.zone)
	if err != nil {
		return err
	}
	m.on = true
	return nil
}

func (m *Milight) Off() error {
	err := m.controller.ZoneOff(m.zone)
	if err != nil {
		return err
	}
	m.on = false
	return nil
}

func (m *Milight) SetBrightness(n int) error {
	return nil
}

func (m *Milight) SetWarmth(n int) error {
	return nil
}

func (m *Milight) SetColor(r, g, b int) error {
	return errColorUnsupported
}

func (m *Milight) GetBrightness() (n int, err error) {
	return 0, err
}

func (m *Milight) GetWarmth() (n int, err error) {
	return 0, err
}

func (m *Milight) GetColor() (r, g, b int, err error) {
	return 0, 0, 0, errColorUnsupported
}

func NewMilight(args map[string]string) (m *Milight, err error) {
	m = &Milight{}
	a, ok := args["addr"]
	if !ok || a == "" {
		return nil, errors.New("milight: 'addr' missing")
	}
	m.addr = a
	z, ok := args["zone"]
	if !ok || z == "" {
		return nil, errors.New("milight: 'zone' missing")
	}
	m.zone, err = strconv.Atoi(z)
	if err != nil {
		return nil, err
	}
	if m.zone > 4 || m.zone < 1 {
		return nil, errors.New("milight: 'zone' must be between 1 - 4")
	}
	m.controller = milight.NewController(a)
	return m, nil
}
