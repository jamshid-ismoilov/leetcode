package easy

/*
169. Majority Element
Solved
Easy
Topics
Companies

Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times. You may assume that the majority element always exists in the array.



Example 1:

Input: nums = [3,2,3]
Output: 3

Example 2:

Input: nums = [2,2,1,1,1,2,2]
Output: 2



Constraints:

    n == nums.length
    1 <= n <= 5 * 104
    -109 <= nums[i] <= 109

*/

func majorityElement(nums []int) int {
	majorities := make(map[int]int)
	for _, element := range nums {
		majorities[element] += 1
	}
	majority, appears := 0, 0
	for key, value := range majorities {
		if value > appears {
			appears = value
			majority = key
		}
	}
	return majority
}
