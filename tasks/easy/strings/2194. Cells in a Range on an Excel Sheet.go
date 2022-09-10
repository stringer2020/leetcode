func cellsInRange(s string) []string {
	var startPosition int
	var endPosition int
	var characterset byte
	var alphabet []string
	var out []string

	reSymbols, reNumbers := regexp.MustCompile(`[A-z]+`), regexp.MustCompile(`[0-9]+`)
	symbol, number := reSymbols.FindAllString(s, -1), reNumbers.FindAllString(s, -1)

	for characterset = 'A'; characterset <= 'Z'; characterset++ {
		alphabet = append(alphabet, string(characterset))
	}

	endNumer, _ := strconv.Atoi(number[1])

	for i, char := range alphabet {
		if string(symbol[0]) == string(symbol[1]) && string(symbol[0]) == string(char) {
			startPosition, endPosition = i, i
		} else if string(symbol[0]) == string(char) {
			startPosition = i
		} else if string(symbol[1]) == string(char) {
			endPosition = i
		}
	}
	fmt.Println(startPosition, endPosition)
	alphabet = alphabet[startPosition : endPosition+1]
	startNumber, _ := strconv.Atoi(number[0])

	for _, char := range alphabet {
		fmt.Println(startNumber, endNumer)
		for i := startNumber; i <= endNumer; i++ {
			out = append(out, char+strconv.Itoa(i))

		}
	}
	fmt.Println(out)
	return out
}