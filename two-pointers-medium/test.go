package main

import "fmt"

func maxAreaTest() {
	// 11. container with most water
	height := []int{1, 2, 1}
	fmt.Println(maxArea(height))
}

func threeSumTest() {
	// 15. 3 sum
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func threeSumClosestTest() {
	// 15. 3 sum
	nums := []int{2, 3, 4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func main() {
	maxAreaTest()
	threeSumTest()
	threeSumClosestTest()
}
