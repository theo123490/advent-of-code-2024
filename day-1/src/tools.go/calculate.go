package tools

import (
	"slices"
)

func CalculateDistance(listA []int, listB []int) int {
	totalDistance := 0
	slices.Sort(listA)
	slices.Sort(listB)
	for i, _ := range listA {
		distance := listA[i] - listB[i]
		if distance < 0 {
			distance = distance * -1
		}
		totalDistance += distance
	}
	return totalDistance
}

func CalculateOccurance(listA []int, listB []int) int {
	totalOccurance := 0
	occuranceMap := map[int]int{}
	for _, elementA := range listA {
		if _, ok := occuranceMap[elementA]; !ok {
			counter := 0
			for _, elementB := range listB {
				if elementA == elementB {
					counter += 1
				}
			}
			occuranceMap[elementA] = counter * elementA
		}
		totalOccurance += occuranceMap[elementA]
	}
	return totalOccurance
}
