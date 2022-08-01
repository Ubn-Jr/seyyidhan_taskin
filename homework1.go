package main

import "fmt"

func main() {

	visaexamscore := 80
	finalexamscore := 60
	bring := average(visaexamscore, finalexamscore)

	fmt.Println(bring)

}

func average(x int, y int) int {

	calculatex := (x * 40) / 100
	calculatey := (y * 60) / 100

	return calculatex + calculatey

}
