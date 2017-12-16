## BibTidy
BibTidy is a bibliography file clean up application. A `.bib` file often contains many kinds of items that exported from different acadamic document databases. Such a file in a mess is difficult to read, modify and use. BibTidy can help you to clean up your `.bib` file such that you can use it directly or with slight modification.

## Quick start
(TODO) Download the package and replace the path of your input `.bib` file and the path of your output `.bib` file, and then build it to get executable file `bibtidy.exe`. Run this executable file and the clean up file will be obtained.

## Ducumentation
(TODO) This package needs improvement.

### How to build?
- Create a Go project in your workspace `xxx`, refer to https://golang.org/doc/code.html.
- Dowload the package or clone the repository into your `src` folder.
- Set your `GOPATH` with `export GOPATH=$HOME/xxx/`
- `cd` into `xxx/src/fp/` and run `go build`
- `cd` into `xxx/src/bibtidy/` and run `go install`
- Check your `bin` folder, and `bibtidy.exe` will be there.

### How to improve?
After you test your codes and install it in your device successfully, you can pull a request to me (See Pull request rules). I will response as soon as possible.

## Example

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
- After processing

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

## Pull request rules

- Add:     `add: Filename, Functional description (based on which version).`
- Update:  `update: Filename, Functional description (based on which version).`
- Delete:  `delete: Filename, why? (based on which version).`
- Others:  `change: Description of what changes you made.`

Another format is `add: xxx; update: xxx.`