func truncateSentence(s string, k int) string {
	newS := strings.Split(s, " ")
	return strings.Join(newS[:k], " ")
}