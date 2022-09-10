func mostWordsFound(sentences []string) int {
	var sLenght []int

	for _, s := range sentences {
		sReplace := strings.Split(strings.ReplaceAll(s, " ", ","), ",")
		sLenght = append(sLenght, len(sReplace))
	}

	max := sLenght[0]
	min := sLenght[0]

	for _, len := range sLenght {
		if len < min {
			min = len
		}
		if len > max {
			max = len
		}
	}
	return max
}