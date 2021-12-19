package main

import "fmt"

func testRemoveDuplicates() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	k := removeDuplicates(nums) // Calls your implementation
	fmt.Println(k, nums)
}

func main() {
	testRemoveDuplicates()
}
