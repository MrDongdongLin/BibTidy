package fp

import (
	"testing"
)

func Test_bibRead(t *testing.T) {
	pin := "../../data/ieee.bib"
	pout := "../../out.bib"
	content := BibRead(pin)
	article := EditArticle(content)
	WriteArticle(article, pout)
}
