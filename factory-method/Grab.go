package factory_method

import "fmt"

type Grab struct {
}

func (g Grab) CreateTaxi() ITaxi {
	return &GrabTaxi{}
}

type GrabTaxi struct {
}

func (g GrabTaxi) Drive() {
	fmt.Printf("Grab Taxi is driving")
}

func (g GrabTaxi) Pickup() {
	fmt.Printf("Grab Taxi is picking up passenger")
}

func (g GrabTaxi) GetFare() int32 {
	return 123
}
