func balancedStringSplit(s string) int {
	var balanced int
	var counter int
	for _, symbol := range s {
		fromRune := string(symbol)
		if fromRune == "L" {
			balanced++
		} else {
			balanced--
		}
		if balanced == 0 {
			counter++
		}
	}
	fmt.Println(counter)
	return counter
}