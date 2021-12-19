package main

import "fmt"

func testMaxArea() {
	// 11. container with most water
	height := []int{1, 2, 1}
	fmt.Println(maxArea(height))
}

func testThreeSum() {
	// 15. 3 sum
	nums := []int{0, 0, 0, 0}
	fmt.Println(threeSum(nums))
}

func testThreeSumClosest() {
	// 16. 3 sum closest
	nums := []int{2, 3, 4}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}

func main() {
	testMaxArea()
	testThreeSum()
	testThreeSumClosest()
}
