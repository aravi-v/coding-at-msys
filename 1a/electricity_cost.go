package main

import (
	"fmt"
)

func calculatePowerCost(solar_house []int, connections [][]int) map[int]int {
	adjacent_node := make(map[int][]int)
	power_cost_map := make(map[int]int)
	for _, connection := range connections {
		house1, house2 := connection[0], connection[1]
		adjacent_node[house1] = append(adjacent_node[house1], house2)
		adjacent_node[house2] = append(adjacent_node[house2], house1)
		power_cost_map[house1],power_cost_map[house2] = 0,0
	}

	queue := solar_house
	visited := make(map[int]bool)

	for len(queue) > 0 {
		curr_house := queue[0]
		queue = queue[1:]

		visited[curr_house] = true

		for _, neighbor := range adjacent_node[curr_house] {
			if !visited[neighbor] {
				power_cost_map[neighbor] = power_cost_map[curr_house] + 1
				queue = append(queue, neighbor)
			}
		}
	}
		
	return power_cost_map
}

func main() {
	solar_house := []int{1, 4}
	connections := [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}}
	power_cost := calculatePowerCost(solar_house, connections)
	fmt.Printf("Input : Solar Houses - %d and Connections : %d \n",solar_house,connections)
	fmt.Println("Output", power_cost)
}
