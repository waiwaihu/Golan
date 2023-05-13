# 前奏

给grep，egrep加上颜色`alias grep='grep --color=auto'`

`alias egrep='egrep --color=auto'`

注意系统的字符集`en_US.UTF-8`，如果出现问题，就修改字符集为C：`export LANG=C`

# 三剑客

| 命令 | 特点                    | 场景                                           |
| ---- | ----------------------- | ---------------------------------------------- |
| grep | 过滤                    | grep命令过滤速度是最快的                       |
| sed  | 替换,修改文件内容，取行 | 如果要进行替换/修改文件内容,取出某个范围的内容 |
| awk  | 取列,统计计算           | 取列，统计计算，对比/比较                      |

# grep

```shell
1、-A：显示搜索内容的那一行和后面的行情（-A后面可以带数字，代表要显示后面几行）
# grep -A 1 root /etc/passwd

2、-B：显示搜索内容的那一行和前一行
# grep -B 1 root /etc/passwd

3、-C：显示搜索内容的那一行和前后行
# grep -C 1 root /etc/passwd

4、-c：显示符合搜索内容的行数
# grep -c root /etc/passwd

5、-i：忽略大小写
# grep -i root /etc/passwd

6、-n：输出符合搜索内容所在的行号
# grep -n root /etc/passwd

7、-v：不包含关键字的行
# grep -v root /etc/passwd

可以过滤掉带有注释符合#的行
# grep -v '^#' root /etc/ssh/sshd_config

再过滤掉空的行
# grep -v '^#' root /etc/ssh/sshd_config | grep -v '^$'
```

### 实例

```shell
1、在当前目录中，查找后缀有 .cpp 字样的文件中包含 test 字符串的文件，并打印出该字符串的行。此时，可以使用如下命令：
grep test *.cpp

2、以递归的方式查找符合条件的文件。例如，查找指定目录/etc/acpi 及其子目录（如果存在子目录的话）下所有文件中包含字符串"update"的文件，并打印出该字符串所在行的内容，使用的命令为：
grep -r update /etc/acpi 

3、反向查找。前面各个例子是查找并打印出符合条件的行，通过"-v"参数可以打印出不符合条件行的内容。
查找文件名中包含 test 的文件中不包含test 的行，此时，使用的命令为：
grep -v test *test*

grep的高级用法是开启-E选项。

4、grep -E 同时匹配多个关键字–或关系
grep -E "word1|word2|word3" file.txt

5、-E和-v可以一起使用
grep -E -v "test|good" *.cpp

6、grep会打印包含匹配pattern的一整行，如果我们期望只打印匹配的pattern，使用-o选项

7、-w可以实现完整单词匹配，也很有用。例如：
$ grep -w "hello" *.cpp

包含helloworld的行是不会显示的，因为只会匹配含有hello单词的行，单词意味着其前后要有空格间隔。

8、-A n A是after的意思 除了显示匹配行，还会显示匹配行后面n行。
-B n B是before的意思 除了显示匹配行，还会显示匹配行前面n行。
-C n C是after的意思 显示匹配行上下各n行。

9、–exclude-dir=GLOB 递归搜索时排除匹配模式的目录

10、–exclude-from=FILE 排除匹配模式的文件。

11、-n 顺带输出匹配行的行号
$ grep -n "test" *.cpp                                                                            127
3:test good

12、以某个单词开头/结尾的行
 grep -E "^hello" *.cpp 
 grep -E "world$" *.cpp 

13、不是以某个字母开头/结尾的行
grep -E "[^#]" *.cpp

14、.是特殊字符，表示匹配一个单词。
下面的正则表达式查找 IP 地址 192.168.1.254 将不能获得预期的结果：
grep '192.168.1.254' /etc/hosts

其中三个点都需要被转义：
grep '192\.168\.1\.254' /etc/hosts
```

# sed

### 1、sed 工具概述

sed（Stream EDitor）是一个强大而简单的文本解析转换工具，可以读取文本，并根据指定的条件对文本内容进行编辑（删除、替换、添加、移动等），最后输出所有行或者仅输出处理的某些行。
sed也可以在无交互的情况下实现相当复杂的文本处理操作，被广泛应用于Shell脚本中，用以完成各种自动化处理任务。sed 依赖于正则表达式

### 2、sed 的工作流程

主要包括读取、执行和显示三个过程。

读取：sed 从输入流（文件、管道、标准输入）中读取一行内容并存储到临时的缓冲区中（又称模式空间，pattern space）。
执行：默认情况下，所有的 sed 命令都在模式空间中顺序地执行，除非指定了行的地址，否则 sed 命令将会在所有的行上依次执行。
显示：发送修改后的内容到输出流。在发送数据后，模式空间将会被清空。
在所有的文件内容都被处理完成之前，上述过程将重复执行，直至所有内容被处理完。

### 3、sed 命令格式

sed [选项] '操作' 参数

| -e     | 表示用指定命令或脚本来处理输入的文本文件,-e : 可以在同一行里执行多条命令 |
| ------ | ------------------------------------------------------------ |
| -f     | 表示用指定的脚本文件来处理输入的文本文件                     |
| -h     | 或--help 显示帮助                                            |
| -n     | 表示显示处理后的结果                                         |
| -i     | 直接编辑文本文件                                             |
| -i.bak | 备份文本文件                                                 |
| -r，-E | 使用扩展正则表达式                                           |
| -s     | 将多个文件视为独立文件，而不是单个连续的长文件流             |

**sed命令的常用操作**

“操作”用于指定对文件操作的动作行为，也就是 sed 的命令

| a    | 增加，在当前行下面增加一行指定内容                           |
| ---- | ------------------------------------------------------------ |
| c    | 替换，将选定行替换为指定内容                                 |
| d    | 删除，删除选定的行                                           |
| i    | 插入，在选定行上面插入一行指定内容                           |
| p    | 打印，如果同时指定行，表示打印指定行；如果不指定行，则表示打印所有内容 |
| s    | 替换，替换指定字符                                           |
| y    | 字符转换                                                     |

```shell
a\ 在当前行下面插入文本。
i\ 在当前行上面插入文本。
c\ 把选定的行改为新的文本。
d 删除，删除选择的行。
D 删除模板块的第一行。
s 替换指定字符
h 拷贝模板块的内容到内存中的缓冲区。
H 追加模板块的内容到内存中的缓冲区。
g 获得内存缓冲区的内容，并替代当前模板块中的文本。
G 获得内存缓冲区的内容，并追加到当前模板块文本的后面。
l 列表不能打印字符的清单。
n 读取下一个输入行，用下一个命令处理新的行而不是用第一个命令。
N 追加下一个输入行到模板块后面并在二者间嵌入一个新行，改变当前行号码。
p 打印模板块的行。
P(大写) 打印模板块的第一行。
q 退出Sed。
b lable 分支到脚本中带有标记的地方，如果分支不存在则分支到脚本的末尾。
r file 从file中读行。
t label if分支，从最后一行开始，条件一旦满足或者T，t命令，将导致分支到带有标号的命令处，或者到脚本的末尾。
T label 错误分支，从最后一行开始，一旦发生错误或者T，t命令，将导致分支到带有标号的命令处，或者到脚本的末尾。
w file 写并追加模板块到file末尾。
W file 写并追加模板块的第一行到file末尾。
! 表示后面的命令对所有没有被选定的行发生作用。
= 打印当前行号码。
# 把注释扩展到下一个换行符以前
```

```shell
sed -i 就是直接对文本文件进行操作的。

sed -i 's/原字符串/新字符串/' /home/1.txt
sed -i 's/原字符串/新字符串/g' /home/1.txt

sed -e : 可以在同一行里执行多条命令
sed -e 's/11/00/g' -e 's/22/99/g' /home/1.txt
```

### sed元字符集

```shell
^ 匹配行开始，如：/^sed/匹配所有以sed开头的行。 
$ 匹配行结束，如：/sed$/匹配所有以sed结尾的行。 
^$ 空白行
. 匹配一个非换行符的任意字符，如：/s.d/匹配s后接一个任意字符，最后是d。 
* 匹配0个或多个字符，如：/*sed/匹配所有模板是一个或多个空格后紧跟sed的行。 
[] 匹配一个指定范围内的字符，如/[Ss]ed/匹配sed和Sed。 
[^] 匹配一个不在指定范围内的字符，如：/[^A-RT-Z]ed/匹配不包含A-R和T-Z的一个字母开头，紧跟ed的行。 
\(..\) 匹配子串，保存匹配的字符，如s/\(love\)able/\1rs，loveable被替换成lovers。 
& 保存搜索字符用来替换其他字符，如s/love/**&**/，love这成**love**。 
\< 匹配单词的开始，如:/\ 匹配单词的结束，如/love\>/匹配包含以love结尾的单词的行。 
x\{m\} 重复字符x，m次，如：/0\{5\}/匹配包含5个0的行。 
x\{m,\} 重复字符x，至少m次，如：/0\{5,\}/匹配至少有5个0的行。 
x\{m,n\} 重复字符x，至少m次，不多于n次，如：/0\{5,10\}/匹配5~10个0的行。
```

### sed用法实例

**1. 替换操作：s命令**

```shell
sed 's/book/books/' file         # 将file文件中每一行的第一个book替换为books
```

**2. 全面替换标记g**

```shell
sed 's/book/books/g' file      # 替换file文件每一行中的所有匹配
# 当需要从第N处匹配开始替换时，可以使用/Ng
sed 's/book/books/2g' file    # 从file每一行第二个匹配的开始替换
```

**3. 定界符**

**以上命令中字符 / 在sed中作为定界符使用，也可以使用其他的任意定界符：**

```shell
sed 's:test:TEXT:g' file             # 以：作为定界符
sed 's|test|TEXT|g' file             # 以｜作为定界符
# 定界符出现在式样内部时，需要进行转义
sed 's/\/bin\/bash/bin/g' file          # 将/bin/bash替换为bin
```

**4. 删除操作：d命令**

```shell
sed '/^$/d' file          #　删除空白行
sed '2d' file             #　删除第二行
sed '$d' file             #　删除最后一行
sed '2,$d' file           #　删除第二行到最后一行
sed '/^test/d'            #　删除所有以test开头的行
```

**5. 已匹配字符串标记＆**

正则表达式\w\+匹配到的每一个单词，使用[＆]替换他，＆对应于之前匹配到的单词

```shell
echo this is a test line | sed 's/\w\+/[&]/g'
[this] [is] [a] [test] [line]

# 所有以192.168.01开头的行都会被替换成自己加上localhost:
sed 's/^192.168.01/&localhost/' file
192.168.01localhost
```

**6. 子串匹配标记\1**
匹配给定样式其中的一部分

```shell
echo this is digit 7 number | sed 's/digit \([0-9]\)/\1/'
this  is 7 number
# 命令中digit 7 被替换成了７。样式匹配到的子串是７,\(\)用于匹配子串，
# 对于匹配到的子串第一个就标记为\1，以此类推，匹配到的第二个子串是\2
echo aaa bbb | sed 's/\([a-z]\+\) \([a-z]\+\)/\2 \1/'
bbb aaa

sed -n 's/\(love\)able/\1rs/p' file 　　　# 将file中的loveable替换为lovers,并打印出来
lovers
```

**7. 组合多个表达式**

```shell
sed '表达式' | sed '表达式'　<==> sed '表达式；表达式'
```

**8. 引用**

sed 表达式可以使用单引号来引用，但是如果表达式内部包含变量字符串，就需要使用双引号

```shell
test=hello
echo hello WORLD | sed "s/$test/HELLO/"
HELLO WORLD
```

**9. 选定行的范围：，(逗号)**

```shell
# 所有在模板test和check所确定的范围内的行都被打印
sed -n '/test/,/check/p' file

# 打印第五行开始到第一个包含以test开头的行之间的所有行
sed -n '5,/^test/p' file

# 对于模板test和west之间的行，每行的末尾用字符aaa bbb替换
sed '/test/,/west/s/$/aaa bbb/' file
```

**10. 多点编辑：e命令 **

**-e选项允许在同一行里执行多条命令**

```shell
sed -e '1,5d' -e 's/test/check/' file
# 第一条命令删除１至５行，第二条命令用check替换test。
# 命令执行顺序对结果有影响，如果两个都是替换命令那么第一个替换命令将影响第二个替换命令的结果
```

**11. 从文件读入：r命令**

```shell
# file里的内容被读进来，显示在与test匹配的行后面，如果匹配多行，则file的内容将显示在所有匹配行的下面
sed '/test/r file' filename
```

**写入文件：ｗ命令**

```shell
# 在example中所有包含test的行都被写入file里
sed -n '/test/w file' example
```

**12. 追加(行下)：a\命令**

```shell
# 将this is a test line追加到以test开头的行的后面
sed '/^test/a\this is a test line' file

# 在test.conf文件第二行之后插入this is a test line
sed -i '2a\this is a test line' file
```

**13. 插入(行上)：i\命令**

```shell
# 将this is a test line 插入到以test开头的行的前面
sed '/^test/i\this is a test line' file

# 在第五行之前插入this is a test line
sed '5i\this is a test line' file
```

**14. 下一个：ｎ命令**

```shell
# 如果test被匹配，则移动到匹配行的下一行，替换这一行的aa,变为bb,并打印该行
sed '/test/{ n; s/aa/bb/; }' file
```

**15. 变形：ｙ命令**

```shell
# 把1-10行内所有的abcde转变为大写，注意，正则表达式元字符不能使用这个命令
sed '1,10y/abcde/ABCDE/' file
```

**16. 退出：ｑ命令**

```shell
sed '10q' file           # 打印完第十行后，退出sed
```

**17. 保持和获取：h命令和G命令**
在sed处理文件的时候，每一行都被保存在一个叫模式空间的临时缓冲区中，除非行被删除或者输出被取消，否则所有被处理的行都将 打印在屏幕上。接着模式空间被清空，并存入新的一行等待处理

```shell
sed -e '/test/h' -e '$G' file
# 在这个例子里，匹配test的行被找到后，将存入模式空间，h命令将其复制并存入一个称为保持缓存区的特殊缓冲区内。
# 第二条语句的意思是，当到达最后一行后，G命令取出保持缓冲区的行，然后把它放回模式空间中，且追加到现在已经存在于模式空间中的行的末尾。
# 在这个例子中就是追加到最后一行。简单来说，任何包含test的行都被复制并追加到该文件的末尾。
```

**18. 保持和互换：h命令和ｘ命令**

```shell
sed -e 'test/h' -e '/check/x' file         # 把包含test与check的行互换
```

**19. 脚本scriptfile**
sed脚本是一个sed的命令清单，启动Sed时以-f选项引导脚本文件名。Sed对于脚本中输入的命令非常挑剔，在命令的末尾不能有任何空白或文本，如果在一行中有多个命令，要用分号分隔。以#开头的行为注释行，且不能跨行。

```shell
sed [options] -f scriptfile file
```

**20. 打印奇数或偶数行**

```shell
方法１：
sed -n 'p;n' test          # 奇数行
sed -n 'n;p' test          # 偶数行

方法２：
sed -n '1~2p' test         # 奇数行
sed -n '2~2p' test         # 偶数行
```

**21. 打印匹配字符串的下一行**

```shell
grep -A 1 SCC URFILE
sed -n '/SCC/{n;p}' URFILE
awk '/SCC/{getline;print}' URFILE
```

# EOF

EOF是END Of File的缩写,表示自定义终止符.既然自定义,那么EOF就不是固定的,可以随意设置别名,在linux按ctrl-d就代表EOF.
EOF一般会配合cat能够多行文本输出.

其用法如下:
<<EOF     //开始
....
EOF      //结束

还可以自定义，比如自定义：
<<BBB     //开始
....
BBB       //结束

通过cat配合重定向能够生成文件并追加操作,在它之前先熟悉几个特殊符号:
<　　  :输入重定向
\> 　　 :输出重定向
\>>　　 :输出重定向,进行追加,不会覆盖之前内容
<< 　　:标准输入来自命令行的一对分隔号的中间内容.

示例：

```shell
cat<<EOF > test.sh

12345
aaaaa
bbbbb

EOF
```

生成 test.sh，文本内容为：

```shell
12345
aaaaa
bbbbb
```

EOF 就是一个标记值，标记多行文本的开始、结束位置。你可以使用任意自定义字符来做标记。

比如改写上面的代码：

```shell
cat<<lalala > test.sh

12345
aaaaa
bbbbb
lalala
```

**【注意重点！！】**

**cat<<EOF 中间没有空格，没有空格，没有空格！有空格就不好使了！**

# tee

[tee](https://so.csdn.net/so/search?q=tee&spm=1001.2101.3001.7020) 的功能是从标准输入读取，再写入标准输出和文件。

```shell
-a, --append                        追加到文件
-i, --ignore-interrupts           忽略中断信号
-p                                         诊断写入非管道的错误
--output-error[=MODE]        设置输出错误的方式，MODE 的选项在下边
--help                                   帮助文档
--version                              版本信息
MODE:
warn                   写入遇到错误时诊断
warn-nopipe       写入非管道遇到错误时诊断
exit                     写入遇到错误时退出
exit-nopipe         写入非管道遇到错误时退出
```

使用示例：

默认功能和追加功能：

```shell
[root@server dir]# echo 'This is a sentence.' | tee output
This is a sentence.
 
[root@server dir]# cat output
This is a sentence.
 
[root@server dir]# echo 'This is another sentence.' | tee -a output
This is another sentence.
 
[root@server dir]# cat output
This is a sentence.
This is another sentence.
 
[root@server dir]# echo 'This is a unique sentence.' | tee output
This is a unique sentence.
 
[root@server dir]# cat output
This is a unique sentence.
```

同时写入两个文件：

```shell
[root@server dir]# tee a b
they have the same content
they have the same content
^C
[root@server dir]# cat a
they have the same content
[root@server dir]# cat b
they have the same content
```

