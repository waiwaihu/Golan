# ElasticSearch

官网：https://www.elastic.co/cn/

```html
下载地址：https://www.elastic.co/cn/downloads/elasticsearch

Run bin/elasticsearch (or bin\elasticsearch.bat on Windows) to start Elasticsearch with security enabled.

ES 7.x 及之前版本，选择 Java 8
ES 8.x 及之后版本，选择 Java 17 或者 Java 18，建议 Java 17，因为对应版本的 Logstash 不支持 Java 18
Java 9、Java 10、Java 12 和 Java 13 均为短期版本，不推荐使用
M1（Arm） 系列 Mac 用户建议选择 ES 7.8.x 以上版本，因为考虑到 ELK 不同产品自身兼容性，7.8.x以上版本原生支持 Arm 原生 JDK

各大版本的jdk国内镜像：https://www.injdk.cn/
ARM JDK:https://www.azul.com/downloads/?version=java-8-lts&os=macos&package=jdk

Azul JDK（支持ARM原生：苹果M1、M2系列）
地址：https://www.azul.com/downloads/
```

# Kibana

```html
下载地址：https://www.elastic.co/cn/downloads/kibana

Run bin/kibana (or bin\kibana.bat on Windows)
```

# Linux 安装JDK

```shell
tar -zxvf test.tar.gz -C test
注释：上面的命令将 test.tar.gz 这个压缩包解压到当前目录下的 test 目录下。

 -C 选项的作用是：指定需要解压到的目录。
```

```shell
#centos7.9 安装jdk1.8
/usr/local/src 上传jdk文件
#mkdir java
#tar -zxvf jdk-8u202-linux-x64.tar.gz -c java

#vi /etc/profile
在profile中最后添加如下内容：
#set java environment环境变量
JAVA_HOME=/usr/local/src/java/jdk1.8.0_202
JRE_HOME=/usr/local/src/java/jdk1.8.0_202/jre
CLASS_PATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JRE_HOME/lib
PATH=$PATH:$JAVA_HOME/bin:$JRE_HOME/bin
export JAVA_HOME JRE_HOME CLASS_PATH PATH

#source /etc/profile重启配置文件
	
#java -version
javac
```

# Elasticsearch目录结构

| 目录名称    | 描述                                                         |
| ----------- | ------------------------------------------------------------ |
| bin         | 可执行脚本文件，包括启动elastic服务、插件管理、函数命令等    |
| config      | 配置文件目录，如elastic配置、角色配置、jvm配置等             |
| lib         | elastic所依赖的java库                                        |
| data        | 默认的数据存放目录，包含节点、分片、索引、文档的所有数据，生产环境要求必须修改 |
| logs        | 默认的日志文件存储路径，生产环境务必修改                     |
| modules     | 包含所有的elastic模块，如cluster、discovery、indices等       |
| plugins     | 已经安装的插件的目录                                         |
| jdk/jdk.app | 7.x 以后特有，自带的 java 环境，8.x版本自带 jdk 17           |

### 基础配置

```shell
cluster.name：集群名称，节点根据集群名称确定是否是同一个集群。默认名称为 elasticsearch，但应将其更改为描述集群用途的适当名称。不要在不同的环境中重用相同的集群名称。否则，节点可能会加入错误的集群

node.name：节点名称，集群内唯一，默认为主机名。，但可以在配置文件中显式配置

network.host： 节点对外提供服务的地址以及集群内通信的ip地址，例如127.0.0.1和 [::1]。

http.port：对外提供服务的端口号，默认 9200

transport.port：节点通信端口号，默认 9300
```

### 开发模式和生产模式

```shell
开发模式：开发模式是默认配置（未配置集群发现设置），如果用户只是出于学习目的，而引导检查会把很多用户挡在门外，所以ES提供了一个设置项discovery.type=single-node。此项配置为指定节点为单节点发现以绕过引导检查。

生产模式：当用户修改了有关集群的相关配置会触发生产模式，在生产模式下，服务启动会触发ES的引导检查或者叫启动检查（bootstrap checks），所谓引导检查就是在服务启动之前对一些重要的配置项进行检查，检查其配置值是否是合理的。引导检查包括对JVM大小、内存锁、虚拟内存、最大线程数、集群发现相关配置等相关的检查，如果某一项或者几项的配置不合理，ES会拒绝启动服务，并且在开发模式下的某些警告信息会升级成错误信息输出。引导检查十分严格，之所以宁可拒绝服务也要阻止用户启动服务是为了防止用户在对ES的基本使用不了解的前提下启动服务而导致的后期性能问题无法解决或者解决起来很麻烦。因为一旦服务以某种不合理的配置启动，时间久了之后可能会产生较大的性能问题，但此时集群已经变得难以维护和扩展，ES为了避免这种情况而做出了引导检查的设置，本来在开发模式下为警告的启动日志会升级为报错（Error）。这种设定虽然增加了用户的使用门槛，但是避免了日后产生更大的问题。
```

### 创建ES服务账号

```shell
ES不允许使用root账号启动服务，如果你当前账号是root，则需要创建一个专有账户（以下命令均在root账户下执行，windows系统在power shell下执行）。如果你的账号不是root账号，此步骤可以跳过
useradd elastic
passwd elastic
```

### 单节点集群

##### 启动命令

| 启动     | windows                                 | linux                                   | macos                                   |
| -------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| 命令行   | cd elasticsearch\bin .\elasticsearch -d | cd elasticsearch/bin ./elasticsearch -d | cd elasticsearch/bin ./elasticsearch -d |
| 图形界面 | 在bin目录下双击elasticsearch.bat        | -                                       | 在bin目录下双击elasticsearch            |
| shell    | start \bin\elasticsearch.bat            | -                                       | open bin/elasticsearch                  |

##### 启动日志

```shell
ES在 7.x 版本时，控制台输出 started 时代表服务启动成功，和 7.x 版本不同，

ES 8.x 启动之后会输出以下信息，此时服务已经启动成功了。
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/c17b8cc1a13e46d588b16186d656871a.jpeg#pic_center)

```shell
首次启动 Elasticsearch 时，会自动进行以下安全配置：

为传输层和 HTTP 层生成 TLS 证书和密钥。
TLS 配置设置被写入elasticsearch.yml.
为 elastic 用户生成密码。
为 Kibana 生成一个注册令牌。
红框1：ES为我们生成的elastic账户的默认密码，重要，需要复制记下来

红框2：CA证书的密钥信息，暂时先不管

红框3：ES为Kibana生成的访问令牌，Kibana访问ES服务需要用到。（有效期为 30 分钟）

红框4：ES位其他节点加入集群生成的访问令牌，当前集群中需要加入新节点时，需要携带此令牌（有效期为 30 分钟）

然后您可以启动 Kibana 并输入有效期为 30 分钟的注册令牌。此令牌自动应用 Elasticsearch 集群中的安全设置，使用内置kibana服务帐户向 Elasticsearch 进行身份验证，并将安全配置写入kibana.yml
```

##### 修改账号密码

在 ES 8.x版本以后，elasticsearch-setup-passwords设置密码的工具已经被弃用删除，此命令为7.x之前第一次生成密码时使用，8.x在第一次启动的时候会自动生密码。

注意：上述（2.5.3截图）内容仅在第一次启动时显示，如果需要修改账户密码，需进行以下操作
```shell
bin/elasticsearch-reset-password

[-a, --auto] [-b, --batch] [-E <KeyValuePair]
[-f, --force] [-h, --help] [-i, --interactive]
[-s, --silent] [-u, --username] [--url] [-v, --verbose]
```

```shell
使用此命令重置本地领域中的任何用户或任何内置用户的密码。默认情况下，系统会为您生成一个强密码。要显式设置密码，请使用 以交互模式运行该工具-i。该命令在 文件领域中生成（并随后删除）一个临时用户，以运行更改用户密码的请求。

-a, --auto

将指定用户的密码重置为自动生成的强密码。（默认）

-b, --batch

运行重置密码过程而不提示用户进行验证。

-E

配置标准 Elasticsearch 或 X-Pack 设置。

-f, --force

强制命令针对不健康的集群运行。

-h, --help

返回所有命令参数。

-i, --interactive

提示输入指定用户的密码。使用此选项显式设置密码。

-s --silent

在控制台中显示最小输出。

-u, --username

本机领域用户或内置用户的用户名。

–url

指定工具用于向 Elasticsearch 提交 API 请求的基本 URL（本地节点的主机名和端口）。默认值由 elasticsearch.yml文件中的设置确定。如果xpack.security.http.ssl.enabled设置为true，则必须指定 HTTPS URL。

-v --verbose

在控制台中显示详细输出。

比如：

为elastic账号自动生成新密码，输出至控制台
bin/elasticsearch-reset-password -u elastic

手工指定user1的新密码
bin/elasticsearch-reset-password --username elastic -i

指定服务地址和账户名
bin/elasticsearch-reset-password --url "https://172.0.0.3:9200" --username elastic -i
```

### 验证服务状态

```shell
在7.x的版本是通过如下地址访问ES服务：http://localhost:9200/
Elastic 8 默认开启了 SSL，将默认配置项由true改为false即可
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/8bbda36e561f480f9941b75bdff4ef94.jpeg#pic_center)

```shell
关闭SSL虽然可以访问服务了，但这本质上是在规避问题而非解决问题，更推荐的做法是使用https协议进行访问：

https://localhost:9200/
输入账号密码验证：

此时输入账号，也就是在 2.5.4 的启动日志中，红框一内的内容，确定即可访问ES服务，至此，单节点ES服务启动成功。
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/4cda4acc08fe4caaa72d2d2c56fed916.jpeg#pic_center)

### 构建基于Security的本地集群

```shell
向集群中加入新节点

默认情况下，要集群中添加新节点，需要通过令牌来完成节点之间的通信，在第一个节点启动的时候，控制台会输出令牌信息（2.5.3 启动日志中的红框4中的Token），注意启动的时候

bin/elasticsearch --enrollment-token <token> 

//启动的时候替换<token>，不要待带括号 比如：
bin/elasticsearch --enrollment-token eyJ2ZXIiOiI4LjEuMCIsImFkciI6WyIxOTIuMTY4LjMuMTAwOjkyMDEiXSwiZmdyIjoiMWJkMTE0OWMzMTJjYzc5MGU1ZWU1YTgzZjlhZWRjMmU2MDkyN2Y2MWVkZDA0ZWU0YTAxZTk2MTVlYzJkODhlYiIsImtleSI6Ilo3elc0bjhCVk1ESEdsMmFzdDVNOkZTWVhrMHV3UjgyUzNlTFFERFlkdncifQ==

```

**如下图为一个基于ES 8.x 的三节点集群：**

![在这里插入图片描述](https://img-blog.csdnimg.cn/25a7cd31d094465c9d95ebadcb8e4bcd.jpeg#pic_center)

### 部署Kibana

```shell
直接启动Kibana服务，控制台显示以下信息
然后我们访问控制台输出的Kibana的服务地址，在页面提示中输入 7.2.2 红框三中的访问令牌，确定即可
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/303bcb6bbaf1461d985fed0ce6ec229c.jpeg#pic_center)

### 服务的安装和启动（基于Security关闭）

```shell
ES 8 默认是开启Security的，我们现在需要修改器配置文件使其关闭Security。打开 Config 目录，修改 elasticsearch.yml 文件，删除文件内所有内容，配置以下信息：
xpack.security.enabled: false

此时，配置文件中仅一行代码：
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/ccbadbeb2d3348c6a4760c2406140e86.jpeg#pic_center)



### 单项目多节点启动

```shell
操作系统	命令
LinuxMacOS	节点1：./elasticsearch -E path.data=data1 -E path.logs=log1 -E node.name=node1 -E cluster.name=elastic.org.cn节点2：./elasticsearch -E path.data=data2 -E path.logs=log2 -E node.name=node2 -E cluster.name=elastic.org.cn节点N：… …

Windows	节点1：.\elasticsearch.bat -E path.data=data1 -E path.logs=log1 -E node.name=node1 -E cluster.name=elastic.org.cn节点2：.\elasticsearch.bat -E path.data=data2 -E path.logs=log2 -E node.name=node1 -E cluster.name=elastic.org.cn节点N：… …
```

###  多项目多节点启动

```shell
操作系统	脚本
MacOS	open /node1/bin/elasticsearch open /node2/bin/elasticsearch open /node3/bin/elasticsearch

windows	start D:\node1\bin\elasticsearch.bat start D:\node2\bin\elasticsearch.bat start D:\node3\bin\elasticsearch.bat
```

### 部署 Kibana

**示例****：**下图中包含一个3节点集群，每个节点都是独立的SDK文件

![在这里插入图片描述](https://img-blog.csdnimg.cn/1e60f5ff92ba47f1818b96f04a9b94e7.jpeg#pic_center)

**验证服务状态：**浏览器执行 `http://localhost:9200/_cat/nodes`（**注意，这里和单节点启动方式不同**）

**优点**：配置简单，一劳永逸

**缺点**：占用较多磁盘空间，因为每个节点都有一套独立的SDK文件，大约几百MB。

### 推荐安装的几款浏览器插件

```shell
Elasticsearch Head	img	方便查看集群节点数据方便管理和索引、分片支持同时连接多集群
https://github.com/mobz/elasticsearch-head github下载
https://www.elastic.org.cn/archives/es-head 安装教程

Elasticvue	img	功能强大对国人友好
https://microsoftedge.microsoft.com/addons/search/elasticvue?hl=zh-CN  Edge下载
```

# 狂神windows教程

电脑内存小的可以修改jvm.options中占用内存大小

```shell
-Xms1g  --256M
-Xmx1g  --512M
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200909180913899.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

### 启动

```shell
双击elasticsearch.bat
http://localhost:9200/
```

### 安装Node

```shell
https://nodejs.org/en/
node -v

#第一种
https://github.com/mobz/elasticsearch-head 下载插件
https://gitcode.net/mirrors/mobz/elasticsearch-head?utm_source=csdn_github_accelerator#这里下载的版本比较多

cd elasticsearch-head

设置npm全局安装路径与缓存路径,CMD输入
npm config set prefix "D:\nodejs\node_global"
npm config set cache "D:\nodejs\node_cache"
npm install -g cnpm --registry=https://registry.npm.taobao.org
在系统变量下新建”NODE_PATH”，输入”d:\nodejs\node_global\node_modules
用户变量”PATH”修改为“D:\nodejs\node_global\”
cnpm -v

cnpm install
npm run start
open http://localhost:9100/

#第二种
https://git-scm.com/download/win 下载Git
git --version
git clone git://github.com/mobz/elasticsearch-head.git
cd elasticsearch-head
cnpm install
cmd命令行
npm run start
open http://localhost:9100/
```

### 跨域问题

```shell
打开glasticsearch-head-master文件夹，修改Gruntfile.js文件，添加hostname:'*', 如图
hostname:'*'
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/2a31538b2e9e4255a44f137099f252ec.png?x-oss-process=image/watermark,type_d3F5LXplbmhlaQ,shadow_50,text_Q1NETiBARWxhc3RpY-W8gOa6kOekvuWMug==,size_20,color_FFFFFF,t_70,g_se,x_16)

```shell
#无法发现ES节点，尝试在ES配置文件中设置允许跨域
cd E:\elasticsearch-7.17.8-windows-x86_64\config
vim elasticsearch.yml .在最后一行添加如下内容:
http.cors.enabled: true
http.cors.allow-origin: "*"
再次重启服务
```

![image-20230103140713218](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20230103140713218.png)

### Kibana安装

```shell
cd E:\kibana-7.17.8-windows-x86_64\config\kibana.yml

#直接在最后一行添加
i18n.locale: "zh-CN"            # 中文汉化

#这一段是解释，没有添加
server.port: 5601
server.host: "0.0.0.0"
elasticsearch.hosts: ["http://192.168.1.30:9201"]
kibana.index: ".kibana" 
```

```shell
#启动测试
双击 kibana.bat
http://localhost:5601/
```

### ES核心概念

- 索引
- 字段类型（mapping）
- 文档（document）

![img](https://img-blog.csdnimg.cn/20200909192956336.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

### IK分词器

```shell
下载地址：https://github.com/medcl/elasticsearch-analysis-ik/releases?after=v7.8.0

cd plugins/ #解压放入到es对应的plugins下即可，重启观察ES，发现ik插件被加载了
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/2020091009234593.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```shell
也可以通过E:\elasticsearch-7.17.7-windows-x86_64\bin目录下输入cmd进入命令行模式，输入elasticsearch-plugin list 查看已经加载的插件
```

如果要使用中文，建议使用ik分词器，IK提供了两个分词算法：ik_smart和ik_max_word，其中ik_smart为最少切分，ik_max_word为最细颗粒度划分

```shell
ik_smart: 会做最粗粒度的拆分
ik_max_word: 会将文本做最细粒度的拆分,穷尽词库的可能， 字典！

POST /_analyze
{
  "text":"中华民族共和国国歌",
  "analyzer":"ik_smart"
}

POST /_analyze
{
  "text":"中华民族共和国国歌",
  "analyzer":"ik_max_word"
}

GET _analyze
{
  "text":"中华民族共和国国歌",
  "analyzer":"ik_max_word"
}
```

#### 扩展词和关键词

- 扩展词：就是有些词并不是关键词，但是也希望被ES用来作为检索的关键词，可以将这些词加入扩展词典
- 停用词：就是有些关键词，我们并不想让他被检索到，可以放入停用词典中

```shell
设置扩展词典和停用词典在es容器中的config/analysis-ik目录下的IKAnalyzer.cfg.xml中
```

```shell
1. 修改vim IKAnalyzer.cfg.xml

    <?xml version="1.0" encoding="UTF-8"?>
    <!DOCTYPE properties SYSTEM "http://java.sun.com/dtd/properties.dtd">
    <properties>
        <comment>IK Analyzer 扩展配置</comment>
        <!--用户可以在这里配置自己的扩展字典 -->
        <entry key="ext_dict">ext_dict.dic</entry>
         <!--用户可以在这里配置自己的扩展停止词字典-->
        <entry key="ext_stopwords">ext_stopword.dic</entry>
    </properties>

2. 在es容器中`config/analysis-ik`目录下中创建ext_dict.dic文件   编码一定要为UTF-8才能生效
	vim ext_dict.dic 加入扩展词即可

3. 在es容器中`config/analysis-ik`目录中创建ext_stopword.dic文件 
	vim ext_stopword.dic 加入停用词即可
	
4.重启es生效
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910101747942.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

s本身也提供了一些常用的扩展词典和停用词典可以直接使用

![[外链图片转存失败,源站可能有防盗链机制,建议将图片保存下来直接上传(img-QEbsSqov-1655732475477)(ElasticSearch.assets/image-20220410150826964.png)]](https://img-blog.csdnimg.cn/ab9c96764cce42088ef301c3fc876301.png)

### Restful风格说明

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910102407169.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

#### 创建一个索引！

```shell
PUT /索引名/~类型名~/文档id
{请求体}
 
# PUT 创建命令  test1 索引 type1 类型 1 id
PUT test1/type1/1
{
  "name": "xiaofan",
  "age": 28
}
 
# 返回结果
# 警告信息： 不支持在文档索引请求中的指定类型
# 而是使用无类型的断点(/{index}/_doc/{id}, /{index}/_doc, or /{index}/_create/{id}).
{
  "_index" : "test1",   # 索引
  "_type" : "type1",    # 类型（已经废弃）
  "_id" : "1",          # id
  "_version" : 1,       # 版本
  "result" : "created", # 操作类型
  "_shards" : {         # 分片信息
    "total" : 2,
    "successful" : 1,
    "failed" : 0
  },
  "_seq_no" : 0,
  "_primary_term" : 1
}
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910105816144.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

#### 指定字段的类型（创建规则）

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910113123847.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

#### 获取具体的索引规则

```shell
# GET test2
 
{
  "test2" : {
    "aliases" : { },
    "mappings" : {
      "properties" : {
        "age" : {
          "type" : "integer"
        },
        "birthday" : {
          "type" : "date"
        },
        "name" : {
          "type" : "text"
        }
      }
    },
    "settings" : {
      "index" : {
        "creation_date" : "1599708623941",
        "number_of_shards" : "1",
        "number_of_replicas" : "1",
        "uuid" : "ANWnhwArSMSl8k8iipgH1Q",
        "version" : {
          "created" : "7080099"
        },
        "provided_name" : "test2"
      }
    }
  }
}
 
# 查看默认的规则
PUT /test3/_doc/1
{
  "name": "狂神说Java",
  "age": 28,
  "birthday": "1997-01-05"
}
 
# GET test3
 
{
  "test3" : {
    "aliases" : { },
    "mappings" : {
      "properties" : {
        "age" : {
          "type" : "long"
        },
        "birthday" : {
          "type" : "date"
        },
        "name" : {
          "type" : "text",
          "fields" : {
            "keyword" : {
              "type" : "keyword",
              "ignore_above" : 256
            }
          }
        }
      }
    },
    "settings" : {
      "index" : {
        "creation_date" : "1599708906181",
        "number_of_shards" : "1",
        "number_of_replicas" : "1",
        "uuid" : "LzPLCDgeQn6tdKo3xBBpbw",
        "version" : {
          "created" : "7080099"
        },
        "provided_name" : "test3"
      }
    }
  }
}
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910114250259.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```shell
GET _cat/indices?v
```

#### 修改索引 POST

```shell
# 只会修改指定项，其他内容保证不变
POST /test3/_doc/1/_update
{
  "doc": {
    "name":"暴徒狂神"
  }
}
 
# GET test3/_doc/1
 
{
  "_index" : "test3",
  "_type" : "_doc",
  "_id" : "1",
  "_version" : 2,
  "_seq_no" : 1,
  "_primary_term" : 1,
  "found" : true,
  "_source" : {
    "name" : "暴徒狂神",
    "age" : 28,
    "birthday" : "1997-01-05"
  }
}
```

#### 删除索引

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910115620763.png#pic_center)

### 关于文档的基本操作

#### 基本操作（简单的查询）

```shell
put /kuangshen/user/1
{
  "name": "狂神说",
  "age": 23,
  "desc": "一顿操作猛如虎，一看工资2500",
  "tags": ["码农", "技术宅", "直男"]
}
 
put /kuangshen/user/2
{
  "name": "张三",
  "age": 28,
  "desc": "法外狂徒",
  "tags": ["旅游", "渣男", "交友"]
}
 
put /kuangshen/user/3
{
  "name": "李四",
  "age": 30,
  "desc": "不知道怎么描述",
  "tags": ["旅游", "靓女", "唱歌"]
}
 
GET kuangshen/user/1
 
GET kuangshen/user/_search?q=name:狂神
```

#### 复杂操作(排序、分页、高亮、模糊查询、标准查询！)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910140701583.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```shell
# 模糊查询
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "name": "狂神"
    }
  }
}
 
# 对查询结果进行字段过滤
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "name": "狂神"
    }
  },
  "_source": ["name", "desc"]
}
 
# 排序
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "name": "狂神"
    }
  },
  "sort":[{
    "age": "asc"
  }]
}
# 排序2
{
  "query": {
    "match": {
      "name": "狂神"
    }
  },
  "sort":[
    {
      "age":{
        "order":"desc"
      }
    }
  ]
}

 
# 分页
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "name": "狂神"
    }
  },
  "sort":[{
    "age": "asc"
  }], 
  "from": 0, #从哪里开始
  "size": 2 #每页显示的数量
}
```

#### 布尔值条件查询

```shell
# 多条件查询 must 相当于and
GET kuangshen/user/_search
{
  "query": {
    "bool": {
      "must": [
        {"match": {
          "name": "狂神"
        }},
        {"match": {
          "age": 23
        }}
      ]
    }
  }
}
 
# 多条件查询 should 相当于or
GET kuangshen/user/_search
{
  "query": {
    "bool": {
      "should": [
        {"match": {
          "name": "狂神说"
        }},
        {"match": {
          "age": 25
        }}
      ]
    }
  }
}
 
# 多条件查询 must_not 相当于 not
GET kuangshen/user/_search
{
  "query": {
    "bool": {
      "must_not": [
        {"match": {
          "age": 25
        }}
      ]
    }
  }
}
 
 
# 过滤查询1 age > 24
GET kuangshen/user/_search
{
  "query": {
    "bool": {
      "must": [
        {"match": {
          "name": "狂神"
        }}
      ],
      "filter": [
        {"range": {
          "age": {
            "gt": 24
          }
        }}
      ]
    }
  }
}
 
# 过滤器2  22<age<30 
GET kuangshen/user/_search
{
  "query": {
    "bool": {
      "must": [
        {"match": {
          "name": "狂神"
        }}
      ],
      "filter": [#过滤
        {"range": { #范围
          "age": {
            "lt": 30,
            "gt": 22
          }
        }}
      ]
    }
  }
}
#gt 大于
#gte 大于等于
#lt 小于
#lte 小于等于
#eq 等于
#neq 不等于
```

#### 多条件查询

```shell
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "tags": "技术 男"
    }
  }
}
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200910144807544.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

#### 精确查询

![img](https://img-blog.csdnimg.cn/2020091014524041.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```shell
text keyword类型不会被分词器解析
term: 精确匹配
```

```shell
# 定义类型
PUT xiaofan_test_db
{
  "mappings": {
    "properties": {
      "name": {
        "type": "text"
      },
      "desc": {
        "type": "keyword"
      }
    }
  }
}
 
 
PUT /xiaofan_test_db/_doc/1
{
  "name": "小范说Java Name",
  "desc": "小范说Java Desc"
}
 
PUT /xiaofan_test_db/_doc/2
{
  "name": "小范说Java Name",
  "desc": "小范说Java Desc 2"
}
 
# 按照keyword类型精准匹配
GET xiaofan_test_db/_search
{
  "query": {
    "term": {
      "desc": "小范说Java Desc"
    }
  }
}
# 结果：
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 1,
      "relation" : "eq"
    },
    "max_score" : 0.6931471,
    "hits" : [
      {
        "_index" : "test_db",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.6931471,
        "_source" : {
          "name" : "小范说Java Name",
          "desc" : "小范说Java Desc"
        }
      }
    ]
  }
}
 
# 按照text类型匹配
GET xiaofan_test_db/_search
{
  "query": {
    "term": {
      "name": "小"
    }
  }
}
 
# 结果：
{
  "took" : 0,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : 0.18232156,
    "hits" : [
      {
        "_index" : "test_db",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 0.18232156,
        "_source" : {
          "name" : "小范说Java Name",
          "desc" : "小范说Java Desc"
        }
      },
      {
        "_index" : "test_db",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 0.18232156,
        "_source" : {
          "name" : "小范说Java Name",
          "desc" : "小范说Java Desc 2"
        }
      }
    ]
  }
}
```

#### 多个值匹配精确查询

```shell
PUT /test_db/_doc/3
{
  "t1": "22",
  "t2": "2020-09-10"
}
 
PUT /test_db/_doc/4
{
  "t1": "33",
  "t2": "2020-09-11"
}
 
GET test_db/_search
{
  "query": {
    "bool": {
      "should": [
        {
          "term": {
            "t1": "22"
          }
        },
         {
          "term": {
            "t1": "33"
          }
        }
      ]
    }
  }
}
```

#### 高亮查询

```shell
GET kuangshen/user/_search
{
  "query": {
    "match": {
      "name": "狂神"
    }
  },
  "highlight": {
    "pre_tags": "<p class='key' style='color:red'>", #前缀
    "post_tags": "</p>", #后缀
    "fields": {
      "name": {}
    }
  }
}
 
# 结果显示：
#! Deprecation: [types removal] Specifying types in search requests is deprecated.
{
  "took" : 1,
  "timed_out" : false,
  "_shards" : {
    "total" : 1,
    "successful" : 1,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : {
      "value" : 2,
      "relation" : "eq"
    },
    "max_score" : 1.3862942,
    "hits" : [
      {
        "_index" : "kuangshen",
        "_type" : "user",
        "_id" : "1",
        "_score" : 1.3862942,
        "_source" : {
          "name" : "狂神说",
          "age" : 23,
          "desc" : "一顿操作猛如虎，一看工资2500",
          "tags" : [
            "码农",
            "技术宅",
            "直男"
          ]
        },
        "highlight" : {
          "name" : [
            "<p class='key' style='color:red'>狂</p><p class='key' style='color:red'>神</p>说"
          ]
        }
      },
      {
        "_index" : "kuangshen",
        "_type" : "user",
        "_id" : "4",
        "_score" : 1.0892314,
        "_source" : {
          "name" : "狂神说前端",
          "age" : 25,
          "desc" : "大王叫我来巡山",
          "tags" : [
            "码农1",
            "技术宅1",
            "直男1"
          ]
        },
        "highlight" : {
          "name" : [
            "<p class='key' style='color:red'>狂</p><p class='key' style='color:red'>神</p>说前端"
          ]
        }
      }
    ]
  }
}
```

# Docker安装Es、ik分词器、kabana

注意：老高的教程失败，没启动成功

### 1、安装Es

```shell
#1、安装Es
docker pull elasticsearch:7.17.7

docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.17.7

-d：后台启动 --name：容器名称 -p：端口映射 -e：设置环境变量 discovery.type=single-node：单机运行 ec0817395263：镜像id/或者镜像名 


如果启动不了，可以减小内存设置：-e ES_JAVA_OPTS=“-Xms64m -Xmx512m”
docker run -d --name es -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.17.7

/*这种容器内有vi/vim可以操作
    docker exec -it es /bin/bash
    cd config
    vim elasticsearch.yml
*/
#一般选择这种
mkdir -p /usr/local/elasticsearch/
docker cp es:/usr/share/elasticsearch/config/elasticsearch.yml /usr/local/elasticsearch/

vim /usr/local/elasticsearch/elasticsearch.yml 
# 末尾加入跨域配置,elasticsearch-head-master互通
http.cors.enabled: true
http.cors.allow-origin: "*"

docker cp /usr/local/elasticsearch/elasticsearch.yml es:/usr/share/elasticsearch/config

#重启容器
docker restart es
```

### 2、安装IK分词器

```shell
https://github.com/medcl/elasticsearch-analysis-ik/releases
IK版本一定要和es版本一致 比如：7.17.7

#离线安装
下载和es一样的版本：https://github.com/medcl/elasticsearch-analysis-ik/releases
#进入容器，创建目录
docker exec -it es /bin/bash
cd plugins
mkdir ik

#将文件压缩包移动到ik中
docker cp elasticsearch-analysis-ik-7.17.7.zip es:/usr/share/elasticsearch/plugins/ik

# 进入ik目录，解压压缩包
cd /usr/share/elasticsearch/plugins/ik
unzip elasticsearch-analysis-ik-7.17.7.zip
rm -rf elasticsearch-analysis-ik-7.17.7.zip

#退出容器
exit
docekr restart es #重启es
postman发送请求时：把 Content-Type 设置为 application/json # 测试ik分词器
POST http://192.168.222.132:9200/_analyze
{
  "text":"中华民族共和国国歌",
  "analyzer":"ik_max_word"
}
```

### 3、部署 ElasticSearch-Head

```shell
#3、部署 ElasticSearch-Head
docker pull mobz/elasticsearch-head:5
docker run -d --name es_head -p 9100:9100 mobz/elasticsearch-head:5
```

### 4、老高安装ES集群

```shell
#1、虚拟机内存设为4G
#2、修改limits.conf，limits.conf 可以设置用户最大可创建文件的数量,不设置ES可能会报错
vi /etc/security/limits.cong 文末添加
    * soft nofile 65536
    * hard nofile 131072
    
#3、修改sysctl.conf  可以设置最大虚拟内存sysctl -p 查看运行时动态的修改内核的运行参数
vi /etc/sysctl.conf 文末添加
	vm.max_map_count=655360

#4、创建目录，node_1
mkdir -p /usr/local/es_cluster/node_1/data
mkdir -p /usr/local/es_cluster/node_1/logs
mkdir -p /usr/local/es_cluster/node_1/plugins/ik

#4、2 node_2
mkdir -p /usr/local/es_cluster/node_2/data
mkdir -p /usr/local/es_cluster/node_2/logs
mkdir -p /usr/local/es_cluster/node_2/plugins/ik

#4、3 node_3
mkdir -p /usr/local/es_cluster/node_3/data
mkdir -p /usr/local/es_cluster/node_3/logs
mkdir -p /usr/local/es_cluster/node_3/plugins/ik

#4.4 目录结构
yum install -y tree
tree /usr/local/es_cluster

$5、分词器，将压缩包上传linux的/opt目录下
yum install -y unzip
unzip /opt/elasticsearch-analysis-ik-7.17.7.zip -d /usr/local/es_cluster/node_1/plugins/ik

unzip /opt/elasticsearch-analysis-ik-7.17.7.zip -d /usr/local/es_cluster/node_2/plugins/ik

unzip /opt/elasticsearch-analysis-ik-7.17.7.zip -d /usr/local/es_cluster/node_3/plugins/ik

#删除文件
rm -f /opt/elasticsearch-analysis-ik-7.17.7.zip

#6、安装Es
docker pull elasticsearch:7.17.7

#6.1、创建并启动ES容器
"cluster.name" #ES集群名称
"node.name" #节点名称
"node.master=true" # 是否可以成为master节点
"node.data=true" #是否允许该节点存储数据，默认开启
"netwotk.host=0.0.0.0" #绑定主机的ip地址
"transport.tcp.port=9300" # 设置集群节点之间交互的tcp端口
"http.port=9200" #设置http端口
"cluster.initial_master_nodes=node-1" #设置那些节点参与第一次master几点选举，其值可以是ES的ip地址
"discovery.seed_hosts=192.168.222.99:9301,192.168.222.99:9302,192.168.222.99:9303" #设置当前节点与那些ES节点建立连接，其值可以是127.0.0.1
"gateway.auto_import_dangling_indices=true" #是否自动引入dangling索引，默认flase
"http.cors.enabled=true" #开启cors跨域访问，默认flase
"http.cors.allow-origin=*" #跨域访问允许的域名地址，允许所有
"ES_JAVA_OPTS=-Xms512m -Xms512m" #设置jvm最小内存，默认2G
"TAKE_FILE_OWNERSHIP=true"

#node1
docker run -d --name=es_node_1 --restart=always -p 9201:9200 -p 9301:9300 --privileged=true \
-v /usr/local/es_cluster/node_1/data:/usr/share/elasticsearch/data \
-v /usr/local/es_cluster/node_1/logs:/usr/share/elasticsearch/logs \
-v /usr/local/es_cluster/node_1/plugins:/usr/share/elasticsearch/plugins \
-e "cluster.name=my-cluster" \
-e "node.name=node-1" \
-e "node.master=true" \
-e "node.data=true" \
-e "netwotk.host=0.0.0.0" \
-e "transport.tcp.port=9300" \
-e "http.port=9200" \
-e "cluster.initial_master_nodes=node-1" \
-e "discovery.seed_hosts=192.168.222.99:9301,192.168.222.99:9302,192.168.222.99:9303" \
-e "gateway.auto_import_dangling_indices=true" \
-e "http.cors.enabled=true" \
-e "http.cors.allow-origin=*" \
-e "ES_JAVA_OPTS=-Xms512m -Xms512m" \
-e "TAKE_FILE_OWNERSHIP=true" \
elasticsearch:7.17.7

#node2
docker run -d --name=es_node_2 --restart=always -p 9202:9200 -p 9302:9300 --privileged=true \
-v /usr/local/es_cluster/node_2/data:/usr/share/elasticsearch/data \
-v /usr/local/es_cluster/node_2/logs:/usr/share/elasticsearch/logs \
-v /usr/local/es_cluster/node_2/plugins:/usr/share/elasticsearch/plugins \
-e "cluster.name=my-cluster" \
-e "node.name=node-2" \
-e "node.master=true" \
-e "node.data=true" \
-e "netwotk.host=0.0.0.0" \
-e "transport.tcp.port=9300" \
-e "http.port=9200" \
-e "cluster.initial_master_nodes=node-1" \
-e "discovery.seed_hosts=192.168.222.99:9301,192.168.222.99:9302,192.168.222.99:9303" \
-e "gateway.auto_import_dangling_indices=true" \
-e "http.cors.enabled=true" \
-e "http.cors.allow-origin=*" \
-e "ES_JAVA_OPTS=-Xms512m -Xms512m" \
-e "TAKE_FILE_OWNERSHIP=true" \
elasticsearch:7.17.7

#node3
docker run -d --name=es_node_3 --restart=always -p 9203:9200 -p 9303:9300 --privileged=true \
-v /usr/local/es_cluster/node_3/data:/usr/share/elasticsearch/data \
-v /usr/local/es_cluster/node_3/logs:/usr/share/elasticsearch/logs \
-v /usr/local/es_cluster/node_3/plugins:/usr/share/elasticsearch/plugins \
-e "cluster.name=my-cluster" \
-e "node.name=node-3" \
-e "node.master=true" \
-e "node.data=true" \
-e "netwotk.host=0.0.0.0" \
-e "transport.tcp.port=9300" \
-e "http.port=9200" \
-e "cluster.initial_master_nodes=node-1" \
-e "discovery.seed_hosts=192.168.222.99:9301,192.168.222.99:9302,192.168.222.99:9303" \
-e "gateway.auto_import_dangling_indices=true" \
-e "http.cors.enabled=true" \
-e "http.cors.allow-origin=*" \
-e "ES_JAVA_OPTS=-Xms512m -Xms512m" \
-e "TAKE_FILE_OWNERSHIP=true" \
elasticsearch:7.17.7

#验证是否开启
docker ps

#7、开放端口
firewall-cmd --add-port=9201/tcp --add-port=9202/tcp --add-port=9203/tcp --permanent
firewall-cmd --add-port=9301/tcp --add-port=9302/tcp --add-port=9303/tcp --permanent
firewall-cmd --reload #重新加载
firewall-cmd --list-ports#查看放行端口

#4.4数据包转发
vi /usr/lib/sysctl.d/50-default.conf

#文末添加内容
net.ipv4.ip_forward = 1

systemctl restart network

#4.5、验证
http://192.168.222.99:9201/_cat/nodes?pretty #验证失败，浏览器打不开
```

### 5、老高安装kabana

```dockerfile
#4.1安装
docker pull kibana:7.17.7

#4.2创建目录
mkdir -p /usr/local/kibana
vi /usr/local/kibana/kibana.yml

#添加配置内容
server.name: "kibana"
server.host: "0.0.0.0"
server.port: 5601
server.shutdownTimeout: "5s"
elasticsearch.requestTimeout:120000
elasticsearch.hosts: [ "http://172.17.0.2:9200" ]
monitoring.ui.container.elasticsearch.enabled: true
i18n.locale: "zh-CN" # 中文汉化

#后者上面ES地址换成设置ES集群地址
elasticsearch.hosts: ["http://192.168.222.132:9200","http://192.168.222.132:9201","http://192.168.222.132:9202"]

#4.3创建并启动容器
docker run -d -p 5601:5601 --name=kibana --restart=always -v /usr/local/kibana/kibana.yml:/usr/share/kibana/config/kibana.yml kibana:7.17.7

#4.4开放端口
firewall-cmd --add-port=5601/tcp --permanent
firewall-cmd --reload #重新加载
firewall-cmd --list-ports#查看放行端口

#4.5验证
http://localhost:5601/app/dev_tools#/console #验证失败，浏览器打不开
```

# Golang客户端

```shell
https://www.elastic.co/guide/en/elasticsearch/client/index.html
#选择go
https://www.elastic.co/guide/en/elasticsearch/client/go-api/7.17/index.html

require github.com/elastic/go-elasticsearch/v7 7.16 #7版本
require github.com/elastic/go-elasticsearch/v8 8.5 #8版本
```

# Dockerb部署Es、kibana、Logstash

```shell
#创建网络
docker network create somenetwork

#1、安装es
mkdir -p /usr/local/es/data
chmod 777 /usr/local/es/data
docker pull elasticsearch:7.17.7
docker run -d -p 9201:9200 -p 9301:9300 --name=es --net somenetwork -v /usr/local/es/data:/usr/share/elasticsearch/data -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.17.7

#2、安装ik分词器
docker exec -it es /bin/bash
cd plugins
mkdir ik
#将文件压缩包移动到ik中
docker cp elasticsearch-analysis-ik-7.17.7.zip es:/usr/share/elasticsearch/plugins/ik

# 进入ik目录，解压压缩包
cd /ik
unzip elasticsearch-analysis-ik-7.17.7.zip
rm -rf elasticsearch-analysis-ik-7.17.7.zip

docker restart es
docker logs es --since 30m
http://ip:9201/

#3、安装kibana
docker pull kibana:7.17.7

#容器间互联的方法：--link（单方向的互联）
docker run -d -p 5602:5601 --name=kibana --net somenetwork --link es:elasticsearch kibana:7.17.7
docker logs -f kibana

#第一种：修改配置文件
docker exec -it kibana /bin/bash
cd config
vi kibana.yml

#第二种：容器vi命令不能修改的方法
mkdir -p /usr/local/kibana
docker cp kibana:/usr/share/kibana/config/kibana.yml /usr/local/kibana/
docker cp /usr/local/kibana/kibana.yml kibana:/usr/share/kibana/config

文末添加：i18n.locale: "zh-CN" # 中文汉化
重启kibana docker restart kibana
http://ip:5602/

#4\安装logstash
docker pull logstash:7.17.7
docker run -d -p 5045:5044 --name=logstash logstash:7.17.7

#第二种：容器vi命令不能修改的方法
mkdir -p /usr/local/logstash
docker cp logstash:/usr/share/logstash/config/logstash.yml /usr/local/logstash
docker cp /usr/local/logstash/logsrash.yml logstash:/usr/share/logstash/config/

docker cp logstash:/usr/share/logstash/pipeline/logstash.conf /usr/local/logstash
docker cp /usr/local/logstash/logsrash.yml logstash:/usr/share/logstash/pipeline/

vi logstash.yml
vi /usr/share/logstash/pipeline/logstash.conf
```

