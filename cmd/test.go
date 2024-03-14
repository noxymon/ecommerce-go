package main

import "fmt"

func main() {
	testdata := map[string]int{
		"test":  1,
		"test2": 2,
		"test3": 3,
		"test4": 4,
		"test5": 5,
	}
	test2 := map[string]int{
		"test6": 6,
		"test7": 7,
	}

	testdata = test2
	fmt.Print(testdata)
}
