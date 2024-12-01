package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func day1Task1() (int, error) {
	// import text file
	data, err := os.ReadFile("input_day1.txt")
	if err != nil {
		return 0, err
	}

	list := strings.Split(string(data), "\n")

	leftList := make([]int, len(list))
	rightList := make([]int, len(list))
	totalDiff := 0
	for i, values := range list {
		if len(values) == 0 {
			continue
		}
		strings := strings.Split(values, " ")
		if len(strings) > 2 {
			return 0, errors.New("split didn't work properly")
		}

		//fmt.Println(strings[0], ",", strings[1])

		leftList[i], err = strconv.Atoi(strings[0])
		if err != nil {
			return 0, err
		}
		rightList[i], err = strconv.Atoi(strings[1])
		if err != nil {
			return 0, err
		}

	}

	diff := make([]int, len(list))
	slices.Sort(leftList)
	slices.Sort(rightList)

	for j := range leftList {
		currentDiff := leftList[j] - rightList[j]
		if currentDiff < 0 {
			currentDiff = -currentDiff
		}
		diff[j] = currentDiff
		totalDiff += currentDiff
		fmt.Printf("%v, %v, diff=%v, total: %v\n", leftList[j], rightList[j], currentDiff, totalDiff)
	}

	return totalDiff, nil
}

//That was the correct answer, 2164381
