package fp

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// This package includes some struct of BibTeX item and some file process functions.

// TODO: This package is continual update. There are many types of bib item, such as "article", "book", "inbook"
// and so on. All these types can be found through the software "WinEdt" (https://www.baidu.com/link?url=3esIj9SvXmByNPMbHZD51jApdwQUwGfOZy4x1xWw_Ue&wd=&eqid=eac100740000222e000000065a3229a0).
// You can check these items by "Insert" -> "BibTeX Items".

// ArticleBib is the objective struct of type "article" in output xxx.bib file.
type Article struct {
	Author  string
	Title   string
	Journal string
	Year    string
	Volume  string
	Number  string
	Pages   string
}

// InprocBib is the objective struct of type "inproceedings" in output xxx.bib file.
type Inproc struct {
	Author    string
	Title     string
	Booktitle string
	Series    string
	Volume    string
	Year      string
	Pages     string
}

// InBookBib is the objective struct of type "inbook" in output xxx.bib file.
type InBook struct {
	Author    string
	Editor    string
	Title     string
	Publisher string
	Series    string
	Volume    string
	Year      string
	Pages     string
}

// bibRead read the original bib file and return its content with a string.
/*
 * vargin:  path of the original bib, string
 * vargout: content of the bib file, string
 * author:  Dongdong Lin
 * time:    2017-12-14 16:00
 */
func BibRead(path string) string {
	fin, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	content := string(fin)
	return content
}

func EditArticle(str string) Article {
	r, _ := regexp.Compile("author={.*}")
	author := r.FindString(str)
	r, _ = regexp.Compile("title={.*}")
	title := r.FindString(str)
	r, _ = regexp.Compile("journal={.*}")
	journal := r.FindString(str)
	r, _ = regexp.Compile("year={.*}")
	year := r.FindString(str)
	r, _ = regexp.Compile("volume={.*}")
	volume := r.FindString(str)
	r, _ = regexp.Compile("number={.*}")
	number := r.FindString(str)
	r, _ = regexp.Compile("pages={.*}")
	pages := r.FindString(str)

	article := Article{Author: author, Title: title, Journal: journal, Year: year, Volume: volume, Number: number, Pages: pages}

	return article
}

func WriteArticle(article Article, pout string) {
	fout, err := os.Create(pout)
	defer fout.Close()
	if err != nil {
		fmt.Println(pout, err)
	}
	fout.WriteString("@article{ieee,\r\n")
	v := strings.Split(article.Author, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Title, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Journal, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Year, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Volume, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Number, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(article.Pages, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	fout.WriteString("}\r\n\r\n")
}
