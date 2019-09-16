/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */
 import "sort"
 func threeSum(nums []int) [][]int {
	 if len(nums) < 3 {
		 return [][]int{}
	 }
	 sort.Ints(nums)
	 if nums[0] > 0 || nums[len(nums)-1] < 0 {
		 return [][]int{}
	 }
	 res := [][]int{}
	 for i := 0; i < len(nums); i++ {
		 if nums[i] > 0 {
			 break
		 }
		 if i > 0 && nums[i] == nums[i-1] {	// !!去重
			 continue
		 }
		 l := i+1
		 r := len(nums)-1
		 for l < r {
			 sum := nums[i] + nums[l] + nums[r]
			 if sum == 0 {
				 res = append(res, []int{nums[i], nums[l], nums[r]})
				 for l < r && nums[l] == nums[l+1]	{	// !!去重
					 l++
				 }
				 for l < r && nums[r] == nums[r-1] {	// !!去重
					 r--
				 }
				 l++
				 r--
			 } else if sum < 0 {
				 l++
			 } else {
				 r--
			 }
		 }
	 }
	 return res
 }

