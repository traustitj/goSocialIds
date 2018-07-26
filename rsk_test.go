package main

import (
	"fmt"
	"testing"

	"./socialid"
)

func TestFirst(t *testing.T) {
	socialid.Tested()
}

func TestLeapYears(t *testing.T) {
	ret := false
	troubleYears := []int{1700, 1800, 1900, 2100, 2200, 2300, 2500, 2600}
	goodYears := []int{1600, 2000, 2400, 1988}
	for _, element := range troubleYears {
		ret = socialid.IsLeapYear(element)
		if ret == true {
			t.Errorf("Element %d is a leap year", element)
		}
	}

	for _, element := range goodYears {
		ret = socialid.IsLeapYear(element)
		if ret == false {
			t.Errorf("Element %d is not a leap year", element)
		}
	}
}

func TestLeapYearDate(t *testing.T) {
	i := socialid.NumberOfDaysInMonth(2, 1700)
	if i != 28 {
		t.Errorf("February 1700 should have 28 days, but is %v", i)
	}
	i = socialid.NumberOfDaysInMonth(2, 1600)
	if i != 29 {
		t.Errorf("February 1600 should have 29 days but is %v", i)
	}
}

func TestPaddingOfNumber(t *testing.T) {
	fmt.Printf("Should be printed as 02 : %02d", 2)
	fmt.Printf("Should be printed as 10 : %02d", 10)
}

func TestGenerateSocialIds(t *testing.T) {
	s := socialid.GenerateYear(1974)
	found := false
	fmt.Printf("s is of length : %v", len(s))
	for _, i := range s {
		if i == "030574-3039" {
			found = true
		}
	}

	if !found {
		t.Errorf("030574-3039 is not found")
	}
}

func TestParity(t *testing.T) {
	parity, _ := socialid.GetParity("12016033")

	if parity != 8 {
		t.Errorf("Parity for 12016033 should be 8 but is %d", parity)
	}

	parity, _ = socialid.GetParity("03057430")
	if parity != 3 {
		t.Errorf("Parity for 03057430 should be 3 but is %d", parity)
	}
	parity, valid := socialid.GetParity("311200-95")
	if valid {
		t.Errorf("Parity for 311200-95 should not be valid %d", parity)
	}
}

func TestYear(t *testing.T) {
	year := socialid.GenerateYear(2000)
	for _, i := range year {
		fmt.Println(i)
	}
	fmt.Println("========================")
	fmt.Printf("I generated %d social id numbers for year 2000", len(year))
}
