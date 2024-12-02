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

func isSafe(report []int) bool {
	descSafe, ascSafe := false, false
	for i := 0; i < len(report)-1; i++ {
		diff := report[i+1] - report[i]
		if diff >= 1 && diff <= 3 && !descSafe {
			ascSafe = true
		} else if diff >= -3 && diff <= -1 && !ascSafe {
			descSafe = true
		} else {
			return false
		}
	}

	return ascSafe || descSafe
}

// This gives the correct answer = 432
func findSafeReportsDay2Task1() (int, error) {
	reports, err := getInputListDay2()
	if err != nil {
		return 0, nil
	}
	totalSafeReports := 0
	for i, report := range reports {
		safe := isSafe(report)
		fmt.Printf("Report %v: %v - %v\n", i, report, safe)
		if safe {
			totalSafeReports++
		}
	}

	return totalSafeReports, err
}
