package easy

/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

For example, 2 is written as II in Roman numeral, just two one's added together. 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

    I can be placed before V (5) and X (10) to make 4 and 9.
    X can be placed before L (50) and C (100) to make 40 and 90.
    C can be placed before D (500) and M (1000) to make 400 and 900.

Given an integer, convert it to a roman numeral.



Example 1:

Input: num = 3
Output: "III"
Explanation: 3 is represented as 3 ones.

Example 2:

Input: num = 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.

Example 3:

Input: num = 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.



Constraints:

    1 <= num <= 3999
*/

// my solution
func romanToInt(s string) int {
	var nums []int
	for _, roman := range s {
		switch roman {
		case 'I':
			nums = append(nums, 1)
		case 'V':
			nums = append(nums, 5)
		case 'X':
			nums = append(nums, 10)
		case 'L':
			nums = append(nums, 50)
		case 'C':
			nums = append(nums, 100)
		case 'D':
			nums = append(nums, 500)
		case 'M':
			nums = append(nums, 1000)
		}
	}
	var resp int
	for i := range nums {
		if i == len(nums)-2 {
			if nums[i] < nums[i+1] {
				resp -= nums[i]
			} else {
				resp += nums[i]
			}
			continue
		}
		if i == len(nums)-1 {
			resp += nums[i]
			continue
		}
		if nums[i] < nums[i+1] {
			resp -= nums[i]
		} else {
			resp += nums[i]
		}
	}
	return resp
}

// best response
func romanToInt2(s string) int {
	sum := 0
	temp := 0
	var roman = map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	for _, value := range s {
		current := roman[value]
		sum += current

		if current > temp {
			sum -= temp * 2
		}
		temp = current
	}
	return sum
}
