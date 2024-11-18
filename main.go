package main

import "design-pattern/singleton"

func main() {
	go singleton.GetInstance()
}
