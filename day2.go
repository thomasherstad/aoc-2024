package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInputListDay2() ([][]int, error) {
	data, err := os.ReadFile("input_day2.txt")
	if err != nil {
		return [][]int{}, err
	}

	dataList := strings.Split(string(data), "\n")

	var reports [][]int
	for _, str := range dataList {
		strs := strings.Split(str, " ")
		var currentList []int
		for _, valStr := range strs {
			if valStr == "" {
				continue
			}
			val, err := strconv.Atoi(valStr)
			if err != nil {
				return [][]int{}, err
			}
			currentList = append(currentList, val)
		}
		if len(currentList) == 0 {
			continue
		}
		reports = append(reports, currentList)
	}

	return reports, nil
}

func findSafeReportsDay2() (int, error) {
	reports, err := getInputListDay2()
	if err != nil {
		return 0, nil
	}

	for i, report := range reports {
		fmt.Printf("Report %v: %v\n", i, report)
	}

	return 0, err
}
