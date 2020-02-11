package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	m := make(map[int][][]int)
	tempMap := make(map[int]map[int]int)
	for k1, num1 := range nums  {
		for k2, num2 := range nums {
			if tempMap[k2 + k1] == nil{
				tempMap[k2 + k1] = make(map[int]int)
			}
			if tempMap[k2 + k1][k2] == k1 || tempMap[k1 + k2][k1] == k2 {continue}
			if k1 == k2 {continue}
			m[num1 + num2] = append(m[num1 + num2], []int{num1, num2})
			tempMap[k1 + k2][k2] = k1
			tempMap[k1 + k2][k1] = k2
		}
	}
	tempMap2 := make(map[int]int)
	for _, num := range nums  {
		sub := 0 - num
		if _, ok := tempMap2[num]; ok {continue}
		if _, ok := m[sub]; ok{
			for _, twoNums := range m[sub]{
				threeNums := append(twoNums, num)
				result = append(result, threeNums)
			}
			tempMap2[num] = 1
		}
	}
	return result
}

func main(){
	nums := []int{-1,0,1,2,-1,-4}
	fmt.Println(threeSum(nums))
}