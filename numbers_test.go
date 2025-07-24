package netescape

import (
	"testing"
)

func TestNumbers(t *testing.T) {
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

	iterations := 1000

	for ri := 0; ri < iterations; ri++ {

		for _, payload := range payloads {
			numbers, err := ToNumbers(&payload)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			plaintext, err := FromNumbers(&numbers)
			if err != nil {
				t.Log(err)
				t.FailNow()
			}

			if plaintext != payload {
				t.Log("plaintext and payload does not match")
				t.Log(plaintext)
				t.Log(payload)
				t.FailNow()
			}
		}
	}
}
