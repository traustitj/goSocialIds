package main

import (
	"fmt"
	"./socialid"
)

func generateForYear(year int, c chan []string) {
	ids := socialid.GenerateYear(year)
	c <- ids
}

func main() {
	c := make(chan []string)

	for i := 1900; i < 2018; i++ {
		go generateForYear(i, c)
	}
	total := 0
	for i := 1900; i < 2018; i++ {
		x := <-c
		fmt.Printf("%d ids for year %d\n", len(x), i)
		total += len(x)
	}

	fmt.Println("Totally generated numbers : ", total)
}
