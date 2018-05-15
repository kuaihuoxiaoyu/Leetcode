package leetcode

/*
1.给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。

你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。

示例:
给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	length := len(nums)
	for n := 0; n < length; n++ {
		m[n] = nums[n]
	}

	i := 0
	j := 0
	diff := 0
LABEL1:
	for i < length {
		diff = target - nums[i]

		for k, v := range m {
			if v == diff {
				j = k
				break LABEL1
			}
		}
		i++
	}

	return []int{i, j}
}
