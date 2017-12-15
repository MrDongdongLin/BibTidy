package main

import (
	"fp"
)

func main() {
	pin := "../data/ieee.bib"
	pout := "../out.bib"
	content := fp.BibRead(pin)
	article := fp.EditArticle(content)
	fp.WriteArticle(article, pout)
}
