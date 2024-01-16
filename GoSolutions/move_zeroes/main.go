package main

func MoveZeroes(nums []int) {
	lZ := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[lZ], nums[i] = nums[i], nums[lZ]
			lZ++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	MoveZeroes(nums)
}
