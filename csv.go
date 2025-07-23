package netescape

import (
	"bytes"
	"encoding/csv"
	"math/rand"
	"slices"
	"strings"

	"github.com/zarkones/netescape/lists"
)

func ToCsv(input *string) (output string, err error) {
	columns := rand.Intn(16) + 4

	rows := [][]string{}

	headers := []string{}
	for i := 0; i < columns; i++ {
		headers = append(headers, lists.Rand(&lists.WordsTop850))
	}

	rows = append(rows, headers)

	if len(*input) == 0 {
		return strings.Join(rows[0], ","), nil
	}

	row := make([]string, columns)
	rowIndex := 0
	cell := ""
	for i := 0; i < len(*input); i++ {
		cell += string((*input)[i])

		if len(cell) > rand.Intn(30)+1 {
			row[rowIndex] = cell
			cell = ""
			rowIndex++
		}

		if rowIndex == columns {
			rowIndex = 0
			rows = append(rows, row)
			row = make([]string, columns)
		}
	}

	if len(cell) != 0 {
		if rowIndex == columns {
			row := make([]string, columns)
			row[0] = cell
		} else {
			row[rowIndex] = cell
		}
		rows = append(rows, row)
		row = make([]string, columns)
	}

	if slices.Compare(row, make([]string, columns)) != 0 {
		rows = append(rows, row)
	}

	return sliceToCSV(rows)
}

func FromCSV(input *string) (output string, err error) {
	if len(*input) == 0 {
		return "", nil
	}

	rows, err := csvToSlice(*input)
	if err != nil {
		return "", err
	}

	for i, row := range rows {
		if i == 0 {
			continue
		}
		output += strings.Join(row, "")
	}

	return output, nil
}

func csvToSlice(csvData string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(csvData))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	return records, nil
}

func sliceToCSV(data [][]string) (string, error) {
	var buf bytes.Buffer
	writer := csv.NewWriter(&buf)

	for _, row := range data {
		if err := writer.Write(row); err != nil {
			return "", err
		}
	}
	writer.Flush()

	if err := writer.Error(); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}
