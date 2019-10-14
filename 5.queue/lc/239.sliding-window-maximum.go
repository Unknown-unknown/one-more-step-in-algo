/*
 * @lc app=leetcode id=239 lang=golang
 *
 * [239] Sliding Window Maximum
 *
 * https://leetcode.com/problems/sliding-window-maximum/description/
 *
 * algorithms
 * Hard (39.23%)
 * Likes:    2072
 * Dislikes: 121
 * Total Accepted:    184.1K
 * Total Submissions: 468.4K
 * Testcase Example:  '[1,3,-1,-3,5,3,6,7]\n3'
 *
 * Given an array nums, there is a sliding window of size k which is moving
 * from the very left of the array to the very right. You can only see the k
 * numbers in the window. Each time the sliding window moves right by one
 * position. Return the max sliding window.
 * 
 * Example:
 * 
 * 
 * Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
 * Output: [3,3,5,5,6,7] 
 * Explanation: 
 * 
 * Window position                Max
 * ---------------               -----
 * [1  3  -1] -3  5  3  6  7      3
 * ⁠1 [3  -1  -3] 5  3  6  7       3
 * ⁠1  3 [-1  -3  5] 3  6  7       5
 * ⁠1  3  -1 [-3  5  3] 6  7       5
 * ⁠1  3  -1  -3 [5  3  6] 7       6
 * ⁠1  3  -1  -3  5 [3  6  7]      7
 * 
 * 
 * Note: 
 * You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty
 * array.
 * 
 * Follow up:
 * Could you solve it in linear time?
 */

// v1, brute force
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) * k == 0 {
		return []int{}
	}
	maxes := []int{}
	for i := 0; i < len(nums)-k+1; i++ {
		max := MinInt
		for j := i; j < i+k; j++ {
			if nums[j] > max {
				max = nums[j]
			}
		}
		maxes = append(maxes, max)
	}
	return maxes
}

// v2, deque(double ended queue)
// todo: didn't understand it...
func maxSlidingWindow(nums []int, k int) []int {
	result := make([]int, 0, len(nums)-k)
	window := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if len(window) > 0 && window[0] <= i-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < nums[i] {
			window = window[:len(window)-1]
		}
		window = append(window, i)
		if i >= k-1 {
			result = append(result, nums[window[0]])
		}
	}
	return result
}
