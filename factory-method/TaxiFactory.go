package factory_method

type ITaxiFactory interface {
	CreateTaxi() ITaxi
}
