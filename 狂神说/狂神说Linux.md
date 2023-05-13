# Linux

- `~`：表示目前所在目录为家目录，其中 `root` 用户的家目录是 `/root` ；普通用户的家目录在 `/home` 下；
- `#`：指示你所具有的权限（ `root` 用户为 `#` ，普通用户为 `$` ）
- 执行 `whoami` 命令可以查看当前用户名；

![image.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/226f8a87a9804141802d5ba0a55bd1f1~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp)

### 开关机

```shell
sync #将数据由内存同步到硬盘中
shutdown #关机指令
shutdown -h 10 #10分钟后关机
shutdown -h now #立马关机
shutdown -h 20:25 #在当天的20:25点关机
shutdown -h +10 #10分钟后关机
shutdown -r now #系统立马重启
shutdown -r +10 #系统10分钟后重启
reboot #重启
halt #关闭系统
poweroff #关机
```

### 目录管理

```shell
#绝对路径：
路径的写法，由根目录 / 写起，例如：/usr/share/doc 这个目录

#相对路径：
路径的写法，不是由 / 写起，例如由 /usr/share/doc 要到 /usr/share/man 底下时，可以写成：cd ../man 

#使用 man [命令] 来查看各个命令的使用文档，如 ：man cp
ls: 列出目录

cd：切换目录

pwd：显示目前的目录

mkdir：创建一个新的目录
    -m ：配置文件的权限喔！直接配置，不需要看默认权限 (umask) 的脸色～
    mkdir -m 711 test2
    
    -p ：帮助你直接将所需要的目录(包含上一级目录)递归创建起来！
	mkdir -p test1/test2/test3/test4
	
rmdir：删除一个空的目录
	-p ：连同上一级『空的』目录也一起删除
	rmdir -p test1/test2/test3/test4

cp: 复制文件或目录
	-a：相当於 -pdr 的意思，至於 pdr 请参考下列说明；(常用)
    -p：连同文件的属性一起复制过去，而非使用默认属性(备份常用)；
    -d：若来源档为连结档的属性(link file)，则复制连结档属性而非文件本身；
    -r：递归持续复制，用於目录的复制行为；(常用)
    -f：为强制(force)的意思，若目标文件已经存在且无法开启，则移除后再尝试一次；
    -i：若目标档(destination)已经存在时，在覆盖时会先询问动作的进行(常用)
    -l：进行硬式连结(hard link)的连结档创建，而非复制文件本身。
    -s：复制成为符号连结档 (symbolic link)，亦即『捷径』文件；
    -u：若 destination 比 source 旧才升级 destination ！
    
    # 复制 root目录下的install.sh 到 home目录下
	[root@kuangshen home]# cp /root/install.sh /home
	
	# 再次复制，加上-i参数，增加覆盖询问？
    [root@kuangshen home]# cp -i /root/install.sh /home
    cp: overwrite ‘/home/install.sh’? y # n不覆盖，y为覆盖

rm: 移除文件或目录
    -f ：就是 force 的意思，忽略不存在的文件，不会出现警告信息；
    -i ：互动模式，在删除前会询问使用者是否动作
	[root@kuangshen home]# rm -i install.sh
    rm: remove regular file ‘install.sh’? y
    -r ：递归删除啊！最常用在目录的删除了！这是非常危险的选项！！！

mv: 移动文件与目录，或修改文件与目录的名称
    -f ：force 强制的意思，如果目标文件已经存在，不会询问而直接覆盖；
    -i ：若目标文件 (destination) 已经存在时，就会询问是否覆盖！
    -u ：若目标文件已经存在，且 source 比较新，才会升级 (update)
    # 将文件移动到我们创建的目录
    [root@kuangshen home]# mv install.sh test
    
    # 将文件夹重命名
    [root@kuangshen home]# mv test mvtest
```

### 基本属性

在Linux中第一个字符代表这个文件是目录、文件或链接文件等等

```shell
当为[ d ]则是目录
当为[ - ]则是文件；
若是[ l ]则表示为链接文档 ( link file )；
若是[ b ]则表示为装置文件里面的可供储存的接口设备 ( 可随机存取装置 )；
若是[ c ]则表示为装置文件里面的串行端口设备，例如键盘、鼠标 ( 一次性读取装置 )
#接下来的字符中，以三个为一组，且均为『rwx』 的三个参数的组合。
其中，[ r ]代表可读(read)、[ w ]代表可写(write)、[ x ]代表可执行(execute)。
要注意的是，这三个权限的位置不会改变，如果没有权限，就会出现减号[ - ]而已。

第0位确定文件类型，第1-3位确定属主（该文件的所有者root）拥有该文件的权限。第4-6位确定属组（所有者的同组用户group）拥有该文件的权限，第7-9位确定其他用户(user)拥有该文件的权限。
```

```shell
chgrp [-R] 属组名 文件名 #更改文件属组
	-R：递归更改文件属组
	chgrp group1 install.log
	chgrp bar file.txt	--> file.txt文件的群组修改为bar
	
chown [–R] 属主名 文件名 #更改文件属主，也可以同时更改文件属组
chown [-R] 属主名：属组名 文件名
    -c : 显示更改的部分的信息
    -f : 忽略错误信息
    -h :修复符号链接
    -v : 显示详细的处理信息
    -R : 处理指定目录以及其子目录下的所有文件
	chown caticUser -R /home/catic
	// 将catic下的所有文件设置为caticUser拥有
	chown lion file.txt	--> 把其它用户创建的file.txt转让给lion用户
	chown lion:bar file.txt	--> 把file.txt的用户改为lion，群组改为bar

chmod：更改文件的属性
	chmod [-R] xyz 文件或目录
	r:4  w:2  x:1
	每种身份(owner/group/others)各自的三个权限(r/w/x)分数是需要累加的，例如当权限为：[-rwxrwx---] 分数则是：
    owner = rwx = 4+2+1 = 7
    group = r_x = 4+0+1 = 5
    group = rw_ = 4+2+0 = 6
    others= --- = 0+0+0 = 0
    chmod 770 filename
```

### chmod命令详解

Linux/Unix 的文件调用权限分为三级 :文件所有者（Owner）、用户组（Group）、其它用户（Other Users）

![img](https://www.runoob.com/wp-content/uploads/2014/08/file-permissions-rwx.jpg)

![img](https://www.runoob.com/wp-content/uploads/2014/08/rwx-standard-unix-permission-bits.png)

### 文件内容

```shell
语法：chmod (u g o a) (+ - =) (r w x) (文件名)

u user 表示该文件的所有者
g group 表示与该文件的所有者属于同一组( group )者，即用户组
o other 表示其它用户组
a all 表示这三者皆是

+增加权限
- 撤销权限
= 设定权限

r read 表示可读取，对于一个目录，如果没有r权限，那么就意味着不能通过ls查看这个目录的内容。
w write 表示可写入，对于一个目录，如果没有w权限，那么就意味着不能在目录下创建新的文件。
x excute 表示可执行，对于一个目录，如果没有x权限，那么就意味着不能通过cd进入这个目录。

chmod u=rwx /etc/hosts
chmod u+r /etc/hosts
chmod u+rx /etc/hosts
chmod u-x /etc/hosts
chmod u-x,o+rw /etc/hosts

chmod u+rx file	--> 文件file的所有者增加读和运行的权限
chmod g+r file	--> 文件file的群组用户增加读的权限
chmod o-r file	--> 文件file的其它用户移除读的权限
chmod g+r o-r file	--> 文件file的群组用户增加读的权限，其它用户移除读的权限
chmod go-r file	--> 文件file的群组和其他用户移除读的权限
chmod +x file	--> 文件file的所有用户增加运行的权限
chmod u=rwx,g=r,o=- file	--> 文件file的所有者分配读写和执行的权限，群组其它用户分配读的权限，其他用户没有任何权限

#操作文件夹需要加入 -R 参数
chmod -R 707 文件夹名称 

#-R 可以递归地修改文件访问权限
chown -R lion:lion /home/frank
```

| #    | 权限           | rwx  | 二进制 |
| :--- | :------------- | :--- | :----- |
| 7    | 读 + 写 + 执行 | rwx  | 111    |
| 6    | 读 + 写        | rw-  | 110    |
| 5    | 读 + 执行      | r-x  | 101    |
| 4    | 只读           | r--  | 100    |
| 3    | 写 + 执行      | -wx  | 011    |
| 2    | 只写           | -w-  | 010    |
| 1    | 只执行         | --x  | 001    |
| 0    | 无             | ---  | 000    |

### 查看文件

```shell
cat 由第一行开始显示文件内容
    -A ：相当於 -vET 的整合选项，可列出一些特殊字符而不是空白而已；
    -b ：列出行号，仅针对非空白行做行号显示，空白行不标行号！
    -E ：将结尾的断行字节 $ 显示出来；
    -n ：列印出行号，连同空白行也会有行号，与 -b 的选项不同；
    -T ：将 [tab] 按键以 ^I 显示出来；
    -v ：列出一些看不出来的特殊字符
    #cat /etc/redis.conf

tac 从最后一行开始显示，可以看出 tac 是 cat 的倒着写！
	#tac /etc/redis.conf

nl  显示的时候，顺道输出行号！
    -b ：指定行号指定的方式，主要有两种：-b a ：表示不论是否为空行，也同样列出行号(类似 cat -n)；-b t ：如果有空行，空的那一行不要列出行号(默认值)；
    -n ：列出行号表示的方法，主要有三种：-n ln ：行号在荧幕的最左方显示；-n rn ：行号在自己栏位的最右方显示，且不加 0 ；-n rz ：行号在自己栏位的最右方显示，且加 0 ；
    -w ：行号栏位的占用的位数。
    #nl /etc/redis.conf
    
more 一页一页的显示文件内容
    空白键 (space)：代表向下翻一页；
    Enter     ：代表向下翻『一行』；
    /字串     ：代表在这个显示的内容当中，向下搜寻『字串』这个关键字；
    :f      ：立刻显示出档名以及目前显示的行数；
    q       ：代表立刻离开 more ，不再显示该文件内容。
    b 或 [ctrl]-b ：代表往回翻页，不过这动作只对文件有用，对管线无用。
    #more /etc/csh.login

less 与 more 类似，但是比 more 更好的是，他可以往前翻页！
    空白键  ：向下翻动一页；
    [pagedown]：向下翻动一页；
    [pageup] ：向上翻动一页；
    /字串   ：向下搜寻『字串』的功能；
    ?字串   ：向上搜寻『字串』的功能；
    n     ：重复前一个搜寻 (与 / 或 ? 有关！)
    N     ：反向的重复前一个搜寻 (与 / 或 ? 有关！)
    q     ：离开 less 这个程序；
	#less /etc/csh.login

head 只看头几行
	-n 后面接数字，代表显示几行的意思！
	# head -n 20 /etc/csh.login

tail 只看尾巴几行
	-f 循环读取
    -n ：后面接数字，代表显示几行的意思
    # tail -n 20 /etc/csh.login
    # tail -f notes.log
```

### 用户管理

```shell
hostname 主机名
hostname kuangshen #修改当前主机名

useradd 选项 用户名
    -c comment 指定一段注释性描述。
    -d 目录 指定用户主目录，如果此目录不存在，则同时使用-m选项，可以创建主目录。
    -g 用户组 指定用户所属的用户组。
    -G 用户组，用户组 指定用户所属的附加组。
    -m　使用者目录如不存在则自动建立。
    -s Shell文件 指定用户的登录Shell。
    -u 用户号 指定用户的用户号，如果同时有-o选项，则可以重复使用其他用户的标识号。
    用户名 :指定新账号的登录名
    # useradd -m kuangshen
    增加用户账号就是在/etc/passwd文件中为新用户增加一条记录，同时更新其他系统文件如/etc/shadow, /etc/group等
    
    1.切换用户的命令为：su username 【username是你的用户名哦】
    2.从普通用户切换到root用户，还可以使用命令：sudo su
    3.在终端输入exit或logout或使用快捷方式ctrl+d，可以退回到原来用户，其实ctrl+d也是执行的exit命令
    4.在切换用户时，如果想在切换用户之后使用新用户的工作环境，可以在su和username之间加-，例如：【su - root】

$表示普通用户

#表示超级用户，也就是root用户
    
userdel 选项 用户名
	-r，它的作用是把用户的主目录一起删除。
	# userdel -r kuangshen
	
usermod 选项 用户名
    常用的选项包括-c, -d, -m, -g, -G, -s, -u以及-o等，这些选项的意义与useradd命令中的选项一样，可以为用户指定新的资源值。
    # usermod -s /bin/ksh -d /home/z –g developer kuangshen

passwd 选项 用户名
    -l 锁定口令，即禁用账号。
    -u 口令解锁。
    -d 使账号无口令。
    -f 强迫用户下次登录时修改口令。
    # passwd kuangshen
    # passwd -d kuangshen
    # passwd -l kuangshen
```

### 用户组管理

```shell
groupadd 选项 用户组
    -g GID 指定新用户组的组标识号（GID）。
    -o 一般与-g选项同时使用，表示新用户组的GID可以与系统已有用户组的GID相同。
    # groupadd group1
    # groupadd -g 101 group2 增加了一个新组group2，同时指定新组的组标识号是101

groupdel 用户组
	# groupdel group1

groupmod 选项 用户组
    -g GID 为用户组指定新的组标识号。
    -o 与-g选项同时使用，用户组的新GID可以与系统已有用户组的GID相同。
    -n新用户组 将用户组的名字改为新名字
    # 此命令将组group2的组标识号修改为102。
    groupmod -g 102 group2
    # 将组group2的标识号改为10000，组名修改为group3。
    groupmod –g 10000 -n group3 group2

newgrp root #切换组,将当前用户切换到root用户组
```

### 磁盘管理

```shell
Linux磁盘管理常用命令为 df、du。
    df ：列出文件系统的整体磁盘使用量
    du：检查磁盘空间使用量

df [可选命令行] [目录或文件名]
    -a ：列出所有的文件系统，包括系统特有的 /proc 等文件系统；
    -k ：以 KBytes 的容量显示各文件系统；
    -m ：以 MBytes 的容量显示各文件系统；
    -h ：以人们较易阅读的 GBytes, MBytes, KBytes 等格式自行显示；
    -H ：以 M=1000K 取代 M=1024K 的进位方式；
    -T ：显示文件系统类型, 连同该 partition 的 filesystem 名称 (例如 ext3) 也列出；
    -i ：不用硬盘容量，而以 inode 的数量来显示
    # df
    # df -h /etc
    # df -h

du [可选命令行] 文件或目录名称
    -a ：列出所有的文件与目录容量，因为默认仅统计目录底下的文件量而已。
    -h ：以人们较易读的容量格式 (G/M) 显示；
    -s ：列出总量而已，而不列出每个各别的目录占用容量；
    -S ：不包括子目录下的总计，与 -s 有点差别。
    -k ：以 KBytes 列出容量显示；
    -m ：以 MBytes 列出容量显示；
    # du
    # du -a
    # du -sm /*

磁盘挂载与卸除
mount [-t 文件系统] [-L Label名] [-o 额外选项] [-n] 装置文件名 挂载点
    # 将 /dev/hdc6 挂载到 /mnt/hdc6 上面！
    [root@www ~]# mkdir /mnt/hdc6
    [root@www ~]# mount /dev/hdc6 /mnt/hdc6

umount [-fn] 装置文件名或挂载点
    -f ：强制卸除！可用在类似网络文件系统 (NFS) 无法读取到的情况下；
    -n ：不升级 /etc/mtab 情况下卸除。
    # umount /dev/hdc6
```

### 进程管理

```shell
ps 查看当前系统中正在执行的各种进程信息
	-a 显示当前终端运行的所有的进程信息
	-u 以用户的信息显示进程
	-x 显示后台运行进程的参数
	
ps -aux 查看所有的进程
	ps -aux|grep mysql
	
ps -ef 查看父进程的信息
	ps -ef|grep mysql

pstree -pu
	-p 显示父id
	-u 显示用户组
	
kill -9 进程id #结束进程
    kill 956 # 结束进程号为956的进程
    kill 956 957 # 结束多个进程
    kill -9 7291 # 强制结束进程
-efH 以乔木状列举出所有进程;
-aux --sort -pcpu 按 CPU 使用降序排列， -aux --sort -pmem 表示按内存使用降序排列;
-axjf 以树形结构显示进程， ps -axjf 它和 pstree 效果类似。
```

### jdk安装

```shell
# 安装完成后配置环境变量 文件：/etc/profile
JAVA_HOME=/usr/java/jdk1.8.0_221-amd64
CLASSPATH=%JAVA_HOME%/lib:%JAVA_HOME%/jre/lib
PATH=$PATH:$JAVA_HOME/bin:$JAVA_HOME/jre/bin
export PATH CLASSPATH JAVA_HOME
# 保存退出

# 让新增的环境变量生效！
source /etc/profile

[root@kuangshen java]# java -version
java version "1.8.0_221"
```

# 稀土掘金Linux

### which

```
查看命令的可执行文件所在路径
[root@localhost ~]# which python
/usr/bin/python
```

### du

```shell
列举目录大小信息。
【常用参数】
-h 适合人类阅读的；
-a 同时列举出目录下文件的大小信息；
-s 只显示总计大小，不显示具体信息。
[root@localhost ~]# du -h nginx-1.23.3
48K     nginx-1.23.3/auto/cc
```

### touch

```shell
创建一个文件
touch new_file
```

### ln

```shell
表示创建链接,Linux 下有两种链接类型：硬链接和软链接。
ln file1 file2 	--> 创建 file2 为 file1 的硬链接

使链接的两个文件共享同样文件内容，就是同样的 inode ，一旦文件1和文件2之间有了硬链接，那么修改任何一个文件，修改的都是同一块内容，它的缺点是，只能创建指向文件的硬链接，不能创建指向目录的（其实也可以，但比较复杂）而软链接都可以，因此软链接使用更加广泛。
如果我们用 rm file1 来删除 file1 ，对 file2 没有什么影响，对于硬链接来说，删除任意一方的文件，共同指向的文件内容并不会从硬盘上删除。只有同时删除了 file1 与 file2 后，它们共同指向的文件内容才会消失。

软链接,软链接就类似 windows 下快捷方式。
ln -s file1 file2
其实 file2 只是 file1 的一个快捷方式，它指向的是 file1 ，所以显示的是 file1 的内容，但其实 file2 的 inode  与 file1 并不相同。如果我们删除了 file2 的话， file1 是不会受影响的，但如果删除 file1 的话， file2 就会变成死链接，因为指向的文件不见了。
```

### sudo

```shell
以 root 身份运行命令
sudo date  --> 当然查看日期是不需要sudo的这里只是演示，sudo 完之后一般还需要输入用户密码的
```

### su

```shell
切换用户，需要 root 用户权限
sudo su	--> 切换为root用户（exit 命令或 CTRL + D 快捷键都可以使普通用户切换为 root 用户）
su lion	--> 切换为普通用户
su -	--> 切换为root用户
```

### locate

```shell
搜索包含关键字的所有文件和目录。后接需要查找的文件名，也可以用正则表达式
yum -y install mlocate	--> 安装包
updatedb	--> 更新数据库

locate file.txt
locate fil*.txt
[注意] locate 命令会去文件数据库中查找命令，而不是全磁盘查找，因此刚创建的文件并不会更新到数据库中，所以无法被查找到，可以执行 updatedb 命令去更新数据库。
```

### find

```shell
用于查找文件，它会去遍历你的实际硬盘进行查找，而且它允许我们对每个找到的文件进行后续操作，功能非常强大。
find <何处> <何物> <做什么>
何处：指定在哪个目录查找，此目录的所有子目录也会被查找。
何物：查找什么，可以根据文件的名字来查找，也可以根据其大小来查找，还可以根据其最近访问时间来查找。
做什么：找到文件后，可以进行后续处理，如果不指定这个参数， find 命令只会显示找到的文件。
```

#### 根据文件名查找

```shell
find -name "file.txt"	--> 当前目录以及子目录下通过名称查找文件
find . -name "syslog"	--> 当前目录以及子目录下通过名称查找文件
find / -name "syslog"	--> 整个硬盘下查找syslog
find /var/log -name "syslog"	--> 在指定的目录/var/log下查找syslog文件
find /var/log -name "syslog*"	--> 查找syslog1、syslog2 ... 等文件，通配符表示所有
find /var/log -name "*syslog*"	--> 查找包含syslog的文件 
[注意] find 命令只会查找完全符合 “何物” 字符串的文件，而 locate 会查找所有包含关键字的文件。
```

#### 根据文件大小查找

```shell
find /var -size +10M	--> /var 目录下查找文件大小超过 10M 的文件
find /var -size -50k	--> /var 目录下查找文件大小小于 50k 的文件
find /var -size +1G	--> /var 目录下查找文件大小查过 1G 的文件
find /var -size 1M	--> /var 目录下查找文件大小等于 1M 的文件
```

#### 根据文件最近访问时间查找

```shell
find -name "*.txt" -atime -7 	--> 近 7天内访问过的.txt结尾的文件
```

#### 仅查找目录或文件

```shell
find . -name "file" -type f 	--> 只查找当前目录下的file文件
find . -name "file" -type d 	--> 只查找当前目录下的file目录
```

#### 操作查找结果

```shell
find -name "*.txt" -printf "%p - %u\n"	--> 找出所有后缀为txt的文件，并按照 %p - %u\n 格式打印，其中%p=文件名，%u=文件所有者
find -name "*.jpg" -delete	--> 删除当前目录以及子目录下所有.jpg为后缀的文件，不会有删除提示，因此要慎用
find -name "*.c" -exec chmod 600 {} \;	--> 对每个.c结尾的文件，都进行 -exec 参数指定的操作，{} 会被查找到的文件替代，\; 是必须的结尾
find -name "*.c" -ok chmod 600 {} \;	--> 和上面的功能一直，会多一个确认提示
```

# 软件仓库

`Linux` 下软件是以包的形式存在，一个软件包其实就是软件的所有文件的压缩包，是二进制的形式，包含了安装软件的所有指令。 `Red Hat` 家族的软件包后缀名一般为 `.rpm` ， `Debian` 家族的软件包后缀是 `.deb` 。

`Linux` 的包都存在一个仓库，叫做软件仓库，它可以使用 `yum` 来管理软件包， `yum` 是 `CentOS` 中默认的包管理工具，适用于 `Red Hat` 一族。可以理解成 `Node.js` 的 `npm` 。

### yum 常用命令

```shell
`yum update | yum upgrade` 更新软件包 yum -y update
`yum search xxx` 搜索相应的软件包
`yum install xxx` 安装软件包
`yum remove xxx` 删除软件包
```

### 切换 CentOS 软件源

有时候 `CentOS` 默认的 `yum` 源不一定是国内镜像，导致 `yum` 在线安装及更新速度不是很理想。这时候需要将 `yum` 源设置为国内镜像站点。国内主要开源的镜像站点是网易和阿里云。

1、首先备份系统自带 `yum` 源配置文件 `mv /etc/yum.repos.d/CentOS-Base.repo /etc/yum.repos.d/CentOS-Base.repo.backup` 

2、下载阿里云的 `yum` 源配置文件到 `/etc/yum.repos.d/CentOS7` 

 wget -*O*选项表示将下载的内容进行重命名

```shell
wget -O /etc/yum.repos.d/CentOS-Base.repo http://mirrors.aliyun.com/repo/Centos-7.repo
```

3、生成缓存

```shell
yum makecache
```

### rpm

RPM（[Redhat](https://so.csdn.net/so/search?q=Redhat&spm=1001.2101.3001.7020) Package Manager）是用于Redhat、CentOS、Fedora等Linux 分发版（distribution）的常见的软件包管理器。rpm工具可以用来制作源码安装包和二进制安装包

#### 一、rpm是什么

　　rpm是一种安装包的格式。就像在Windows系统上我们常见的安装包格式是exe和msi一样，在linux上常见的安装包格式是deb和rpm。一般在红帽系列的系统上，不支持deb，所以我们需要将程序打包成rpm安装。

#### 二、rpm的打包工具

　　rpm的打包，我们需要用到：rpmbuild 和 rpmdevtools

　　有的系统预装的 rpmbuild，这样我们就不需要安装这个了，可以用使用命令检查系统是否有安装 rpmbuild

```shell
# 检查是否有 rpmbuild
rpmbuild --version

# 安装 rpmbuild
yum install rpm-build
 
# 安装 rpmdevtools
yum install rpmdevtools
```

#### 三、创建打包文件夹

　    创建打包文件夹有两种方法：

　　方法一：使用命令 rpmdev-setuptree 来创建，但是创建的文件夹是在用户主目录（home），我们可以将rpmbuild整个文件夹拷贝到项目文件夹，这样可以方便我们项目管理。

　　方法二：手动的去创建文件夹，具体文件夹结构可以参考下面的目录树

```shell
# 方法一：使用命令创建
rpmdev-setuptree

# 目录树
rpmbuild
    - BUILD           // 编译时用到的暂存目录
    - RPMS            // 打包后生成的 rpm 包会放在这里
    - SOURCES         // 源码压缩包
    - SPECS           // 放 xx.spec 文件
    - SRPMS           // 打包后生成的 srpm 包会放在这里
```

#### 四、创建 spec 文件

*：spec文件这个文件非常重要，控制整个rpm包安装卸载等全部过程

```shell
# 使用命令可以创建中 x.spec 文件模板
rpmdev-newspec eloam.spec
```

#### 五、spec 文件详解

```shell
Name:           名称
Version:        版本号
Release:        release版本
Summary:        对包的描述

License:        开源协议
URL:            项目主页
Source0:        源码包

# 没有用到，所以注释掉了
#BuildRequires:  
#Requires:       

%description
# 详细描述，多行，每行小于等于80个字符，否则算新的一段

%prep
# 静默模式解压，并进入解压后的目录，常用：%setup -q

%build
# 编译过程

%install
# 安装过程

%files
# 要打包的文件
```

##### eloam.spec 模板示例

```shell
Name:           eloamwss
Version:        1.0
Release:        1
Summary:        web rpm package

License:        GPL
URL:            http://sdk.eloam.net
Source0:        %{name}-%{version}.tar.gz
BuildRoot:      %{mktemp -ud %{_tmppath}/%{name}-%{version}-%{release}-XXXXX}

%description
THIS IS A WEB ELOAM PACKAGE

%prep
#%setup -n %{name}
%setup -q

%build

%install
mkdir -p %{buildroot}/usr/local/lib
mkdir -p %{buildroot}/etc/ld.so.conf.d/
mkdir -p %{buildroot}/etc/udev
mkdir -p %{buildroot}/opt/eloamwss
mkdir -p %{buildroot}/opt/eloamwss/lib
mkdir -p %{buildroot}/opt/eloamwss/doc
mkdir -p %{buildroot}/opt/eloamwss/Image
mkdir -p %{buildroot}/opt/eloamwss/icon

install eloamwss.conf %{buildroot}/etc/ld.so.conf.d/
install eloamwss/eloam.rules %{buildroot}/etc/udev
install -m 0755 eloamwss/*.sh %{buildroot}/opt/eloamwss
install -m 0755 eloamwss/eloamwss %{buildroot}/opt/eloamwss
install eloamwss/lib/* %{buildroot}/opt/eloamwss/lib
install eloamwss/doc/* %{buildroot}/opt/eloamwss/doc
install eloamwss/icon/* %{buildroot}/opt/eloamwss/icon



%define _unpackaged_files_terminate_build 0

%clean
rm -rf $RPM_BUILD_ROOT


%files
%defattr(-,root,root,-)
/opt/eloamwss/*
/etc/ld.so.conf.d/eloamwss.conf
/etc/udev/eloam.rules


%post
chmod 755 /opt/eloamwss/*.sh
chmod 755 /opt/eloamwss/eloamwss
ldconfig

%preun
/opt/eloamwss/shutdown.sh

%define __debug_install_post \
%{_rpmconfigdir}/find-debuginfo.sh %{?_find_debuginfo_opts} "%{_builddir}/%{?buildsubdir}"\
%{nil}
```

```shell
# 1.信息定义阶段
### 软件名字，要与spec的文件名一致 
Name:           apache-zookeeper
### 软件主版本号
Version:        3.6.1
### 发行编号，每打包一次值递增，主版本号发布新版后需重置该值
Release:        1
### 一行简短的软件简介，结尾不要加标点
Summary:        Zookeeper is a highly reliable distributed coordination service
Group:          
### 软件许可
License:        Apache 2.0
### 软件项目主页
URL:           https://zookeeper.apache.org/
### 放置在SOUIRCES目录的软件源码包名，可以指定多个:source1、source2等
#Source0:        %{name}.%{version}.tar.gz  
### 在 install 阶段的测试安装目录，方便写files
#buildroot:      %_topdir/BUILDROOT
#BuildRequires:  go
### 安装软件包时所需的依赖包列表，可以指定版本如 bash >= 1.1.1
#Requires:       readline-devel,pcre-devel,openssl-devel
### 程序的详细多行描述，每行必须小于等于 80 个字符，空行表示开始新段
%description
Zookeeper 3.6.1

# 2.准备阶段
%prep
## 静默模式解压并进入解压后的目录,也常用：%autosetup -n %{name}
 
# 3. 编译阶段
%build

# 4.安装阶段
%install
### 删除之前的残留文件

## rpm安装前制行的脚本
%pre
### $1==1 代表的是第一次安装，2代表是升级，0代表是卸载
#if [ $1 == 1 ];then     
#     /usr/sbin/useradd -r %{nginx_user} 2> /dev/null 
#fi

## rpm安装后制行的脚本
%post

###卸载前执行的脚本
%preun
###卸载后执行的脚本
%postun
rm -rf opt/zookeeper

# 5.清理阶段
%clean
### 删除buildroot目录
rm -rf %{buildroot}

#  6.文件设置阶段
%files
### 设定默认权限，如果下面没有指定权限，则继承默认
%defattr (-,root,root)
###要打包的文件和目录，在执行完rpmbuild -bi后，参考%{rootbuild}下生成的文件和目录
/opt/zookeeper


#编写完 SEPC 文件后，可以通过 rpmlint 检查是否有配置错误
rpmlint motan-go.spec

#执行打 rpm 包命令
cd /root/rpmbuild/SPECS
rpmbuild -ba zookeeper.spec

#安装测试RPM包
rpm -Uvh apache-zookeeper-3.6.1-1.x86_64.rpm
```

#### 六、打包

```shell
rpmbuild -bb  SPECS/elaom.spec --define="_topdir `pwd`"
```

#### 七、RPM 的安装 & 卸载

```shell
# rpm 安装     sudo rpm -ivh 包名 --nodeps
# rpm 卸载     sudo rpm -e --nodeps 包名
```

#### 八、 ** 特别重要 · 核心 **

2. 所有文件尽量不要使用记事本打开，可以使用 vi 来操作

3.我们的打包方法是将需要打包的文件，先打包成 tar.gz ，然后在使用 rpm 的打包工具对压缩包解压，解压后再次打成 rpm 包

4.压缩包的名称一定要和 spec 文件中名称（Source0）一致

5.rpm的打包过程中会遇到各种各样的问题，要看报错，慢慢分析，一步一个坑，习惯了就好，可以有效抑制头发生长

打压缩包 & 压缩包结构

```shell
# 打压缩包命令
tar -czvf eloamwss-1.0.tar.gz   eloamwss-1.0

# 压缩包结构
eloamwss-1.0.tar.gz
    - eloamwss-1.0
        - eloamwss.conf
        - eloamwss
            - doc
            - lib
            - eloamwss
```

#### 九、RPM 查询命令

查询已安装的RPM软件信息
格式：rpm -q[子选项] [软件名]
用法：结合不同子选项完成不同查询

-qa：查看系统中已安装的所有RPM软件包列表
-qi：查看指定软件的详细信息(information)
-ql：列出该软件所有的文件与目录所在的完整文件名(list)
-qc：列出该软件所有的配置文件(找出在/etc下的文件)
-qd：列出该软件所有的说明文件(找出与man相关的文件)
-qR：列出与该软件有关的依赖软件所含的文件(Required)
查询文件/目录属于哪个RPM软件

rpm -qf 文件或目录名
查询未安装的RPM包文件
格式：rpm -qp[子选项] [RPM包文件]
用法：-qp后接的所有参数与上面说明的一致，但用途仅在找出某个RPM文件内的信息，而非已安装的软件信息

-qpi：通过.rpm包文件查看该软件的详细信息
-qpl：查看.rpm安装包内所包含的目录、文件列表
-qpc：查看.rpm安装包内包含的配置文件列表
-qpd：查看.rpm安装包内包含的文档文件列表

```shell
[root@kafka01 ~]# rpm -qa
[root@kafka01 ~]# rpm -qi nginx
[root@kafka01 etc]# rpm -ql nginx
[root@kafka01 etc]# rpm -qf dhcp
dhcp-client-4.3.6-44.0.1.el8.x86_64
 
[root@zabbix_server ~]# rpm -qpl /mnt/Packages/zsh-5.0.2-34.el7_8.2.x86_64.rpm |more
 
# 查看openssh的安装包
[root@kafka01 ~]# rpm -qa|grep openssh
openssh-8.0p1-5.el8.x86_64
openssh-server-8.0p1-5.el8.x86_64
openssh-clients-8.0p1-5.el8.x86_64
 
[root@kafka01 ~]# rpm -ql openssh-server
/etc/pam.d/sshd
```

#### 十、RPM 软件包的安装、卸载

选项与参数：
-i：install 安装
-v：查看更详细的安装信息
-h：显示安装进度
-e：erase 卸载清除安装包

依赖关系
安装有依赖关系的多个软件时，被依赖的软件包需要先安装 --> 需要同时指定多个.rpm包文件进行安装
卸载有依赖关系的多个软件时，依赖其他程序的软件包需要先卸载 --> 同时指定多个软件吗进行卸载

yum可以自动解决依赖关系，但rpm安装需要我们自己去解决依赖关系
结合 "--nodeps"可以忽略依赖关系，但可能导致软件异常

![img](https://img-blog.csdnimg.cn/df759a6928e746179844cfd62ff09b7a.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBAeGlhb3hpZV9jb2Rpbmc=,size_20,color_FFFFFF,t_70,g_se,x_16)

```shell
# 卸载
rpm -e zsh
```

辅助选项
--force：强制安装所指定的rpm软件包

--nodeps：安装、升级或卸载软件时，忽略依赖关系

--replacefiles：在安装过程中出现某个文件已经被安装在你的系统上的信息，或出现版本不合的信息，可以用这个参数直接覆盖文件

--replacekgs：重新安装某个已经安装过的软件。防止安装RPM文件时因为某软件已安装导致无法继续安装

--test：测试软件是否可以被安装到用户的Linux环境中

建议：尽量安装时直接使用 -ivh 就好，尽量不要使用--force暴力安装，否则可能会发生很多不可预期的问题。
#### 十一、RPM 升级与更新
格式：rpm [选项] [RPM包文件]…

选项与参数

-i：安装一个新的rpm软件包
-U：升级某个rpm软件，若原本未装，则进行安装
需要自己提供高版本的软件包，不会自动去帮助到哪里下载
-F：更新某个rpm软件，若原本未装，则放弃安装

```shell
rpm -ivh file.rpm 　＃[安装新的rpm]--install--verbose--hash
rpm -Uvh file.rpm    ＃[升级一个rpm]--upgrade
rpm -e file.rpm      ＃[删除一个rpm包]--erase
```

#### 十二、导入签名：

[root@localhost RPMS]# rpm --import 签名文件

举例：

```shell
[root@localhost fc40]# rpm --import RPM-GPG-KEY
[root@localhost fc40]# rpm --import RPM-GPG-KEY-fedora
```

### man

```shell
安装更新 man
sudo yum install -y man-pages	--> 安装
sudo mandb	--> 更新

输入 man + 数字 + 命令/函数，可以查到相关的命令和函数，若不加数字， man 默认从数字较小的手册中寻找相关命令和函数

man 3 rand 	--> 表示在手册的第三部分查找 rand 函数
man ls 			--> 查找 ls 用法手册
```

# 挂载

### 查看设备挂载情况

lsblk

```shell
lsblk -f 查看详细的设备挂载情况，显示文件系统信息
```

# Linux 进阶

### grep

在文件中查找关键字，并显示关键字所在行。

```shell
grep text file # text代表要搜索的文本，file代表供搜索的文件

# 实例
[root@lion ~]# grep path /etc/profile
```

#### 常用参数

```shell
-i 忽略大小写， grep -i path /etc/profile 
-n 显示行号，grep -n path /etc/profile
-v 只显示搜索文本不在的那些行，grep -v path /etc/profile
-r 递归查找， grep -r hello /etc ，Linux 中还有一个 rgrep 命令，作用相当于 grep -r
```

#### 高级用法

```shell
grep 可以配合正则表达式使用
grep -E path /etc/profile --> 完全匹配path
grep -E ^path /etc/profile --> 匹配path开头的字符串
grep -E [Pp]ath /etc/profile --> 匹配path或Path
```

### sort

```shell
对文件的行进行排序
sort name.txt # 对name.txt对文本内容进行排序

-o 将排序后的文件写入新文件， sort -o name_sorted.txt name.txt
-r 倒序排序， sort -r name.txt
-R 随机排序， sort -R name.txt
-n 对数字进行排序，默认是把数字识别成字符串的，因此 138 会排在 25 前面，如果添加了 -n 数字排序的话，则 25 会在 138 前面
```

### wc

```shell
用于文件的统计。它可以统计单词数目、行数、字符数，字节数等
[root@lion ~]wc name.txt # 统计name.txt
13 13 91 name.txt
    第一个13，表示行数
    第二个13，表示单词数
    第三个91，表示字节数

-l 只统计行数， wc -l name.txt
-w 只统计单词数， wc -w name.txt
-c 只统计字节数， wc -c name.txt
-m 只统计字符数， wc -m name.txt
```

### uniq

```shell
删除文件中的重复内容
uniq name.txt # 去除name.txt重复的行数，并打印到屏幕上
uniq name.txt uniq_name.txt # 把去除重复后的文件保存为 uniq_name.txt

【注意】它只能去除连续重复的行数。
-c 统计重复行数， uniq -c name.txt ；
-d 只显示重复的行数， uniq -d name.txt 。
```

### cut

```shell
剪切文件的一部分内容
cut -c 2-4 name.txt # 剪切每一行第二到第四个字符

-d 用于指定用什么分隔符（比如逗号、分号、双引号等等） cut -d , name.txt
-f 表示剪切下用分隔符分割的哪一块或哪几块区域， cut -d , -f 1 name.txt
```

### 重定向

#### 输出重定向 `>`

```shell
`>` 表示重定向到新的文件， `cut -d , -f 1 notes.csv > name.csv` ，它表示通过逗号剪切 `notes.csv` 文件（剪切完有3个部分）获取第一个部分，重定向到 `name.csv` 文件。

【注意】使用 `>` 要注意，如果输出的文件不存在它会新建一个，如果输出的文件已经存在，则会覆盖。因此执行这个操作要非常小心，以免覆盖其它重要文件。
```

#### 输出重定向 `>>` 

```shell
表示重定向到文件末尾，因此它不会像 > 命令这么危险，它是追加到文件的末尾（当然如果文件不存在，也会被创建）。

再次执行 cut -d , -f 1 notes.csv >> name.csv ，则会把名字追加到 name.csv 里面。
我们平时读的 log 日志文件其实都是用这个命令输出的。
```

#### 输出重定向 `2>`

```shell
标准错误输出
cat not_exist_file.csv > res.txt 2> errors.log
当我们 cat 一个文件时，会把文件内容打印到屏幕上，这个是标准输出；
当使用了 > res.txt 时，则不会打印到屏幕，会把标准输出写入文件 res.txt 文件中；
2> errors.log 当发生错误时会写入 errors.log 文件中。
```

#### 输出重定向 `2>>`

```shell
标准错误输出（追加到文件末尾）同 >> 相似
```

#### 输出重定向 `2>&1`

```shell
标准输出和标准错误输出都重定向都一个地方
cat not_exist_file.csv > res.txt 2>&1  # 覆盖输出
cat not_exist_file.csv >> res.txt 2>&1 # 追加输出
```

#### 输入重定向 `<`

```shell
< 符号用于指定命令的输入
cat < name.csv # 指定命令的输入为 name.csv

虽然它的运行结果与 cat name.csv 一样，但是它们的原理却完全不同。

cat name.csv 表示 cat 命令接收的输入是 notes.csv 文件名，那么要先打开这个文件，然后打印出文件内容。
cat < name.csv 表示 cat 命令接收的输入直接是 notes.csv 这个文件的内容， cat 命令只负责将其内容打印，打开文件并将文件内容传递给 cat 命令的工作则交给终端完成。
```

#### 输入重定向 `<<` 

```shell
将键盘的输入重定向为某个命令的输入。

sort -n << END # 输入这个命令之后，按下回车，终端就进入键盘输入模式，其中END为结束命令（这个可以自定义）

wc -m << END # 统计输入的单词
```

### 管道 `|` 

```shell
把两个命令连起来使用，一个命令的输出作为另外一个命令的输入，英文是 pipeline ，可以想象一个个水管连接起来，管道算是重定向流的一种。
```

![未命名文件 (1).png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d46b974834864da2a734b42f5703d65c~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp)

```shell
cut -d , -f 1 name.csv | sort > sorted_name.txt 
# 第一步获取到的 name 列表，通过管道符再进行排序，最后输出到sorted_name.txt

du | sort -nr | head 
# du 表示列举目录大小信息
# sort 进行排序,-n 表示按数字排序，-r 表示倒序
# head 前10行文件

grep log -Ir /var/log | cut -d : -f 1 | sort | uniq
# grep log -Ir /var/log 表示在log文件夹下搜索 /var/log 文本，-r 表示递归，-I 用于排除二进制文件
# cut -d : -f 1 表示通过冒号进行剪切，获取剪切的第一部分
# sort 进行排序
# uniq 进行去重
```

### 流

```shell
流并非一个命令，在计算机科学中，流 stream 的含义是比较难理解的，记住一点即可：流就是读一点数据, 处理一点点数据。其中数据一般就是二进制格式。 上面提及的重定向或管道，就是把数据当做流去运转的。
```

### w

帮助我们快速了解系统中目前有哪些用户登录着，以及他们在干什么

```shell
[root@lion ~]# w
 06:31:53 up 25 days,  9:53,  1 user,  load average: 0.00, 0.01, 0.05
USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
root     pts/0    118.31.243.53    05:56    1.00s  0.02s  0.00s w
 
06:31:53：表示当前时间
up 25 days, 9:53：表示系统已经正常运行了“25天9小时53分钟”
1 user：表示一个用户
load average: 0.00, 0.01, 0.05：表示系统的负载，3个值分别表示“1分钟的平均负载”，“5分钟的平均负载”，“15分钟的平均负载”

 USER：表示登录的用于
 TTY：登录的终端名称为pts/0
 FROM：连接到服务器的ip地址
 LOGIN@：登录时间
 IDLE：用户有多久没有活跃了
 JCPU：该终端所有相关的进程使用的 CPU 时间，每当进程结束就停止计时，开始新的进程则会重新计时
 PCPU：表示 CPU 执行当前程序所消耗的时间，当前进程就是在 WHAT 列里显示的程序
 WHAT：表示当下用户正运行的程序是什么，这里我运行的是 w
```

# 文件压缩解压

- 打包：是将多个文件变成一个总的文件，它的学名叫存档、归档。
- 压缩：是将一个大文件（通常指归档）压缩变成一个小文件。
- 我们常常使用 `tar` 将多个文件归档为一个总的文件，称为 `archive` 。 然后用 `gzip` 或 `bzip2` 命令将 `archive` 压缩为更小的文件。

![未命名文件.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0d87434a4c414defb180b05d9bfca4c4~tplv-k3u1fbpfcp-zoom-in-crop-mark:4536:0:0:0.awebp)

### tar

```shell
创建一个 tar 归档
-cvf 表示 create（创建）+ verbose（细节）+ file（文件），创建归档文件并显示操作细节；
-tf 显示归档里的内容，并不解开归档；
-rvf 追加文件到归档， tar -rvf archive.tar file.txt ；
-xvf 解开归档， tar -xvf archive.tar 。

tar -cvf sort.tar sort/ # 将sort文件夹归档为sort.tar
tar -cvf archive.tar file1 file2 file3 # 将 file1 file2 file3 归档为archive.tar
```

### gzip/gunzip

```shell
“压缩/解压”归档，默认用 gzip 命令，压缩后的文件后缀名为 .tar.gz 
gzip archive.tar # 压缩
gunzip archive.tar.gz # 解压
```

### tar 归档+压缩

```shell
可以用 tar 命令同时完成归档和压缩的操作，就是给 tar 命令多加一个选项参数，使之完成归档操作后，还是调用 gzip 或 bzip2 命令来完成压缩操作。

tar -zcvf archive.tar.gz archive/ # 将archive文件夹归档并压缩
tar -zxvf archive.tar.gz # 将archive.tar.gz归档压缩文件解压
```

### zcat、zless、zmore

```shell
之前讲过使用 cat less more 可以查看文件内容，但是压缩文件的内容是不能使用这些命令进行查看的，而要使用 zcat、zless、zmore 进行查看。

zcat archive.tar.gz
```

### zip/unzip

```shell
“压缩/解压” zip 文件（ zip 压缩文件一般来自 windows 操作系统）

# Red Hat 一族中的安装方式
yum install zip 
yum install unzip 

unzip archive.zip # 解压 .zip 文件
unzip -l archive.zip # 不解开 .zip 文件，只看其中内容

zip -r sort.zip sort/ # 将sort文件夹压缩为 sort.zip，其中-r表示递归
```

### ifconfig

```shell
查看 ip 网络相关信息，如果命令不存在的话， 执行命令 yum install -y net-tools 安装。

eth0 对应有线连接（对应你的有线网卡），就是用网线来连接的上网。 eth 是 Ethernet 的缩写，表示“以太网”。有些电脑可能同时有好几条网线连着，例如服务器，那么除了 eht0 ，你还会看到 eth1 、 eth2 等。
lo 表示本地回环（ Local Loopback 的缩写，对应一个虚拟网卡）可以看到它的 ip 地址是 127.0.0.1 。每台电脑都应该有这个接口，因为它对应着“连向自己的链接”。这也是被称之为“本地回环”的原因。所有经由这个接口发送的东西都会回到你自己的电脑。看起来好像并没有什么用，但有时为了某些缘故，我们需要连接自己。例如用来测试一个网络程序，但又不想让局域网或外网的用户查看，只能在此台主机上运行和查看所有的网络接口。例如在我们启动一个前端工程时，在浏览器输入 127.0.0.1:3000 启动项目就能查看到自己的 web 网站，并且它只有你能看到。
wlan0 表示无线局域网（上面案例并未展示）。
```

### 备份

#### scp

它是 `Secure Copy` 的缩写，表示安全拷贝。 `scp` 可以使我们通过网络，把文件从一台电脑拷贝到另一台电脑。
`scp` 是基于 `ssh` 的原理来运作的， `ssh` 会在两台通过网络连接的电脑之间创建一条安全通信的管道， `scp` 就利用这条管道安全地拷贝文件。

```shell
scp source_file destination_file # source_file 表示源文件，destination_file 表示目标文件
```

其中 `source_file` 和 `destination_file` 都可以这样表示： `user@ip:file_name` ， `user` 是登录名， `ip` 是域名或 `ip` 地址。 `file_name` 是文件路径。

```shell
scp file.txt root@192.168.1.5:/root # 表示把我的电脑中当前文件夹下的 file.txt 文件拷贝到远程电脑
scp root@192.168.1.5:/root/file.txt file.txt # 表示把远程电脑上的 file.txt 文件拷贝到本机
```

#### rsync

`rsync` 命令主要用于远程同步文件。它可以同步两个目录，不管它们是否处于同一台电脑。它应该是最常用于“增量备份”的命令了。它就是智能版的 `scp` 命令。

```shell
yum install -y rsync

rsync -arv Images/ backups/ # 将Images 目录下的所有文件备份到 backups 目录下
rsync -arv Images/ root@192.x.x.x:backups/ # 同步到服务器的backups目录下

-a 保留文件的所有信息，包括权限，修改日期等；
-r 递归调用，表示子目录的所有文件也都包括；
-v 冗余模式，输出详细操作信息。

默认地， rsync 在同步时并不会删除目标目录的文件，例如你在源目录中删除一个文件，但是用 rsync 同步时，它并不会删除同步目录中的相同文件。如果向删除也可以这么做： rsync -arv --delete Images/ backups/
```

# **centOS7下的静态Ip的配置**

设置vm虚拟机的网路配置，将网卡设置为NAT 模式

设置虚拟机的网卡类型为NAT 模式
查看NAT模式下的虚拟网卡信息。

```shell
子网ip：192.168.192.0 也就是centOS的ip可以设置为192.168.192.0~192.168.192.255之间。
子网掩码： 255.255.255.0
网关:   192.168.192.2
```

修改配置文件
ifconfig查看网卡信息：

查看默认网卡信息.

一般网卡信息在/etc/sysconfig/network-scripts/ 的目录下的ifcfg-ens33文件中。

使用命令 vi /etc/sysconfig/network-scripts/ifcfg-ens33

查看默认的网卡信息如下：

```shell
TYPE="Ethernet"   # 网卡类型： 这里默认是以太网
PROXY_METHOD="none"  # 代理方式
BROWSER_ONLY="no"    
BOOTPROTO="no"   # 静态ip
DEFROUTE="yes"       # 默认路由
IPV4_FAILURE_FATAL="no"   # 是否开启IPV4致命错误检测：否
IPV6INIT="yes"        # ipv6是否初始化：是
IPV6_AUTOCONF="yes" 
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"   #网卡物理设备名称
UUID="b4af2f5f-f889-40d6-9058-2eff9f29539b"  # 网卡信息通用唯一识别码
DEVICE="ens33"   # 网卡设备名称，必须哈`NAME` 相同
ONBOOT=no      # 是否开机启动，默认：no
```

```shell
先修改网络
ONBOOT=yes
#重新启动
```

设置网卡引导协议为静态

```shell 
BOOTPROTO=static
```

```shell
TYPE="Ethernet"
PROXY_METHOD="none"
BROWSER_ONLY="no"
BOOTPROTO="static"
DEFROUTE="yes"
IPV4_FAILURE_FATAL="no"
IPV6INIT="yes"
IPV6_AUTOCONF="yes"
IPV6_DEFROUTE="yes"
IPV6_FAILURE_FATAL="no"
IPV6_ADDR_GEN_MODE="stable-privacy"
NAME="ens33"
UUID="b4af2f5f-f889-40d6-9058-2eff9f29539b"
DEVICE="ens33"

ONBOOT="yes"
IPADDR=192.168.222.131
NETMASK=255.255.255.0
GATEWAY=192.168.192.2
DNS1=119.29.29.29
```

```shell
systemctl restart network​​ 重启网络
```

