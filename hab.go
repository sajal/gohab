package gohab

import (
	"errors"
	"log"
)

var (
	ErrDevicenotfound        = errors.New("Device not found")
	ErrNamenotfound          = errors.New("Thing does not have a name!")
	ErrPresencenotconfigured = errors.New("No presence module detected!")
)

//The main instance of the hab
type Hab struct {
	config   *Config
	lights   map[string]Light
	presence map[string]Presence
}

//NewHab creates a new Hab
func NewHab(c *Config) (*Hab, error) {
	//Initialize sane defaults
	if c == nil {
		c = &Config{}
	}
	if c.ListenAddr == "" {
		c.ListenAddr = ":80"
	}
	h := &Hab{
		config:   c,
		lights:   make(map[string]Light),
		presence: make(map[string]Presence),
	}
	//Initialize things...
	for _, thing := range c.Things {
		if thing.Name == "" {
			return nil, ErrNamenotfound
		}
		if thing.Type == "light" {
			if thing.Module == "milight" {
				log.Println("Initializing new light using module milight")
				m, err := NewMilight(thing.Args)
				if err != nil {
					return nil, err
				}
				h.lights[thing.Name] = m
			} else {
				return nil, errors.New("module " + thing.Module + " not found")
			}
		} else if thing.Type == "presence" {
			if thing.Module == "arp" {
				a, err := NewArp(thing.Args)
				if err != nil {
					return nil, err
				}
				h.presence[thing.Name] = a
			} else {
				return nil, errors.New("module " + thing.Module + " not found")
			}
		} else {
			return nil, errors.New("type " + thing.Type + " not found")
		}
	}
	return h, nil
}

func (h *Hab) LightOn(name string) error {
	l, ok := h.lights[name]
	if !ok {
		return ErrDevicenotfound
	}
	return l.On()
}

func (h *Hab) LightOff(name string) error {
	l, ok := h.lights[name]
	if !ok {
		return ErrDevicenotfound
	}
	return l.Off()
}

//OccupiedPresence checks a single named instance
func (h *Hab) OccupiedPresence(name string) (bool, error) {
	l, ok := h.presence[name]
	if !ok {
		return false, ErrDevicenotfound
	}
	return l.Occupied()
}

//Occupied checks all Presence modules and returns true if any of them are true
func (h *Hab) Occupied() (res bool, err error) {
	if len(h.presence) == 0 {
		err = ErrPresencenotconfigured
		return
	}
	var r bool
	for _, l := range h.presence {
		r, err = l.Occupied()
		if err != nil {
			return
		}
		res = res || r
	}
	//TODO: Implement grace: check state somewhere in a loop and cache it
	return
}
