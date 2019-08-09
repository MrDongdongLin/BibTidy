# version 6.2.0
author: Donny
email:  dongdonglin8@gmail.com

## 更新：

1. 调整了各个域的顺序
```tex
@ARTICLE{,
  author =       {},
  title =        {},
  journal =      {},
  year =         {},
  volume =       {},
  number =       {},
  pages =        {},
  month =        {}, 
  doi =          {},
}

@PROCEEDINGS{,
  title =        {},
  year =         {},
  editor =       {},
  publisher =    {},
  volume =       {},
  number =       {},
  series =       {},
  organization = {},
  address =      {},
  month =        {},
}
```
2. 月份改成首字母大写
3. 修复了article后没有逗号的问题
4. 修复了注释行未被忽略的问题
5. 修复了STRING域没有全部写出的问题

## 存在bug

+++++
新版更加灵活，可添加不同类型参看文献条目的格式化输出功能。
目前支持article和proceddings的格式化输出，如果bib文件中包含其他类型的条目（比如misc），那么会多出一个'}'结尾的符号。

解决该bug的方法很简单，在polish.py文件中第44行后添加相应代码即可

```python
article = re.findall(r"(@article.*?,)", str(buffer), re.IGNORECASE)
proceed = re.findall(r"(@proceedings.*?,)", str(buffer), re.IGNORECASE)
string  = re.findall(r"(@STRING.*?\"})", str(buffer))
# 在此处添加更多参考文献类型以格式化输出，例如：
# misc = re.findall(r"(@misc.*?,)", str(buffer), re.IGNORECASE)
```

+++++
正则匹配可能会出现多个匹配的情况，比如

```tex
@article{cqli:meet:ISAP19,
  author =       {Chengqing Li and Yun Zhang and Eric Yong Xie},
  title =        {{When an attacker meets a cipher-image in 2018: A Year in Review}},
  journal =      {Journal of Information Security and Applications},
  volume =       {48},
  pages =        {art. no. 102361},
  year =         {2019},
  doi =          {10.1016/j.jisa.2019.102361},
}
```

不区分大小写匹配`year`会产生两组匹配成功: `year = ['Year in Review}},', 'year =         {2019},']`，一般而言，关键字`year`不会出现大写的情况，但是如果其他地方也出现了小写的`year`，那么就会出现bug。比如上述例子写到文件就会变成

```tex
@article{cqli:meet:ISAP19,
  author =       {Chengqing Li and Yun Zhang and Eric Yong Xie},
  title =        {{When an attacker meets a cipher-image in 2018: A Year in Review}},
  journal =      {Journal of Information Security and Applications},
  volume =       {48},
  pages =        {art. no. 102361},
  Year in Review}},
  doi =          {10.1016/j.jisa.2019.102361},
}
```

`year`字段写入错误！

解决方案：考虑到year字段总是排在后面的位置（在doi之前），只要`doi`字段不出现字符串`year`，那么就可以使用向后匹配，或者使用匹配到的字符串中`['Year in Review}},', 'year =         {2019},']`中的最后一组`year[-1]`。


-------------------------------------------------------------------------
# version 6.1
author	:ChangzhouLi
tel	:15521268393
mail	:494691998@qq.com

更新：
1.处理mode1的文件时，请将待处理文件放入mode1_res文件夹中(secu_full.bib也请放在此文件夹)，输出的文件会存在mode1_out文件夹中
2.处理mode2的文件时，请将文件路径（程序文件路径为当前路径）逐一输入，以空格分隔，输出文件会存在mode2_out文件夹中

-------------------------------------------------------------------------
# version 6.0
author	:ChangzhouLi
tel	:15521268393
mail	:494691998@qq.com

修复问题：
1.修复一行中出现多个"="只保存第二个"="前字符的BUG

更新：处理后bib文件名将以原文件名保存至"handle/"文件夹中
-------------------------------------------------------------------------

