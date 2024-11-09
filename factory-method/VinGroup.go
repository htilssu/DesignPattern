package factory_method

type VinGroup struct {
}

type VinGroupTaxi struct {
}

func (v VinGroupTaxi) Drive() {
	println("VinGroupTaxi is driving")
}

func (v VinGroupTaxi) Pickup() {
	println("VinGroupTaxi is picking up passenger")
}

func (v VinGroupTaxi) GetFare() int32 {
	return 456
}

func (v VinGroup) CreateTaxi() ITaxi {
	return &VinGroupTaxi{}
}
