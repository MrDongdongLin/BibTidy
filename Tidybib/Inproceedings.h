#pragma once
#include "BaseInf.h"

//or incollection
class Inproceedings: public BaseInf
{
  private:
	string booktitle;

  public:
    Inproceedings() {
       setType("inproceedings");//
    }

    void setBooktitle(string _booktitle) {
         booktitle = tidyBooktitle(_booktitle);
    }

    string getBooktitle() {
         return booktitle;
    }

    string tidyBooktitle(string str) {
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
           if (title[i].find(string("IEEE"))!=string::npos) {
              continue;
           }
           string tmp = title[i];
            //special case
           if (tmp[0] == '&') {
              tmp.insert(0, "\\");
           }
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
       string _secondKey = "";
       string ret = "";
       for (int i = 0; i < size; i++) {
           ret += title[i];
           if (i < size-1) {
              ret += " ";
           }
           if (title[i].size()>3 && isalpha(title[i][0])) {
             if (isupper(title[i][0])) {
               _secondKey += title[i][0];
             } else {
               _secondKey += (title[i][0]-32);
             }
           }
       }
       setSecondKey(_secondKey);
       if (isalpha(ret[0])&&islower(ret[0])) {
         ret[0] -= 32;
       }
       return ret;
    }

};
