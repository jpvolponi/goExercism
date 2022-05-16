package main

import "fmt"

func FavoriteCards() []int {
	return []int{2, 6, 9}
}

func GetItem(s []int, index int) int {
	if index < len(s)-len(s) || index > len(s)-1 {
		return -1
	} else {
		return s[index]
	}
}

func SetItem(s []int, index int, newCard int) []int {
	if index < len(s)-len(s) || index > len(s)-1 {
		s = append(s, newCard)
	} else {
		s[index] = newCard
	}
	return s
}

func PrependItems(s []int, x ...int) []int {
	var temp []int
	if x != nil {
		temp = append(temp, x...)
		s = append(temp, s...)
		return s
	} else {
		return s
	}
}

func RemoveItem(s []int, index int) []int {
	var temp []int
	if index < len(s)-len(s) || index > len(s)-1 {
		return s
	} else {
		temp = append(s[:index], s[index+1:]...)
		s = temp
	}
	return s
}

func main() {
	//fmt.Println(GetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 8))
	//fmt.Println(SetItem([]int{5, 2, 10, 6, 8, 7, 0, 9}, 4, 1))
	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 2))
	fmt.Println(RemoveItem([]int{3, 2, 6, 4, 8}, 11))
}
