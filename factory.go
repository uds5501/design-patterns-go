package main

import "log"

// Factory pattern - In Factory pattern, we create object without exposing the creation logic to the client and refer to
// newly created object using a common interface.

// Shape is the common interface which each further shape
// has to implement.
type Shape interface {
	// draw will be called by all the shapes differently
	draw()
}

// Circle will be implementing Shape interface
type Circle struct{}

// draw is the Circle's implementation of draw
func (c Circle) draw() {
	log.Println("Hello from a circle")
}

// Rectangle will be implementing Shape interface
type Rectangle struct{}

// draw is the Rectangle's implementation of draw
func (r Rectangle) draw() {
	log.Println("Hello from a rectangle")
}

// Square will be implementing Shape interface
type Square struct{}

// draw is the Square's implementation of draw
func (s Square) draw() {
	log.Println("Hello from a Square")
}

// ShapeFactory is the interface which will return Shape
type ShapeFactory interface {
	create(string) Shape
}

// shapeFactory implements ShapeFactory
type shapeFactory struct{}

// create returns appropriate Shape Object
func (sf shapeFactory) create(s string) Shape {
	switch s {
	case "square":
		return Square{}
	case "rect":
		return Rectangle{}
	case "circle":
		return Circle{}
	default:
		return nil
	}
}

func main() {
	factory := shapeFactory{}
	obj := factory.create("square")
	obj.draw()
}
