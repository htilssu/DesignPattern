package factory_method

type ITaxi interface {
	Drive()
	Pickup()
	GetFare() int32
}
