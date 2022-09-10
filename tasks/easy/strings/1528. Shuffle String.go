func restoreString(s string, indices []int) string {
	res := make([]byte, len(s))
	for i, v := range []byte(s) {
		res[indices[i]] = v
	}
	return string(res)
}
