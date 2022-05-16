//https://exercism.org/tracks/go/exercises/vehicle-purchase

package main

import (
	"fmt"
	"strings"
)

func NeedsLicense(kind string) bool {
	return strings.ToLower(kind) == "car" || strings.ToLower(kind) == "truck"
}

func ChooseVehicle(option1, option2 string) string {
	var resp string
	if option1 > option2 {
		resp = fmt.Sprintf("%s is clearly the better choice.", option2)
	} else {
		resp = fmt.Sprintf("%s is clearly the better choice.", option1)
	}
	return resp
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	var resp float64
	if age >= 10 {
		resp = originalPrice * 0.5
	} else if age >= 3 && age < 10 {
		resp = originalPrice * (0.7)
	} else {
		resp = originalPrice * (0.8)
	}

	return resp
}

func main() {
	fmt.Println(NeedsLicense("CAR"))
	fmt.Println(ChooseVehicle("Wuling Hongguang", "Toyota Corolla"))
	fmt.Println(ChooseVehicle("Volkswagen Beetle", "Volkswagen Golf"))
	fmt.Println(CalculateResellPrice(1000, 1))
	fmt.Println(CalculateResellPrice(1000, 5))
	fmt.Println(CalculateResellPrice(1000, 15))

}
