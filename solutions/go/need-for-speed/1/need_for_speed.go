package speed

// TODO: define the 'Car' type struct

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery:      100,
		distance:     0,
		batteryDrain: batteryDrain,
		speed:        speed,
	}
}

// TODO: define the 'Track' type struct

// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if !(car.batteryDrain > car.battery) { // if can run
		car.distance += car.speed
		car.battery -= car.batteryDrain
		return car
	}

	return car // if can't run
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	rotCom := track.distance / car.speed
	batteryCom := car.battery / car.batteryDrain

	return rotCom < batteryCom || rotCom == batteryCom
}
