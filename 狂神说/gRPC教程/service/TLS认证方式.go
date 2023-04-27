package service

//官网下载 https://www.openssl.org/source/
//其他人做的便携版本 http://slproweb.com/products/Win32OpenSSL.html
//我们使用便携版,选择这个版本EXE: Win64 OpenSSL v3.0.7 EXE | MSI	140MB Installer 一直下一步即可
//配置环境变量 D:\\OpenSSL-Win64\bin,命令行测试openssl

/*
key：服务器上的私钥文件，用于对发送给客户端数据的加密，以及对从客户端接收到数据的加密
csr：证书签名请求文件，用于提交给证书颁发机构(CA)对证书签名
crt：由证书颁发机构签名后的证书，或者是开发者自签名的证书，包含证书持有人的信息，持有人的公钥，以及签署者的签名等信息
pem：是基于base64编码的证书格式，扩展名包括PEM,CRT和CER。

cd server
cd key

1、生成私钥
openssl genrsa -out server.key 2048

2、生成证书 全部回车即可 可以不填
openssl req -new -x509 -key server.key -out server.crt -days 36500

下面可以不填，一直回车
国家名称
Countty Name (2 letter code) [AU]:CN

省名称
State or Province Name (full name) [Some-State]:GuangDong

城市名称
Locality Name (eg, city) []:MeiZhou

公司组织名称
Organization Name (eg, company) [Internet Widgits Pty Ltd]:Xuexiangban

部门名称
Organization Unit Name (eg, section) []:go

服务器or网站名称
Common Name (e.g. server FQDN or YOUR name) []:kuangstudy

邮件
Email Address []:1632976236@qq.com

3、生成csr证书
openssl req -new -key server.key -out server.csr
下面信息也可以不填，一路回车
*/

/*
更改openssl.cnf(linux 是openssl.cfg)
D:\OpenSSL-Win64\bin\openssl.cfg
1、复制一份你安装的openssl的bin目录里面的openssl.cnf文件到你项目所在的目录
2、找到[ CA_default ],打开 copy_extensions = copy (把这句话前面的#去掉)
3、找到[ req ],打开 req_extensions = v3_req #The extensions to add to a certificate request(把这句话前面的#去掉)
4、找到[ v3_req ],在下面添加 subjectAltName = @alt_names
5、添加新的标签 [ alt_names ],和标签字段
DNS.1 = *.kuangstudy.com
*/

/*
1、生成证书私钥test.key
openssl genpkey -algorithm RSA -out test.key

2、通过私钥test.key生成证书请求文件test.csr(注意cfg和cnf)
openssl req -new -nodes -key test.key -out test.csr -days 3650 -subj "/C=cn/OU=myorg/O=mycomp/CN=myname" -config ./openssl.cfg -extensions v3_req
test.csr是上面生成的证书请求文件，ca.crt/server.key是CA证书文件和key，用来对test.csr进程签名认证，这两个文件在第一部分生成

3、生成SAN证书 pem
openssl x509 -req -days 365 -in test.csr -out test.pem -CA server.crt -CAkey server.key -CAcreateserial -extfile ./openssl.cfg -extensions v3_req
*/
