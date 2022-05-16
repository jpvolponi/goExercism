/*
card	value	card	value
ace	11	eight	8
two	2	nine	9
three	3	ten	10
four	4	jack	10
five	5	queen	10
six	6	king	10
seven	7	other	0
*/

package main

import (
	"fmt"
	"strings"
)

func ParseCard(s string) int {
	var resp int
	switch strings.ToLower(s) {
	case "ace":
		resp = 11
	case "two":
		resp = 2
	case "three":
		resp = 3
	case "four":
		resp = 4
	case "five":
		resp = 5
	case "six":
		resp = 6
	case "seven":
		resp = 7
	case "eight":
		resp = 8
	case "nine":
		resp = 9
	case "ten":
		resp = 10
	case "jack":
		resp = 10
	case "queen":
		resp = 10
	case "king":
		resp = 10
	default:
		resp = 0
	}

	return resp
}
func FirstTurn(card1, card2, dealerCard string) string {
	var resp string
	if ParseCard(card1) == 11 && ParseCard(card2) == 11 {
		resp = "P"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) < 10 {
		resp = "W"
	} else if ParseCard(card1)+ParseCard(card2) == 21 && ParseCard(dealerCard) >= 10 {
		resp = "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 17 && ParseCard(card1)+ParseCard(card2) <= 20 {
		resp = "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) < 7 {
		resp = "S"
	} else if ParseCard(card1)+ParseCard(card2) >= 12 && ParseCard(card1)+ParseCard(card2) <= 16 && ParseCard(dealerCard) >= 7 {
		resp = "H"
	} else if ParseCard(card1)+ParseCard(card2) <= 11 {
		resp = "H"
	}

	return resp

}

func main() {
	//fmt.Println(ParseCard("ace") + ParseCard("ace"))

	fmt.Println(FirstTurn("ace", "ace", "jack"))
	fmt.Println(FirstTurn("ace", "king", "ace"))
	fmt.Println(FirstTurn("five", "queen", "ace"))
	fmt.Println(FirstTurn("ace", "queen", "two"))
}
