package main

import "fmt"

type RoadVehicle interface {
	Accelerate()
	Brake()
	Turn(direction string)
	CurrentSpeed() int
}

type Car struct {
	Make        string
	Model       string
	Colour      string
	TopSpeed    int
	ZeroToSixty float32
}

type Lorry struct {
	Make        string
	Model       string
	MaxLoadTons float32
	Articulated bool
}

func main() {
	myCar := Car{
		Make:   "Honda",
		Model:  "Jazz",
		Colour: "Black",
	}

	PrintInfo(&myCar)

}

func PrintInfo(rv RoadVehicle) {
	fmt.Println("The current speed of this vehicle is ", rv.CurrentSpeed())
}

func (c *Car) Accelerate() {
	fmt.Println("Accelerateing")
}

func (c *Car) Brake() {
	fmt.Println("Braking")
}

func (c *Car) CurrentSpeed() int {
	return 30 // Returing fake value for testing
}

func (c *Car) Turn(direction string) {
	fmt.Println("Turning", direction)
}
