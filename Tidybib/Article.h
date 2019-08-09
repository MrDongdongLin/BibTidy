#pragma once
#include "BaseInf.h"

class Article: public BaseInf
{
  private:
	string journal;
	string number;
	string volume;

  public:
    Article() {
        setType("article");
    }

    void setJournal(string _journal) {
         journal = tidyJournal(_journal);
    }
    string getJournal() {
         return journal;
    }

    void setNumber(string _number) {
         number = getString(_number);
    }
    string getNumber() {
         return number;
    }

    void setVolume(string _volume) {
         volume = getString(_volume);
    }
    string getVolume() {
         return volume;
    }

    string getCiteKeyWords() {
    }

    string tidyJournal(string str) {
       str = getString(str);
       vector<string> journal;
       size_t pos = 0, find;
       while ((find = str.find(" ", pos)) !=  string::npos) {
             journal.push_back(str.substr(pos, find-pos));
             pos = find + 1;
       }
       journal.push_back(str.substr(pos, str.size()-pos+1));
       int size = journal.size();
       for (int i = 0; i < size; i++) {
           if (journal[i]==string("IEEE")) {
              continue;
           }
           string tmp = journal[i];
           if (tmp[0] == '&') {
              tmp.insert(0, "\\");
           }
           int idx = 0;
           int len = tmp.size();
           if (len > 3) {
             idx = 1;
           }

           for (int j = idx; j < len; j++) {
               if (j > 0 && tmp[j-1] == '-') {
                 if (islower(tmp[j-1])) {
                   tmp[j] -= 32;
                 }
               } else
               if (isupper(tmp[j])) {
                  tmp[j] += 32;
               }

           }
           journal[i] = tmp;
       }
       string _secondKey = "";
       string ret = "";
       for (int i = 0; i < size; i++) {
           ret += journal[i];
           if (i < size-1) {
              ret += " ";
           }
           if (journal[i].size()>3 && isupper(journal[i][0])) {
              _secondKey += journal[i][0];
           }
       }
       setSecondKey(_secondKey);
       return ret;
    }

};

