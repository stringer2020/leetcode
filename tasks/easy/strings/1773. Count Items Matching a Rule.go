func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	var searchIdx int

	if ruleKey == "type" {
		searchIdx = 0
	} else if ruleKey == "color" {
		searchIdx = 1
	} else {
		searchIdx = 2
	}

	var searchCoutner int

	for _, arr := range items {
		if ruleValue == arr[searchIdx] {
			searchCoutner++
		}
	}
	return searchCoutner
}