package main

import "design-pattern/prototype"

func main() {
	mclaren := prototype.McLaren{
		Name: "McLaren 22",
	}
	newMclaren := mclaren.Clone().(*prototype.McLaren)
	print(newMclaren.Name)
}
