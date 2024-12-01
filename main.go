package main

import "fmt"

func main() {

	number, err := day1Task1()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(number)
}
