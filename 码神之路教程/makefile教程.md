# makefile教程

我们注意到在k8s的源码根目录下有一个makefile文件，是k8s的构建方式，我们会分析一下k8s的构建过程，所以我们学习一下makefile。

> makefile，命令执行时，需要一个makefile文件，以告诉make命令需要怎么样的去编译和链接程序

执行：

```bash
yum install -y gcc gcc-c++
```

## 1、语法规则

一个Makefile由一组规则组成，规则通常如下所示：

```makefile
targets: prerequisites
	command
	command
	command
```

- 目标 `targets` 是文件名，以空格分开，通常，每条规则只有一个。
- 命令 `command` 是生成目标的一系列步骤，以制表符Tab开头，不能用空格开头
- 先决条件 `prerequisites` 是文件名(也称为依赖项)，以空格分隔，这些文件需要在执行命令之前存在

上述其实是描述了一个文件依赖关系，即生成一个或多个targets依赖于prerequisites，生成规则定义在

command中。

实例：

```makefile
[root@localhost ~]# mkdir mak
[root@localhost ~]# vi makefile

# 注释
blah: blah.o
	cc blah.o -o blah #第三步
blah.o: blah.c
	cc -c blah.c -o blah.o #第二步
blah.c: 
	echo "int main() { return 0; }" > blah.c #第一步
```

执行

```bash
[root@localhost mak]# make blah
```

## 2、all

如果makefile中定义了多个目标，通常定义一个all目标来生成所有目标

```makefile
all: one two three
one: 
	touch one
two: 
	touch two
three: 
	touch three
clean: 
	rm -f one two three
```

执行

```bash
[root@localhost mak]# make all
[root@localhost mak]# make clean
```

> makefile 中的第一个目标会被作为其默认目标，直接使用make即可，无需指定target

### 2.1、伪目标

```makefile
clean: 
	rm *.o
```

上述的clean并不会真的生成一个clean文件，所以我们称为clean为伪目标，伪目标的取名不能和文件名重名

为了避免和文件重名的这种情况，我们可以使用一个特殊的标记`".PHONY"` 来显示的指明一个目标是伪目标，向make说明，不管是否有这个文件，这个目标就是伪目标

示例：

```makefile
.PHONY: clean
clean: 
	rm -f one two three
```

>  `".PHONY"` 不写其实也会推导出clean是一个伪目标，不去生成文件，但显示的表明clean是一个伪目标是一个好的习惯
>
> 目标可以成为依赖，伪目标同样可以成为依赖

### 2.2、多行

当命令太长时，反斜杠`("\")`字符使我们能够使用多行

```makefile
some_file: 
	echo This line is too long,so \
	it is broken up into multiple lines
```

## 3、变量

变量是把一个名字和任意长的字符串关联起来，基本语法如下：

```makefile
MY = A text string
MK := K text string
```

使用 ${}或$() 来引用变量

```makefile
MA = A text string
all: 
	echo ${MA}
	echo $(MA)
```

```makefile
MB = file1.c file2.c
all: 
	echo ${MB}
	echo $(MB)
```

### 3.1、求值时机

- `=` 仅在使用命名时解析变量值
- `:= ` 在定义时立即解析变量值

```makefile
# 在下面的echo命令执行时在求值，输出 "later"
one = one ${later_variable}

#简单的扩展变量，由于 "later_variable" 未定义，下面不会输出 "later"
two := two ${later_variable}

later_variable = later

all: 
	echo ${one}
	echo ${two}
```

`=` 被称为是递归变量，因为其在使用命名时解析，所以不能进行递归定义

`:=` 被称为简单扩展变量

```makefile
#简单的扩展变量,若将 := 改为 = 则产生无限循环错误
one = ${one} there
all: 
	echo ${one}
```

```bash
[root@localhost mak]# make
makefile:1: *** Recursive variable `one' references itself (eventually).  Stop.
```

### 3.2、是否覆盖变量

`?=` 仅在尚未设置变量时设置变量

```makefile
one = hello
one ?= will not be set
two ?= will be set
all: 
	echo ${one}
	echo ${two}
```

### 3.3、空格

行尾的空格不会被删除，但开头的空格会被删除

```makefile
with_spaces = hello   #with_spaces 变量是 hello 末尾有三个空格
after = ${with_spaces}there

nullstring = 
space = ${nullstring} #创建只有一个空格的变量

all: 
	echo "$(after)"
	echo start"$(space)"end
```

> 未定义的变量实际上是一个空字符串

```makefile
all: 
	#未定义的变量实际上是一个空字符串
	echo $(nowhere)
```

### 3.4、追加

`+=` 用于追加

```makefile
foo := start
foo += more
all: 
	echo ${foo}
```

### 3.5 自动变量

makefile定义了一些自动变量，用于自动获取一些值，比如

```makefile
all: f1.o f2.o

f1.o f2.o: 
	echo $@ #比较常用 相当于一个集合，依次取出并执行命令
#相当于：
f1.o: 
	echo f1.o
f2.o: 
	echo f2.o
```

- `$@` 规则目标的文件名
- `$< ` 第一个先决条件的名称
- `$?` 比目标新的所有先决条件的名称，他们之间有空格
- `$^` 所有先诀条件的名称，他们之间有空格

```makefile
hey: one two
	echo $@ #输出hey
	echo $? #所有比目标新的先决条件 one two
	echo $^ #所有先决条件
	touch hey
one: 
	touch one
two: 
	touch two
clean: 
	rm -f hey one two
```

## 4、通配符

`*` 称为通配符

- `*` 可以在目标，先决条件或 wildcard函数（查找指定目录下指定的类型文件）中使用
- `*` 不能在变量中直接使用
- 当 `*` 没有匹配到文件时，保持原样（除非在 wildcard函数 中执行）

```makefile
thing_wrong := *.o #Don't do this output *.o
thing_right := $(wildcard *.o)
```

```makefile
thing_wrong := *.o
thing_right := $(wildcard *.o)
all: one two three four
one: 
	echo $(thing_wrong)
two: *.o
	echo $^
three: 
	echo $(thing_right)
four: 
	echo $(wildcard *.o)
```

如果我们的文件名中有通配符，如：`*`，那么可以用转义字符 `\` ,如 `\*` 来表示真实的`*`字符，而不是任意的字符串

## 5、隐式规则

make会有一些默认的约定，来帮助我们简化书写

```makefile
blah: blah.o
	cc blah.o -o blah #第三步
blah.o: blah.c
	cc -c blah.c -o blah.o #第二步
blah.c: 
	echo "int main() { return 0; }" > blah.c #第一步
```

之前我们会经过`blah.o`这一步，接下来我们省略这一步

```makefile
blah: blah.o
	cc blah.o -o blah #第三步
blah.c: 
	echo "int main() { return 0; }" > blah.c #第一步
```

```bash
[root@localhost mak]# make
echo "int main() { return 0; }" > blah.c #第一步
cc    -c -o blah.o blah.c
cc blah.o -o blah #第三步
```

执行后，我们发现多了一步`cc    -c -o blah.o blah.c`，和我们之前写的一样，这种根据依赖自动推导的规则就是`隐式规则`

隐式规则：

- 编译`C`程序时，会自动使用c的编译命令`$(CC) -c $(CPPFLAGS) $(CFLAGS)` 来生成`.O`,比如遇到blah.o依赖，那么会自动找到blah.c运行命令进行生成
- 编译`C++`程序时，会自动使用c的编译命令`$(CXX) -c $(CPPFLAGS) $(CXXFLAGS)`来生成`.O`,比如遇到blah.o依赖，那么会自动找到blah.c或者blah.cpp运行命令进行生成

如果我们将上述的程序修改为：

```makefile
blah: blah.o
blah.c: 
	echo "int main() { return 0; }" > blah.c #第一步
```

```makefile
[root@localhost mak]# make
echo "int main() { return 0; }" > blah.c #第一步
cc    -c -o blah.o blah.c
cc   blah.o   -o blah
```

我们发现还能正常执行，这又是另一个隐式规则

- `<n>` 目标依赖于`<n>.o`,通过运行C的编译器来运行链接程序生成，其生成命令是：`$(CC) $(LDFLAGS) <n>.o $(LOADLIBES) $(LDLIBS)`

## 6、静态模式规则

静态模式可以更容易的定义多目标规则，可以让我们的规则变得更加的有弹性和灵活

静态模式规则的语法：

```makefile
targets...:= target-pattern: prereq-patterns ...
	commands
```

匹配`target-pattern` 生成 `targets`,匹配`prereq-patterns`生成`target-pattern`

比如我们编译一系列的.c文件到.o文件

```makefile
objects = foo.o bar.o all.o
all: $(objects)
foo.o: foo.c
bar.o: bar.c
all.o: all.c
all.c: 
	echo "int main() {return 0;}" > all.c
%.c: 
	touch $@
clean: 
	rm -f *.c *.o all
```

使用静态规则模式后：

```makefile
objects = foo.o bar.o all.o
all: $(objects)
$(objects): %.o: %.c
all.c: 
	echo "int main() {return 0;}" > all.c
%.c: 
	touch $@
clean: 
	rm -f *.c *.o all
```

### 6.1、filter

filter函数可用于静态模式规则以匹配正确的文件

```makefile
obj_files = foo.result bar.o lose.o
src_files = foo.raw bar.c lose.c

.PHONY: all
all: $(obj_files)

$(filter %.o, $(obj_files)): %.o: %.raw
	echo "target: $@ prereq: $<"
$(filter %.result, $(obj_files)): %.result: %.raw
	echo "target: $@ prereq: $<"
%.c %.raw: 
	touch $@
clean: 
	rm -f $(src_files)
```

## 7、双冒号规则

双冒号规则很少使用，但允许为同一个目标定义多个规则

如果这些是单冒号，则会打印一条警告，并且只会运行第二组命令

```makefile
all: blah
blah:: 
	echo "hello"
blah:: 
	echo "hello again"
```

## 8、命令

### 8.1、显示与隐藏

在命令之前添加一个 `@` 以阻止它被打印

你也可以运行 make 时使用 `-s` 参数，这将为每一行命令添加一个`@`

```makefile
all: 
	@echo "This make line will not be printed"
	echo "But this will"
```

### 8.2、执行

每个命令都在一个新的shell中运行(或者至少效果是这样的)

```makefile
all: 
	cd ..
	echo `pwd` #cd不会影响pwd 因为不在一行
	cd ..;echo `pwd`  #cd会影响pwd 因为在一行
	cd ..; \
	echo `pwd`
```

### 8.3、默认shell

默认shell是 `/bin/sh`,你可以通过更改变量SHELL来更改它

```makefile
SHELL=/bin/bash
```

### 8.4、递归使用make

要递归调用makefile，请使用`$(MAKE)`变量代替`make`

它会为你传递make标志并且本身不会受到它们的影响

```makefile
new_-contents = "hello:\n\ttouch inside_file"
all: 
	mkdir -p subdir
	printf $(new_-contents) | sed -e 's/^ //' > subdir/Makefile
	cd subdir && %(MAKE)
clean: 
	rm -f subdir 
```

### 8.5、创建多个目标

```makefile
make clean run test 运行 clean目标，然后run，然后test
```

### 8.6 define

define 实际上就是命令列表

```makefile
one = export blah="I was set!"; echo $$blah
define two
export blah=set
echo $$blah
endef
all: 
	@echo "This prints 'I was set!'"
	@$(one)
	@echo "This does not print 'I was set!' because each command runs in a separate shell"
	@$(two)
```

### 8.7、特定目标变量

```makefile
all: one = cool
all: 
	echo one is defined: $(one)
other: 
	echo one is nothing: $(one)
```

```makefile
%.c: one = cool
blah.c: 
	echo one is defined: $(one)
other: 
	echo one is nothing: $(one)
```

## 9、条件

### 9.1、if else

```makefile
foo = ok
all: 
ifeq ($(foo), ok)
	echo "foo equals ok"
else
	echo "noqe"
endif
```

判空

```makefile
nullstring = 
foo = $(nullstring)
all: 
ifeq ($(strip $(foo)),)
	echo "foo is empty after being stripped"
endif
```

判断变量是否定义

```makefile
bar = 
foo = $(bar)
all: 
ifdef foo
	echo "foo is defined"
endif
```

## 10、函数

函数主要用于文本处理，使用 `$(fn, arguments)` 或 `${fn,arguments}` 调用函数

### 10.1、subst

用法是 `$(subst FROM,TO,TEXT)`, 即将`TEXT`中的东西从`FROM`变为`TO`

```makefile
bar := ${subst not, totally, "I am superman"}
all: 
	@echo $(bar)
```

```makefile
comma := ,
empty := 
space := $(empty) $(empty)
foo := a b c
bar := $(subst $(space),$(comma),$(foo)) #注意$(foo)前面不要有空格 否则会当做字符串的一部分
all: 
	@echo $(bar)
```

### 10.2、pathsubst

格式字符串替换函数

格式：$(patsubst <pattern>,<replacement>,<text>)

查找<text>中单词是否符合模式<pattern>,如果匹配的话，则以<replacement>替换

替换引用${text:pattern=replacement}是对此的简写

```makefile
foo := a.o b.o l.a c.o
one := $(patsubst %.o,%.c,$(foo)) # %代表任意字符
two := $(foo:%.o=%.c)
three := $(foo:.o=.c)
all: 
	echo $(one)
	echo $(two)
	echo $(three)
```

### 10.3、foreach

$(foreach var,list,text)：将一个单词列表(由空格分隔)转换为另一个单词列表

list 代表单词列表，var 设置列表中每个单词， text 针对每个单词进行扩展

```makefile
foo := wha are you
bar := $(foreach wrd,$(foo),$(wrd)!)
all: 
	@echo $(bar)
```

### 10.4、if

if 检查第一个参数是否为非空，如果是，则运行第二个参数，否则运行第三个

```makefile
foo := $(if this-is-not-empty,then!,else!)
empty := 
bar := $(if $(empty),then!,else!)
all: 
	@echo $(foo)
	@echo $(bar)
```

### 10.5、call

$(call VARIABLE,PARAM,PARAM,...) ：在执行时，将它的参数 "PARAM"依次赋给VARIABLE中的临时变量 “$(1) ","$(2)”等

$(0)是获取VARIABLE变量名称

```makefile
sweet_new_fn = Variable Name: $(0) First: $(1), Second: $(2) Empty Variable: $(3)
all: 
	@echo $(call sweet_new_fn, go, tigers)
```

### 10.6、shell

shell函数就是调用shell

```makefile
all: 
	@echo $(shell ls -la)
```

## 11、包含

includde 指令告诉 make 读取一个或多个其他makefile

```makefile
include filenames...
```

## 12、vpath

使用 vpath 指定某些先决条件存在的位置

格式：vpath <pattern> <directories, space/colon separated>

<pattern>可以有一个%，它匹配任何零个或多个字符

```makefile
vpath $.h ../headers
```

代表要求make在"../headers"目录下搜索所有以`.h`结尾的文件。前提是当前目前没有找到

```makefile
vpath $.h ../headers
some_binary: blah.h
	touch some_binary
blah: 
	mkdir ../headers
	touch ../headers/blah.h
clean: 
	rm -rf ../headers
	rm -f some_binary
```

