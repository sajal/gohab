package gohab

type Light interface {
	On() error                          //Turn on light
	Off() error                         //Turn on light
	SetBrightness(n int) error          //Set brightness 1 - 10
	SetWarmth(n int) error              //Set warmth 1 - 10
	SetColor(r, g, b int) error         //Set color
	GetBrightness() (n int, err error)  //Get current brightness 1 - 10
	GetWarmth() (n int, err error)      //Get current warmth 1 - 10
	GetColor() (r, g, b int, err error) //Get current color
}
