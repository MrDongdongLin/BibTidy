package main

import (
	"fp"
)

func main() {
	// pin := flag.String("i", "", "Path of the input bib file.")
	// pout := flag.String("o", "/", "Path of the output bib file.")
	pin := "../../data/ieee.bib"
	pout := "../../data/out.bib"
	fp.BibScan(pin, pout)
}
