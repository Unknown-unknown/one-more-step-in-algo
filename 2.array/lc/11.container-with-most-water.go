/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water/description/
 *
 * algorithms
 * Medium (46.58%)
 * Likes:    4076
 * Dislikes: 479
 * Total Accepted:    464.6K
 * Total Submissions: 986.2K
 * Testcase Example:  '[1,8,6,2,5,4,8,3,7]'
 *
 * Given n non-negative integers a1, a2, ..., an , where each represents a
 * point at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 * 
 * Note: You may not slant the container and n is at least 2.
 * 
 * 
 * 
 * 
 * 
 * The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In
 * this case, the max area of water (blue section) the container can contain is
 * 49. 
 * 
 * 
 * 
 * Example:
 * 
 * 
 * Input: [1,8,6,2,5,4,8,3,7]
 * Output: 49
 * 
 */

// 1.枚举，left bar x, right bar y, (y-x)*min(x,y)，O(n^2)
// !! 2.左右边界i，j，向中间收敛：左右夹逼，O(n)

// @lc code=start
// v1, 暴力枚举，时间复杂度 O(n^2)
func maxArea(height []int) int {
	max := 0
    for i := 0; i < len(height)-1; i++ {
		for j := i+1; j < len(height); j++ {
			var area int
			if height[j] > height[i] {
				area = (j-i) * height[i]
			} else {
				area = (j-i) * height[j]
			}
			if area > max {
				max = area
			}
		}
	}
	return max
}

// v2, 左右夹逼，时间复杂度 O(n)
func maxArea(height []int) int {	
	max := 0
	for i, j := 0, len(height) - 1; i < j; {
		area := 0
		if height[i] < height[j] {
			area = height[i] * (j - i)
			i++
		} else {
			area = height[j] * (j - i)
			j--
		}
		if area > max {
			max = area
		}
	}
	return max
}

// @lc code=end

