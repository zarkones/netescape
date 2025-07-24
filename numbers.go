package netescape

import (
	"encoding/hex"
	"errors"

	"github.com/zarkones/netescape/lists"
)

var (
	ErrInvalidHexChar = errors.New("invalid hex character")
)

var hexToNumbersMap = map[string][]string{
	"0": {"10", "11", "12", "13"},
	"1": {"14", "15", "16", "17"},
	"2": {"18", "19", "20", "21"},
	"3": {"22", "23", "24", "25"},
	"4": {"26", "27", "28", "29"},
	"5": {"30", "31", "32", "33"},
	"6": {"34", "35", "36", "37"},
	"7": {"38", "39", "40", "41"},
	"8": {"42", "43", "44", "45"},
	"9": {"46", "47", "48", "49"},
	"a": {"50", "51", "52", "53"},
	"b": {"54", "55", "56", "57"},
	"c": {"58", "59", "60", "61"},
	"d": {"62", "63", "64", "65"},
	"e": {"66", "67", "68", "69"},
	"f": {"70", "71", "72", "73"},
}

var numbersToHexMap = func() (m map[string]string) {
	m = map[string]string{}
	for hexN, obfNs := range hexToNumbersMap {
		for _, obfN := range obfNs {
			m[string(obfN)] = hexN
		}
	}
	return m
}()

func ToNumbers(data *string) (output string, err error) {
	hexEncoded := hex.EncodeToString([]byte(*data))

	for i := 0; i < len(hexEncoded); i++ {
		c := string((hexEncoded)[i])
		chars := hexToNumbersMap[c]
		output += lists.Rand(&chars)
	}

	return output, nil
}

func FromNumbers(input *string) (output string, err error) {
	if len(*input) == 0 {
		return "", nil
	}

	hexEncoded := ""

	n := ""
	for i := 0; i < len(*input); i++ {
		n += string((*input)[i])
		if len(n) == 2 {
			hN, ok := numbersToHexMap[n]
			if !ok {
				return "", ErrInvalidHexChar
			}
			hexEncoded += hN
			n = ""
		}
	}

	hexDecoded, err := hex.DecodeString(hexEncoded)
	if err != nil {
		return "", err
	}

	return string(hexDecoded), nil
}
