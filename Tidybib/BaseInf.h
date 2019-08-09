#pragma once
#include "Object.h"

class BaseInf: public Object
{
  private:
	string title;
	string author;
	string pages;
	string year;
	string month;
	string doi;
  string firstKey;
  string secondKey;

  public:
    void setTitle(string _title) {
         title = tidyTitle(_title);
    }
    string getTitle() {
         return title;
    }

    void setAuthor(string _author) {
         author = tidyAuthor(_author);
    }
    string getAuthor() {
         return author;
    }

    void setPages(string _pages) {
         pages = getString(_pages);
    }
    string getPages() {
         return pages;
    }

    void setFristKey(string _firstKey) {
      firstKey = _firstKey;
    }

    string getFirstKey() {
      return firstKey;
    }

    void setSecondKey(string _secondKey) {
      secondKey = _secondKey;
    }

    string getSecondKey() {
      return secondKey;
    }

    void setYear(string _year) {
         year = getString(_year);
    }

    string getYear() {
         return year;
    }

    void setMonth(string _month) {
         month = getString(_month);
    }

    string getMonth() {
         return month;
    }

    void setDoi(string _doi) {
         doi = getString(_doi);
    }

    string getDoi() {
         return doi;
    }

    string getCite() {
         return firstKey + ":" + secondKey + ":" + year;
    }

    string getAuthorString(string str) {//get string between { }
       size_t begin = str.find("{") + 1;
       size_t end = str.rfind("}");
       return str.substr(begin, end-begin);
    }

    string tidyAuthor(string str) {
       str = getAuthorString(str);
       vector<string> author;
       size_t pos = 0, find;
       while ((find = str.find(" and ", pos)) !=  string::npos) {
             author.push_back(str.substr(pos, find-pos));
             pos = find + 5;
       }
       author.push_back(str.substr(pos, str.size()-pos+1));
       int size = author.size();
       for (int i = 0; i < size; i++) {
           string tmp = author[i];
           find = tmp.find(", ");
           author[i] = tmp.substr(find+2) + " " + tmp.substr(0, find);
           if (0==i) {
             firstKey = tmp.substr(0, find);
           }
       }
       string ret = "";
       for (int i = 0; i < size; i++) {
           ret += author[i];
           if (i < size-1) {
              ret += " and ";
           }
       }

       vector<string> key;
       pos = 0, find;
       while ((find = firstKey.find(" ", pos)) !=  string::npos) {
             key.push_back(firstKey.substr(pos, find-pos));
             pos = find + 1;
       }
       key.push_back(firstKey.substr(pos, firstKey.size()-pos+1));
       firstKey = "";
       size = key.size();
       for (int i =0; i < size; i++) {
          firstKey += key[i];
       }
       return ret;
    }

    string tidyTitle(string str) {
       str = getString(str);
       vector<string> title;
       size_t pos = 0, find;
       while ((find = str.find(" ", pos)) !=  string::npos) {
             title.push_back(str.substr(pos, find-pos));
             pos = find + 1;
       }
       title.push_back(str.substr(pos, str.size()-pos+1));
       int size = title.size();
       for (int i = 0; i < size; i++) {
           string tmp = title[i];
           /* //special case
           if (tmp[0] == '&') {
              tmp.insert(0, "\\");
           }*/
           int len = tmp.size();
           if (isalpha(tmp[0])) {
             for (int j = 0; j < len; j++) {
                 if (isupper(tmp[j])) {
                    tmp[j] += 32;
                 }
             }
             title[i] = tmp;
           }
       }

       string ret = "";
       for (int i = 0; i < size; i++) {
           ret += title[i];
           if (i < size-1) {
              ret += " ";
           }
       }
       if (isalpha(ret[0])) {
         ret[0] -= 32;
       }
       return ret;
    }

};
