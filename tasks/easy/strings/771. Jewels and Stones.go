func numJewelsInStones(jewels string, stones string) int {
	counter := 0
	for _, s1 := range jewels {
		for _, s2 := range stones {
			if s1 == s2 {
				counter++
			}
		}
	}
	return counter
}