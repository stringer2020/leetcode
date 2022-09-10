func finalValueAfterOperations(operations []string) int {
	var x int = 0
	for _, operator := range operations {
		switch strings.Contains(operator, "++") {
		case true:
			x++
		case false:
			x--
		}
	}
	return x
}