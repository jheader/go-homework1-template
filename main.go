package main

import "fmt"

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++ {
		if nums[slow] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}

	}
	return slow + 1
}

func main() {

	a := []int{1, 2, 2, 3, 3, 3}
	le := RemoveDuplicates(a)

	fmt.Println(le)
	fmt.Println(a[:])

}
