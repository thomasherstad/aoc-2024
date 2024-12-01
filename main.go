package main

import "fmt"

func main() {

	number, err := day1Task1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Total Difference:", number)

	similarityScore, err := day1Task2()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Similarity Score:", similarityScore)
}
