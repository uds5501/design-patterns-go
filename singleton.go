package main

import (
	"log"
	"sync"
)

// Singleton pattern involves a single class which is responsible to create an object while making sure that only single
// object gets created. This class provides a way to access its only object which can be accessed directly
// without need to instantiate the object of the class.

var instance *Calculator
var once sync.Once

type Calculator struct{}

func GetInstance() *Calculator {
	once.Do(func() {
		log.Println("Initializing Calculator")
		instance = &Calculator{}
	})
	return instance
}

func (c Calculator) Hello() {
	log.Println("Hello!")
}
func main() {
	x := GetInstance()
	y := GetInstance()
	z := GetInstance()
	x.Hello()
	y.Hello()
	z.Hello()
}
