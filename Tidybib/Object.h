#pragma once
#include "ComHead.h"

class Object
{
    private:
       string type;

    public:
       void setType(string _type) {
            type = _type;
       }
       string getType() {
            return type;
       }
       string getString(string str) {//get string between {{ }}
         size_t begin = str.find("{{") + 2;
         size_t end = str.rfind("}}");
         return str.substr(begin, end-begin);
       }

};
