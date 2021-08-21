package main

import (
	"fmt"
	"./socialid"
	"runtime"
)

func generateForYear(year int, c chan []string) {
	ids := socialid.GenerateYear(year)
	c <- ids
}

func main() {
	numcpu := runtime.NumCPU()
	from := 1700
	to := 2018
	c := make(chan []string)

	for i := from; i < to; i++ {
		go generateForYear(i, c)
	}
	total := 0
	for i := from; i < to; i++ {
		x := <-c
		fmt.Printf("%d ids for year %d\n", len(x), i)
		total += len(x)
	}

	fmt.Println("Totally generated numbers : ", total)
	fmt.Println("Number of cpu's : ", numcpu)
}
