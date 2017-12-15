// This package includes some struct of BibTeX item and some file process functions.

// TODO: This package is continual update. There are many types of bib item, such as "article", "book", "inbook"
// and so on. All these types can be found through the software "WinEdt" (https://www.baidu.com/link?url=3esIj9SvXmByNPMbHZD51jApdwQUwGfOZy4x1xWw_Ue&wd=&eqid=eac100740000222e000000065a3229a0).
// You can check these items by "Insert" -> "BibTeX Items".

package fp

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

// item types
var (
	rArticle = regexp.MustCompile("@(?i:article)")
	rBook    = regexp.MustCompile("@book")
	rInbook  = regexp.MustCompile("@inbook")
	rProc    = regexp.MustCompile("@proceedings")
	rInporc  = regexp.MustCompile("@inproceedings")
	rInclec  = regexp.MustCompile("@incollection")
	rBooklet = regexp.MustCompile("@booklet")
	rManual  = regexp.MustCompile("@manual")
	rReport  = regexp.MustCompile("@techreport")
	rConfer  = regexp.MustCompile("@conference")
	rPhd     = regexp.MustCompile("@phdthesis")
	rMaster  = regexp.MustCompile("@masterthesis")
	rMisc    = regexp.MustCompile("@misc")
	rUnpub   = regexp.MustCompile("@unpublished")
)

var (
	rAuthor    = regexp.MustCompile("author={.*}")
	rTitle     = regexp.MustCompile("title={.*}")
	rBookTitle = regexp.MustCompile("booktitle={.*}")
	rSeries    = regexp.MustCompile("series={.*}")
	rEditor    = regexp.MustCompile("editor={.*}")
	rPublisher = regexp.MustCompile("publisher={.*}")
	rJournal   = regexp.MustCompile("journal={.*}")
	rYear      = regexp.MustCompile("year={.*}")
	rVolume    = regexp.MustCompile("volume={.*}")
	rNumber    = regexp.MustCompile("number={.*}")
	rPages     = regexp.MustCompile("pages={.*}")
)

var rLast = regexp.MustCompile(",}")

// BibScan read the original bib file and return its content with a string.
/*
 * author:  Dongdong Lin
 * time:    2017-12-15 22:00
 */
func BibScan(pin string, pout string) {
	// Open a bib file
	fin, err := os.Open(pin)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	// Create a bib file
	fout, err := os.Create(pout)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	// Scan the file line by line
	scanner := bufio.NewScanner(fin)
	for scanner.Scan() {
		line := string(scanner.Text())
		if sArticle := rArticle.FindString(line); sArticle != "" {
			fout.WriteString(strings.ToLower(sArticle) + "\r\n")
			for scanner.Scan() {
				line = string(scanner.Text())
				if sAuthor := rAuthor.FindString(line); sAuthor != "" {
					v := strings.Split(sAuthor, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sTitle := rTitle.FindString(line); sTitle != "" {
					v := strings.Split(sTitle, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sJournal := rJournal.FindString(line); sJournal != "" {
					v := strings.Split(sJournal, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sYear := rYear.FindString(line); sYear != "" {
					v := strings.Split(sYear, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sVolume := rVolume.FindString(line); sVolume != "" {
					v := strings.Split(sVolume, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sNumber := rNumber.FindString(line); sNumber != "" {
					v := strings.Split(sNumber, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sPages := rPages.FindString(line); sPages != "" {
					v := strings.Split(sPages, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if slast := rLast.FindString(line); slast != "" {
					fout.WriteString("}\r\n")
				} else {
					continue
				}
			}
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
