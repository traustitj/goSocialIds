package socialid

import (
	"fmt"
	"strconv"
	"strings"
)

func Tested() bool {
	fmt.Println("Social Hello World")
	return false
}

// some comment
func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0) || (year%400 == 0) {
		return true
	} else {
		return false
	}
}

// some comment
func NumberOfDaysInMonth(month, year int) int {
	var months = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	if month == 2 && IsLeapYear(year) {
		return 29
	}

	return months[month-1]
}

func GetParity(theid string) (int, bool) {
	theid = strings.Replace(theid, "-", "", -1)
	chars := strings.Split(theid, "")
	numbers := make([]int, 8)
	sums := []int{3, 2, 7, 6, 5, 4, 3, 2}
	parity := 0
	for index, value := range chars {
		numbers[index], _ = strconv.Atoi(value)
		number := (numbers[index] * sums[index])
		parity += number
	}

	parity = parity % 11

	if parity == 0 {
		return 0, true
	}

	parity = 11 - parity

	if parity > 9 {

		return parity, false
	}

	return parity, true
}

func GenerateYear(year int) []string {
	ret := make([]string, 999)
	century := year / 100
	decade := year % 100
	var decadeEnd int
	if century == 19 {
		decadeEnd = 9
	} else {
		decadeEnd = 0
	}

	for m := 1; m < 13; m++ {
		for d := 1; d < NumberOfDaysInMonth(m, year)+1; d++ {
			s1 := fmt.Sprintf("%02d%02d%02d", d, m, decade)
			for i := 20; i < 100; i++ {
				forParity := fmt.Sprintf("%v-%02d", s1, i)
				parity, valid := GetParity(forParity)
				if valid {
					s2 := fmt.Sprintf("%v%d%d", forParity, parity, decadeEnd)
					ret = append(ret, s2)
				}
			}
		}
	}

	return ret
}
