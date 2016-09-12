package gohab

import (
	"errors"
	"github.com/mostlygeek/arp"
	"log"
)

type Arp struct {
	mac string
	ip  string
}

//Occupied checks if local Arp cache is aware of this device
func (a *Arp) Occupied() (present bool, err error) {
	tbl := arp.Table()
	for k, v := range tbl {
		if a.mac != "" && v == a.mac {
			return true, nil
		} else if k == a.ip && a.mac == "" {
			return true, nil
		}
	}
	return false, nil
}

func NewArp(args map[string]string) (a *Arp, err error) {
	a = &Arp{}
	ip := args["ip"]
	m := args["mac"]

	a.ip = ip
	a.mac = m

	if ip == "" && m == "" {
		return nil, errors.New("arp: 'ip' or 'mac' needed")
	}
	log.Println("Arp presence module initialized ip: ", ip, " mac: ", m)

	return a, nil
}
