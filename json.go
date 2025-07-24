package netescape

import (
	"encoding/json"
	"math/rand"
	"strconv"

	"github.com/zarkones/netescape/lists"
)

type elem struct{ key, val string }

type object []elem

// MarshalJSON converts object type into JSON structure.
func (o object) MarshalJSON() (out []byte, err error) {
	if o == nil {
		return []byte(`null`), nil
	}
	if len(o) == 0 {
		return []byte(`{}`), nil
	}

	out = append(out, '{')
	for _, e := range o {
		key, err := json.Marshal(e.key)
		if err != nil {
			return nil, err
		}
		val, err := json.Marshal(e.val)
		if err != nil {
			return nil, err
		}
		out = append(out, key...)
		out = append(out, ':')
		out = append(out, val...)
		out = append(out, ',')
	}
	// replace last ',' with '}'
	out[len(out)-1] = '}'
	return out, nil
}

func splitString(s string) []string {
	var substrings []string
	for len(s) > 0 {
		chunkSize := rand.Intn(10) + 1
		if chunkSize > len(s) {
			chunkSize = len(s)
		}
		substrings = append(substrings, s[:chunkSize])
		s = s[chunkSize:]
	}
	return substrings
}

func generateJSON(substrings []string, index *int, maxDepth int) (interface{}, int) {
	if maxDepth <= 0 || *index >= len(substrings) {
		if *index < len(substrings) {
			str := substrings[*index]
			*index++
			return str, *index
		}
		return rand.Intn(100), *index
	}

	choice := rand.Intn(4)
	switch choice {
	case 0:
		if *index < len(substrings) {
			str := substrings[*index]
			*index++
			return str, *index
		}
		return rand.Intn(100), *index
	case 1:
		if *index < len(substrings) {
			numStr, err := ToNumbers(&substrings[*index])
			if err == nil {
				if num, err := strconv.Atoi(numStr); err == nil {
					return num, *index
				}
			}
		}
		return rand.Intn(100), *index
	case 2:
		arrLen := rand.Intn(3) + 1
		arr := make([]interface{}, 0, arrLen)
		for i := 0; i < arrLen && *index < len(substrings); i++ {
			var elem interface{}
			elem, *index = generateJSON(substrings, index, maxDepth-1)
			arr = append(arr, elem)
		}
		return arr, *index
	case 3:
		key := lists.Rand(&lists.WordsTop850)
		var val interface{}
		val, *index = generateJSON(substrings, index, maxDepth-1)
		return map[string]interface{}{key: val}, *index
	}
	return nil, *index
}

func collectStrings(val interface{}) string {
	switch v := val.(type) {
	case string:
		return v
	case []interface{}:
		var result string
		for _, elem := range v {
			result += collectStrings(elem)
		}
		return result
	case map[string]interface{}:
		var result string
		for _, value := range v {
			result += collectStrings(value)
		}
		return result
	default:
		return ""
	}
}

func FromJSON(input *string) (output string, err error) {
	var jsonVal interface{}
	err = json.Unmarshal([]byte(*input), &jsonVal)
	if err != nil {
		return "", err
	}
	return collectStrings(jsonVal), nil
}

func ToJSON(input *string) (output string, err error) {
	substrings := splitString(*input)
	index := 0
	maxDepth := 3

	arr := make([]interface{}, 0)
	for index < len(substrings) {
		elem, newIndex := generateJSON(substrings, &index, maxDepth)
		arr = append(arr, elem)
		index = newIndex
	}

	jsonBytes, err := json.Marshal(arr)
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}
