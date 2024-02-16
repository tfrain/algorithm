func titleToNumber(columnTitle string) int {
	title := []byte(columnTitle)
	res, posNum := 0, 1
	for i := len(title) - 1; i >= 0; i-- {
		res += int(title[i]-'A'+1) * posNum
		posNum *= 26
	}
	return res
}

func titleToNumber(s string) int {
	var num int
	for i := 0; i < len(s); i++ {
		num = num*26 + int(s[i]-'A')+1
	}
	return num
}