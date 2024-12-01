package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

func getInputLists() ([]int, []int, error) {
	// import text file
	data, err := os.ReadFile("input_day1.txt")
	if err != nil {
		return []int{}, []int{}, err
	}

	list := strings.Split(string(data), "\n")

	leftList := make([]int, len(list))
	rightList := make([]int, len(list))
	for i, values := range list {
		if len(values) == 0 {
			continue
		}
		strings := strings.Split(values, " ")
		if len(strings) > 2 {
			return []int{}, []int{}, errors.New("split didn't work properly")
		}

		//fmt.Println(strings[0], ",", strings[1])

		leftList[i], err = strconv.Atoi(strings[0])
		if err != nil {
			return []int{}, []int{}, err
		}
		rightList[i], err = strconv.Atoi(strings[1])
		if err != nil {
			return []int{}, []int{}, err
		}

	}

	return leftList, rightList, nil
}

func day1Task1() (int, error) {

	leftList, rightList, err := getInputLists()
	if err != nil {
		return 0, err
	}

	totalDiff := 0

	diff := make([]int, 0)
	slices.Sort(leftList)
	slices.Sort(rightList)

	for j := range leftList {
		currentDiff := leftList[j] - rightList[j]
		if currentDiff < 0 {
			currentDiff = -currentDiff
		}
		diff = append(diff, currentDiff)
		totalDiff += currentDiff
		//fmt.Printf("%v, %v, diff=%v, total: %v\n", leftList[j], rightList[j], currentDiff, totalDiff)
	}

	return totalDiff, nil
}

//Day 1 Task 1: That gave the correct answer, 2164381

func day1Task2() (int, error) {
	leftList, rightList, err := getInputLists()
	if err != nil {
		return 0, err
	}

	rightFrequencyMap := make(map[int]int)

	for _, rightVal := range rightList {
		if _, ok := rightFrequencyMap[rightVal]; !ok {
			rightFrequencyMap[rightVal] = 1
		} else {
			rightFrequencyMap[rightVal]++
		}
	}

	for key, val := range rightFrequencyMap {
		fmt.Println(key, ":", val)
	}

	totalSimilarityScore := 0
	for _, leftVal := range leftList {
		val, ok := rightFrequencyMap[leftVal]
		if ok {
			totalSimilarityScore += leftVal * val
		}
	}
	return totalSimilarityScore, nil
}

//Day 1 Task 2: That gave the correct answer, 20719933
