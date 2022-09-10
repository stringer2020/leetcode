func interpret(command string) string {
	new := strings.ReplaceAll(command, "()", "o")
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		return "error"
	}
	out := re.ReplaceAllString(new, "")
	return out
}