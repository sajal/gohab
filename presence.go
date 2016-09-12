package gohab

//Presence is a type of boolean device that indicates occupancy
type Presence interface {
	Occupied() (present bool, err error)
}
