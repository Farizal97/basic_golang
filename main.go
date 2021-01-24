package main

import (
"fmt"
"basic_golang/calculation")

func main() {
	fmt.Println("hello world")

	result := calculation.Add(8,9)
	fmt.Println(result)
}
