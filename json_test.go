package netescape

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	payloads := []string{
		"",
		" ",
		"           ",
		"           \t",
		"\t",
		"hello world",
		"asdasd asd asd asf saffsa asf asf as as g er treg er erg weefwqefawe fwef awef awef awef ewa fawefawe fawef wae",
		"drwxr-xr-x    - user 23 Jul 23:49 .git		drwxr-xr-x    - user 23 Jul 23:34 lists.rw-r--r-- 1.9k user 23 Jul 23:41 csv.go",
		"drwxr-xr-x    - user 23 Jul 23:49 .git		drwxr-xr-x    - user 23 Jul \n23:34 lists.rw-r--r-- 1.9k user 23 Jul 23:41 csv.go",
	}

	iterations := 1

	for ri := 0; ri < iterations; ri++ {

		for _, payload := range payloads {
			payload := hex.EncodeToString([]byte(payload))

			serialized, err := ToJSON(&payload)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			deserialized, err := FromJSON(&serialized)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if payload != deserialized {
				t.Log("not the same")
				t.Log(payload)
				t.Log(deserialized)
			}

			fmt.Println(serialized)
		}
	}
}
