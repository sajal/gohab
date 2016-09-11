package gohab

//Main config object
type Config struct {
	ListenAddr string      `yaml:"listen"` //HTTP address to listen on, default ":80"
	Things     []ThingConf `yaml:"things"` //List of things we want to monitor/control
}

//Thing is something that can be controlled and/or queried by the hab
type ThingConf struct {
	Name   string            `yaml:"name"`   //Friendly name for human convienence
	Type   string            `yaml:"type"`   //What type of device is this?
	Module string            `yaml:"module"` //Which Module to use
	Args   map[string]string `yaml:"args"`   //Arbitary key-value data which the relavent module can understand
}
