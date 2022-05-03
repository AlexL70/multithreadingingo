package common

type Train struct {
	Id          int
	TrainLength int
	Front       int
}

type Intersection struct {
	Id       int
	LockedBy int
}

type Crossing struct {
	Position     int
	Intersection *Intersection
}
