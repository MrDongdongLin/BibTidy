package fp

import (
	"testing"
)

func Test_bibRead(t *testing.T) {
	pin := "../../data/ieee.bib"
	pout := "../../data/out.bib"
	BibScan(pin, pout)

	// fin, err := os.Open(pin)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer fin.Close()

	// scanner := bufio.NewScanner(fin)
	// for scanner.Scan() {
	// 	line := string(scanner.Text())
	// 	v := rLast.FindString(line)
	// 	if v != "" {
	// 		fmt.Println(v)
	// 	}
	// }
}
