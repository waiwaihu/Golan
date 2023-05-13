# Nginx

### Linux安装

```shell
Nginx开源版
http://nginx.org/

Nginx plus开源版
https://www.nginx.com/

Openresty
http://openresty.org

Tengine
http://tengine.taobao.org/

下载地址：http://nginx.org/en/download.html
# windows建议在D:\nginx-1.23.3\目录下输入CMD，进入命令行模式，在输入nginx.exe启动
```

```shell
# Linux
#先看看本机有没有安装nginx
whereis nginx

#1、安装依赖环境
yum -y install gcc-c++ pcre pcre-devel zlib zlib-devel openssl openssl-devel
pcre用于解析正则表达式
zlib压缩和解压缩依赖
ssl安全的加密的套接字协议层

#2、解压
##cd /usr/local/ #上传文件
tar -zxvf nginx-1.23.3.tar.gz #解压

#3、编译
编译之前，先创建nginx的临时目录，如果不建，在启动nginx的过程会报错
mkdir -p /var/temp/nginx

cd nginx-1.23.3 #进入文件

#4、配置，在nginx目录下，输入如下命令进行配置，目的是为了创建makefile文件
./configure \
--prefix=/usr/local/nginx \
--pid-path=/var/run/nginx/nginx.pid \
--lock-path=/var/lock/nginx.lock \
--error-log-path=/var/log/nginx/error.log \
--http-log-path=/var/log/nginx/access.log \
--with-http_gzip_static_module \
--http-client-body-temp-path=/var/temp/nginx/client \
--http-proxy-temp-path=/var/temp/nginx/proxy \
--http-fastcgi-temp-path=/var/temp/nginx/fastcgi \
--http-uwsgi-temp-path=/var/temp/nginx/uwsgi \
--http-scgi-temp-path=/var/temp/nginx/scgi      
`\`代表换行
`--prefix      指定nginx安装目录`
`--pid-path    指定nginx的pid`
`--lock-path   锁定安装文件，防止被恶意篡改或误操作`
`--error-log   错误日志`
`--http-log-path   http日志`
`--with-http_gzip_static_module   启用gzip模块，在线实时压缩输出数据流`
`--http-client-body-temp-path    设定客户端请求的临时目录`
`--http-proxy-temp-path     设定http代理临时目录`
`--http-fastcgi-temp-path   设定fastcgi临时目录`
`--http-uwsgi-temp-path     设定uwsgi临时目录`
`--http-scgi-temp-path      设定scgi临时目录`

[root@localhost nginx-1.23.3]#ls 查看是否生成了makefile文件

#5、make编译安装
make && make install

#6、进入sbin目录启动nginx
cd /usr/local/nginx
cd sbin
./nginx
`./nginx -s stop 停止`
`./nginx -s reload 重新加载`
`./nginx -s quit 安全退出`
`ps aux|grep nginx 查看nginx进程 `

#7、访问虚拟机的ip地址就可打开nginx默认页面
	云服务器安装，需要开启默认的nginx端口:80
	在虚拟机安装，需要关闭防火墙systemctl stop firewalld.service
	本地win或mac需要关闭防火墙
firewall-cmd --zone=public --add-port=80/tcp --permanent 开放端口
firewall-cmd --reload 重启防火墙
curl http://localhost:80

#8、设置开机启动
vim /usr/lib/systemb/nginx.service 创建服务脚本

[Unit]
Description=nginx -web service
After=network.target remote-fs.target nss-lookup.target
 
[Service]
Type=forking
PIDFile=/usr/local/nginx/logs/nginx.pid
ExecStartPre=/usr/local/nginx/sbin/nginx -t -c /usr/local/nginx/conf/nginx.conf
ExecStart=/usr/local/nginx/sbin/nginx -c /usr/local/nginx/conf/nginx.conf
ExecReload=/usr/local/nginx/sbin/nginx -s reload
ExecStop=/usr/local/nginx/sbin/nginx -s stop
ExecQuit=/usr/local/nginx/sbin/nginx -s quit
PrivateTmp=true

[Install]
WantedBy=multi-user.target
#Description:描述服务
#After:描述服务类别
#[Service]服务运行参数的设置
#Type=forking是后台运行的形式
#ExecStart为服务的具体运行命令
#ExecReload为重启命令
#ExecStop为停止命令
#PrivateTmp=True表示给服务分配独立的临时空间
#注意：[Service]的启动、重启、停止命令全部要求使用绝对路径
#[Install]运行级别下服务安装的相关设置，可设置为多用户，即系统运行级别为3
#保存退出

#9、重新加载系统服务
systemctl daemon-reload

#启动nginx服务
systemctl start nginx.service　 
 
#停止服务
systemctl stop nginx.service　      
    
 #重新启动服务
systemctl restart nginx.service　     
 
#查看所有已启动的服务
systemctl list-units --type=service    
 
#查看服务当前状态
systemctl status nginx.service         
 
 #设置开机自启动
systemctl enable nginx.service     
    
 # 停止开机自启动
systemctl disable nginx.service     
```

### 常见的命令

./nginx -h 查看nginx所有的命令参数

| options       | 说明                                                         |
| :------------ | :----------------------------------------------------------- |
| -?,-h         | this help                                                    |
| -v            | 显示nginx的版本号                                            |
| -V            | 显示nginx的版本号和编译信息                                  |
| -t            | 检查nginx配置文件的正确性                                    |
| -T            | 检查nginx配置文件的正确定及配置文件的详细配置内容            |
| -q            | suppress non-error messages during configuration testing     |
| -s signal     | 向主进程发送信号，如:./nginx -s reload 配置文件变化后重新加载配置文件并重启nginx服务 |
| -p prefix     | 设置nginx的安装路径                                          |
| -c filename   | 设置nginx配置文件的路径                                      |
| -g directives | 设置配置文件之外的全局指令                                   |

### nginx.conf https 配置

```nginx
    server {
        listen       443 ssl;
        server_name  localhost;
        ssl_certificate      cert.pem;#根证书地址（默认把证书放在conf目录）
        ssl_certificate_key  cert.key;#证书秘钥（默认把证书放在conf目录）
        ssl_session_cache    shared:SSL:1m;
        ssl_session_timeout  5m;
        ssl_ciphers  HIGH:!aNULL:!MD5;
        ssl_prefer_server_ciphers  on;
        location / {
            root   html;
            index  index.html index.htm;
        }
    }
```

### 将 http 重定向 https

```nginx
server {
    listen 80;
    server_name localhost;
    #将请求转成https
    rewrite ^(.*) https://$server_name$1 permanent;
}
```

### nginx相关挂载

```nginx
configure命令用来配置nginx编译环境。 该命令定义了系统各方面功能，包括允许nginx使用的连接处理方式。 其执行结果是生成一个Makefile文件。 configure命令支持如下参数：

--prefix=path — 定义服务器文件的完整路径，该路径同时也是configure命令设置的 相对路径（除类库源文件外）以及nginx.conf文件定义的相对路径的基准路径。其默认 值是/usr/local/nginx。

--sbin-path=path — 设置nginx可执行文件的完整路径，该路径仅在安装期间使用， 默认路径为prefix/sbin/nginx。

--conf-path=path — 设置配置文件nginx.conf的完整路径。如有必要，总是可以 在nginx启动时通过命令行参数-cfile指定一个不同的配置文件路径。 默认路径为prefix/conf/nginx.conf。

--pid-path=path — 设置nginx.pid文件的完整路径，该文件存储了主进程的进程ID。安装完成后， 该路径总是可以在nginx.conf文件中用 pid指令来修改。 默认路径为prefix/logs/nginx.pid。

--error-log-path=path — 设置记录主要错误、警告以及调试信息日志的完整路径。安装完成后， 该路径总是可以在nginx.conf文件中用 error_log 指令来修改。 默认路径为prefix/logs/error.log。

--http-log-path=path — 设置记录HTTP服务器主要请求日志的完整路径。安装完成后， 该路径总是可以在nginx.conf文件用 access_log 指令来修改。 默认路径为prefix/logs/access.log

--user=name — 设置工作进程以某非特权用户的身份运行。安装完成后，该用户名总是可以在 nginx.conf文件用user指令来修改。 其默认值为nobody。

--group=name — 设置工作进程以某非特权用户的身份运行。安装完成后，该用户名总是可以在 nginx.conf用user指令来修改。 其默认名称与未授权用户名称相同。

--with-select_module
--without-select_module — 设置是否将select()方法模块编译进nginx中。如果系统平台不支持kqueue、epoll、rtsig或/dev/poll等更合适的方法， 该模块会被自动编译。

--with-poll_module
--without-poll_module — 设置是否将poll()方法模块编译进nginx中。如果系统平台不支持kqueue、epoll、rtsig或/dev/poll等更合适的方法， 该模块会被自动编译。

--without-http_gzip_module — 不编译http_gzip_module模块。该模块可以压缩HTTP服务器的响应，该模块需要zlib库才能编译和运行。

--without-http_rewrite_module — 不编译http_rewrite_module模块。该模块允许HTTP服务器重定向请求，改变请求的URI地址 。创建并运行该模块需要PCRE库支持。

--without-http_proxy_module — 不编译HTTP服务器的代理模块。

--with-http_ssl_module — 为HTTP服务器编译HTTPS协议支持的模块。该模块默认是不编译的。它需要OpenSSL库才能编译和运行。

--with-pcre=path — 设置PCRE库源文件的路径地址。PCRE库的发行版(version 4.4 — 8.30)需要先从PCRE站点下载并解压缩。 剩下的安装工作由nginx的./configure和make命令来完成。该库应用于location 指令的正则表达式支持以及ngx_http_rewrite_module模块。

--with-pcre-jit — 编译PCRE库时增加“实时编译（pcre_jit）”支持。

--with-zlib=path — 设置zlib库源文件的路径地址。zlib库的发行版(version 1.1.3 — 1.2.5)需要先从zlib站点下载并解压缩。 剩下的安装工作由nginx的./configure和make命令来完成。该库应用于 ngx_http_gzip_module模块。

--with-cc-opt=parameters — 设置将会添加额外参数到CFLAGS变量中。当在FreeBSD使用系统PCRE库时，需要指定 --with-cc-opt="-I /usr/local/include"。 如果需要增加select()方法所支持的文件数量，也可以参照如下方式指定： --with-cc-opt="-D FD_SETSIZE=2048"。

--with-ld-opt=parameters — 设置将会在链接（linking）过程中使用的额外参数。当在FreeBSD使用系统PCRE库时，需要指定 --with-ld-opt="-L /usr/local/lib"。

以下是参数使用示例（注意不应有换行）：

./configure
    --sbin-path=/usr/local/nginx/nginx
    --conf-path=/usr/local/nginx/nginx.conf
    --pid-path=/usr/local/nginx/nginx.pid
    --with-http_ssl_module
    --with-pcre=../pcre-4.4
    --with-zlib=../zlib-1.1.3
```

### nginx.conf

![Nginx配置文件（nginx.conf）详解，轻松掌握nginx~_nginx_02](https://s2.51cto.com/images/202205/a927e74473974b8f3bc099eae5396117f917d2.png?x-oss-process=image/watermark,size_14,text_QDUxQ1RP5Y2a5a6i,color_FFFFFF,t_30,g_se,x_10,y_10,shadow_20,type_ZmFuZ3poZW5naGVpdGk=,x-oss-process=image/resize,m_fixed,w_1184/format,webp)

#### 全局块

全局块是默认配置文件从开始到events块之间的内容。主要设置nginx整体运行的配置指令，这些指令的作用域是全局

```nginx
#定义Nginx运行的用户和用户组
user www www;
 
#nginx进程数，建议设置为等于CPU总核心数。
worker_processes 8;
 
#全局错误日志定义类型，[ debug | info | notice | warn | error | crit ]
error_log /usr/local/nginx/logs/error.log info;
 
#进程pid文件
pid /usr/local/nginx/logs/nginx.pid;
 
#指定进程可以打开的最大描述符：数目
#工作模式与连接数上限
#这个指令是指当一个nginx进程打开的最多文件描述符数目，理论值应该是最多打开文件数（ulimit -n）与nginx进程数相除，但是nginx分配请求并不是那么均匀，所以最好与ulimit -n 的值保持一致。
#现在在linux 2.6内核下开启文件打开数为65535，worker_rlimit_nofile就相应应该填写65535。
#这是因为nginx调度时分配请求到进程并不是那么的均衡，所以假如填写10240，总并发量达到3-4万时就有进程可能超过10240了，这时会返回502错误。
worker_rlimit_nofile 65535;
```

#### events块

ewents块的指令主要影响nginx服务器和用户的网络连接，对性能影响较大。

```nginx
events
{
    #参考事件模型，use [ kqueue | rtsig | epoll | /dev/poll | select | poll ]; epoll模型
    #是Linux 2.6以上版本内核中的高性能网络I/O模型，linux建议epoll，如果跑在FreeBSD上面，就用kqueue模型。
    #补充说明：
    #与apache相类，nginx针对不同的操作系统，有不同的事件模型
    #A）标准事件模型
    #Select、poll属于标准事件模型，如果当前系统不存在更有效的方法，nginx会选择select或poll
    #B）高效事件模型
    #Kqueue：使用于FreeBSD 4.1+, OpenBSD 2.9+, NetBSD 2.0 和 MacOS X.使用双处理器的MacOS X系统使用kqueue可能会造成内核崩溃。
    #Epoll：使用于Linux内核2.6版本及以后的系统。
    #/dev/poll：使用于Solaris 7 11/99+，HP/UX 11.22+ (eventport)，IRIX 6.5.15+ 和 Tru64 UNIX 5.1A+。
    #Eventport：使用于Solaris 10。 为了防止出现内核崩溃的问题， 有必要安装安全补丁。
    use epoll;
 
    #单个进程最大连接数（最大连接数=连接数*进程数）
    #根据硬件调整，和前面工作进程配合起来用，尽量大，但是别把cpu跑到100%就行。每个进程允许的最多连接数，理论上每台nginx服务器的最大连接数为。
    worker_connections 65535;
 
    #keepalive超时时间。
    keepalive_timeout 60;
 
    #客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求头的大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。
    #分页大小可以用命令getconf PAGESIZE 取得。
    #[root@web001 ~]# getconf PAGESIZE
    #4096
    #但也有client_header_buffer_size超过4k的情况，但是client_header_buffer_size该值必须设置为“系统分页大小”的整倍数。
    client_header_buffer_size 4k;
 
    #这个将为打开文件指定缓存，默认是没有启用的，max指定缓存数量，建议和打开文件数一致，inactive是指经过多长时间文件没被请求后删除缓存。
    open_file_cache max=65535 inactive=60s;
 
    #这个是指多长时间检查一次缓存的有效信息。
    #语法:open_file_cache_valid time 默认值:open_file_cache_valid 60 使用字段:http, server, location 这个指令指定了何时需要检查open_file_cache中缓存项目的有效信息.
    open_file_cache_valid 80s;
 
    #open_file_cache指令中的inactive参数时间内文件的最少使用次数，如果超过这个数字，文件描述符一直是在缓存中打开的，如上例，如果有一个文件在inactive时间内一次没被使用，它将被移除。
    #语法:open_file_cache_min_uses number 默认值:open_file_cache_min_uses 1 使用字段:http, server, location  这个指令指定了在open_file_cache指令无效的参数中一定的时间范围内可以使用的最小文件数,如果使用更大的值,文件描述符在cache中总是打开状态.
    open_file_cache_min_uses 1;
    
    #语法:open_file_cache_errors on | off 默认值:open_file_cache_errors off 使用字段:http, server, location 这个指令指定是否在搜索一个文件时记录cache错误.
    open_file_cache_errors on;
}
```

常用到的配置指令案例

```nginx
events {            # events块开始
    worker_connections  1024;  #每个工作进程的最大连接数量（根据硬件调整，和前面工作进程配合起来用，尽量大，但是别把cpu跑到100%就行。）
    use epoll;    # 使用epoll的I/O 模型。linux建议epoll，FreeBSD建议采用kqueue，window下不指定。
    accept_mutex on;   #开启网络连接的序列化(防止多个进程对连接的争抢）
    multi_accept  on;  #允许同时接收多个网络连接(默认关闭），工作进程都有能力同时接收多个新到达的网络连接
} 
```

#### http块

http块是nginx服务配置中的重要部分，代理、缓存、日志定义等很多的功能指令都可以放在http块中

```nginx
#文件扩展名与文件类型映射表
include mime.types;
 
#默认文件类型
default_type application/octet-stream;
 
#默认编码
#charset utf-8;
 
#服务器名字的hash表大小
#保存服务器名字的hash表是由指令server_names_hash_max_size 和server_names_hash_bucket_size所控制的。参数hash bucket size总是等于hash表的大小，并且是一路处理器缓存大小的倍数。在减少了在内存中的存取次数后，使在处理器中加速查找hash表键值成为可能。如果hash bucket size等于一路处理器缓存的大小，那么在查找键的时候，最坏的情况下在内存中查找的次数为2。第一次是确定存储单元的地址，第二次是在存储单元中查找键 值。因此，如果Nginx给出需要增大hash max size 或 hash bucket size的提示，那么首要的是增大前一个参数的大小.
server_names_hash_bucket_size 128;
 
#客户端请求头部的缓冲区大小。这个可以根据你的系统分页大小来设置，一般一个请求的头部大小不会超过1k，不过由于一般系统分页都要大于1k，所以这里设置为分页大小。分页大小可以用命令getconf PAGESIZE取得。
client_header_buffer_size 32k;
 
#客户请求头缓冲大小。nginx默认会用client_header_buffer_size这个buffer来读取header值，如果header过大，它会使用large_client_header_buffers来读取。
large_client_header_buffers 4 64k;

#设定通过nginx上传文件的大小
client_max_body_size 8m;
 
#开启高效文件传输模式，sendfile指令指定nginx是否调用sendfile函数来输出文件，对于普通应用设为 on，如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络I/O处理速度，降低系统的负载。注意：如果图片显示不正常把这个改成off。
#sendfile指令指定 nginx 是否调用sendfile 函数（zero copy 方式）来输出文件，对于普通应用，必须设为on。如果用来进行下载等应用磁盘IO重负载应用，可设置为off，以平衡磁盘与网络IO处理速度，降低系统uptime。
sendfile on;
 
#开启目录列表访问，合适下载服务器，默认关闭。
autoindex on;
```

```nginx
#实例一
http {     # http块开始
    include       mime.types;   #定义MIME-Type(网络资源的媒体类型），nginx作为web服务器必须能够识别前端请求的资源类型。引用外部文件mime.types。
    default_type  application/octet-stream;  #配置用于处理前端请求的MIME类型（默认为text/plain）
    access_log  logs/access.log  main;  #配置服务日志的存放路径、日志格式、临时存放日志的内存缓存区大小；
    log_format  main  xxx;  #专门定义服务日志的格式；
    sendfile        on;  #允许sendfile方式传输文件，
    sendfile_max_chunk  xxx;   #每个工作进程调用sendfile()传输的数据量最大值（0为不限制）
    tcp_nopush on; #该指令必须在sendfile打开的状态下才会生效，主要是用来提升网络包的传输'效率'
    tcp_nodelay on; #该指令必须在keep-alive连接开启的情况下才生效，来提高网络包传输的'实时性'
    keepalive_timeout  65;  # 连接超时时间，与用户建立连接会话后nginx服务器保持会话的时间；
    gzip  on;   # 开启Gzip功能，对响应数据进行在线实时压缩；
    ...
}   # http块结束
```

server块 必须包含在http之下，server可单独拆分为一个文件 在nginx下http块下引用即可

![img](https://img-blog.csdnimg.cn/c74dd1a8e20a483784d4fba25d50c637.png)

```nginx
server {
        listen       443 ssl;
        server_name  freephp.us freephp.us;
 
        ssl_certificate      freephp.us.pem;
        ssl_certificate_key  freephp.us-key.pem;
 
        ssl_session_cache    shared:SSL:1m;
        ssl_session_timeout  5m;
 
        ssl_ciphers  HIGH:!aNULL:!MD5;
        ssl_prefer_server_ciphers  on;
 
 
        root   "C:\phpStudy\PHPTutorial\WWW\freephp";
        location / {
            index  index.html index.htm index.php;
            #autoindex  on;
        }
        location ~ \.php(.*)$ {
            fastcgi_pass   127.0.0.1:9000;
            fastcgi_index  index.php;
            fastcgi_split_path_info  ^((?U).+\.php)(/?.+)$;
            fastcgi_param  SCRIPT_FILENAME  $document_root$fastcgi_script_name;
            fastcgi_param  PATH_INFO  $fastcgi_path_info;
            fastcgi_param  PATH_TRANSLATED  $document_root$fastcgi_path_info;
            include        fastcgi_params;
        }
}
```

```nginx
#实例一
server {    # server块开始
        keepalive_requests 100;   # 单连接请求上限，限制用户通过某一连接想服务端发送请求的次数
        listen       80;    # 设置网络监听（分3种监听方式：IP地址、端口、Socket)
        server_name  localhost;  #虚拟主机配置（基于主机名、基于IP、基于域名）
        location / {
            root   html;   #配置请求的根目录
            index  index.html index.htm;   #设置网站的默认首页
        }
        error_page   500 502 503 504  /50x.html;    # 设置网站的错误页面
        location = /50x.html {
            root   html;   
   }   # server块结束
```

### Nginx日志切割

现有的日志都会存在于access.log文件中时间久了日志越来越多

#### 手动

1、创建一个shell可执行文件：cut_my_log.sh

```shell
cd /usr/local/nginx/sbin
vim cut_my_log.sh

#!/bin/bash
LOG_PATH="/var/log/nginx/"
RECORD_TIME=$(date -d "yesterday" + %Y-%m -%d + %H:%M)
PID=/var/run/nginx/nginx.pid
mv ${LOG_PATH}/access.log ${LOG_PATH}/access.${RECORD_TIME}.log
mv ${LOG_PATH}/error.log ${LOG_PATH}/error.${RECORD_TIME}.log

#向nginx主进程发送信号，用于重新打开日志文件
kill -USR1 `cat $PID`
```

2、为cut_my_log.sh添加可执行的权限

```shell
chmod +x cut_my_log.sh
```

3、测试日志切割后的结果

```shell
./cut_my_log.sh
cd /var/log/nginx
ls
```

#### 自动(定时)

1、安装定时任务

```shell
yum install -y crontabs # ps aux|grep crontabs先看看本地是否存在
```

2、crontab -e编辑并且添加一行新的任务

```shell
*/1 * * * * /usr/local/nginx/sbin/cut_my_log.sh
```

3、重启定时任务

```shell
service crond restart
```

附：常用定时任务命令

```shell
crontab -l  #查看任务列表
crontab -e  #编辑任务
service crond start #启动服务
service crond stop  #关闭服务
service crond restart  #重启服务
service crond reload  #重新载入配置
```

4、crontab时间格式内容

```shell
*    *    *    *    *    command
M    H    D    m    d    command
分   时   日   月   周   命令
第1列表示分钟1～59 每分钟用*或者 */1表示
第2列表示小时1～23（0表示0点）
第3列表示日期1～31
第4列表示月份1～12
第5列标识号星期0～6（0表示星期天）
第6列要运行的命令或脚本内容
```

每分钟执行

```shell
*/1 * * * *
* * * * *
```

每日凌晨(每天晚上23:59)执行

```shell
59 23 * * *
```

每日凌晨1点执行

```shell
0 1 * * *
```

每天定时为数据库备份

```html
https://www.cnblogs.com/leechenxiang/p/7110382.html
```

### Nginx静态资源

