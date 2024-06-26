# Key-Words-Replace
一句话简介：允许用户将".txt"格式文件中的某一关键词替换为另一关键词

### 使用场景

当我们在盗版网站下载小说时，常会发现小说中的某些关键词被网站替换为其名称，以彰显其“神通”。这些关键词通常是其他大型盗版网站的名称，例如：词汇“书包”，它经常会被替换为其他网站的名称，这些网站的本意我们可以理解，但是过于简单粗暴的替换会严重影响我们的阅读体验。于是有聪明的家伙就会在文本编辑器里打开文件，将关键词全部替换。这不失为一种方法，但对程序员来说太不优雅了，因此开发本程序，允许通过完全命令行的方式来完成关键词替换的工作。

### 使用语言及其版本

`go version go1.22.1 darwin/arm64`  

郑重声明：本人不对任何版本差异性问题负责，请与 GO 官方反馈

### 使用方法

##### 基础命令：

在可执行文件所在的目录下运行终端，执行命令：`$ kwr file.txt old_name new_name` 程序会在同级目录下生成文件 `new_yourFileName.txt` ，这就是替换关键词之后的文件

##### 进阶命令：

如果希望将某段文字完全删去，可以不输入 `new_name` 参数，程序会自动执行删除关键词的操作（尚未实现）

### 未来功能

0. 预计实现文件批量处理，完成对上述手动操作过程的效率碾压～～
1. 预计实现对所有文本类型的支持（着急的可以自己修改一下源码）但不会包括 `.docx .xlsx .pptx` 等格式

### 分支说明

Branch of `main` controls stable version, which will be updated when a new but stable feature is created.

Branch of `develop` controls develop version, which will be updated as any new line is written. 

### 编译说明

如需自行编译，请使用 `go1.22.1` 版本的编译器进行编译，建议编译 `main` 分支，`develop` 分支不保证可运行性
本程序暂时没有引入内核相关库
