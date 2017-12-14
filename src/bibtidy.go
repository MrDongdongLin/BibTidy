package main

import (
	fp "github.com/fp"
)

func main() {
	pin := "../../data/ieee.bib"
	pout := "../../out.bib"
	content := fp.bibRead(pin)
	article := fp.editArticle(content)
	fp.writeArticle(article, pout)
}
