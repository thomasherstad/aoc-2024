package main

import (
	"fmt"
	"log"
)

func main() {

	//number, err := day1Task1()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Total Difference:", number)

	//similarityScore, err := day1Task2()
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println("Similarity Score:", similarityScore)

	safeReportsAmount, err := findSafeReportsDay2()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Safe reports:", safeReportsAmount)
}
