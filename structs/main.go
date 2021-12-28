package main

import (
	"fmt"
	"log"
)

// Define a type of Point
type Point struct {
	x int
	y int
}

// Define a type of Square
type Square struct {
	Centre *Point
	Length int
}

// The Area method of the Square receiver returns its area
func (s *Square) Area() int {
	area := s.Length * s.Length
	return area
}

// The Move method of the Square receiver
func (s *Square) Move(x int, y int) {
	s.Centre.x = x
	s.Centre.y = y
}

// Get the positio of our square
func (s *Square) position() Point {
	return *s.Centre
}

// The NewSquare fuction creates a new Square in position x y with sides of length.
func NewSquare(x int, y int, length int) (*Square, error) {

	if length <= 0 {
		return nil, fmt.Errorf("length must be greater than 0")
	}

	p := &Point{
		x: x,
		y: y,
	}

	s := &Square{
		Centre: p,
		Length: length,
	}
	return s, nil
}

func main() {
	// Create a square
	mySquare, err := NewSquare(10, 10, 0)
	if err != nil {
		log.Fatalf("ERROR: unable to create Square. Reason: %v", err)

	}

	// Get the area and position
	area := mySquare.Area()
	pos := mySquare.position()
	fmt.Printf("AREA: %v  POSITION: %v\n", area, pos)

	// move the square
	mySquare.Move(21, 39)

	// now check the new position
	newPos := mySquare.position()
	fmt.Printf("POSITION: %v\n", newPos)

}
