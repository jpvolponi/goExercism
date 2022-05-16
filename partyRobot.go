package main

import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
	resp := fmt.Sprintf("Welcome to my party, %s!", name)
	return resp
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	resp := fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
	return resp
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
	resp := fmt.Sprintf("Welcome to my party, %s!\nYou have been assigned to table %.3d. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", name, table, direction, distance, neighbor)
	return resp
}

func main() {

	fmt.Println(Welcome("Jean"))
	fmt.Println(HappyBirthday("Jean", 32))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
