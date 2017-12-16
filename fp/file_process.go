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
	rArticle = regexp.MustCompile("@(?i:article){.*")
	rBook    = regexp.MustCompile("@(?i:book){.*")
	rInbook  = regexp.MustCompile("@(?i:inbook){.*")
	rProc    = regexp.MustCompile("@(?i:proceedings){.*")
	rInporc  = regexp.MustCompile("@(?i:inproceedings){.*")
	rInclec  = regexp.MustCompile("@(?i:incollection){.*")
	rBooklet = regexp.MustCompile("@(?i:booklet){.*")
	rManual  = regexp.MustCompile("@(?i:manual){.*")
	rReport  = regexp.MustCompile("@(?i:techreport){.*")
	rConfer  = regexp.MustCompile("@(?i:conference){.*")
	rPhd     = regexp.MustCompile("@(?i:phdthesis){.*")
	rMaster  = regexp.MustCompile("@(?i:masterthesis){.*")
	rMisc    = regexp.MustCompile("@(?i:misc){.*")
	rUnpub   = regexp.MustCompile("@(?i:unpublished){.*")
)

var (
	rAuthor    = regexp.MustCompile("(?i:author)={.*}")
	rTitle     = regexp.MustCompile("(?i:title)={.*}")
	rBookTitle = regexp.MustCompile("(?i:booktitle)={.*}")
	rSeries    = regexp.MustCompile("(?i:series)={.*}")
	rEditor    = regexp.MustCompile("(?i:editor)={.*}")
	rPublisher = regexp.MustCompile("(?i:publisher)={.*}")
	rJournal   = regexp.MustCompile("(?i:journal)={.*}")
	rYear      = regexp.MustCompile("(?i:year)={.*}")
	rVolume    = regexp.MustCompile("(?i:volume)={.*}")
	rNumber    = regexp.MustCompile("(?i:number)={.*}")
	rPages     = regexp.MustCompile("(?i:pages)={.*}")
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
					break
				} else {
					continue
				}
			}
		} else if sInproc := rInporc.FindString(line); sInproc != "" {
			fout.WriteString(strings.ToLower(sInproc) + "\r\n")
			for scanner.Scan() {
				line = string(scanner.Text())
				if sAuthor := rAuthor.FindString(line); sAuthor != "" {
					v := strings.Split(sAuthor, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sBookTitle := rBookTitle.FindString(line); sBookTitle != "" {
					v := strings.Split(sBookTitle, "=")
					fout.WriteString("\t" + v[0] + "\t" + "= " + v[1] + ",\r\n")
				} else if sTitle := rTitle.FindString(line); sTitle != "" {
					v := strings.Split(sTitle, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sYear := rYear.FindString(line); sYear != "" {
					v := strings.Split(sYear, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sVolume := rVolume.FindString(line); sVolume != "" {
					v := strings.Split(sVolume, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sSeries := rSeries.FindString(line); sSeries != "" {
					v := strings.Split(sSeries, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if sPages := rPages.FindString(line); sPages != "" {
					v := strings.Split(sPages, "=")
					fout.WriteString("\t" + v[0] + "\t\t" + "= " + v[1] + ",\r\n")
				} else if slast := rLast.FindString(line); slast != "" {
					fout.WriteString("}\r\n")
					break
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
