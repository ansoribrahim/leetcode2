package strings

func longestCommonPrefix(strs []string) string {
	result := strs[0]
	for _, str := range strs[1:] {
		i := 0
		for i < len(result) && i < len(str) && result[i] == str[i] {
			i++
		}
		result = result[:i]
	}
	return result
}
