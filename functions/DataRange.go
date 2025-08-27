package function

type room struct {
	name string
	x    int
	y    int
}

type graph struct {
	start string
	end   string
	links map[string][]string
}

func DataRange(data []string) graph{
	var Start string
	var End string

	for i := 0; i < len(data); i++ {
		if data[i] == "##start" && i != len(data)-1 {
			Start = string(data[i+1][0])
		}
		if data[i] == "##end" && i != len(data)-1 {
			End = string(data[i][0])
		}

	}
	return graph{start: Start, end: End}
}
