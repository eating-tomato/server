package main

import "sort"

func threeSum2(nums []int) [][]int{
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := range nums {
		left, right := i + 1, len(nums) - 1

		// 如果最小的也是大于0的或者最大的也是小于0的 就没必要继续比下去了
		if nums[i] > 0 || nums[right] < 0 {
			return res
		}

		// 如果和上一个一样 就会出现重复 就要跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for left < right {
			if nums[left] + nums[right] + nums[i] > 0 {
				right--
			} else if nums[left] + nums[right] + nums[i] < 0 {
				left++
			} else {
				res = append(res, []int{nums[i], nums[left], nums[right]})

				// left边和right边可能有很多重复的 如果重复使用结果中也会出现重复的，所以使指针一直移动到重复部分的末端
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}
		}


	}
	return res
}
