package main

import "strconv"

type Solution struct{}

func decode(str string) string {
	shift := 3
	by := []byte(str)
	result := make([]byte, len(by))

	for _, val := range by {
		byt := byte((int(val) - shift + 256) % 256)
		result = append(result, byt)

	}

	return string(result)
}

func encode(str string) string {
	shift := 3
	by := []byte(str)
	result := make([]byte, len(by))

	for _, val := range by {
		byt := byte((int(val) + shift) % 256)
		result = append(result, byt)

	}

	return string(result)
}

func (s *Solution) Encode(strs []string) string {
	var result string

	for _, val := range strs {
		e := encode(val)
		result += strconv.Itoa(len(e)) + "#" + e
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	var result []string
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#' {
			j++
		}
		length, _ := strconv.Atoi(encoded[i:j])
		start := j + 1
		result = append(result, decode(encoded[start:start+length]))
		i = start + length
	}
	return result
}
