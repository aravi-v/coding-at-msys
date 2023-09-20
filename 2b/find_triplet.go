package main

import (
	"fmt"
)

func main() {
	sum := 10
	input := []int{1, 5, 3, 2}
	is_success, triplet := FindTriplet(input, sum)
	fmt.Printf("input : array - %v and Target - %d \n",input, sum)
	fmt.Println("Output",is_success, triplet)
}

func FindTriplet(arr []int, sum int) (bool, []int) {
	s_map := make(map[int]struct{})
	res := []int{}
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr); j++ {
			if _, ok := s_map[sum-(arr[i]+arr[j])]; ok {
				res = []int{arr[i], sum - (arr[i] + arr[j]), arr[j]}
				return true, res
			}
			s_map[arr[j]] = struct{}{}
		}

	}
	return false, res
}
