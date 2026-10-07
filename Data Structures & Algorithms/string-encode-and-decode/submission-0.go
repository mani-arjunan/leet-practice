type Solution struct{}

func decode(str string) string {
	shift := 3
	var result string

	for _, val := range str {
		str := byte((int(val) - shift + 256) % 256)
		result += string(str)
	}

	return result
}

func encode(str string) string {
	shift := 3
	var result string

	for _, val := range str {
		str := byte((int(val) + shift) % 256)
		result += string(str)
	}

	return result
}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var result string

	for _, val := range strs {
		str := "#" + encode(val) + "#"
		result += str
	}

	return result
}

func (s *Solution) Decode(encoded string) []string {
	if encoded == "" {
		return []string{}
	}
	var result []string

	for i := 1; i < len(encoded); i++ {
		var str string
		for string(encoded[i]) != "#" {
			str += string(encoded[i])
			i++
		}
		de := decode(str)
		result = append(result, de)
		str = ""
		i++
	}

	return result
}
