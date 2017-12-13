package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// TODO
type OutBib struct {
	Author  string
	Title   string
	Journal string
	Year    string
	Volume  string
	Number  string
	Pages   string
}

func main() {
	s := "data/ieee.bib"
	outbib := editIEEEBib(s)

	outfile := "ieee.bib"
	fout, err := os.Create(outfile)
	defer fout.Close()
	if err != nil {
		fmt.Println(outfile, err)
	}
	fout.WriteString("@article{ieee,\r\n")
	v := strings.Split(outbib.Author, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Title, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Journal, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Year, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Volume, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Number, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	v = strings.Split(outbib.Pages, "=")
	fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
	fout.WriteString("}\r\n\r\n")
}

func editIEEEBib(s string) OutBib {
	fieee, err := ioutil.ReadFile(s)
	if err != nil {
		fmt.Print(err)
	}
	str := string(fieee)
	// fmt.Println(str)
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

	outbib := OutBib{Author: author, Title: title, Journal: journal, Year: year, Volume: volume, Number: number, Pages: pages}

	return outbib
}
