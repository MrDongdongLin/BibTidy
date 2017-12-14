package fp

import (
	"testing"
)

func Test_bibRead(t *testing.T) {
	pin := "../../data/ieee.bib"
	pout := "../../out.bib"
	content := bibRead(pin)
	article := editArticle(content)
	writeArticle(article, pout)
}
