#pragma once
#include "Article.h"
#include "Inproceedings.h"

class ParseBib
{
    private:
       string sourceBibFile;
       string tidyedBibFile;
       vector <Object *> obj;
       map<string, string>b;   //ÓÃÓÚÓ³ÉäÌæ»»Journal

    private:
       void readBibFile();
       void writeBibFile();
       void getMapReplace();
       void printItem(string, string);
       bool isExist(const string&);

    public:
       ParseBib(string _sourceBibFile, string _tidyedBibFile) {
          sourceBibFile = _sourceBibFile;
          tidyedBibFile = _tidyedBibFile;
          //readBibFile();
       }

       void setSourceBibFile(string _sourceBibFile) {
            sourceBibFile = _sourceBibFile;
       }

       void setTidyedBibFile(string _tidyedBibFile) {
            tidyedBibFile = _tidyedBibFile;
       }

       void saveTidyedBibFile();
};

void ParseBib::readBibFile() {
    freopen(sourceBibFile.c_str(), "r", stdin);
    char line[200];

    while (NULL != gets(line)) {

       if (line[0]!='@') continue;
       BaseInf *baseInf;
       if (line[1] == 'a') {
          baseInf = new Article();
       } else
       if (line[1] == 'i') {
          baseInf = new Inproceedings();
          if (line[3] == 'c') {
             baseInf->setType("incollection");
          }
       }
       string str = "";
       while (NULL != gets(line)) {
             int len = strlen(line);
             int idx = 0;
             while (' ' == line[len-1]) {
                   line[len-1] = '\0';
                   len--;
             }
             while (' ' == line[idx]) {
                   idx ++;
             }
             if (strlen(line+idx) == 1 && '}' == line[idx]) {
                break;
             }
             str += string(line+idx);
             if ('}' == line[len-2] && ',' == line[len-1]) {

                if ( (str[0] == 'A' || str[0] == 'a') && str[1] == 'u' ) { //author
                   baseInf->setAuthor(str);
                } else
                if ( (str[0] == 'T' || str[0] == 't') && (str[1] == 'i') ) { //title
                   baseInf->setTitle(str);
                } else
                if ( baseInf->getType() == "article" && (str[0] == 'V' || str[0] == 'v') && (str[1] == 'o') ) { //Volume
                   Article *art = static_cast<Article*>(baseInf);
                   art->setVolume(str);
                } else
                if ( (str[0] == 'P' || str[0] == 'p') && (str[1] == 'a') ) { //pages
                   baseInf->setPages(str);
                } else
                if ( (str[0] == 'Y' || str[0] == 'y') && (str[1] == 'e') ) { //year
                   baseInf->setYear(str);
                } else
                if ( (str[0] == 'M' || str[0] == 'm') && (str[1] == 'o') ) { //month
                   baseInf->setMonth(str);
                } else
                if ( (str[0] == 'D' || str[0] == 'd') && (str[1] == 'O' || str[1] == 'o') ) { //doi
                   baseInf->setDoi(str);
                } else
                if ( (str[0] == 'J' || str[0] == 'j') && (str[1] == 'o') ) { //journal
                   Article *art = static_cast<Article*>(baseInf);
                   art->setJournal(str);
                } else
                if ( baseInf->getType() == "article" && (str[0] == 'N' || str[0] == 'n') && (str[1] == 'u') ) { //number
                   Article *art = static_cast<Article*>(baseInf);
                   art->setNumber(str);
                } else
                if ( (str[0] == 'A' || str[0] == 'a') && (str[1] == 'r') ) { //Article-number
                   str.insert(str.find("{{")+2, "Article number ");
                   //freopen("CON", "w", stdout);
                   //printf("%s\n", str.c_str());
                   baseInf->setPages(str);
                } else
                if ( (str[0] == 'B' || str[0] == 'b') && (str[1] == 'o') ) { //book
                   Inproceedings *inp = static_cast<Inproceedings*>(baseInf);
                   inp->setBooktitle(str);
                }
                //printString(str);
                str = "";
             } else {
               str += string(" ");
             }
       }
       //printf("push_back() +1\n");
       obj.push_back(baseInf);
    }
    fclose(stdin);
}

void ParseBib::printItem(string type, string str) {
     if(str.length() == 0)return;
     int size = type.size();
     printf("  %s =", type.c_str());
     for (int i = size + 2; i < 15; i++) {
         printf(" ");
     }
     if(type == "journal"){
        if(isExist(str))printf("%s,\n", b[str].c_str());
        else printf("{%s},\n", str.c_str());
        return;
     }
     printf("{%s},\n", str.c_str());
}
/*
void ParseBib::printString(string str) {
     int len = str.length();
     printf("    ");
     int cnt = 0;
     for (int i = 0; i < len; i++) {
         if ( str[i] == '=' ) {
            for (int j = cnt; j < 11; j++ ) {
                printf(" ");
            }
            printf("=   ");
         } else {
           printf("%c", str[i]);
           cnt++;
         }
     }
     printf(" \n");
}
*/
void ParseBib::getMapReplace(){
    freopen("secu_full.bib", "r", stdin);
    char s[200];
    b.clear();
    while(NULL != gets(s)){
        if(s[0] != '@')continue;
        string s1 = "", s2 = "";
        int i;
        for(i = 8; s[i] != ' '; i++)s1 += s[i];
        while(s[i++] != '"');
        for(i; s[i] != '"'; i++)s2 += s[i];
        b[s2] = s1;
    }
    fclose(stdin);
}
bool ParseBib::isExist(const string& keyName){
    return b.find(keyName) != b.end();
}
void ParseBib::writeBibFile() {
     readBibFile();
     getMapReplace();
     freopen(tidyedBibFile.c_str(), "w", stdout);
     //printf("%%Tidybib@Yuansheng Liu of Xiangtan University\n");
     printf("\n");
     int size = obj.size();
     for (int i = 0; i < size; i++) {
         BaseInf *baseInf = static_cast<BaseInf*>(obj[i]);
         printf("@%s{%s,\n", baseInf->getType().c_str(), baseInf->getCite().c_str());

         printItem("author", baseInf->getAuthor());
         printItem("title", baseInf->getTitle());

         if (baseInf->getType() == "inproceedings" || baseInf->getType() == "incollection") {
             Inproceedings *inp = static_cast<Inproceedings*>(baseInf);
             printItem("booktitle", inp->getBooktitle());
         }

         if (baseInf->getType() == "article") {
             Article *art = static_cast<Article*>(baseInf);
             printItem("journal", art->getJournal());
         }

         printItem("year", baseInf->getYear());
         printItem("month", baseInf->getMonth());

         if (baseInf->getType() == "article") {
             Article *art = static_cast<Article*>(baseInf);
             printItem("volume", art->getVolume());
         }

         if (baseInf->getType() == "article") {
             Article *art = static_cast<Article*>(baseInf);
             printItem("number", art->getNumber());
         }

         printItem("pages", baseInf->getPages());
         printItem("doi", baseInf->getDoi());
         printf("}\n\n");

     }
     fclose(stdout);
}

void ParseBib::saveTidyedBibFile(){
     try {
          writeBibFile();
     }catch(exception e) {
     }
}

