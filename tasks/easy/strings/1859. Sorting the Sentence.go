func sortSentence(s string) string {
	var sortedWords []string

	data := make(map[int]string)
	words := strings.Split(s, " ")
	regex := regexp.MustCompile(`[0-9]+`)

	for _, word := range words {
		index, _ := strconv.Atoi(regex.FindString(word))
		data[index] = regex.ReplaceAllString(word, "")
	}

	counter := 0

	for i := 1; i <= len(data); i++ {
		counter++
		sortedWords = append(sortedWords, data[counter])
	}

	return strings.Join(sortedWords, " ")
}