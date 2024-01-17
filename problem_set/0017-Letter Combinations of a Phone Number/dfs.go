func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	var res []string
	n := len(digits)
	letters := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var bf func(start int, sub []byte)
	bf = func(start int, sub []byte) {
		if start == n {
			res = append(res, string(sub))
			return
		}
		letter := letters[digits[start]-'2']
		for i := range letter {
			bf(start+1, append(sub, letter[i]))
		}
	}
	bf(0, nil)
	return res
}
