# BibTidy
## 输入格式
待处理bib文件。这里bib文件可分为两种：

- 鱼龙混杂型
什么类型的都有，有从IEEE上导出的（见下方格式举例），有从ACM上导出的，也有从Elsevier上导出的等等。

- 统一型
绝大部分bibtex都能从Web of Science上导出，这类bib文件格式是统一的。

## 文件处理

- 正则匹配到对应字段的，就写到文件，否则删去该字段（字段的定义见下方输出格式）。

## 输出格式

使用结构体保存匹配到的字段，结构体（article）的定义为
```go
type OutBib struct {
	Author  string
	Title   string
	Journal string
	Year    string
	Volume  string
	Number  string
	Pages   string
}
```

## 关键字
整理bib文件；字符串；正则表达式；Go语言

# IEEE
## 格式举例

```bib
@ARTICLE{7999153, 
author={C. Li and D. Lin and J. L\"u}, 
journal={IEEE MultiMedia}, 
title={Cryptanalyzing an Image-Scrambling Encryption Algorithm of Pixel Bits}, 
year={2017}, 
volume={24}, 
number={3}, 
pages={64-71}, 
abstract={Position scrambling (permutation) is widely used in multimedia encryption schemes and some international encryption standards, such as the Data Encryption Standard and the Advanced Encryption Standard. In this article, the authors re-evaluate the security of a typical image-scrambling encryption algorithm (ISEA). Using the internal correlation remaining in the cipher image, they disclose important visual information of the corresponding plain image in a ciphertext-only attack scenario. Furthermore, they found that the real scrambling domain - the position-scrambling scope of ISEA's scrambled elements - can be used to support an efficient known or chosen-plaintext attack on it. Detailed experimental results have verified these points and demonstrate that some advanced multimedia processing techniques can facilitate the cryptanalysis of multimedia encryption algorithms.}, 
keywords={cryptography;image processing;multimedia systems;ISEA;advanced encryption standard;chosen-plaintext attack;cipher image;ciphertext-only attack scenario;cryptanalysis;data encryption standard;image-scrambling encryption algorithm security;internal correlation;international encryption standards;multimedia encryption schemes;multimedia processing techniques;permutation;pixel bits;position scrambling;position-scrambling scope;visual information;Algorithm design and analysis;Ciphers;Cryptography;Encryption;Image processing;Mathematical model;Multimedia communication;Visualization;ciphertext-only attack;cryptanalysis;cryptography;graphics;image encryption;known-plaintext attack;multimedia;security;template matching}, 
doi={10.1109/MMUL.2017.3051512}, 
ISSN={1070-986X}, 
month={},}
```

## 处理后

```bib
@article{ieee,
	author		= {C. Li and D. Lin and J. L\"u},
	title		= {Cryptanalyzing an Image-Scrambling Encryption Algorithm of Pixel Bits},
	journal		= {IEEE MultiMedia},
	year		= {2017},
	volume		= {24},
	number		= {3},
	pages		= {64-71},
}
```

# PR规范

- 添加文件：`add: Filename, Functional description (based on which version).`
- 更新文件：`update: Filename, Functional description (based on which version).`
- 删除文件：`delete: Filename, why? (based on which version).`
- 文件结构变动：`change: Description of what changes you made.`

如有多种更改，格式可以是`add: xxx; update: xxx.`