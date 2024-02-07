package main

import "fmt"

func InsertionSort(nums []int) []int {
	var temp, j int
	for i := 1; i < len(nums); i++ {
		j = i
		for j > 0 && nums[j] < nums[j-1] {
			temp = nums[j]
			nums[j] = nums[j-1]
			nums[j-1] = temp
			j = j - 1
		}
	}
	return nums
}

func main() {
	slice := []int{8, 5, 4, 3, 2}
	sorted := InsertionSort(slice)

	fmt.Println(sorted)
}
