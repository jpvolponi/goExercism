package main

type Car struct {
	battery      int
	batteryDrain int
	speed        int
	distance     int
}

type Track struct {
	distance int
}

// TODO: define the 'Car' type struct

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{speed: speed, batteryDrain: batteryDrain, battery: 100, distance: 0}
}

// TODO: define the 'Track' type struct

// NewTrack created a new track
func NewTrack(distance int) Track {
	return Track{distance: distance}
}

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func Drive(c Car) Car {
	if c.battery < c.batteryDrain {
		return c
	} else {
		c.battery = c.battery - c.batteryDrain
		c.distance += c.speed
	}

	return c
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, raceTrack Track) bool {
	batterySafe := (raceTrack.distance / car.speed) * car.batteryDrain
	var resp bool
	if car.battery <= 100 && batterySafe <= car.battery && raceTrack.distance >= 0 {
		resp = true
	} else {
		resp = false
	}

	return resp
}
