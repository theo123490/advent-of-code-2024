package tools

import (
	"fmt"
)

func IsStable(graphData []int) bool {
	if graphData[1] > graphData[0] {
		for i := 1; i < len(graphData); i++ {
			dif := graphData[i] - graphData[i-1]
			if dif > 3 || dif < 1 {
				return false
			}

		}
		return true
	}

	if graphData[1] < graphData[0] {
		for i := 1; i < len(graphData); i++ {
			dif := graphData[i-1] - graphData[i]
			if dif > 3 || dif < 1 {
				return false
			}
		}
		return true
	}

	if graphData[1] == graphData[0] {
		return false
	}
	panic(fmt.Sprintf("Invalid input: %v", graphData))
}

func IsStableWithDampener(graphData []int) bool {
	if IsStable(graphData) {
		return true
	} else {
		for n := 0; n < len(graphData); n++ { //should optimize this but I misunderstood the problem for too long
			newGraphData := popSlice(graphData, n)
			if IsStable(newGraphData) {
				return true
			}
		}
		return false
	}
}

func popSlice(slice []int, removeIndex int) []int {
	newGraph := make([]int, len(slice)-1)
	modifier := 0
	for i := 0; i < len(slice); i++ {
		if i == removeIndex {
			modifier = 1
			continue
		}
		newGraph[i-modifier] = slice[i]
	}
	return newGraph
}
