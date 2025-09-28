package lemin

import (
	"sort"
)

func BestGroup(allPathGroups [][][]room) {
	var allGroupTurns []int
	for _, group := range allPathGroups {
		sort.Slice(group, func(i, j int) bool {
			return len(group[i]) < len(group[j])
		})
	}
	for _,group :=range allPathGroups {
		var turnslice,sliceNUmber []int
		var turn int
		antNumber := colony.antNumber
		for _,path := range group {
			number := len(group[len(group)-1]) - len(path)
			antNumber -= number
			sliceNUmber = append(sliceNUmber, number)
		}
		
		antsPerPath := antNumber / len(group)
		overflowAnts := antNumber % len(group)

		for i,path := range group {
			
			turn = len(path)+sliceNUmber[i]+antsPerPath-1
			if i < overflowAnts {
				turn++
			}
			turnslice = append(turnslice, turn)
		}
		sort.Ints(turnslice)
		turn = turnslice[len(turnslice)-1]
		allGroupTurns = append(allGroupTurns, turn-1)
	}

	var index int
	var compare int = allGroupTurns[0]
	for i, turn := range allGroupTurns {
		if turn < compare {
			index = i
			compare = turn
		}
	}

	Print(allPathGroups[index])
}