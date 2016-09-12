package main

import (
	"github.com/sajal/gohab"
	"log"
)

func main() {
	arg := make(map[string]string)
	arg["addr"] = "192.168.5.62:8899"
	arg["zone"] = "4"
	arparg := make(map[string]string)
	arparg["mac"] = "a0:39:f7:39:0c:6e"
	//TODO: Load yaml
	cfg := &gohab.Config{
		ListenAddr: ":1337",
		Things: []gohab.ThingConf{
			gohab.ThingConf{
				Name:   "bedroom",
				Type:   "light",
				Module: "milight",
				Args:   arg,
			},
			gohab.ThingConf{
				Name:   "arpphone",
				Type:   "presence",
				Module: "arp",
				Args:   arparg,
			},
		},
	}
	hab, err := gohab.NewHab(cfg)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(hab)
	log.Println(hab.LightOff("bedroom"))
	log.Println(hab.Occupied())
}
