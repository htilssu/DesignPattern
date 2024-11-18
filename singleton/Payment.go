package singleton

import "sync"

type Payment struct {
}

var lock = &sync.Mutex{}
var payment *Payment

func GetInstance() *Payment {
	if payment == nil {
		lock.Lock()
		payment = &Payment{}
		println("Create new payment")

		defer lock.Unlock()
	} else {
		println("Payment already created")
	}
	return payment
}
