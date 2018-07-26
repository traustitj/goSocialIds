package main

import (
	"./socialid"
)

func generateForYear(year int, c chan []string) {
	ids := socialid.GenerateYear(year)
	c <- ids
}

func main() {
	// size := 2018 - 1900
	c := make(chan []string, 2018-1900+1)

	for i := 1900; i < 2018; i++ {
		go generateForYear(1974, c)
	}
	for range c {
		<-c
		// fmt.Printf("Generated %d ids for year %d\n", len(ids), i)
	}
}
