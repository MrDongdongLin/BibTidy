#include "ParseBib.h"
#include <windows.h>

map<string, string>b;
vector<string> bibFile;

bool IsRoot(string Path)
{
    string Root;
    Root = Path.at(0)+":\\";
    if (Root == Path)
       return true;
    else
       return false;
}

void FindAllBibFile(string Path)
{
    string szFind;
    szFind = Path;
    if (!IsRoot(szFind))
       szFind += "\\";
    szFind += "*.bib";
    WIN32_FIND_DATA FindFileData;
    HANDLE hFind=FindFirstFile(szFind.c_str(),& FindFileData);
    if(hFind==INVALID_HANDLE_VALUE)
       return ;
    do
    {
       if(!(FindFileData.dwFileAttributes & FILE_ATTRIBUTE_DIRECTORY)) //find .bib file
       {
           //cout<<FindFileData.cFileName<<endl;
           bibFile.push_back(FindFileData.cFileName);
       }
    } while (FindNextFile(hFind,& FindFileData));
    FindClose(hFind);
}

void getMapReplace(){
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

int main(int argc, char* args[]) {
    cout << "mode:\n    1. *.bib->Tidy*.bib    2. rearrange bib file already Tidy\nplease select one mode:";
    int num;
    cin >> num;
    if(num == 1){
    if (argc == 1 ) {
       char cf[1000];
       GetCurrentDirectory(1000, cf);
       string cc = cf;
       cc += "\\mode1_res\\";
       FindAllBibFile(cc);
       int size = bibFile.size();
       if (size == 0) {
         freopen("CON", "w", stdout);
         printf("*.bib file not found.\n");
       }
       system("md mode1_out");
       for (int i = 0; i < size; i++) {
            if(bibFile[i] == "secu_full.bib")continue;
           if (bibFile[i].substr(0, 6) != "tiyed_")  {
              ParseBib p(string("mode1_res\\") + bibFile[i], string("mode1_out\\") + bibFile[i]);
              p.saveTidyedBibFile();
              freopen("CON", "w", stdout);
              printf("%s ---> %s\n", (string("mode1_res\\") + bibFile[i]).c_str(), (string("mode1_out\\") + bibFile[i]).c_str());
           } else {
                getMapReplace();
                freopen(bibFile[i].c_str(), "r", stdin);
                char line[200];
                remove("temp.bib");
                freopen("temp.bib", "w", stdout);
                while(NULL != gets(line)){
                    if((line[2] == 'j' || line[2] == 'J') && (line[17] == '{' || line[17] == '"')){
                        printf("  journal =      ");
                        int len = strlen(line);
                        int ind = 18;
                        string s = "";
                        len -= 2;
                        while(ind < len){
                            s += line[ind++];
                        }
                        if(b.find(s) != b.end())printf("%s,\n", b[s].c_str());
                        else printf("{%s},\n", s.c_str());
                    }
                    else printf("%s\n", line);
                }
              freopen("CON", "w", stdout);
              freopen("CON", "r", stdin);
              printf("%s has replaced.\n", bibFile[i].c_str());
              remove(bibFile[i].c_str());
                rename("temp.bib", bibFile[i].c_str());
           }
       }
    } else
    if (argc == 3 ) {
       ParseBib p(args[1], args[2]);
       p.saveTidyedBibFile();
       freopen("CON", "w", stdout);
       printf("%s --->  %s\n", args[1], args[2]);
    }
    }
    else if(num == 2){
        string cmd = "python polish.py ";
        cout << "\nplease input the files you want to deal, separated by space:\n";
        char fileName[200];
        getchar();
        cin.getline(fileName, 100);
        system((cmd+fileName).c_str());
    }
    else{
        cout << "the number you input is illegal!" << endl;
    }
    freopen("CON", "r", stdin);
    freopen("CON", "w", stdout);
    printf("\nPress Enter to continue.\n");
    getchar();
    return 0;
}
