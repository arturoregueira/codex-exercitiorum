package jedlik // We finished this one.w-;l

import (
	"fmt"
)

// TODO: define the 'Drive()' method
func (mycar *Car) Drive() {

	newBattery := mycar.battery - mycar.batteryDrain
	newDistance := mycar.speed

	if !(newBattery < 0) {
		mycar.battery = newBattery
		mycar.distance = newDistance
	}

}

// TODO: define the 'DisplayDistance() string' method
func (mycar *Car) DisplayDistance() string {
	output := fmt.Sprintf("Driven %d meters", mycar.distance)

	return output
}

// TODO: define the 'DisplayBattery() string' method
func (mycar *Car) DisplayBattery() string {
	output := fmt.Sprintf("Battery at %d", mycar.battery) + "%"

	return output
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (mycar *Car) CanFinish(trackDistance int) bool {
	runsToFinish := float64(trackDistance) / float64(mycar.speed)
	//fmt.Println(runsToFinish)
	runsThatCanBeDone := float64(mycar.battery) / float64(mycar.batteryDrain)
	//fmt.Println(runsThatCanBeDone)

	return runsToFinish <= runsThatCanBeDone
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
