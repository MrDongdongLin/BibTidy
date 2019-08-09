#-*- coding:utf-8 -*-
import numpy as np
import os
import sys
import string
import re
# reload(sys)
# sys.setdefaultencoding('utf-8')

num = len(sys.argv)
if num == 1:
    print("please input the file you want to handle....")
else:
    print(num-1, 'file need to be dealt...')
    os.system("md mode2_out")
    for i in range(1,num):
        save_file = "mode2_out/" + sys.argv[i]
        item = sys.argv[i]
        if os.path.isfile(item):
            fin = open(item, 'r')
            fout = open(save_file, 'w')

            line = True
            flag_end = False
            buffer = [] # 一次缓冲一条ref的所有内容
            # for line in fin:
            while line:
                line = fin.readline()
                # line = line.strip()
                # print(type(line))
                if len(line) == 0:
                    continue
                if line[0] == '%': # 如果是注释行，直接写入文件
                    fout.write(line)
                    continue
                if line[0] == '@': # 如果是第一行，缓存这一行
                    buffer.append(line)
                    continue
                if line[0] == '}': # 如果是最后一行，为了和以'}'结尾的行有所区分，这里用标志flag_end标识
                    flag_end = True
                    continue

                if line == '\n': # 如果遇到空行，说明缓存完毕，开始正则匹配以及写入文件
                    article = re.findall(r"(@article.*?,)", str(buffer), re.IGNORECASE)
                    proceed = re.findall(r"(@proceedings.*?,)", str(buffer), re.IGNORECASE)
                    string  = re.findall(r"(@STRING.*?\"})", str(buffer))
                    # 在此处添加更多参考文献类型以格式化输出，例如：
                    # misc = re.findall(r"(@misc.*?,)", str(buffer), re.IGNORECASE)

                    author  = re.findall(r"(author.*?},)", str(buffer))
                    title   = re.findall(r"(title.*?},)", str(buffer))
                    journal = re.findall(r"(journal.*?},)", str(buffer))
                    year    = re.findall(r"(year.*?},)", str(buffer))
                    volume  = re.findall(r"(volume.*?},)", str(buffer))
                    number  = re.findall(r"(number.*?},)", str(buffer))
                    pages   = re.findall(r"(pages.*?},)", str(buffer))
                    month   = re.findall(r"(month.*?},)", str(buffer))
                    doi     = re.findall(r"(doi.*?},)", str(buffer))

                    editor    = re.findall(r"(editor.*?},)", str(buffer))
                    publisher = re.findall(r"(publisher.*?},)", str(buffer))
                    series    = re.findall(r"(series.*?},)", str(buffer))
                    organization = re.findall(r"(organization.*?},)", str(buffer))
                    address   = re.findall(r"(address.*?},)", str(buffer))

                    # 月份第一个字母大写
                    if month:
                        month = month[0]
                        month = month.split('=')
                        month[1] = month[1].title()
                        month = month[0] + '=' + month[1]
                    
                    if string:
                        for i in range(len(string)):
                            fout.write(string[i]+'\n')
                        fout.write('\n')

                    if article:
                        article = article[0]
                        fout.write(article+'\n')
                        if author:
                            author = author[0]
                            fout.write('  '+author+'\n')
                        if title:
                            title = title[0]
                            fout.write('  '+title+'\n')
                        if journal:
                            journal = journal[0]
                            fout.write('  '+journal+'\n')
                        if year:
                            year = year[0]
                            fout.write('  '+year+'\n')
                        if volume:
                            volume = volume[0]
                            fout.write('  '+volume+'\n')
                        if number:
                            number = number[0]
                            fout.write('  '+number+'\n')
                        if pages:
                            pages = pages[0]
                            fout.write('  '+pages+'\n')
                        if month:
                            # month = month[0]
                            fout.write('  '+month+'\n')
                        if doi:
                            doi = doi[0]
                            fout.write('  '+doi+'\n')
                    elif proceed:
                        proceed = proceed[0]
                        fout.write(proceed+'\n')
                        if title:
                            title = title[0]
                            fout.write('  '+title+'\n')
                        if year:
                            year = year[0]
                            fout.write('  '+year+'\n')
                        if editor:
                            editor = editor[0]
                            fout.write('  '+editor+'\n')
                        if publisher:
                            publisher = publisher[0]
                            fout.write('  '+publisher+'\n')
                        if volume:
                            volume = volume[0]
                            fout.write('  '+volume+'\n')
                        if number:
                            number = [0]
                            fout.write('  '+number+'\n')
                        if series:
                            series = series[0]
                            fout.write('  '+series+'\n')
                        if organization:
                            organization = organization[0]
                            fout.write('  '+organization+'\n')
                        if address:
                            address = address[0]
                            fout.write('  '+address+'\n')
                        if month:
                            # month = month[0]
                            fout.write('  '+month+'\n')

                    # print(str(buffer))
                    # fout.write(line)
                    if flag_end:
                        fout.write('}\n\n')
                    buffer = [] # 清空缓冲区中的字符串
                    flag_end = False
                
                else: # 非空行，说明读到的是每个条目，则对它们进行格式化处理
                    _line = line.split('=')
                    ss = '  ' + _line[0].strip() + ' ='
                    while(len(ss) < 17):
                        ss += ' '
                    ss += _line[1].strip()
                    for i in _line[2:]:##若后面还有'=',需要继续补全防止有遗漏的数据
                        ss = ss + "=" + i.strip()
                    ss += "\n"
                    buffer.append(ss)

            fin.close()
            fout.close()
            print(item, 'is dealt.')
        else:
            print(item, ' No such file!!')

