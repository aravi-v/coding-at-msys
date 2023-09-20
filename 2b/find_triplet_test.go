package main

import (
	"reflect"
	"testing"
)

func TestFindTriplet(t *testing.T) {
	tests := []struct {
		arr  []int
		sum  int
		want_bool  bool
		want []int
	}{
		{[]int{5, 3, 10, 1, 15, 4, 8}, 10,true, []int{5,1,4}},      
		{[]int{1, 5, 3, 2}, 10,true, []int{5,2,3}},
		{[]int{5, 3, 10, 1, 15, 4, 8}, 25,false, []int{}}, 
		{[]int{10, 20, 30, 40}, 70,true, []int{10, 20, 40}},   
	}

	for i, tt := range tests {
		gotSuccess, gotTriplet := FindTriplet(tt.arr, tt.sum)

		if tt.want_bool != gotSuccess || !reflect.DeepEqual(gotTriplet, tt.want) {
			t.Errorf("Test %d - Failed. Expected triplet: %v and Got triplet %v | Expected status: %t and Got Status: %t",i+1,tt.want,gotTriplet,tt.want_bool,gotSuccess)
		}
	}
}
