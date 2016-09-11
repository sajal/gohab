package gohab

import (
	"errors"
	"log"
)

var (
	ErrDevicenotfound = errors.New("Device not found")
	ErrNamenotfound   = errors.New("Thing does not have a name!")
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
		config: c,
		lights: make(map[string]Light),
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
