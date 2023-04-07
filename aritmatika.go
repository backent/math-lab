package mathlab

func Adding(args ...int) int {
	i := 0
	for _, val := range args {
		i = i + val
	}

	return i
}
