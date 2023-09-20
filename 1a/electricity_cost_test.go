package main

import (
	"reflect"
	"testing"
)

func TestCalculatePowerCost(t *testing.T) {

	solar_house := []int{1, 4}
	connections := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	expected_cost := map[int]int{
		1:0,
		2:1,
		3:2,
		4:0,
		5:1,
	}
	actual_cost := calculatePowerCost(solar_house, connections)

	if !reflect.DeepEqual(actual_cost, expected_cost) {
		t.Errorf("Expected power cost: %v, but got: %v", expected_cost, actual_cost)
	}

	negative_solar_souse := []int{0, 6}
	expected_negative_cost := map[int]int{
		1:0,
		2:0,
		3:0,
		4:0,
		5:0,
	}
	actual_negative_cost := calculatePowerCost(negative_solar_souse, connections)

	if !reflect.DeepEqual(actual_negative_cost, expected_negative_cost) {
		t.Errorf("Expected power cost for invalid solar house: %v, but got: %v", expected_negative_cost, actual_negative_cost)
	}
}
