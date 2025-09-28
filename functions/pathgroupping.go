package lemin

func GrouppingPaths(allFoundPaths [][]room) {
	groupedPathsSets := [][][]room{}
	for i := 0; i < len(allFoundPaths); i++ {
		currentGroup := [][]room{allFoundPaths[i]}
		for j := 0; j < len(allFoundPaths); j++ {
			if i == j  {
				continue
			}
			hasOverlap := false
			for _, groupedPath := range currentGroup {
				if overpLap(groupedPath, allFoundPaths[j]) {
					hasOverlap = true
					break
				}
			}
			if !hasOverlap {
				currentGroup = append(currentGroup, allFoundPaths[j])
			}
		}
		groupedPathsSets = append(groupedPathsSets, currentGroup)

	}
	BestGroup(groupedPathsSets)
}

func overpLap(path1, path2 []room) bool {
	for _, room1 := range path1[1 : len(path1)-1] {
		for _, room2 := range path2[1 : len(path2)-1] {
			if room1 == room2 {
				return true
			}
		}
	}
	return false
}