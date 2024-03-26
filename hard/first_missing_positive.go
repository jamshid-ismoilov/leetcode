/*
Given an unsorted integer array nums. Return the smallest positive integer that is not present in nums.

You must implement an algorithm that runs in O(n) time and uses O(1) auxiliary space.

Example 1:

Input: nums = [1,2,0]
Output: 3
Explanation: The numbers in the range [1,2] are all in the array.

Example 2:

Input: nums = [3,4,-1,1]
Output: 2
Explanation: 1 is in the array but 2 is missing.

Example 3:

Input: nums = [7,8,9,11,12]
Output: 1
Explanation: The smallest positive integer 1 is missing.

Constraints:

	1 <= nums.length <= 105
	-231 <= nums[i] <= 231 - 1
*/
package hard

func firstMissingPositive(nums []int) int {
	numsMap := make(map[int]bool)
	for _, element := range nums {
		numsMap[element] = true
	}

	lenNum := len(nums)
	for i := 1; i <= lenNum+1; i++ {
		exists := numsMap[i]
		if exists {
			continue
		} else {
			return i
		}
	}
	if lenNum == 1 {
		return 2
	}
	return 1
}

/*EXPLANATION to solution:
if you want to find number from list of numbers, map is the fastest for this. So I converted slice of nums to map, to make it faster to find number from list.

Since I need to return first missing positive(in other words, smallest unmentioned positive number)
I started with 1 (smallest positive number) and increased it every time until I find the missing one.

*/
