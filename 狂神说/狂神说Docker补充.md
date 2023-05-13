# Docker

### 安装Docker

```dockerfile
# 使用 cat /etc/redhat-release 命令，查看当前 Linux 的发行版本
[root@localhost ~]# cat /etc/redhat-release 
CentOS Linux release 7.0.1406 (Core) 

# Docker 安装要求 CentOS 系统的内核版本高于 3.10，通过 uname -r 命令查看当前的内核版本
[root@localhost ~]# uname -r
3.10.0-123.el7.x86_64

#官网 https://docs.docker.com/engine/install/centos/
#1.卸载旧版本
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine
#2.需要的安装包
yum update -y
yum install -y gcc gcc-c++ vim wget yum-utils device-mapper-persistent-data lvm2
#3.设置镜像的仓库
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo
#上述方法默认是从国外的，不推荐

#推荐使用国内阿里云镜像
yum-config-manager \
    --add-repo \
    https://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo
    
#国内清华镜像
yum-config-manager \
    --add-repo \
    https://mirrors.tuna.tsinghua.edu.cn/docker-ce/linux/centos/docker-ce.repo    
    
#更新yum软件包索引
yum makecache fast

#4.安装docker相关的 docker-ce 社区版 而ee是企业版
yum install -y docker-ce docker-ce-cli containerd.io docker-compose-plugin # 这里我们使用社区版即可

#yum install -y docker-ce-20.10.9 docker-ce-cli-20.10.9 containerd.io docker-compose-plugin

#5.启动docker
systemctl start docker

#6. 使用docker version查看是否按照成功
docker version

#7. 测试
docker run hello-world


#8.查看已经下载的镜像(从这里可以查看已有镜像的id)
[root@iz2zeak7sgj6i7hrb2g862z ~]# docker images
REPOSITORY            TAG                 IMAGE ID            CREATED             SIZE
hello-world           latest              bf756fb1ae65        4 months ago      13.3kB

#9.开机自启
systemctl enable docker

#10.查看已启动的服务
systemctl list-units -lype=service

#11.查看是否设置开机启动
systemctl list-unit-files | grep enable

#12.关闭开机启动
systemctl disable docker.service
```

### 配置阿里云加速

```dockerfile
#1.创建一个目录
sudo mkdir -p /etc/docker

#2.编写配置文件
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://i3hqskj1.mirror.aliyuncs.com"]
}
EOF

#3.重启服务
sudo systemctl daemon-reload
sudo systemctl restart docker
```

### daemon.json配置文件详解

```shell
docker- daemon.json各配置详解
{
    "allow-nondistributable-artifacts": [], #不对外分发的产品提交的registry仓库
    
    “api-cors-header”: "" ,         #在引擎API中设置CORS标头
    
    “authorization-plugins”:[],    #要加载的授权插件
    
    “bridge”: "" ,           #将容器附加到网桥
    
    “cgroup-parent”: "" ,           #为所有容器设置父cgroup
    
    “cluster-store”: "" ,           #分布式存储后端的URL
    
    “cluster-store- opts”:{},       #设置集群存储选项（默认map []）
    
    “cluster-advertise”: "" ,       #要通告的地址或接口名称
    
    “data-root”: " /var/lib/docker " ,           #Docker运行时使用的根路径，默认/var/lib/ docker
    
    “debug”:  true ,           #启用调试模式，启用后，可以看到很多的启动信息。默认false
    
    “default-gateway”: "" ,          #容器默认网关IPv4地址
    
    “default-gateway-v6”: "" ,       #容器默认网关IPv6地址
    
    “default-runtime”:“runc”,       #容器的默认OCI运行时（默认为“ runc”）
    
    “default-ulimits”:{},           #容器的默认ulimit（默认[]）
    
    “dns”: [],       #设定容器DNS的地址，在容器的 /etc/ resolv.conf文件中可查看。
    
    “dns-opts”: [],                  #容器 /etc/ resolv.conf 文件，其他设置
    
    “dns-search”: [],         #设定容器的搜索域，当设定搜索域为 .example.com 时，在搜索一个名为 host 的 主机时，DNS不仅搜索host，还会搜索host.example.com 。 注意：如果不设置， Docker 会默认用主机上的  /etc/ resolv.conf 来配置容器。
    
    “exec-opts”: [],          #运行时执行选项
    
    “exec-root”: "" ,          #执行状态文件的根目录（默认为’/var/run/ docker‘）
    
    “fixed-cidr”: "" ,         #固定IP的IPv4子网
    
    “fixed-cidr-v6”: "" ,      #固定IP的IPv6子网
    
    “group”: “”,           #UNIX套接字的组（默认为“docker”）
    
    "graph":"/var/lib/docker",  #已废弃，使用data-root代替，查看docker版本
    
    “hosts”: [],           #设置容器hosts
    
    “icc”:  false ,        #启用容器间通信（默认为true）
    
    “insecure-registries”: [“120.123.122.123:12312”],           #设置私有仓库地址可以设为http
    
    “ip”:“ 0.0.0.0 ”,    #绑定容器端口时的默认IP（默认0.0.0.0 ）
    
    “iptables”:  false ,   #启用iptables规则添加（默认为true）
    
    “ipv6”:  false ,       #启用IPv6网络
    
    “ip-forward”:  false ,       #默认true, 启用 net.ipv4.ip_forward ,进入容器后使用 sysctl -a |  grepnet.ipv4.ip_forward 查看
    
    “ip-masq”: false ,           #启用IP伪装（默认为true）
    
    “labels”:[“nodeName =node-121 ”],           #docker主机的标签，很实用的功能,例如定义：–label nodeName=host- 121 
    
    “live-restore”:  true ,     #在容器仍在运行时启用docker的实时还原
    
    “log-driver”: "" ,          #容器日志的默认驱动程序（默认为“ json- file ”）
    
    “log-level”: "" ,           #设置日志记录级别（“调试”，“信息”，“警告”，“错误”，“致命”）（默认为“信息”）
    
    “max-concurrent-downloads”: 3,         #设置每个请求的最大并发下载量（默认为3）
    
    “max-concurrent-uploads”: 5,           #设置每次推送的最大同时上传数（默认为5）
    
    “mtu”:  0 ,              #设置容器网络MTU
    
    “oom-score-adjust”: -500 ,          #设置守护程序的oom_score_adj（默认值为- 500 ）
    
    “pidfile”: “”,           #Docker守护进程的PID文件
    
    “raw-logs”:  false ,           #原始日志、全时间戳机制
    
    “registry-mirrors”: [“https://192.168.2.23:89”],   #设置镜像加速地址
    
    “selinux-enabled”:  false ,    #默认  false ，启用selinux支持
    
    “storage-driver”: "" ,         #要使用的存储驱动程序
    
    “swarm-default-advertise-addr”: "" ,          #设置默认地址或群集广告地址的接口
    
    “tls”:  true ,           #默认  false , 启动TLS认证开关
    
    “tlscacert”: “”,         #默认  ~/.docker/ ca.pem，通过CA认证过的的certificate文件路径
    
    “tlscert”: “”,           #默认  ~/.docker/ cert.pem ，TLS的certificate文件路径
    
    “tlskey”: “”,            #默认 ~/.docker/ key.pem，TLS的key文件路径
    
    “tlsverify”:true ,           #默认false，使用TLS并做后台进程与客户端通讯的验证
    
    “userland-proxy”:false ,    #使用userland代理进行环回流量（默认为true）
    
    “userns-remap”: "" ,        #用户名称空间的用户/ 组设置
    
    “bip”:“192.168.88.0/22 ”,          #指定网桥IP
    
    “storage-opts”: {
    “overlay2.override_kernel_check = true ”,
    “overlay2.size = 15G”
    },         #存储驱动程序选项
 "labels":["nodeName=node-121"],        #docker主机的标签
 "live-restore": true,
 "log-driver":"",
 "log-level":"",
 "log-opts": {},
 "max-concurrent-downloads":3,
 "max-concurrent-uploads":5,
 "mtu": 0,
 "oom-score-adjust":-500,
 
“log-opts”: {
    “max-file ”: “ 3 ”,
    “max-size”: “10m”,
    },         #容器默认日志驱动程序选项
    “iptables”:  false          #启用iptables规则添加（默认为true）
} 
```

### 卸载Docker

```dockerfile
#1. 卸载依赖
yum remove docker-ce docker-ce-cli containerd.io docker-compose-plugin
#2. 删除资源
rm -rf /var/lib/docker
rm -rf /var/lib/containerd
# /var/lib/docker 是docker的默认工作路径！
```

# docker容器设置时区

在本地运行项目时，使用的是本地时区一切正常，但部署到服务器容器中，用的是世界时区差8个小时。

解决办法：

### 1、在Dockerfile中设置镜像时区

```dockerfile
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
```

### 2、创建容器时设置时区

将宿主机与容器的时间进行挂载

```bash
-v /etc/localtime:/etc/localtime
```

### 3、进入容器进行设置

1）进入容器内：

```bash
docker exec -it 容器名 /bin/bash
```

2）设定时区

```bash
rm /etc/localtime
ln -s /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
```

# Docker命令

参考：https://haicoder.net/docker/docker-build.html

### 帮助命令

```dockerfile
docker version  # docker版本信息
docker info     # 系统级别的信息，包括镜像和容器的数量
docker 命令 --help 
```

### run命令

```dockerfile
-a stdin: 指定标准输入输出内容类型，可选 STDIN/STDOUT/STDERR 三项；

-d: 后台运行容器，并返回容器ID；

-i: 以交互模式运行容器，通常与 -t 同时使用；

-P: 随机端口映射，容器内部端口随机映射到主机的端口（大写）

-p: 指定端口映射，格式为：主机(宿主)端口:容器端口（小写）

-t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；

--name="nginx-lb": 为容器指定一个名称；

--dns 8.8.8.8: 指定容器使用的DNS服务器，默认和宿主一致；

--dns-search example.com: 指定容器DNS搜索域名，默认和宿主一致；

-h "mars": 指定容器的hostname；

-e username="ritchie": 设置环境变量；

--env-file=[]: 从指定文件读入环境变量；

--cpuset="0-2" or --cpuset="0,1,2": 绑定容器到指定CPU运行；

-m :设置容器使用内存最大值；

--net="bridge": 指定容器的网络连接类型，支持 bridge/host/none/container: 四种类型；

--link=[]: 添加链接到另一个容器；

--expose=[]: 开放一个端口或一组端口；

--volume , -v: 绑定一个卷
```

### 镜像命令

#### docker images

```docker
docker images 查看所有本地主机上的镜像
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker images
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
hello-world         latest              bf756fb1ae65        7 months ago        13.3kB
 
# 解释
REPOSITORY      # 镜像的仓库
TAG             # 镜像的标签
IMAGE ID        # 镜像的ID
CREATED         # 镜像的创建时间
SIZE            # 镜像的大小
 
# 可选项
--all , -a      # 列出所有镜像
--quiet , -q    # 只显示镜像的id
```

| docker images -a         | 列出本地所有的镜像（含中间映像层，默认情况下，过滤掉中间映像层）。 |
| ------------------------ | ------------------------------------------------------------ |
| docker images --digests  | 显示镜像的摘要信息。                                         |
| docker images -f         | 显示满足条件的镜像。                                         |
| docker images --format   | 指定返回值的模板文件。                                       |
| docker images --no-trunc | 显示完整的镜像信息。                                         |
| docker images -q         | 只显示镜像ID。                                               |

#### docker search

```dockerfile
docker search 查找镜像
```

| docker search --filter   | 设置过滤条件。           |
| ------------------------ | ------------------------ |
| docker search --limit    | 最多展示多少条搜索结果。 |
| docker search --no-trunc | 显示完整的镜像描述。     |

#### docker pull

```dockerfile
docker pull 下拉镜像
```

#### docker rmi

```crystal
docker rmi 删除镜像
docker rmi -f	强制删除镜像。
docker rmi --no-prune	不移除该镜像的过程镜像，默认移除。
```

#### docker run

```dockerfile
docker run [可选参数] image #新建容器并启动
 
# 参数说明
--name=“Name”   容器名字    tomcat01    tomcat02    用来区分容器
-d      后台方式运行
-it     使用交互方式运行，进入容器查看内容
-p(小写)      指定容器的端口     -p 8080:8080
    -p  ip:主机端口：容器端口
    -p  主机端口：容器端口（常用）
    -p  容器端口
    容器端口
-P(大写)      随机指定端口
 
# 测试，启动并进入容器
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker run -it centos /bin/bash
[root@74e82b7980e7 /]# ls   # 查看容器内的centos，基础版本，很多命令是不完善的
bin  etc   lib    lost+found  mnt  proc  run   srv  tmp  var
dev  home  lib64  media       opt  root  sbin  sys  usr
 
# 从容器中退回主机
[root@77969f5dcbf9 /]# exit
exit
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# ls
bin   dev  fanfan  lib    lost+found  mnt  proc  run   srv  tmp  var
boot  etc  home    lib64  media       opt  root  sbin  sys  usr
```

#### docker -P

当使用 -P 标记时，Docker 会随机映射一个 49000~49900 的端口到内部容器开放的网络端口。

使用 docker container ls 可以看到，本地主机的 49155 被映射到了容器的 5000 端口。此时访问本机的 49155 端口即可访问容器内 web 应用提供的界面。

```shell
$ docker run -d -P training/webapp python app.py
 
$ docker container ls -l
CONTAINER ID  IMAGE                   COMMAND       CREATED        STATUS        PORTS                    NAMES
bc533791f3f5  training/webapp:latest  python app.py 5 seconds ago  Up 2 seconds  0.0.0.0:49155->5000/tcp  nostalgic_morse
```

#### docker ps

```dockerfile
# docker ps 命令,列出所有的运行的容器
        # 列出当前正在运行的容器
-a      # 列出正在运行的容器包括历史容器
-n=?    # 显示最近创建的容器
-q      # 只显示当前容器的编号
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker ps
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS              PORTS               NAMES
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker ps -a
CONTAINER ID        IMAGE               COMMAND             CREATED             STATUS                     PORTS               NAMES
77969f5dcbf9        centos              "/bin/bash"         5 minutes ago       Exited (0) 5 minutes ago                       xenodochial_bose
74e82b7980e7        centos              "/bin/bash"         16 minutes ago      Exited (0) 6 minutes ago                       silly_cori
a57250395804        bf756fb1ae65        "/hello"            7 hours ago         Exited (0) 7 hours ago                         elated_nash
392d674f4f18        bf756fb1ae65        "/hello"            8 hours ago         Exited (0) 8 hours ago                         distracted_mcnulty
571d1bc0e8e8        bf756fb1ae65        "/hello"            23 hours ago        Exited (0) 23 hours ago                        magical_burnell
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker ps -qa
77969f5dcbf9
74e82b7980e7
a57250395804
392d674f4f18
571d1bc0e8e8
```

#### 退出容器

```dockerfile
exit            # 直接退出容器并关闭
Ctrl + P + Q    # 容器不关闭退出
```

#### 删除容器

```dockerfile
docker rm -f 容器id                  # 删除指定容器
docker rm -f $(docker ps -aq)       # 删除所有容器
docker rmi -f $(docker images -aq)       # 删除所有的镜像
docker ps -a -q|xargs docker rm -f  # 删除所有的容器
```

#### 启动和停止容器

```dockerfile
docker start 容器id           # 启动容器
docker restart 容器id         # 重启容器
docker stop 容器id            # 停止当前正在运行的容器
docker kill 容器id            # 强制停止当前的容器
```

#### 后台启动容器

```dockerfile
# 命令 docker run -d 镜像名
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker run -d centos
 
# 问题 docker ps， 发现centos停止了
 
# 常见的坑， docker 容器使用后台运行， 就必须要有一个前台进程，docker发现没有应用，就会自动停止
# nginx， 容器启动后，发现自己没有提供服务，就会立即停止，就是没有程序了
```

#### 查看日志

```dockerfile
docker logs -tf --tail number 容器id
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker logs -tf --tail 1 8d1621e09bff
2020-08-11T10:53:15.987702897Z [root@8d1621e09bff /]# exit      # 日志输出
 
# 自己编写一段shell脚本
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker run -d centos /bin/sh -c "while true;do echo xiaofan;sleep 1;done"
a0d580a21251da97bc050763cf2d5692a455c228fa2a711c3609872008e654c2
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
a0d580a21251        centos              "/bin/sh -c 'while t…"   3 seconds ago       Up 1 second                             lucid_black
 
# 显示日志
-tf                 # 显示日志
--tail number       # 显示日志条数
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker logs -tf --tail 10 a0d580a21251

docker logs --details	显示 log 信息的额外的详细信息。
docker logs -f, --follow	跟踪日志输出。
docker logs --since	显示某个开始时间的所有日志。
docker logs --tail	仅列出最新 N 条容器日志。
docker logs -t, --timestamps	显示时间戳。
```

#### 查看容器中进程信息

```dockerfile
# 命令 docker top 容器id
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker top df358bc06b17
UID                 PID                 PPID                C                   STIME               TTY     
root                28498               28482               0                   19:38               ?      
```

#### 查看镜像的元数据

```dockerfile
# 命令
docker inspect 容器id
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker inspect df358bc06b17
[
    {
        "Id": "df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3",
        "Created": "2020-08-11T11:38:34.935048603Z",
        "Path": "/bin/bash",
        "Args": [],
        "State": {
            "Status": "running",
            "Running": true,
            "Paused": false,
            "Restarting": false,
            "OOMKilled": false,
            "Dead": false,
            "Pid": 28498,
            "ExitCode": 0,
            "Error": "",
            "StartedAt": "2020-08-11T11:38:35.216616071Z",
            "FinishedAt": "0001-01-01T00:00:00Z"
        },
        "Image": "sha256:0d120b6ccaa8c5e149176798b3501d4dd1885f961922497cd0abef155c869566",
        "ResolvConfPath": "/var/lib/docker/containers/df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3/resolv.conf",
        "HostnamePath": "/var/lib/docker/containers/df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3/hostname",
        "HostsPath": "/var/lib/docker/containers/df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3/hosts",
        "LogPath": "/var/lib/docker/containers/df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3/df358bc06b17ef44f215d35d9f46336b28981853069a3739edfc6bd400f99bf3-json.log",
        "Name": "/hungry_heisenberg",
        "RestartCount": 0,
        "Driver": "overlay2",
        "Platform": "linux",
        "MountLabel": "",
        "ProcessLabel": "",
        "AppArmorProfile": "",
        "ExecIDs": null,
        "HostConfig": {
            "Binds": null,
            "ContainerIDFile": "",
            "LogConfig": {
                "Type": "json-file",
                "Config": {}
            },
            "NetworkMode": "default",
            "PortBindings": {},
            "RestartPolicy": {
                "Name": "no",
                "MaximumRetryCount": 0
            },
            "AutoRemove": false,
            "VolumeDriver": "",
            "VolumesFrom": null,
            "CapAdd": null,
            "CapDrop": null,
            "Capabilities": null,
            "Dns": [],
            "DnsOptions": [],
            "DnsSearch": [],
            "ExtraHosts": null,
            "GroupAdd": null,
            "IpcMode": "private",
            "Cgroup": "",
            "Links": null,
            "OomScoreAdj": 0,
            "PidMode": "",
            "Privileged": false,
            "PublishAllPorts": false,
            "ReadonlyRootfs": false,
            "SecurityOpt": null,
            "UTSMode": "",
            "UsernsMode": "",
            "ShmSize": 67108864,
            "Runtime": "runc",
            "ConsoleSize": [
                0,
                0
            ],
            "Isolation": "",
            "CpuShares": 0,
            "Memory": 0,
            "NanoCpus": 0,
            "CgroupParent": "",
            "BlkioWeight": 0,
            "BlkioWeightDevice": [],
            "BlkioDeviceReadBps": null,
            "BlkioDeviceWriteBps": null,
            "BlkioDeviceReadIOps": null,
            "BlkioDeviceWriteIOps": null,
            "CpuPeriod": 0,
            "CpuQuota": 0,
            "CpuRealtimePeriod": 0,
            "CpuRealtimeRuntime": 0,
            "CpusetCpus": "",
            "CpusetMems": "",
            "Devices": [],
            "DeviceCgroupRules": null,
            "DeviceRequests": null,
            "KernelMemory": 0,
            "KernelMemoryTCP": 0,
            "MemoryReservation": 0,
            "MemorySwap": 0,
            "MemorySwappiness": null,
            "OomKillDisable": false,
            "PidsLimit": null,
            "Ulimits": null,
            "CpuCount": 0,
            "CpuPercent": 0,
            "IOMaximumIOps": 0,
            "IOMaximumBandwidth": 0,
            "MaskedPaths": [
                "/proc/asound",
                "/proc/acpi",
                "/proc/kcore",
                "/proc/keys",
                "/proc/latency_stats",
                "/proc/timer_list",
                "/proc/timer_stats",
                "/proc/sched_debug",
                "/proc/scsi",
                "/sys/firmware"
            ],
            "ReadonlyPaths": [
                "/proc/bus",
                "/proc/fs",
                "/proc/irq",
                "/proc/sys",
                "/proc/sysrq-trigger"
            ]
        },
        "GraphDriver": {
            "Data": {
                "LowerDir": "/var/lib/docker/overlay2/5af8a2aadbdba9e1e066331ff4bce56398617710a22ef906f9ce4d58bde2d360-init/diff:/var/lib/docker/overlay2/62926d498bd9d1a6684bb2f9920fb77a2f88896098e66ef93c4b74fcb19f29b6/diff",
                "MergedDir": "/var/lib/docker/overlay2/5af8a2aadbdba9e1e066331ff4bce56398617710a22ef906f9ce4d58bde2d360/merged",
                "UpperDir": "/var/lib/docker/overlay2/5af8a2aadbdba9e1e066331ff4bce56398617710a22ef906f9ce4d58bde2d360/diff",
                "WorkDir": "/var/lib/docker/overlay2/5af8a2aadbdba9e1e066331ff4bce56398617710a22ef906f9ce4d58bde2d360/work"
            },
            "Name": "overlay2"
        },
        "Mounts": [],
        "Config": {
            "Hostname": "df358bc06b17",
            "Domainname": "",
            "User": "",
            "AttachStdin": true,
            "AttachStdout": true,
            "AttachStderr": true,
            "Tty": true,
            "OpenStdin": true,
            "StdinOnce": true,
            "Env": [
                "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin"
            ],
            "Cmd": [
                "/bin/bash"
            ],
            "Image": "centos",
            "Volumes": null,
            "WorkingDir": "",
            "Entrypoint": null,
            "OnBuild": null,
            "Labels": {
                "org.label-schema.build-date": "20200809",
                "org.label-schema.license": "GPLv2",
                "org.label-schema.name": "CentOS Base Image",
                "org.label-schema.schema-version": "1.0",
                "org.label-schema.vendor": "CentOS"
            }
        },
        "NetworkSettings": {
            "Bridge": "",
            "SandboxID": "4822f9ac2058e8415ebefbfa73f05424fe20cc8280a5720ad3708fa6e80cdb08",
            "HairpinMode": false,
            "LinkLocalIPv6Address": "",
            "LinkLocalIPv6PrefixLen": 0,
            "Ports": {},
            "SandboxKey": "/var/run/docker/netns/4822f9ac2058",
            "SecondaryIPAddresses": null,
            "SecondaryIPv6Addresses": null,
            "EndpointID": "5fd269c0a28227241e40cd30658e3ffe8ad6cc3e6514917c867d89d36a31d605",
            "Gateway": "172.17.0.1",
            "GlobalIPv6Address": "",
            "GlobalIPv6PrefixLen": 0,
            "IPAddress": "172.17.0.2",
            "IPPrefixLen": 16,
            "IPv6Gateway": "",
            "MacAddress": "02:42:ac:11:00:02",
            "Networks": {
                "bridge": {
                    "IPAMConfig": null,
                    "Links": null,
                    "Aliases": null,
                    "NetworkID": "30d6017888627cb565618b1639fecf8fc97e1ae4df5a9fd5ddb046d8fb02b565",
                    "EndpointID": "5fd269c0a28227241e40cd30658e3ffe8ad6cc3e6514917c867d89d36a31d605",
                    "Gateway": "172.17.0.1",
                    "IPAddress": "172.17.0.2",
                    "IPPrefixLen": 16,
                    "IPv6Gateway": "",
                    "GlobalIPv6Address": "",
                    "GlobalIPv6PrefixLen": 0,
                    "MacAddress": "02:42:ac:11:00:02",
                    "DriverOpts": null
                }
            }
        }
    }
]
```

#### 进入当前正在运行的容器

```dockerfile
# 我们通常容器使用后台方式运行的， 需要进入容器，修改一些配置
 
# 命令
docker exec -it 容器id /bin/bash
 
# 测试
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker exec -it df358bc06b17 /bin/bash
[root@df358bc06b17 /]# ls       
bin  etc   lib    lost+found  mnt  proc  run   srv  tmp  var
dev  home  lib64  media       opt  root  sbin  sys  usr
[root@df358bc06b17 /]# ps -ef
UID        PID  PPID  C STIME TTY          TIME CMD
root         1     0  0 Aug11 pts/0    00:00:00 /bin/bash
root        29     0  0 01:06 pts/1    00:00:00 /bin/bash
root        43    29  0 01:06 pts/1    00:00:00 ps -ef
 
# 方式二
docker attach 容器id
 
# docker exec       # 进入容器后开启一个新的终端，可以在里面操作
# docker attach     # 进入容器正在执行的终端，不会启动新的进程
```

#### 从容器中拷贝文件到主机

```dockerfile
docker cp 容器id：容器内路径 目的地主机路径
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ /]# docker cp 7af535f807e0:/home/Test.java /home
```

#### docker create

```dockerfile
docker create ubuntu
# docker create 创建的容器并未实际启动，还需要执行 docker start 命令或 docker run 命令以启动容器
```

#### docker exec

| –detach, -d      | 后台运行模式，在后台执行命令相关命令。 |
| ---------------- | -------------------------------------- |
| –detach-keys     | 覆盖容器后台运行的一些参数信息。       |
| –env, -e         | 设置环境变量。                         |
| –interactive, -i | 展示容器输入信息 STDIN。               |
| –privileged      | 为命令提供一些扩展权限。               |
| –tty, -t         | 命令行交互模式。                       |
| –user, -u        | 设置用户名。                           |

#### docker pause 

```dockerfile
docker pause 容器Id/容器名 #用来暂停 Docker容器 中所有的进程
```

#### docker unpause

```dockerfile
docker unpause 容器Id/容器名 #用来恢复 Docker容器 中所有的进程
```

#### docker top

```dockerfile
docker top 容器Id/容器名 #用来查看 Docker 中运行的进程信息
```

#### docker wait

```dockerfile
docker wait 容器Id/容器名 #用于阻塞一个或多个 Docker容器 直到容器停止，然后打印退出代码
```

#### docker export

```dockerfile
docker export 容器Id/容器名 #用于将 Docker容器 里的文件系统作为一个 tar 归档文件导出到标准输出
```

#### docker port

```dockerfile
docker port 容器Id/容器名 #用于列出指定的 Docker容器 的端口映射，或者将容器里的端口映射到宿主机。
```

#### docker rename

```dockerfile
docker rename 容器Id/容器名 #用于重命名一个 Docker容器
```

#### docker stats

```dockerfile
docker stats 容器Id/容器名 #用于动态显示 Docker容器 的资源消耗情况，包括：CPU、内存、网络I/O
```

| –all, -a   | 查看所有容器信息(默认显示运行中的)。 |
| ---------- | ------------------------------------ |
| –format    | Go模板展示镜像信息。                 |
| –no-stream | 不展示容器的一些动态信息。           |

#### docker update

```dockerfile
docker update 容器Id/容器名 #用于更新一个或多个 Docker容器 的配置
```

| –cpu-shares    | 更新 cpu-shares。  |
| -------------- | ------------------ |
| –kernel-memory | 更新内核内存限制。 |
| –memory        | 更新内存限制。     |
| –restart       | 更新重启策略。     |

#### docker diff

```dockerfile
docker diff 容器Id/容器名 #用于比较一个 Docker容器 不同版本提交的文件差异
```

#### docker login

```dockerfile
docker login #用于登陆到一个 Docker 镜像仓库
docker login -u username
docker login -u username -p password
```

| docker login -u               | 登陆的用户名。       |
| ----------------------------- | -------------------- |
| docker login -p               | 登陆的密码。         |
| docker login --password-stdin | 从标准输入中读取密码 |

#### docker logout

```dockerfile
docker logout #用于登出一个Docker镜像仓库
docker logout
```

#### docker tag

```dockerfile
docker tag #用于给镜像打标签
docker tag diytomcat tomcat:1.0
```

#### docker load

```dockerfile
docker load #用于从 tar 归档文件或者标准输入流载入镜像
docker load -i	指定导出的文件。
docker load -q	精简输出信息

#docker save -o centos.tar 67fa590cfc1c
#docker load -i centos.tar
```

#### docker save

```dockerfile
docker save #将 Docker镜像 保存成 tar 包。docker save 命令的相对应的命令为 docker load

docker save -o, --output	将归档文件输出到的文件
docker save -o centos.tar 67fa590cfc1c
docker save -o nginx-alpine.tar nginx:1.7 #导出镜像到文件中
```

#### docker import

```dockerfile
docker import #用于从归档文件中创建镜像。docker import 命令的相对应的命令为 docker export

docker import -c, --change	应用 docker 指令创建镜像。
docker import -m, --message	提交时的说明文字。

#举例
docker export -o haicoder.tar haicoder
docker import haicoder.tar haicoder_centos
```

#### docker events

```dockerfile
docker events #用于打印出实时的系统事件

docker events -f	根据条件过滤事件。
docker events --since	从指定的时间戳后显示所有事件。
docker events --until	流水时间显示到指定的时间为止。
#docker events -f "image"="centos"  --since="1564588800"
```

#### docker history

```dockerfile
docker history #用于打印出指定的 Docker镜像 的历史版本信息

docker history -H, --human	以可读的格式打印镜像大小和日期，默认为 true。
docker history --no-trunc	显示完整的提交记录。
docker history -q, --quiet	仅列出提交记录 ID。
#docker history -q centos
```



# 安装案例

### 安装Nginx

```dockerfile
docker pull nginx
docker run -d --name nginx01 -p 3344:80 nginx #后台启动
```

### 安装Tomcat

```dockerfile
docker pull tomcat
docker run -it --rm -p 8888:8080 tomcat #用完即删，一般用来测试
docker run -d --name tomcat01 -p 8888:8080 tomcat #后台启动
docker exec -it tomcat01 /bin/bash
#发现问题：1、linux命令减少了；2、没有webapps，阿里云镜像的原因，默认是最小的镜像，所以不必要的都剔除掉了，保证了最小可运行的环境
cp -r webapps.dist/* webapps #如果是目录，不能直接复制，要加上 -r参数
```

### 安装Elasticsearch

```dockerfile
#es暴露的端口很多，十分的耗内存，一般需要挂载目录
#--net somenetwork 网络配置
docker pull elasticsearch:8.5.3
docker run -d --name elasticsearch01 --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:8.5.3

#修改配置文件
docker run -d --name elasticsearch02 --net somenetwork -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:8.5.3
```

### 安装mysql

```shell
docker run -p 3306:3306 --name mysql -v /usr/mydata/mysql/log:/var/log/mysql -v /usr/mydata/mysql/data:/var/lib/mysql -v /usr/mydata/mysql/conf:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=123456 -d mysql:latest

分析：
    docker run -d mysql:latest             以后台的方式运行 mysql 版本的镜像，生成一个容器。
    --name mysql                           容器名为 mysql
    -e MYSQL_ROOT_PASSWORD=123456          设置登陆密码为 123456，登陆用户为 root
    -p 3306:3306                           将容器内部 3306 端口映射到 主机的 3306 端口，即通过 主机的 3306 可以访问容器的 3306 端口
    -v /usr/mydata/mysql/log:/var/log/mysql    将容器的 日志文件夹 挂载到 主机的相应位置
    -v /usr/mydata/mysql/data:/var/lib/mysql   将容器的 数据文件夹 挂载到 主机的相应位置
    -v /usr/mydata/mysql/conf:/etc/mysql/conf.d   将容器的 自定义配置文件夹 挂载到主机的相应位置
```

```shell
docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=12345678 -d mysql:8.0.33
```

# 可视化面板

### 安装Portainer

Portainer 是一个轻量级的管理 UI ，可让你轻松管理不同的 Docker 环境（Docker 主机或 Swarm 群集）

```dockerfile
docker pull portainer/portainer-ce
#官方用法
docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data --privileged=true portainer/portainer-ce:latest

https://192.168.222.100:9443/ #访问地址
admin1234567 #设置密码
```

### 推荐使用Rancher(CI/CD集群再用)

Rancher是一个开源的企业级容器管理平台。通过Rancher，企业再也不必自己使用一系列的开源软件去从头搭建容器服务平台。Rancher提供了在生产环境中使用的管理Docker和Kubernetes的全栈化容器部署与管理平台

```dockerfile
docker pull rancher/rancher
docker inspect rancher/rancher:latest #查看rancher镜像详细信息
mkdir -p /docker_volume/rancher_home/rancher
mkdir -p /docker_volume/rancher_home/auditlog
docker run -d --privileged --restart=unless-stopped -p 8080:8080 -p 443:443 --name rancher-server -e CATTLE_DB_CATTLE_MYSQL_HOST=192.168.222.100 -e CATTLE_DB_CATTLE_MYSQL_PORT=3306 -e CATTLE_DB_CATTLE_MYSQL_NAME=cattle -e CATTLE_DB_CATTLE_USERNAME=cattle -e CATTLE_DB_CATTLE_PASSWORD=Cattle@123 -v /docker_volume/rancher_home/rancher:/var/lib/rancher -v /docker_volume/rancher_home/auditlog:/var/log/auditlog rancher/rancher
#运行 docker 命令启动容器管理平台应用 rancher，这里假定容器管理数据库的 IP 地址为 192.168.222.100，端口号为 3306，数据库为cattle，用户名为 cattle，密码为 Cattle123

#简单测试版
docker run -d --restart=unless-stopped -p 8080:8080 -p 443:443 --name rancher-server rancher/rancher
#测试启动失败，后续在学习一下
```

### 删除启动的镜像

```
docker rm 镜像ID
```

# commit镜像

```dockerfile
docker commit #提交容器成为一个新的副本
docker commit -m="提交的描述信息" -a="作者" 容器ID 目标镜像名:tag #命令和git类似

#举例
docker commit -a="kuangshen" -m="add webapps app" eb664817b34f tomcaty:1.0
```

# docker 数据管理

数据卷是一个可供一个或多个容器使用的特殊目录，它绕过 UFS，可以提供很多有用的特性：

        1）数据卷可以在容器之间共享和重用；
    
        2）对数据卷的修改会立马生效；
    
        3）对数据卷的更新，不会影响镜像；
    
        4）数据卷 默认会一直存在，即使容器被删除。
注意：数据卷 的使用，类似于 Linux 下对目录或文件进行 mount，镜像中的被指定为挂载点的目录中的文件会隐藏掉，能显示看的是挂载的 数据卷。

### 1.1、创建一个数据卷并查看( docker volume create name)

```shell
$ docker volume create my-vol
```

查看所有的数据卷(docker volume ls)

```shell
$ docker volume ls 
local my-vol
```

查看指定数据卷信息(docker volume [inspect](https://so.csdn.net/so/search?q=inspect&spm=1001.2101.3001.7020) my-vol)：

```shell
$ docker volume inspect my-vol
[
    {
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/my-vol/_data",
        "Name": "my-vol",
        "Options": {},
        "Scope": "local"
    }
]
```

### 1.2、启动一个挂载数据卷的容器(docker run --mount source=,target=)

​        在用 docker run 命令的时候，使用 --mount 标记来将 数据卷 挂载到容器里。在一次 docker run 中可以挂载多个 数据卷。--mount source=my-vol,target=/webapp

```shell
#下面创建一个名为 web 的容器，并加载一 个 数据卷 到容器的 /webapp 目录。
$ docker run -d -P \
    --name web \
    # -v my-vol:/wepapp \
    --mount source=my-vol,target=/webapp \
    training/webapp \
    python app.py
```

### 1.3、查看数据卷的具体信息(docker inspect name)

​    在主机里使用以下命令可以查看 web 容器的信息。

```shell
$ docker inspect web
```

数据卷 信息在 "Mounts" Key 下面。

```shell
"Mounts": [
    {
        "Type": "volume",
        "Name": "my-vol",
        "Source": "/var/lib/docker/volumes/my-vol/_data",
        "Destination": "/app",
        "Driver": "local",
        "Mode": "",
        "RW": true,
        "Propagation": ""
    }
]
```

### 1.4、删除数据卷(docker volume rm name/docker volum prune)

```shell
$ docker volume rm my-vol
```

数据卷 是被设计用来持久化数据的，它的生命周期独立于容器，Docker 不会在容器被删除后自动删除 数据卷，并且也不存在垃圾回收这样的机制来处理没有任何容器引用的 数据卷。如果需要在删除容器的同时移除数据卷。可以在删除容器的时候使用 docker rm -v 这个命令。

```shell
无主的数据卷可能会占据很多空间，要清理请使用以下命令：
$ docker volume prune
```
### 1.5、挂载一个主机目录作为数据卷(--mount type=bind,source=,target=)

​        --mount type=bind,source=/src/webapp,target=/opt/webapp。用 --mount 标记可以指定挂载一个本地主机的目录到容器中去。
```shell
$ docker run -d -P \
    --name web \
    # -v /src/webapp:/opt/webapp \
    --mount type=bind,source=/src/webapp,target=/opt/webapp \
    training/webapp \
    python app.py
```

上面的命令加载主机的 /src/webapp 目录到容器的 /opt/webapp目录。这个功能在进行测试的时候十分方便，比如用户可以放置一些程序到本地目录中，来查看容器是否正常工作。本地目录的路径必须是绝对路径，以前使用 -v 参数时如果本地目录不存在 Docker 会自动为你创建一个文件夹，现在使用 --mount 参数时如果本地目录不存在，Docker 会报错。

Docker 挂载主机目录的默认权限是 读写，用户也可以通过增加 readonly 指定为只读。

```shell
$ docker run -d -P \
    --name web \
    # -v /src/webapp:/opt/webapp:ro \
    --mount type=bind,source=/src/webapp,target=/opt/webapp,readonly \
    training/webapp \
    python app.py
```

 加了 readonly 之后，就挂载为 只读 了。如果你在容器内 /opt/webapp 目录新建文件，会显示如下错误：

```shell
/opt/webapp # touch new.txt
touch: new.txt: Read-only file system
```

### 1.6、查看数据卷的具体信息

​    在主机里使用以下命令可以查看 web 容器的信息：

```shell
$ docker inspect web
```

​    挂载主机目录 的配置信息在 "Mounts" Key 下面：

```shell
"Mounts": [
    {
        "Type": "bind",
        "Source": "/src/webapp",
        "Destination": "/opt/webapp",
        "Mode": "",
        "RW": true,
        "Propagation": "rprivate"
    }
]
```

### 1.7、挂载一个本地主机文件作为数据卷

​    --mount 标记也可以从主机挂载单个文件到容器中，

​    --mount type=bind,source=$HOME/.bash_history,target=/root/.bash_history \：

```shell
$ docker run --rm -it \
   # -v $HOME/.bash_history:/root/.bash_history \
   --mount type=bind,source=$HOME/.bash_history,target=/root/.bash_history \
   ubuntu:18.04 \
   bash
 
root@2affd44b4667:/# history
1  ls
2  diskutil list
```

这样就可以记录在容器输入过的命令了

# 容器数据卷

```dockerfile
docker run -it -p 主机端口:容器内端口 -v 主机目录:容器内目录 -e 环境配置 -P(大写)随机指定端口
#容器目录不可以为相对路径
#宿主机目录如果不存在，则会自动生成

#举例
docker run -d mysql:latest             以后台的方式运行 mysql 版本的镜像，生成一个容器。
    --name mysql                           容器名为 mysql
    -e MYSQL_ROOT_PASSWORD=123456          设置登陆密码为 123456，登陆用户为 root
    -p 3306:3306                           将容器内部 3306 端口映射到 主机的 3306 端口，即通过 主机的 3306 可以访问容器的 3306 端口
    -v /usr/mydata/mysql/log:/var/log/mysql    将容器的 日志文件夹 挂载到 主机的相应位置
    -v /usr/mydata/mysql/data:/var/lib/mysql   将容器的 数据文件夹 挂载到 主机的相应位置
    -v /usr/mydata/mysql/conf:/etc/mysql/conf.d   将容器的 自定义配置文件夹 挂载到主机的相应位置
```

###  匿名目录挂载

```dockerfile
docker volume --help #查看其他的命令
```

```dockerfile
#匿名目录挂载只需要写容器内目录或者文件即可，而宿主机对应的目录会在/var/lib/docker/volumes路径下生成：

//以交互模式运行容器，并使用-v 匿名挂载容器数据卷
docker run -it -v 容器内目录/文件的绝对路径[:rw/ro] -p 主机端口:容器端口 --name=容器名称 镜像ID/镜像名称[:版本号]

//以后台方式运行容器，并使用-v 匿名挂载容器数据卷 (推荐)
docker run -d -v 容器内目录/文件的绝对路径[:rw/ro] -p 主机端口:容器端口 --name=容器名称 镜像ID/镜像名称[:版本号]

//注意：如果出现Docker挂载宿主机目录显示cannot open directory .:Permission denied
解决办法：在挂载目录后面 多加一个--privileged=true参数即可
```

### 具名目录挂载

```dockerfile
#具名目录挂载相对于匿名目录挂载，就是在宿主机生成对应的目录时可以指定该目录的名称，同样目录也会在/var/lib/docker/volumes路径下生成

//以交互模式运行容器，并使用-v 具名挂载容器数据卷
docker run -it -v 目录名称:容器内目录/文件的绝对路径[:rw/ro] -p 主机端口:容器端口 --name=容器名称 镜像ID/镜像名称[:版本号]

//以后台方式运行容器，并使用-v 具名挂载容器数据卷 (推荐)
docker run -d -v 目录名称:容器内目录/文件的绝对路径[:rw/ro] -p 主机端口:容器端口 --name=容器名称 镜像ID/镜像名称[:版本号]

//注意：如果出现Docker挂载宿主机目录显示cannot open directory .:Permission denied
解决办法：在挂载目录后面 多加一个--privileged=true参数即可
```

```docker
#如何确定是匿名挂载还是具名挂载，还是指定路径挂载
-v 容器内路径 #匿名挂载
-v 卷名:容器内路径 #具名挂载
-v /宿主机路径:容器内路径 #指定路径挂载
```

拓展

```dockerfile
#通过 -v 容器内路径:ro/rw 改变读写权限
#ro 只读 只能通过宿主机来操作，容器内部是无法操作的
#rw 可读可写
docker run -d --name nginx01 -p 3344:80 -v juming-nginx:/etc/nginx:ro/rw nginx
```

# 安装vim

```dockerfile
#查看一下你本机已经存在的包，确认一下你的VIM是否已经安装
rpm -qa|grep vim
```

输出结果如下，如无以下输出结果

![img](https://img2020.cnblogs.com/blog/1501477/202005/1501477-20200503112705676-836170760.png)

如果缺少了其中某个，比如说： vim-enhanced这个包少了，则执行

```linux
yum -y install vim-enhanced
```

它会自动下载安装。如果上面三个包一个都没有显示，则直接输入命令：

```linux
yum -y install vim*
```

安装完成后开始配置vim

```linux
vim /etc/vimrc
```

打开文件后，按 **i** 进入编辑模式，在最后 添加如下代码， 添加好了之后，按Esc，然后输入 ：qw 退出并保存即可。

```dockerfile
set nu          # 设置显示行号
set showmode    # 设置在命令行界面最下面显示当前模式等
set ruler       # 在右下角显示光标所在的行数等信息
set autoindent  # 设置每次单击Enter键后，光标移动到下一行时与上一行的起始字符对齐
syntax on       # 即设置语法检测，当编辑C或者Shell脚本时，关键字会用特殊颜色显示
```

![img](https://img2020.cnblogs.com/blog/1501477/202005/1501477-20200503145955224-765100808.png)

# DockerFile

### DockerFile的构建过程

dockerfile就是用来构建docker镜像的构建文件，命令参数脚本

**基础知识：**

1. 每个保留关键字（指令）都是必须大写字母
2. 执行从上到下顺序执行
3. `#` 表示注释
4. 每个指令都会创建提交一个新的镜像层，并提交！

```dockerfile
FROM #基础镜像，一切从这里开始构建
MAINTAINER #镜像是谁写的，名字+邮箱
RUN #镜像构建的时候被需要运行的命令 run是在docker build时运行
ADD #步骤，tomcat镜像，这个tomcat压缩包，添加内容
WORKDIR #镜像的挂载目录
VOLUME #挂载的目录
EXPOST #保留端口配置
CMD #指定这个容器启动的时候要运行的命令，只有之后一个会生效，可被替代 CMD是在docker run时运行
ENTRYPOINT #指定这个容器启动的时候要运行的命令，可以追加命令
COPY #类似ADD，将我们文件拷贝到镜像中
ENV #构建的时候设置环境变量
ONBUILD #当构建一个被继承DockerFile，这个时候就会运行ONBUILD的指令，触发指令
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20210316230210380.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2Nsb3ZlcjY2MQ==,size_16,color_FFFFFF,t_70)

##### FROM

功能为指定基础镜像，并且必须是第一条指令。

如果不以任何镜像为基础，那么写法为：FROM scratch。

同时意味着接下来所写的指令将作为镜像的第一层开始
语法：

```shell
FROM <image>
FROM <image>:<tag>
FROM <image>:<digest> 
三种写法，其中<tag>和<digest> 是可选项，如果没有选择，那么默认值为latest
```

##### MAINTAINER

指定作者

语法：

```shell
MAINTAINER <name>
```

- 新版docker中使用LABEL指明

##### LABEL

功能是为镜像指定标签

语法：

```shell
LABEL <key>=<value> <key>=<value> <key>=<value> ...
一个Dockerfile种可以有多个LABEL，如下：

LABEL "com.example.vendor"="ACME Incorporated"
LABEL com.example.label-with-value="foo"
LABEL version="1.0"
LABEL description="This text illustrates \
that label-values can span multiple lines."
但是并不建议这样写，最好就写成一行，如太长需要换行的话则使用\符号

如下：
LABEL multi.label1="value1" \
multi.label2="value2" \
other="value3"

```

说明：LABEL会继承基础镜像种的LABEL，如遇到key相同，则值覆盖

##### ADD

一个复制命令，把文件复制到镜像中。
如果把虚拟机与容器想象成两台linux服务器的话，那么这个命令就类似于scp，只是scp需要加用户名和密码的权限验证，而ADD不用。
语法如下：

```shell
1. ADD <src>... <dest>
2. ADD ["<src>",... "<dest>"]
```

- 路径的填写可以是容器内的绝对路径，也可以是相对于工作目录的相对路径，推荐写成绝对路径
- 可以是一个本地文件或者是一个本地压缩文件，还可以是一个url
- 如果把写成一个url，那么ADD就类似于wget命令
  示例

```shell
ADD test relativeDir/ 
ADD test /relativeDir
ADD http://example.com/foobar /
ADD centos-7-x86_64-docker.tar.xz /
```

注意事项

- src为一个目录的时候，会自动把目录下的文件复制过去，目录本身不会复制
- 如果src为多个文件，dest一定要是一个目录

##### COPY

看这个名字就知道，又是一个复制命令

语法如下：

```shell
COPY <src>... <dest>
COPY ["<src>",... "<dest>"]
```

与ADD的区别

- COPY的只能是本地文件，其他用法一致

##### EXPOSE

功能为暴漏容器运行时的监听端口给外部

但是EXPOSE并不会使容器访问主机的端口

如果想使得容器与主机的端口有映射关系，必须在容器启动的时候加上 -P参数
语法：

```shell
EXPOSE <port>/<tcp/udp>
```

##### ENV

功能为设置环境变量

语法有两种

```shell
ENV <key> <value>
ENV <key>=<value> ...
```

两者的区别就是第一种是一次设置一个，第二种是一次设置多个

在Dockerfile中使用变量的方式
$varname
${varname}
${varname:-default value}
$(varname:+default value}
第一种和第二种相同
第三种表示当变量不存在使用-号后面的值
第四种表示当变量存在时使用+号后面的值（当然不存在也是使用后面的值）

##### VOLUME

可实现挂载功能，可以将宿主机目录挂载到容器中

说的这里大家都懂了，可用专用的文件存储当作Docker容器的数据存储部分

语法如下：

```shell
VOLUME ["/data"]
```

说明：

[“/data”]可以是一个JsonArray ，也可以是多个值。所以如下几种写法都是正确的

```shell
VOLUME ["/var/log/"]
VOLUME /var/log
VOLUME /var/log /var/db
```

一般的使用场景为需要持久化存储数据时

容器使用的是AUFS，这种文件系统不能持久化数据，当容器关闭后，所有的更改都会丢失。

所以当数据需要持久化时用这个命令。

##### USER

设置启动容器的用户，可以是用户名或UID，所以，只有下面的两种写法是正确的

```shell
USER daemo
USER UID
```

注意：如果设置了容器以daemon用户去运行，那么RUN, CMD 和 ENTRYPOINT 都会以这个用户去运行,
使用这个命令一定要确认容器中拥有这个用户，并且拥有足够权限

##### WORKDIR

语法：

```shell
WORKDIR /path/to/workdir
```

设置工作目录，对RUN,CMD,ENTRYPOINT,COPY,ADD生效。如果不存在则会创建，也可以设置多次。

如：

```shell
WORKDIR /a
WORKDIR b
WORKDIR c
RUN pwd
```

pwd执行的结果是/a/b/c

WORKDIR也可以解析环境变量

如：

```shell
ENV DIRPATH /path
WORKDIR $DIRPATH/$DIRNAME
RUN pwd
```

pwd的执行结果是/path/$DIRNAME

##### 狂神实例1

```dockerfile
[root@node143 dockerfile]# vim mydockerfile

FROM centos

ADD centos-7-x86_64-docker.tar.xz /

MAINTAINER xiaoshimei<12345645@qq.com>

ENV MYPATH /usr/locat
WORKDIR $MYPATH

RUN yum -y install vim
RUN yum -y install net-tools

EXPOSE 80

CMD echo $MYPATH
CMD echo"----end----"
CMD /bin/bash  
```

### 通过这个文件构建镜像

```dockerfile
#命令 docker build -f dockerfile文件路径 -t 镜像名：[tag]版本号，同一个目录下不用-f
docker build -f mydockerfile -t mydockerfile:0.1 .

#构建成功最底部会显示这两行
Successfully built 29f50f45bc0b
Successfully tagged mydockerfile:0.1
```

### 案例，构建自己的centos

#### Build命令

```dockerfile
docker build  -t ImageName:TagName dir
    -t 给镜像加一个Tag
    ImageName 给镜像起的名称
    TagName 给镜像的Tag名
    Dir Dockerfile所在目录
```

#### 构建自己的dockerfile

```dockerfile
# 1. 编写Dockerfile的文件
# mkdir dockerf
# cd dockerf
# cat /etc/redhat-release 查看centos的版本号
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# vim mydockerfile-centos
FROM centos:centos7.9.2009
MAINTAINER xiaofan<594042358@qq.com>
 
ENV MYPATH /usr/local
WORKDIR $MYPATH     # 镜像的工作目录
 
RUN yum -y install vim
RUN yum -y install net-tools
RUN yum -y install glibc.i686

RUN yum -y install wget \
    && wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
    && tar -xvf redis-5.0.3.tar.gz #以 && 符号连接命令，这样执行后，只会创建 1 层镜像
 
EXPOSE 80
 
CMD echo $MYPATH
CMD echo "---end---"
CMD /bin/bash
 
# 2. 通过这个文件构建镜像
# 命令 docker build -f dockerfile文件路径 -t 镜像名:[tag] .
 
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker build -f mydockerfile-centos -t mycentos:0.1 .
 
Successfully built d2d9f0ea8cb2
Successfully tagged mycentos:0.1
```

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813152240210.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)



#### CMD 和ENTRYPOINT区别

```dockerfile
CMD         # 指定这个容器启动的时候要运行的命令，只有最后一个会生效可被替代
ENTRYPOINT  # 指定这个容器启动的时候要运行的命令， 可以追加命令
```

测试CMD

```docker
# 1. 编写dockerfile文件
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# vim dockerfile-cmd-test 
FROM centos
CMD ["ls", "-a"]
 
# 2. 构建镜像
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker build -f dockerfile-cmd-test -t cmdtest .
 
# 3. run运行， 发现我们的ls -a 命令生效
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker run ebe6a52bb125
.
..
.dockerenv
bin
dev
etc
home
lib
lib64
 
# 想追加一个命令 -l 变成 ls -al
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker run ebe6a52bb125 -l
docker: Error response from daemon: OCI runtime create failed: container_linux.go:349: starting container process caused "exec: \"-l\": executable file not found in $PATH": unknown.
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker run ebe6a52bb125 ls -al
 
# cmd的情况下 -l替换了CMD["ls", "-a"]命令， -l不是命令，所以报错了
```

测试ENTRYPOINT

```dockerfile
# 1. 编写dockerfile文件
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# vim dockerfile-entrypoint-test 
FROM centos
ENTRYPOINT ["ls", "-a"]
 
# 2. 构建文件
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker build -f dockerfile-entrypoint-test -t entrypoint-test .
 
# 3. run运行 发现我们的ls -a 命令同样生效
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker run entrypoint-test
.
..
.dockerenv
bin
dev
etc
home
lib
 
# 4. 我们的追加命令， 是直接拼接到ENTRYPOINT命令的后面的！
[root@iZ2zeg4ytp0whqtmxbsqiiZ dockerfile]# docker run entrypoint-test -l
total 56
drwxr-xr-x  1 root root 4096 Aug 13 07:52 .
drwxr-xr-x  1 root root 4096 Aug 13 07:52 ..
-rwxr-xr-x  1 root root    0 Aug 13 07:52 .dockerenv
lrwxrwxrwx  1 root root    7 May 11  2019 bin -> usr/bin
drwxr-xr-x  5 root root  340 Aug 13 07:52 dev
drwxr-xr-x  1 root root 4096 Aug 13 07:52 etc
drwxr-xr-x  2 root root 4096 May 11  2019 home
lrwxrwxrwx  1 root root    7 May 11  2019 lib -> usr/lib
lrwxrwxrwx  1 root root    9 May 11  2019 lib64 -> usr/lib64
drwx------  2 root root 4096 Aug  9 21:40 lost+found
```

#### Dockerfile制作tomcat镜像

1、准备镜像文件 tomcat压缩包，jdk的压缩包！

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813164403261.png#pic_center)

2、编写Dockerfile文件，官方命名`Dockerfile`, build会自动寻找这个文件，就不需要-f指定了！

```dockerfile
[root@iZ2zeg4ytp0whqtmxbsqiiZ tomcat]# vim Dockerfile 
FROM centos
MAINTAINER xiaofan<594042358@qq.com>
 
COPY readme.txt /usr/local/readme.txt
 
ADD jdk-8u202-linux-x64.tar.gz /usr/local/ #ADD自动解压文件
ADD apache-tomcat-10.0.27.tar.gz /usr/local/
 
RUN yum -y install vim
 
ENV MYPATH /usr/local
WORKDIR $MYPATH
 
ENV JAVA_HOME /usr/local/jdk1.8.0_202
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
ENV CATALINA_HOME /usr/local/apache-tomcat-10.0.27
ENV CATALINA_BASH /usr/local/apache-tomcat-10.0.27
ENV PATH $PATH:$JAVA_HOME/bin:$CATALINA_HOME/lib:$CATALINA_HOME/bin
 
EXPOSE 8080
 
CMD /usr/local/apache-tomcat-10.0.27/bin/startup.sh && tail -F /usr/local/apache-tomcat-10.0.27/bin/logs/catalina.out

# 构建镜像
docker build -t diytomcat .

# 启动镜像
# docker run -d -p 3344:8080 --name tomcat1 -v /home/tomcat/test:/usr/local/apache-tomcat-10.0.27/webapps/test -v /home/tomcat/tomcatlogs/:/usr/local/apache-tomcat-10.0.27/logs diytomcat

#在本地挂载目录/home/tomcat/test下百度随便编写一个web.xml和index.jsp文件进行测试
```

web.xml

```xml
<?xml version="1.0" encoding="UTF-8"?>
<web-app version="2.4" 
    xmlns="http://java.sun.com/xml/ns/j2ee" 
    xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
    xsi:schemaLocation="http://java.sun.com/xml/ns/j2ee 
        http://java.sun.com/xml/ns/j2ee/web-app_2_4.xsd">    
</web-app>
```

index.jsp

```jsp
<%@ page language="java" contentType="text/html; charset=UTF-8"
    pageEncoding="UTF-8"%>
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>hello. xiaofan</title>
</head>
<body>
Hello World!<br/>
<%
System.out.println("-----my test web logs------");
%>
</body>
</html>
```

项目部署成功， 可以直接访问ok！

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813175909845.png#pic_center)

#### 制作案例二镜像

```dockerfile
[root@iZ2zeg4ytp0w ]# vim Dockerfile
FROM centos 
MAINTAINER xiaofan<594042358@qq.com>
 
ENV MYPATH /usr/local # 镜像的工作目录
WORKDIR $MYPATH

#指定临时目录/tmp，在主机/var/lib/docker目录下创建了一个临时文件并连接到容器的/tmp
VOLUME /tmp

#cetos8官方会将旧版系统的yum源移动到的https://vault.centos.org中
#/etc/yum.repos中更新repos.d使用vault.centos.org代替mirror.centos.org
RUN sed -i -e "s|mirrorlist=|#mirrorlist=|g" /etc/yum.repos.d/CentOS-*
RUN sed -i -e "s|#baseurl=http://mirror.centos.org|baseurl=http://vault.centos.org|g" /etc/yum.repos.d/CentOS-*

RUN yum -y install vim
RUN yum -y install net-tools
RUN yum -y install glibc.i686
RUN mkdir /usr/local/java

#安装包必须要和Dockerfile文件同一位置
ADD jdk-8u202-linux-x64.tar.gz /usr/local/java/
ENV JAVA_HOME /usr/local/java/jdk1.8.0_202
ENV JRE_HOME $JAVA_HOME/jre
ENV CLASSPATH $JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar:$JAVA_HOME/lib:$CLASSPATH
ENV PATH $JAVA_HOME/bin:$PATH

EXPOSE 80
 
CMD echo $MYPATH
CMD echo "---end---"
CMD /bin/bash

[root@iZ2zeg4ytp0w ]#docker build -t mycentos:1.0 .
docker run -it 镜像id /bin/bash
```

# 数据卷容器

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813115602683.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```dockerfile
docker run -it --name docker001 mydockerfile:0.1
ctrl +p +q 退出当前正在运行的容器
docker attach 容器ID #连接到正在运行中的容器
docker run -it --name docker002 --volumes-from docker001 mydockerfile:0.1
docker run -it --name docker003 --volumes-from docker002 mydockerfile:0.1
002继承了001，完成了同步，001就是数据卷容器
```

# Docker Hub

| docker login -u 用户名        | 登陆的用户名:datutu    |
| ----------------------------- | ---------------------- |
| docker login -p 密码          | 登陆的密码:sswan1314   |
| docker login --password-stdin | 从标准输入中读取密码。 |

![image-20221229143751142](C:\Users\Administrator\AppData\Roaming\Typora\typora-user-images\image-20221229143751142.png)

在我们的服务器上提交自己的镜像

```dockerfile
docker push datutu/diytomcat:1.0

# push镜像的问题？
The push refers to repository [docker.io/datutu/diytomcat]
An image does not exist locally with the tag: datutu/diytomcat
 
# 解决，增加一个tag
docker tag diytomcat datutu/tomcat:1.0

# 提交
docker push datutu/tomcat:1.0
```

# 私有仓库

有时候使用 Docker Hub 这样的公共仓库可能不方便，用户可以创建一个本地仓库供私人使用。

 docker-registry 是官方提供的工具，可以用于构建私有的镜像仓库。本文内容基于 docker-registry v2.x 版本。

### 1、安装运行 docker-registry

docker 查询或获取私有仓库(registry)中的镜像，使用

```shell
$ docker search 192.168.222.10:5000
```

 可以通过获取官方 registry 镜像来运行。

```shell
$ docker run -d -p 5000:5000 --restart=always --name registry registry
```

这将使用官方的 registry 镜像来启动私有仓库。默认情况下，仓库会被创建在容器的 /var/lib/registry 目录下。你可以通过 -v 参数来将镜像文件存放在本地的指定路径。例如下面的例子将上传的镜像放到本地的 /opt/data/registry 目录。

```shell
$ docker run -d \
    -p 5000:5000 \
    -v /opt/data/registry:/var/lib/registry \
    registry
```

### 2、在私有仓库上传、搜索、下载镜像

​    创建好私有仓库之后，就可以使用 docker tag 来标记一个镜像，然后推送它到仓库。例如私有仓库地址为 127.0.0.1:5000。先在本机查看已有的镜像。

```shell
$ docker image ls
REPOSITORY                        TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
ubuntu                            latest              ba5877dc9bec        6 weeks ago         192.7 MB
```

使用 docker tag 将 ubuntu:latest 这个镜像标记为 127.0.0.1:5000/ubuntu:latest。格式为 docker tag IMAGE[:TAG] [REGISTRY_HOST[:REGISTRY_PORT]/]REPOSITORY[:TAG]。
```shell
$ docker tag ubuntu:latest 127.0.0.1:5000/ubuntu:latest
$ docker image ls
REPOSITORY                        TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
ubuntu                            latest              ba5877dc9bec        6 weeks ago         192.7 MB
127.0.0.1:5000/ubuntu:latest      latest              ba5877dc9bec        6 weeks ago         192.7 MB
```

使用 docker push 上传标记的镜像。

```shell
$ docker push 127.0.0.1:5000/ubuntu:latest
The push refers to repository [127.0.0.1:5000/ubuntu]
373a30c24545: Pushed
a9148f5200b0: Pushed
cdd3de0940ab: Pushed
fc56279bbb33: Pushed
b38367233d37: Pushed
2aebd096e0e2: Pushed
latest: digest: sha256:fe4277621f10b5026266932ddf760f5a756d2facd505a94d2da12f4f52f71f5a size: 1568
```

用 curl 查看仓库中的镜像。

```cobol
$ curl 127.0.0.1:5000/v2/_catalog
{"repositories":["ubuntu"]}
```

 这里可以看到 {"repositories":["ubuntu"]}，表明镜像已经被成功上传了。先删除已有镜像，再尝试从私有仓库中下载这个镜像。

```shell
$ docker image rm 127.0.0.1:5000/ubuntu:latest
$ docker pull 127.0.0.1:5000/ubuntu:latest
Pulling repository 127.0.0.1:5000/ubuntu:latest
ba5877dc9bec: Download complete
511136ea3c5a: Download complete
9bad880da3d2: Download complete
25f11f5fb0cb: Download complete
ebc34468f71d: Download complete
2318d26665ef: Download complete
 
$ docker image ls
REPOSITORY                         TAG                 IMAGE ID            CREATED             VIRTUAL SIZE
127.0.0.1:5000/ubuntu:latest       latest              ba5877dc9bec        6 weeks ago         192.7 MB
```

2.3、注意事项
        如果你不想使用 127.0.0.1:5000 作为仓库地址，比如想让本网段的其他主机也能把镜像推送到私有仓库。你就得把例如 192.168.222.10:5000 这样的内网地址作为私有仓库地址，这时你会发现无法成功推送镜像。

        这是因为 Docker 默认不允许非 HTTPS 方式推送镜像。我们可以通过 Docker 的配置选项来取消这个限制，或者查看下一节配置能够通过 HTTPS 访问的私有仓库。
2.4、Ubuntu 14.04, Debian 7 Wheezy
        对于使用 upstart 的系统而言，编辑 /etc/default/docker 文件，在其中的 DOCKER_OPTS 中增加如下内容：

````shell
DOCKER_OPTS="--registry-mirror=https://registry.docker-cn.com --insecure-registries=192.168.199.100:5000"
````

重新启动服务。

```shell
$ sudo service docker restart
```


2.5、Ubuntu 16.04+, Debian 8+, centos 7
        对于使用 systemd 的系统，请在 /etc/docker/daemon.json 中写入如下内容（如果文件不存在请新建该文件）

```json
{
  "registry-mirror": [
    "https://registry.docker-cn.com"
  ],
  "insecure-registries": [
    "192.168.222.10:5000"
  ]
}
```

  注意：该文件必须符合 json 规范，否则 Docker 将不能启动。

# 私有仓库高级配置
​        上一节我们搭建了一个具有基础功能的私有仓库，本小节我们来使用 Docker Compose 搭建一个拥有权限认证、TLS 的私有仓库。

新建一个文件夹，以下步骤均在该文件夹中进行。

### 3.1、准备站点证书

​        如果你拥有一个域名，国内各大云服务商均提供免费的站点证书。你也可以使用 openssl 自行签发证书。

这里假设我们将要搭建的私有仓库地址为 docker.domain.com，下面我们介绍使用 openssl 自行签发 docker.domain.com 的站点 SSL 证书。

第一步创建 CA 私钥。

```shell
$ openssl genrsa -out "root-ca.key" 4096
```

第二步利用私钥创建 CA 根证书请求文件。

```shell
$ openssl req \
          -new -key "root-ca.key" \
          -out "root-ca.csr" -sha256 \
          -subj '/C=CN/ST=Shanxi/L=Datong/O=Your Company Name/CN=Your Company Name Docker Registry CA'
```

以上命令中 -subj 参数里的 /C 表示国家，如 CN；/ST 表示省；/L 表示城市或者地区；/O 表示组织名；/CN 通用名称。

第三步配置 CA 根证书，新建 root-ca.cnf。

```shell
[root_ca]
basicConstraints = critical,CA:TRUE,pathlen:1
keyUsage = critical, nonRepudiation, cRLSign, keyCertSign
subjectKeyIdentifier=hash
```

第四步签发根证书。

```shell
$ openssl x509 -req  -days 3650  -in "root-ca.csr" \
               -signkey "root-ca.key" -sha256 -out "root-ca.crt" \
               -extfile "root-ca.cnf" -extensions \
               root_ca
```

第五步生成站点 SSL 私钥。

```shell
$ openssl genrsa -out "docker.domain.com.key" 4096
```

第六步使用私钥生成证书请求文件。

```shell
$ openssl req -new -key "docker.domain.com.key" -out "site.csr" -sha256 \
          -subj '/C=CN/ST=Shanxi/L=Datong/O=Your Company Name/CN=docker.domain.com'
```

第七步配置证书，新建 site.cnf 文件。

```cobol
[server]
authorityKeyIdentifier=keyid,issuer
basicConstraints = critical,CA:FALSE
extendedKeyUsage=serverAuth
keyUsage = critical, digitalSignature, keyEncipherment
subjectAltName = DNS:docker.domain.com, IP:127.0.0.1
subjectKeyIdentifier=hash
```

第八步签署站点 SSL 证书。

```cobol
$ openssl x509 -req -days 750 -in "site.csr" -sha256 \
    -CA "root-ca.crt" -CAkey "root-ca.key"  -CAcreateserial \
    -out "docker.domain.com.crt" -extfile "site.cnf" -extensions server
```

这样已经拥有了 docker.domain.com 的网站 SSL 私钥 docker.domain.com.key 和 SSL 证书 docker.domain.com.crt 及 CA 根证书 root-ca.crt。

新建 ssl 文件夹并将 docker.domain.com.key docker.domain.com.crt root-ca.crt 这三个文件移入，删除其他文件。

### 3.2、配置私有仓库

​    私有仓库默认的配置文件位于 /etc/docker/registry/config.yml，我们先在本地编辑 config.yml，之后挂载到容器中。

```yml
version: 0.1
log:
  accesslog:
    disabled: true
  level: debug
  formatter: text
  fields:
    service: registry
    environment: staging
storage:
  delete:
    enabled: true
  cache:
    blobdescriptor: inmemory
  filesystem:
    rootdirectory: /var/lib/registry
auth:
  htpasswd:
    realm: basic-realm
    path: /etc/docker/registry/auth/nginx.htpasswd
http:
  addr: :443
  host: https://docker.domain.com
  headers:
    X-Content-Type-Options: [nosniff]
  http2:
    disabled: false
  tls:
    certificate: /etc/docker/registry/ssl/docker.domain.com.crt
    key: /etc/docker/registry/ssl/docker.domain.com.key
health:
  storagedriver:
    enabled: true
    interval: 10s
threshold: 3
```

### 3.3、生成 http 认证文件

```shell
$ mkdir auth
 
$ docker run --rm \
    --entrypoint htpasswd \
    registry \
    -Bbn username password > auth/nginx.htpasswd
```

将上面的 username password 替换为你自己的用户名和密码。

### 3.4、编辑 docker-compose.yml

```yml
version: '3'
 
services:
  registry:
    image: registry
    ports:
      - "443:443"
    volumes:
      - ./:/etc/docker/registry
      - registry-data:/var/lib/registry
 
volumes:
  registry-data:
```

### 3.5、修改 hosts

​    编辑 /etc/hosts：

```shell
127.0.0.1 docker.domain.com
```

### 3.6、启动

```shell
$ docker-compose up -d
```

这样我们就搭建好了一个具有权限认证、TLS 的私有仓库，接下来我们测试其功能是否正常。

### 3.7、测试私有仓库功能

​    由于自行签发的 CA 根证书不被系统信任，所以我们需要将 CA 根证书 ssl/root-ca.crt 移入 /etc/docker/certs.d/docker.domain.com 文件夹中。

```shell
$ sudo mkdir -p /etc/docker/certs.d/docker.domain.com
$ sudo cp ssl/root-ca.crt /etc/docker/certs.d/docker.domain.com/ca.crt
```

登录到私有仓库。

```shell
$ docker login docker.domain.com
```

尝试推送、拉取镜像。

```shell
$ docker pull ubuntu:18.04
$ docker tag ubuntu:18.04 docker.domain.com/username/ubuntu:18.04
$ docker push docker.domain.com/username/ubuntu:18.04
$ docker image rm docker.domain.com/username/ubuntu:18.04
$ docker pull docker.domain.com/username/ubuntu:18.04
```

如果我们退出登录，尝试推送镜像。

```shell
$ docker logout docker.domain.com
$ docker push docker.domain.com/username/ubuntu:18.04
no basic auth credentials
```

发现会提示没有登录，不能将镜像推送到私有仓库中。

### 3.8、注意事项

​    如果你本机占用了 443 端口，你可以配置 Nginx 代理，这里不再赘述。

​    [Nginx代理](https://so.csdn.net/so/search?q=Nginx代理&spm=1001.2101.3001.7020)：[Authenticate proxy with nginx | Docker Documentation](https://docs.docker.com/registry/recipes/nginx/)

# Nexus3.x 的私有仓库
​        使用 Docker 官方的 Registry 创建的仓库面临一些维护问题。比如某些镜像删除以后空间默认是不会回收的，需要一些命令去回收空间然后重启 Registry 程序。在企业中把内部的一些工具包放入 Nexus 中是比较常见的做法，最新版本 Nexus3.x 全面支持 Docker 的私有镜像。所以使用 Nexus3.x 一个软件来管理 Docker , Maven , Yum , PyPI 等是一个明智的选择。

### 4.1、启动 Nexus 容器

```shell
$ docker run -d --name nexus3 --restart=always \
    -p 8081:8081 \
    --mount src=nexus-data,target=/nexus-data \
    sonatype/nexus3
```

等待 3-5 分钟，如果 nexus3 容器没有异常退出，那么你可以使用浏览器打开 http://YourIP:8081 访问 Nexus 了。

​    第一次启动 Nexus 的默认帐号是 admin 密码是 admin123 登录以后点击页面上方的齿轮按钮进行设置。

### 4.2、创建仓库

​        创建一个私有仓库的方法： Repository->Repositories 点击右边菜单 Create repository 选择 docker (hosted)

        Name: 仓库的名称
    
        HTTP: 仓库单独的访问端口
    
        Enable Docker V1 API: 如果需要同时支持 V1 版本请勾选此项（不建议勾选）。
    
        Hosted -> Deployment pollcy: 请选择 Allow redeploy 否则无法上传 Docker 镜像。
    
        其它的仓库创建方法请各位自己摸索，还可以创建一个 docker (proxy) 类型的仓库链接到 DockerHub 上。再创建一个 docker (group) 类型的仓库把刚才的 hosted 与 proxy 添加在一起。主机在访问的时候默认下载私有仓库中的镜像，如果没有将链接到 DockerHub 中下载并缓存到 Nexus 中。

### 4.3、添加访问权限

​        菜单 Security->Realms 把 Docker Bearer Token Realm 移到右边的框中保存。

        添加用户规则：菜单 Security->Roles->Create role 在 Privlleges 选项搜索 docker 把相应的规则移动到右边的框中然后保存。
    
        添加用户：菜单 Security->Users->Create local user 在 Roles 选项中选中刚才创建的规则移动到右边的窗口保存。

### 4.4、NGINX 加密代理

​        证书的生成请参见 私有仓库高级配置 里面证书生成一节。
**NGINX 示例配置如下：**

```json
upstream register
{
    server "YourHostName OR IP":5001; #端口为上面添加的私有镜像仓库是设置的 HTTP 选项的端口号
    check interval=3000 rise=2 fall=10 timeout=1000 type=http;
    check_http_send "HEAD / HTTP/1.0\r\n\r\n";
    check_http_expect_alive http_4xx;
}
 
server {
    server_name YourDomainName;#如果没有 DNS 服务器做解析，请删除此选项使用本机 IP 地址访问
    listen       443 ssl;
 
    ssl_certificate key/example.crt;
    ssl_certificate_key key/example.key;
 
    ssl_session_timeout  5m;
    ssl_protocols TLSv1 TLSv1.1 TLSv1.2;
    ssl_ciphers  HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers   on;
    large_client_header_buffers 4 32k;
    client_max_body_size 300m;
    client_body_buffer_size 512k;
    proxy_connect_timeout 600;
    proxy_read_timeout   600;
    proxy_send_timeout   600;
    proxy_buffer_size    128k;
    proxy_buffers       4 64k;
    proxy_busy_buffers_size 128k;
    proxy_temp_file_write_size 512k;
 
    location / {
        proxy_set_header Host $host;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Port $server_port;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection $connection_upgrade;
        proxy_redirect off;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://register;
        proxy_read_timeout 900s;
 
    }
    error_page   500 502 503 504  /50x.html;
}
```

### 4.5、Docker 主机访问镜像仓库

​    如果不启用 SSL 加密可以通过前面章节的方法添加信任地址到 Docker 的配置文件中然后重启 Docker，使用 SSL 加密以后程序需要访问就不能采用修改配置的访问了。具体方法如下：

```shell
$ openssl s_client -showcerts -connect YourDomainName OR HostIP:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >ca.crt
$ cat ca.crt | sudo tee -a /etc/ssl/certs/ca-certificates.crt
$ systemctl restart docker
```

使用 docker login YourDomainName OR HostIP 进行测试，用户名密码填写上面 Nexus 中生成的。

# 发布到阿里云镜像服务上

1. 登录阿里云

2. 找到容器镜像服务

3. 创建命名空间

   ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813190111625.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

   创建容器镜像

   ![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813190303741.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

2\点击仓库名称，参考官方文档即可

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200813191526549.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

# Docker网络

只要安装了docker，就会有一个网卡 docker0桥接模式，使用的技术是veth-pair技术

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200814091723905.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

所有容器不指定网络的情况下，都是docker0路由的，doucker会给我们的容器分配一个默认的可用IP

### docker network

```dockerfile
docker network ls #查看所有的docker网络

#bridge： 桥接模式，桥接 docker 默认，自己创建的也是用brdge模式
#none： 不配置网络
#host： 和宿主机共享网络
#container：容器网络连通！（用的少， 局限很大）
```

###  测试

```dockerfile
# 我们直接启动的命令默认有一个 --net bridge，而这个就是我们的docker0
docker run -d -P --name tomcat01 --net bridge tomcat
# docker0特点，默认，容器名不能访问， --link可以打通连接！

# 我们可以自定义一个网络！
# --driver bridge
# --subnet 192.168.0.0/16 可以支持255*255个网络 192.168.0.2 ~ 192.168.255.254 子网掩码
# --gateway 192.168.0.1 网关，路由器
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker network create --driver bridge --subnet 192.168.0.0/16 --gateway 192.168.0.1 mynet
26a5afdf4805d7ee0a660b82244929a4226470d99a179355558dca35a2b983ec
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker network ls
NETWORK ID          NAME                DRIVER              SCOPE
30d601788862        bridge              bridge              local
226019b14d91        host                host                local
26a5afdf4805        mynet               bridge              local
7496c014f74b        none                null                local

#创建一个新的 Docker 网络。
#-d 指定模式（默认桥接）
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker network create -d bridge my-net
```

我们自己创建的网络就ok了！--->192.168.0.1

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200814112009570.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

在自己创建的网络里面启动两个容器

```dockerfile
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker run -d -P --name tomcat-net-01 --net mynet tomcat
0e85ebe6279fd23379d39b27b5f47c1e18f23ba7838637802973bf6449e22f5c
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker run -d -P --name tomcat-net-02 --net mynet tomcat
c6e462809ccdcebb51a4078b1ac8fdec33f1112e9e416406b606d0c9fb6f21b5
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker network inspect mynet
[
    {
        "Name": "mynet",
        "Id": "26a5afdf4805d7ee0a660b82244929a4226470d99a179355558dca35a2b983ec",
        "Created": "2020-08-14T11:12:40.553433163+08:00",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": {},
            "Config": [
                {
                    "Subnet": "192.168.0.0/16",
                    "Gateway": "192.168.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {
            "0e85ebe6279fd23379d39b27b5f47c1e18f23ba7838637802973bf6449e22f5c": {
                "Name": "tomcat-net-01",
                "EndpointID": "576ce5c0f5860a5aab5e487a805da9d72f41a409c460f983c0bd341dd75d83ac",
                "MacAddress": "02:42:c0:a8:00:02",
                "IPv4Address": "192.168.0.2/16",
                "IPv6Address": ""
            },
            "c6e462809ccdcebb51a4078b1ac8fdec33f1112e9e416406b606d0c9fb6f21b5": {
                "Name": "tomcat-net-02",
                "EndpointID": "81ecbc4fe26e49855fe374f2d7c00d517b11107cc91a174d383ff6be37d25a30",
                "MacAddress": "02:42:c0:a8:00:03",
                "IPv4Address": "192.168.0.3/16",
                "IPv6Address": ""
            }
        },
        "Options": {},
        "Labels": {}
    }
]
 
# 再次拼连接
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker exec -it tomcat-net-01 ping 192.168.0.3
PING 192.168.0.3 (192.168.0.3) 56(84) bytes of data.
64 bytes from 192.168.0.3: icmp_seq=1 ttl=64 time=0.113 ms
64 bytes from 192.168.0.3: icmp_seq=2 ttl=64 time=0.093 ms
^C
--- 192.168.0.3 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 999ms
rtt min/avg/max/mdev = 0.093/0.103/0.113/0.010 ms
# 现在不使用 --link也可以ping名字了！
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker exec -it tomcat-net-01 ping tomcat-net-02
PING tomcat-net-02 (192.168.0.3) 56(84) bytes of data.
64 bytes from tomcat-net-02.mynet (192.168.0.3): icmp_seq=1 ttl=64 time=0.068 ms
64 bytes from tomcat-net-02.mynet (192.168.0.3): icmp_seq=2 ttl=64 time=0.096 ms
64 bytes from tomcat-net-02.mynet (192.168.0.3): icmp_seq=3 ttl=64 time=0.094 ms
```

我们自定义的网络docker都已经帮我们维护好了对应的关系，推荐我们平时这样使用网络

好处：

redis - 不同的集群使用不同的网络，保证集群时安全和健康的

mysql - 不同的集群使用不同的网络，保证集群时安全和健康的

### 网络连通

测试打通tomcat01 和mynet

![在这里插入图片描述](https://img-blog.csdnimg.cn/2020081411482318.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

![在这里插入图片描述](https://img-blog.csdnimg.cn/20200814114621170.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2ZhbmppYW5oYWk=,size_16,color_FFFFFF,t_70#pic_center)

```dockerfile
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker network connect mynet tomcat01
 
# 连通之后就是讲tomcat01 放到了mynet网路下
# 一个容器两个ip地址：
# 阿里云服务器，公网ip，私网ip
```

```dockerfile
# 连通ok
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker exec -it tomcat01 ping tomcat-net-01
PING tomcat-net-01 (192.168.0.2) 56(84) bytes of data.
64 bytes from tomcat-net-01.mynet (192.168.0.2): icmp_seq=1 ttl=64 time=0.100 ms
64 bytes from tomcat-net-01.mynet (192.168.0.2): icmp_seq=2 ttl=64 time=0.085 ms
^C
--- tomcat-net-01 ping statistics ---
2 packets transmitted, 2 received, 0% packet loss, time 1000ms
rtt min/avg/max/mdev = 0.085/0.092/0.100/0.012 ms
# 依旧无法连通，没有connect
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker exec -it tomcat02 ping tomcat-net-01
ping: tomcat-net-01: Name or service not known
```

### 实战：部署redis集群

```dockerfile
# 创建网卡
docker network create redis --subnet 172.38.0.0/16

# 通过脚本创建六个redis配置
for port in $(seq 1 6); \
do \
mkdir -p /mydata/redis/node-${port}/conf
touch /mydata/redis/node-${port}/conf/redis.conf
cat << EOF >/mydata/redis/node-${port}/conf/redis.conf
port 6379
bind 0.0.0.0
cluster-enabled yes
cluster-config-file nodes.conf
cluster-node-timeout 5000
cluster-announce-ip 172.38.0.1${port}
cluster-announce-port 6379
cluster-announce-bus-port 16379
appendonly yes
EOF
done

# 创建结点1
docker run -p 6371:6379 -p 16371:16379 --name redis-1 \
-v /mydata/redis/node-1/data:/data \
-v /mydata/redis/node-1/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.11 redis:7.0.7 redis-server /etc/redis/redis.conf
 
#创建结点2
docker run -p 6372:6379 -p 16372:16379 --name redis-2 \
-v /mydata/redis/node-2/data:/data \
-v /mydata/redis/node-2/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.12 redis:7.0.7 redis-server /etc/redis/redis.conf

#创建结点3
docker run -p 6373:6379 -p 16373:16379 --name redis-3 \
-v /mydata/redis/node-3/data:/data \
-v /mydata/redis/node-3/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.13 redis:7.0.7 redis-server /etc/redis/redis.conf

#创建结点4
docker run -p 6374:6379 -p 16374:16379 --name redis-4 \
-v /mydata/redis/node-4/data:/data \
-v /mydata/redis/node-4/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.14 redis:7.0.7 redis-server /etc/redis/redis.conf

#创建结点5
docker run -p 6375:6379 -p 16375:16379 --name redis-5 \
-v /mydata/redis/node-5/data:/data \
-v /mydata/redis/node-5/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.15 redis:7.0.7 redis-server /etc/redis/redis.conf
#创建结点6
docker run -p 6376:6379 -p 16376:16379 --name redis-6 \
-v /mydata/redis/node-6/data:/data \
-v /mydata/redis/node-6/conf/redis.conf:/etc/redis/redis.conf \
-d --net redis --ip 172.38.0.16 redis:7.0.7 redis-server /etc/redis/redis.conf
 
# 创建集群
[root@iZ2zeg4ytp0whqtmxbsqiiZ ~]# docker exec -it redis-1 /bin/sh
进来以后redis默认/data 目录

[/data]# ls
appendonly.aof  nodes.conf

[/data ]# redis-cli --cluster create 172.38.0.11:6379 172.38.0.12:6379 172.38.0.13:6379 172.38.0.14:6379 172.38.0.15:6379 172.38.0.16:6379 --cluster-replicas 1
>>> Performing hash slots allocation on 6 nodes...
Master[0] -> Slots 0 - 5460
Master[1] -> Slots 5461 - 10922
Master[2] -> Slots 10923 - 16383
Adding replica 172.38.0.15:6379 to 172.38.0.11:6379
Adding replica 172.38.0.16:6379 to 172.38.0.12:6379
Adding replica 172.38.0.14:6379 to 172.38.0.13:6379
M: 541b7d237b641ac2ffc94d17c6ab96b18b26a638 172.38.0.11:6379
   slots:[0-5460] (5461 slots) master
M: a89c1f1245b264e4a402a3cf99766bcb6138dbca 172.38.0.12:6379
   slots:[5461-10922] (5462 slots) master
M: 259e804d6df74e67a72e4206d7db691a300c775e 172.38.0.13:6379
   slots:[10923-16383] (5461 slots) master
S: 9b19170eea3ea1b92c58ad18c0b5522633a9e271 172.38.0.14:6379
   replicates 259e804d6df74e67a72e4206d7db691a300c775e
S: 061a9d38f22910aaf0ba1dbd21bf1d8f57bcb7d5 172.38.0.15:6379
   replicates 541b7d237b641ac2ffc94d17c6ab96b18b26a638
S: 7a16b9bbb0615ec95fc978fa62fc054df60536f0 172.38.0.16:6379
   replicates a89c1f1245b264e4a402a3cf99766bcb6138dbca
Can I set the above configuration? (type 'yes' to accept): yes
>>> Nodes configuration updated
>>> Assign a different config epoch to each node
>>> Sending CLUSTER MEET messages to join the cluster
Waiting for the cluster to join
...
>>> Performing Cluster Check (using node 172.38.0.11:6379)
M: 541b7d237b641ac2ffc94d17c6ab96b18b26a638 172.38.0.11:6379
   slots:[0-5460] (5461 slots) master
   1 additional replica(s)
M: a89c1f1245b264e4a402a3cf99766bcb6138dbca 172.38.0.12:6379
   slots:[5461-10922] (5462 slots) master
   1 additional replica(s)
S: 7a16b9bbb0615ec95fc978fa62fc054df60536f0 172.38.0.16:6379
   slots: (0 slots) slave
   replicates a89c1f1245b264e4a402a3cf99766bcb6138dbca
S: 061a9d38f22910aaf0ba1dbd21bf1d8f57bcb7d5 172.38.0.15:6379
   slots: (0 slots) slave
   replicates 541b7d237b641ac2ffc94d17c6ab96b18b26a638
M: 259e804d6df74e67a72e4206d7db691a300c775e 172.38.0.13:6379
   slots:[10923-16383] (5461 slots) master
   1 additional replica(s)
S: 9b19170eea3ea1b92c58ad18c0b5522633a9e271 172.38.0.14:6379
   slots: (0 slots) slave
   replicates 259e804d6df74e67a72e4206d7db691a300c775e
[OK] All nodes agree about slots configuration.
>>> Check for open slots...
>>> Check slots coverage...
[OK] All 16384 slots covered.
```

docker搭建redis集群完成！

```dockerfile
[/data]#redis-cli -c(连接集群)
<127.0.0.1:6379>#cluster info(查看集群信息)
<127.0.0.1:6379>#cluster nodes(查看集群节点)
#测试
<127.0.0.1:6379>#set a b
<127.0.0.1:6379>#get a
```

# Docker Compose

### 1、下载安装

```dockerfile
https://docs.docker.com/compose/
https://docs.docker.com/compose/compose-file/compose-file-v3/

curl -SL https://github.com/docker/compose/releases/download/v2.15.1/docker-compose-linux-x86_64 -o /usr/local/bin/docker-compose #如果下载不成功就用下面试试

curl -L "https://get.daocloud.io/docker/compose/releases/download/v2.15.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose

sudo chmod +x /usr/local/bin/docker-compose
docker-compose --version

#卸载
sudo rm /usr/local/bin/docker-compose
```

### 2、Compose 常用命令

```dockerfile
docker-compose -h                  #查看帮助
docker-compose up                  #启动所有docker-compose服务
docker-compose up -d               #启动所有docker-compose并后台运行
docker-compose down                #停止并删除容器、网络、卷、镜像
docker-compose exec yml里面的服务id  #进入容器内部 /bin/bash
docker-compose ps                  #展示当前docker-compose编排过的运行的所有容器
docker-compose top                 #展示当前docker-compose编排过的容器进程
docker-compose logs yml里面的服务id  #查看容器输出日志
docker-compose config              #检查配置
docker-compose config -q           #检查配置，有问题才有输出
docker-compose start               #启动服务
docker-compose stop                #停止服务
docker-compose restart             #重启服务
docker-compose -f docker-compose.yml up -d  # 指定模板
#kill：通过发送 SIGKILL 信号来停止指定服务的容器
docker-compose kill eureka

#pull：下载服务镜像
docker-compose pull eureka

#scale：设置指定服务运气容器的个数，以 service=num 形式指定
docker-compose scale user=3 movie=3
```

### 3、docker-compose.yml 语法说明

#### 1、image
指定为镜像名称或镜像ID。如果镜像不存在，Compose将尝试从互联网拉取这个镜像，例如： image: ubuntu image: orchardup/postgresql image: a4bc65fd

#### 2、build
build 指定 Dockerfile 所在文件夹的路径（可以是绝对路径，或者相对 docker-compose.yml 文件的路径。Compose将会利用他自动构建这个镜像，然后使用这个镜像。 build: ./dir

```yaml
services:
  webapp:
    build: ./       # 从当前目录下寻找Dockerfile文件
```

你也可以使用 context 指令指定 Dockerfile 所在文件夹的路径，使用 dockerfile 指令指定 Dockerfile 文件名，可以使用 args 指令指定构建镜像时往Dockerfile中传入的变量。

```yaml
services:
  webapp:
    build:
      context: ./dir
      dockerfile: Dockerfile_flask
      args:
        buildno: 1
```

#### 3、command

覆盖容器启动后默认执行的命令。 command: bundle exec thin -p 3000

#### 4、links
链接到其他服务容器，使用服务名称(同时作为别名)或服务别名（SERVICE:ALIAS）都可以

```yaml
links:
 - db
 - db:database
 - redis
```

注意：使用别名会自动在服务器中的/etc/hosts 里创建，如：172.17.2.186 db，相应的环境变量也会被创建。

#### 5、external_links

链接到docker-compose.yml外部的容器，甚至并非是Compose管理的容器。参数格式和links类似。 external_links:

```yaml
- redis_1
 - project_db_1:mysql
 - project_db_2:sqlserver
```

#### 6、ports

暴露端口信息。 宿主机器端口：容器端口（HOST:CONTAINER）格式或者仅仅指定容器的端口（宿主机器将会随机分配端口）都可以。

```yaml
ports:
 - "3306"
 - "8080:80"
 - "127.0.0.1:8090:8001"
 
ports:
 - "80:80" # 绑定容器的80端口到主机的80端口
 - "9000:8080" # 绑定容器的8080端口到主机的9000端口
 - "443" # 绑定容器的443端口到主机的任意端口，容器启动时随机分配绑定的主机端口号
```

注意：当使用 HOST:CONTAINER 格式来映射端口时，如果你使用的容器端口小于 60 你可能会得到错误得结果，因为 YAML 将会解析 xx:yy 这种数字格式为 60 进制。所以建议采用字符串格式。

#### 7、expose

暴露端口，与posts不同的是expose只可以暴露端口而不能映射到主机，只供外部服务连接使用；仅可以指定内部端口为参数。

和ports的区别是，expose暴露容器给link到当前容器的容器，不会将端口暴露给主机。

```yaml
expose:
 - "3000"
 - "8000"
```

#### 8、volumes

设置卷挂载的路径。可以设置宿主机路径:容器路径（host:container）或加上访问模式（host:container:ro）ro就是readonly的意思，只读模式。

数据卷所挂载路径设置，挂载数据卷的默认权限是读写（rw）。
你可以在主机上挂载绝对路径，或者挂载相对路径，相对路径是相对于当前正在使用的compose配置文件的目录进行扩展。 相对路径应始终以 . 或者 … 开始。

```yaml
volumes:
 - /var/lib/mysql:/var/lib/mysql
 - /configs/mysql:/etc/configs/:ro
 
volumes:    
	# 挂载绝对路径映射，没有这个路径的话会自动创建
	- /root/docker/composetest/tomcat/app_data:/var/lib/mysql
	# 或者使用相对路径映射，也会自动创建
	- ./app_data:/var/lib/mysql
```

#### 9、volunes_from

挂载另一个服务或容器的所有数据卷。

```yaml
volumes_from:
 - service_name
 - container_name
```

#### 10、environment

设置环境变量。可以属于数组或字典两种格式。 如果只给定变量的名称则会自动加载它在Compose主机上的值，可以用来防止泄露不必要的数据。

```yaml
environment:
 - RACK_ENV=development
 - SESSION_SECRET
 
environment:    # 使用字典格式，类似于 docker run -e MYSQL_ROOT_PASSWORD=root
	MYSQL_ROOT_PASSWORD: root
	
environment:    # 使用数组格式
	- MYSQL_ROOT_PASSWORD=root
```

#### 11、env_file

从文件中获取环境变量，可以为单独的文件路径或列表。 如果通过docker-compose -f FILE指定了模板文件，则env_file中路径会基于模板文件路径。 如果有变量名称与environment指令冲突，则以后者为准。

```yaml
env_file: .env  # 仅使用单个env文件
env_file:       # 通过数组来使用多个env文件
 - ./common.env
 - ./apps/web.env
 - /opt/secrets.env
```

环境变量文件中每一行都必须有注释，支持#开头的注释行。

```yaml
# common.env: Set Rails/Rack environment
RACK_ENV=development

# common.env: Set development environment
MYSQL_ROOT_PASSWORD=root
```

#### 12、extends

基于已有的服务进行服务扩展。例如我们已经有了一个webapp服务，模板文件为common.yml.

```yaml
# common.yml
webapp:
build: ./webapp
environment:
\ - DEBUG=false
\ - SEND_EMAILS=false
```

编写一个新的 development.yml 文件，使用 common.yml 中的 webapp 服务进行扩展。 development.yml

```yaml
web:
extends:
file: common.yml
service: 
  webapp:
    ports:
      \ - "8080:80"
    links:
      \ - db
    envelopment:
      - DEBUG=true
   db:
    image: mysql:5.7
```

后者会自动继承common.yml中的webapp服务及相关的环境变量。

#### 13、net

设置网络模式。使用和docker client 的 --net 参数一样的值。

```yaml
# 容器默认连接的网络，是所有Docker安装时都默认安装的docker0网络.
net: "bridge"
# 容器定制的网络栈.
net: "none"
# 使用另一个容器的网络配置
net: "container:[name or id]"
# 在宿主网络栈上添加一个容器，容器中的网络配置会与宿主的一样
net: "host"

services:
  webapp:
    networks:
      - flask-net
networks:   # 在顶级networks关键字中需要声明，才会在启动时自动创建该网络，否则报错。
  flask-net:
```

Docker会为每个节点自动创建三个网络： 网络名称 作用 bridge 容器默认连接的网络，是所有Docker安装时都默认安装的docker0网络 none 容器定制的网络栈 host 在宿主网络栈上添加一个容器，容器中的网络配置会与宿主的一样 附录： 操作名称 命令 创建网络 docker network create -d bridge mynet 查看网络列表 docker network ls

#### 14、pid

和宿主机系统共享进程命名空间，打开该选项的容器可以相互通过进程id来访问和操作。

```yaml
pid: "host"
```

#### 15、dns

```yaml
配置DNS服务器。可以是一个值，也可以是一个列表。
dns: 8.8.8.8
dns:
 - 8.8.8.8
 - 9.9.9.9
```

#### 16、cap_add，cap_drop

添加或放弃容器的Linux能力（Capability）。

```yaml
cap_add:
 - ALL
cap_drop:
 - NET_ADMIN
 - SYS_ADMIN
```

#### 17、dns_search

配置DNS搜索域。可以是一个值也可以是一个列表。

```yaml
dns_search: example.com
dns_search:
 - domain1.example.com
 \ - domain2.example.com
working_dir, entrypoint, user, hostname, domainname, mem_limit, privileged, restart, stdin_open, tty, cpu_shares
```

这些都是和 docker run 支持的选项类似。

```yaml
cpu_shares: 73
working_dir: /code
entrypoint: /code/entrypoint.sh
user: postgresql
hostname: foo
domainname: foo.com
mem_limit: 1000000000
privileged: true
restart: always
stdin_open: true
tty: true
```

#### 18、depends_on

解决容器的依赖、启动先后的问题。以下例子中会先启动 redis db 再启动 web。
注意：web 服务不会等待 redis、db 完全启动之后才启动。

```yaml
services:
  webapp:
    build: .
    depends_on:
      - db
      - redis
  redis:
    image: redis:latest
  db:
    image: mysql:latest
```

#### 19、restart

restart 指定docker容器（服务）总是运行。

```yaml
services:
  webapp:
    networks:
      - flask-net
    restart: always
```

### 4、docker-compose配置文件的常见参数

```yaml
常用参数：
    version           # 指定 compose 文件的版本
    services          # 定义所有的 service 信息, services 下面的第一级别的 key 既是一个 service 的名称
 
        build                 # 指定包含构建上下文的路径, 或作为一个对象，该对象具有 context 和指定的 dockerfile 文件以及 args 参数值
            context               # context: 指定 Dockerfile 文件所在的路径
            dockerfile            # dockerfile: 指定 context 指定的目录下面的 Dockerfile 的名称(默认为 Dockerfile)
            args                  # args: Dockerfile 在 build 过程中需要的参数 (等同于 docker container build --build-arg 的作用)
            cache_from            # v3.2中新增的参数, 指定缓存的镜像列表 (等同于 docker container build --cache_from 的作用)
            labels                # v3.3中新增的参数, 设置镜像的元数据 (等同于 docker container build --labels 的作用)
            shm_size              # v3.5中新增的参数, 设置容器 /dev/shm 分区的大小 (等同于 docker container build --shm-size 的作用)
 
        command               # 覆盖容器启动后默认执行的命令, 支持 shell 格式和 [] 格式
 
        configs               # 不知道怎么用
 
        cgroup_parent         # 不知道怎么用
 
        container_name        # 指定容器的名称 (等同于 docker run --name 的作用)
 
        credential_spec       # 不知道怎么用
 
        deploy                # v3 版本以上, 指定与部署和运行服务相关的配置, deploy 部分是 docker stack 使用的, docker stack 依赖 docker swarm
            endpoint_mode         # v3.3 版本中新增的功能, 指定服务暴露的方式
                vip                   # Docker 为该服务分配了一个虚拟 IP(VIP), 作为客户端的访问服务的地址
                dnsrr                 # DNS轮询, Docker 为该服务设置 DNS 条目, 使得服务名称的 DNS 查询返回一个 IP 地址列表, 客户端直接访问其中的一个地址
            labels                # 指定服务的标签，这些标签仅在服务上设置
            mode                  # 指定 deploy 的模式
                global                # 每个集群节点都只有一个容器
                replicated            # 用户可以指定集群中容器的数量(默认)
            placement             # 不知道怎么用
            replicas              # deploy 的 mode 为 replicated 时, 指定容器副本的数量
            resources             # 资源限制
                limits                # 设置容器的资源限制
                    cpus: "0.5"           # 设置该容器最多只能使用 50% 的 CPU 
                    memory: 50M           # 设置该容器最多只能使用 50M 的内存空间 
                reservations          # 设置为容器预留的系统资源(随时可用)
                    cpus: "0.2"           # 为该容器保留 20% 的 CPU
                    memory: 20M           # 为该容器保留 20M 的内存空间
            restart_policy        # 定义容器重启策略, 用于代替 restart 参数
                condition             # 定义容器重启策略(接受三个参数)
                    none                  # 不尝试重启
                    on-failure            # 只有当容器内部应用程序出现问题才会重启
                    any                   # 无论如何都会尝试重启(默认)
                delay                 # 尝试重启的间隔时间(默认为 0s)
                max_attempts          # 尝试重启次数(默认一直尝试重启)
                window                # 检查重启是否成功之前的等待时间(即如果容器启动了, 隔多少秒之后去检测容器是否正常, 默认 0s)
            update_config         # 用于配置滚动更新配置
                parallelism           # 一次性更新的容器数量
                delay                 # 更新一组容器之间的间隔时间
                failure_action        # 定义更新失败的策略
                    continue              # 继续更新
                    rollback              # 回滚更新
                    pause                 # 暂停更新(默认)
                monitor               # 每次更新后的持续时间以监视更新是否失败(单位: ns|us|ms|s|m|h) (默认为0)
                max_failure_ratio     # 回滚期间容忍的失败率(默认值为0)
                order                 # v3.4 版本中新增的参数, 回滚期间的操作顺序
                    stop-first            #旧任务在启动新任务之前停止(默认)
                    start-first           #首先启动新任务, 并且正在运行的任务暂时重叠
            rollback_config       # v3.7 版本中新增的参数, 用于定义在 update_config 更新失败的回滚策略
                parallelism           # 一次回滚的容器数, 如果设置为0, 则所有容器同时回滚
                delay                 # 每个组回滚之间的时间间隔(默认为0)
                failure_action        # 定义回滚失败的策略
                    continue              # 继续回滚
                    pause                 # 暂停回滚
                monitor               # 每次回滚任务后的持续时间以监视失败(单位: ns|us|ms|s|m|h) (默认为0)
                max_failure_ratio     # 回滚期间容忍的失败率(默认值0)
                order                 # 回滚期间的操作顺序
                    stop-first            # 旧任务在启动新任务之前停止(默认)
                    start-first           # 首先启动新任务, 并且正在运行的任务暂时重叠
 
            注意：
                支持 docker-compose up 和 docker-compose run 但不支持 docker stack deploy 的子选项
                security_opt  container_name  devices  tmpfs  stop_signal  links    cgroup_parent
                network_mode  external_links  restart  build  userns_mode  sysctls
 
        devices               # 指定设备映射列表 (等同于 docker run --device 的作用)
 
        depends_on            # 定义容器启动顺序 (此选项解决了容器之间的依赖关系， 此选项在 v3 版本中 使用 swarm 部署时将忽略该选项)
            示例：
                docker-compose up 以依赖顺序启动服务，下面例子中 redis 和 db 服务在 web 启动前启动
                默认情况下使用 docker-compose up web 这样的方式启动 web 服务时，也会启动 redis 和 db 两个服务，因为在配置文件中定义了依赖关系
 
                version: '3'
                services:
                    web:
                        build: .
                        depends_on:
                            - db      
                            - redis  
                    redis:
                        image: redis
                    db:
                        image: postgres                             
 
        dns                   # 设置 DNS 地址(等同于 docker run --dns 的作用)
 
        dns_search            # 设置 DNS 搜索域(等同于 docker run --dns-search 的作用)
 
        tmpfs                 # v2 版本以上, 挂载目录到容器中, 作为容器的临时文件系统(等同于 docker run --tmpfs 的作用, 在使用 swarm 部署时将忽略该选项)
 
        entrypoint            # 覆盖容器的默认 entrypoint 指令 (等同于 docker run --entrypoint 的作用)
 
        env_file              # 从指定文件中读取变量设置为容器中的环境变量, 可以是单个值或者一个文件列表, 如果多个文件中的变量重名则后面的变量覆盖前面的变量, environment 的值覆盖 env_file 的值
            文件格式：
                RACK_ENV=development 
 
        environment           # 设置环境变量， environment 的值可以覆盖 env_file 的值 (等同于 docker run --env 的作用)
 
        expose                # 暴露端口, 但是不能和宿主机建立映射关系, 类似于 Dockerfile 的 EXPOSE 指令
 
        external_links        # 连接不在 docker-compose.yml 中定义的容器或者不在 compose 管理的容器(docker run 启动的容器, 在 v3 版本中使用 swarm 部署时将忽略该选项)
 
        extra_hosts           # 添加 host 记录到容器中的 /etc/hosts 中 (等同于 docker run --add-host 的作用)
 
        healthcheck           # v2.1 以上版本, 定义容器健康状态检查, 类似于 Dockerfile 的 HEALTHCHECK 指令
            test                  # 检查容器检查状态的命令, 该选项必须是一个字符串或者列表, 第一项必须是 NONE, CMD 或 CMD-SHELL, 如果其是一个字符串则相当于 CMD-SHELL 加该字符串
                NONE                  # 禁用容器的健康状态检测
                CMD                   # test: ["CMD", "curl", "-f", "http://localhost"]
                CMD-SHELL             # test: ["CMD-SHELL", "curl -f http://localhost || exit 1"] 或者　test: curl -f https://localhost || exit 1
            interval: 1m30s       # 每次检查之间的间隔时间
            timeout: 10s          # 运行命令的超时时间
            retries: 3            # 重试次数
            start_period: 40s     # v3.4 以上新增的选项, 定义容器启动时间间隔
            disable: true         # true 或 false, 表示是否禁用健康状态检测和　test: NONE 相同
 
        image                 # 指定 docker 镜像, 可以是远程仓库镜像、本地镜像
 
        init                  # v3.7 中新增的参数, true 或 false 表示是否在容器中运行一个 init, 它接收信号并传递给进程
 
        isolation             # 隔离容器技术, 在 Linux 中仅支持 default 值
 
        labels                # 使用 Docker 标签将元数据添加到容器, 与 Dockerfile 中的 LABELS 类似
 
        links                 # 链接到其它服务中的容器, 该选项是 docker 历史遗留的选项, 目前已被用户自定义网络名称空间取代, 最终有可能被废弃 (在使用 swarm 部署时将忽略该选项)
 
        logging               # 设置容器日志服务
            driver                # 指定日志记录驱动程序, 默认 json-file (等同于 docker run --log-driver 的作用)
            options               # 指定日志的相关参数 (等同于 docker run --log-opt 的作用)
                max-size              # 设置单个日志文件的大小, 当到达这个值后会进行日志滚动操作
                max-file              # 日志文件保留的数量
 
        network_mode          # 指定网络模式 (等同于 docker run --net 的作用, 在使用 swarm 部署时将忽略该选项)         
 
        networks              # 将容器加入指定网络 (等同于 docker network connect 的作用), networks 可以位于 compose 文件顶级键和 services 键的二级键
            aliases               # 同一网络上的容器可以使用服务名称或别名连接到其中一个服务的容器
            ipv4_address      # IP V4 格式
            ipv6_address      # IP V6 格式
 
            示例:
                version: '3.7'
                services: 
                    test: 
                        image: nginx:1.14-alpine
                        container_name: mynginx
                        command: ifconfig
                        networks: 
                            app_net:                                # 调用下面 networks 定义的 app_net 网络
                            ipv4_address: 172.16.238.10
                networks:
                    app_net:
                        driver: bridge
                        ipam:
                            driver: default
                            config:
                                - subnet: 172.16.238.0/24
 
        pid: 'host'           # 共享宿主机的 进程空间(PID)
 
        ports                 # 建立宿主机和容器之间的端口映射关系, ports 支持两种语法格式
            SHORT 语法格式示例:
                - "3000"                            # 暴露容器的 3000 端口, 宿主机的端口由 docker 随机映射一个没有被占用的端口
                - "3000-3005"                       # 暴露容器的 3000 到 3005 端口, 宿主机的端口由 docker 随机映射没有被占用的端口
                - "8000:8000"                       # 容器的 8000 端口和宿主机的 8000 端口建立映射关系
                - "9090-9091:8080-8081"
                - "127.0.0.1:8001:8001"             # 指定映射宿主机的指定地址的
                - "127.0.0.1:5000-5010:5000-5010"   
                - "6060:6060/udp"                   # 指定协议
 
            LONG 语法格式示例:(v3.2 新增的语法格式)
                ports:
                    - target: 80                    # 容器端口
                      published: 8080               # 宿主机端口
                      protocol: tcp                 # 协议类型
                      mode: host                    # host 在每个节点上发布主机端口,  ingress 对于群模式端口进行负载均衡
 
        secrets               # 不知道怎么用
 
        security_opt          # 为每个容器覆盖默认的标签 (在使用 swarm 部署时将忽略该选项)
 
        stop_grace_period     # 指定在发送了 SIGTERM 信号之后, 容器等待多少秒之后退出(默认 10s)
 
        stop_signal           # 指定停止容器发送的信号 (默认为 SIGTERM 相当于 kill PID; SIGKILL 相当于 kill -9 PID; 在使用 swarm 部署时将忽略该选项)
 
        sysctls               # 设置容器中的内核参数 (在使用 swarm 部署时将忽略该选项)
 
        ulimits               # 设置容器的 limit
 
        userns_mode           # 如果Docker守护程序配置了用户名称空间, 则禁用此服务的用户名称空间 (在使用 swarm 部署时将忽略该选项)
 
        volumes               # 定义容器和宿主机的卷映射关系, 其和 networks 一样可以位于 services 键的二级键和 compose 顶级键, 如果需要跨服务间使用则在顶级键定义, 在 services 中引用
            SHORT 语法格式示例:
                volumes:
                    - /var/lib/mysql                # 映射容器内的 /var/lib/mysql 到宿主机的一个随机目录中
                    - /opt/data:/var/lib/mysql      # 映射容器内的 /var/lib/mysql 到宿主机的 /opt/data
                    - ./cache:/tmp/cache            # 映射容器内的 /var/lib/mysql 到宿主机 compose 文件所在的位置
                    - ~/configs:/etc/configs/:ro    # 映射容器宿主机的目录到容器中去, 权限只读
                    - datavolume:/var/lib/mysql     # datavolume 为 volumes 顶级键定义的目录, 在此处直接调用
 
            LONG 语法格式示例:(v3.2 新增的语法格式)
                version: "3.2"
                services:
                    web:
                        image: nginx:alpine
                        ports:
                            - "80:80"
                        volumes:
                            - type: volume                  # mount 的类型, 必须是 bind、volume 或 tmpfs
                                source: mydata              # 宿主机目录
                                target: /data               # 容器目录
                                volume:                     # 配置额外的选项, 其 key 必须和 type 的值相同
                                    nocopy: true                # volume 额外的选项, 在创建卷时禁用从容器复制数据
                            - type: bind                    # volume 模式只指定容器路径即可, 宿主机路径随机生成; bind 需要指定容器和数据机的映射路径
                                source: ./static
                                target: /opt/app/static
                                read_only: true             # 设置文件系统为只读文件系统
                volumes:
                    mydata:                                 # 定义在 volume, 可在所有服务中调用
 
        restart               # 定义容器重启策略(在使用 swarm 部署时将忽略该选项, 在 swarm 使用 restart_policy 代替 restart)
            no                    # 禁止自动重启容器(默认)
            always                # 无论如何容器都会重启
            on-failure            # 当出现 on-failure 报错时, 容器重新启动
 
        其他选项：
            domainname, hostname, ipc, mac_address, privileged, read_only, shm_size, stdin_open, tty, user, working_dir
            上面这些选项都只接受单个值和 docker run 的对应参数类似
 
        对于值为时间的可接受的值：
            2.5s
            10s
            1m30s
            2h32m
            5h34m56s
            时间单位: us, ms, s, m， h
        对于值为大小的可接受的值：
            2b
            1024kb
            2048k
            300m
            1gb
            单位: b, k, m, g 或者 kb, mb, gb
    networks          # 定义 networks 信息
        driver                # 指定网络模式, 大多数情况下, 它 bridge 于单个主机和 overlay Swarm 上
            bridge                # Docker 默认使用 bridge 连接单个主机上的网络
            overlay               # overlay 驱动程序创建一个跨多个节点命名的网络
            host                  # 共享主机网络名称空间(等同于 docker run --net=host)
            none                  # 等同于 docker run --net=none
        driver_opts           # v3.2以上版本, 传递给驱动程序的参数, 这些参数取决于驱动程序
        attachable            # driver 为 overlay 时使用, 如果设置为 true 则除了服务之外，独立容器也可以附加到该网络; 如果独立容器连接到该网络，则它可以与其他 Docker 守护进程连接到的该网络的服务和独立容器进行通信
        ipam                  # 自定义 IPAM 配置. 这是一个具有多个属性的对象, 每个属性都是可选的
            driver                # IPAM 驱动程序, bridge 或者 default
            config                # 配置项
                subnet                # CIDR格式的子网，表示该网络的网段
        external              # 外部网络, 如果设置为 true 则 docker-compose up 不会尝试创建它, 如果它不存在则引发错误
        name                  # v3.5 以上版本, 为此网络设置名称
文件格式示例：
    version: "3"
    services:
      redis:
        image: redis:alpine
        ports:
          - "6379"
        networks:
          - frontend
        deploy:
          replicas: 2
          update_config:
            parallelism: 2
            delay: 10s
          restart_policy:
            condition: on-failure
      db:
        image: postgres:9.4
        volumes:
          - db-data:/var/lib/postgresql/data
        networks:
          - backend
        deploy:
          placement:
            constraints: [node.role == manager]
```

### 5、docker-compose.yml实例

```yaml
vim docker-compose.yml
version: "2"

#services代表几个服务容器实例
services:
### console服务名
    console:
        build:
            context: ./images/console
            args:
                # console 容器 www-data用户密码
                - USERPASS=root
                - GIT_NAME=yangnan
                - GIT_EMAIL=20706149@qq.com
                - INSTALL_YARN=false
        volumes_from:
            - php-fpm
            - nginx
            - mysql
            - redis
        volumes:
            - ./ssh:/home/www-data/.ssh
        links:
            - redis
            - mysql
        tty: true
 
### php-fpm
    php-fpm:
        build: ./images/php-fpm
        volumes:
            - ./app/:/var/www/
 
### nginx
    nginx:
        image: nginx
        ports:
            - "8081:80"
        volumes_from:
            - php-fpm
        volumes:
            - ./logs/nginx/:/var/log/nginx/
            - ./images/nginx/sites:/etc/nginx/conf.d/
        links:
            - php-fpm
 
### mysql
    mysql:
        image: mysql:5.7
        ports:
            - "7706:3306"
        environment:
            MYSQL_ROOT_PASSWORD: "123456"
            MYSQL_ALLOW_EMPTY_PASSWORD: "no"
            MYSQL_DATABASE: "test"
            MYSQL_USER: "root"
            MYSQL_PASSWORD: "123"
        volumes:
            - ./data/mysql:/var/lib/mysql
        networks:
      		- backend
      	command: --default-authentication-plugin=mysql_native_password #解决外部无法访问
 
### redis
    redis:
        image: redis
        ports:
            - "6379:6379"
        volumes:
            - ./data/redis:/data
```

### 6、docker-compose官方实例

```yaml
version: "3.9"
services:

  redis:
    image: redis:alpine
    ports:
      - "6379"
    networks:
      - frontend
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  db:
    image: postgres:9.4
    volumes:
      - db-data:/var/lib/postgresql/data
    networks:
      - backend
    deploy:
      placement:
        max_replicas_per_node: 1
        constraints:
          - "node.role==manager"

  vote:
    image: dockersamples/examplevotingapp_vote:before
    container_name: "vote01" #容器名
    ports:
      - "5000:80"
    networks:
      - frontend
    depends_on: #依赖于
      - redis
    deploy:
      replicas: 2
      update_config:
        parallelism: 2
      restart_policy:
        condition: on-failure

  result:
    image: dockersamples/examplevotingapp_result:before
    ports:
      - "5001:80"
    networks:
      - backend
    depends_on:
      - db
    deploy:
      replicas: 1
      update_config:
        parallelism: 2
        delay: 10s
      restart_policy:
        condition: on-failure

  worker:
    image: dockersamples/examplevotingapp_worker
    networks:
      - frontend
      - backend
    deploy:
      mode: replicated
      replicas: 1
      labels: [APP=VOTING]
      restart_policy:
        condition: on-failure
        delay: 10s
        max_attempts: 3
        window: 120s
      placement:
        constraints:
          - "node.role==manager"

  visualizer:
    image: dockersamples/visualizer:stable
    ports:
      - "8080:8080"
    stop_grace_period: 1m30s
    volumes:
      - "/var/run/docker.sock:/var/run/docker.sock"
    deploy:
      placement:
        constraints:
          - "node.role==manager"

networks: #docker network create 网络
  frontend:
  backend:

volumes:
  db-data:
```

### 7、实例二

```yaml
version: "3.0"
services:
  mysql: # mysql服务
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password #解决外部无法访问
    ports:
      - "3306:3306"      #容器端口映射到宿主机的端口
    environment:
      MYSQL_ROOT_PASSWORD: 'root'
      MYSQL_ALLOW_EMPTY_PASSWORD: 'no'
      MYSQL_DATABASE: 'docker-compose-boot'
      MYSQL_USER: 'gblfy'
      MYSQL_PASSWORD: '123456'
    networks:
      - gblfy_web
  docker-compose-boot-web: #自己单独的springboot项目
    hostname: gblfy
    build: ./     #需要构建的Dockerfile文件
    ports:
      - "38000:8080"      #容器端口映射到宿主机的端口
    depends_on:      #web服务依赖mysql服务，要等mysql服务先启动
      - mysql
    networks:
      - gblfy_web
networks:  ## 定义服务的桥
  gblfy_web:
```

### 8、实例三

```yaml
version: '3'    #指定compose版本

services:
  log:    #服务名称
    image: vmware/harbor-log    #指定镜像名称
    container_name: harbor-log  #启动后的容器名称
    restart: always    #down掉自动重启
    volumes:    #宿主机和容器关联的目录
      - /var/log/harbor/:/var/log/docker/
    ports:    #映射出来的端口
      - 1514:514

  registry:
    image: library/registry:2.5.0
    container_name: registry
    restart: always
    volumes:
      - /data/registry:/storage
      - ./common/config/registry/:/etc/registry/
    environment:    #设置环境变量
      - GODEBUG=netdns=cgo
    command:    #容器内执行命令
      ["serve", "/etc/registry/config.yml"]
    depends_on:    #依赖关系
      - log
    logging:    #日志设置
      driver: "syslog"    #指定日志设备的容器
      options:
        syslog-address: "tcp://127.0.0.1:1514" #日志连接地址
        tag: "registry"    #日志标签
```

### 9、简易测试版

```yaml
version: "3"
services:
  redis:
    image: redis:6.2.10
    container_name: "redis01"
    ports:
      - "6379:6379"
    networks:
      - frontend
      
  mysql:
    image: mysql:5.7
    command: --default-authentication-plugin=mysql_native_password
    container_name: "mysql"
    ports:
      - "3306:3306"
    environment:
      MYSQL_ROOT_PASSWORD: 'root'
      MYSQL_ALLOW_EMPTY_PASSWORD: 'no'
      MYSQL_DATABASE: 'docker-compose-boot'
      MYSQL_USER: 'gblfy'
      MYSQL_PASSWORD: '123456'
    networks:
      - frontend
      
networks:
  frontend:
```

### 9.3案例四

```yaml
version: "3.5"
services:
  web:
    build: .
    command: python test.py
    port: 
      - target: 5000
        published: 5000
    networks:
      - counter-net
    volumes:
    - type: volume
      source: test-vol
      target: /code
  redis:
    image: "redis:alpine"
    networks:
      test-net:
        
networks:
  test-net:

volumes:
  test-vol: 
```

```shell
version：必须指定项，且必须位于首行。它定义了整个文件的格式。建议保持最新版本。
services：用于定义不同的服务，如上我们定义了两个服务分别是web和redis。Docker Compose会将他们分别部署在各自的容器内。
networks：用来创建一个网络，如果不写的话，Docker Compose会默认创建bridge网络，但这是单机网络，只能实现同一主机上容器的连接。
volumes：用于创建新的卷。主要是实现数据持久化的。


build：”.“ 表示基于当前目录下Dockerfile中定义的指令来构建一个新镜像。
command：python test.py是指定在Docker容器中执行test.py的Python脚本。所以镜像中得存在test.py文件和Python
ports：用于指定端口映射得，将容器内（target这一行）5000端口映射到主机的（published这一行）5000端口。这样访问主机5000端口的流量都会被转到容器内的5000端口。
networks：告知Docker将服务连接到指定的网络，这个网络要么是存在的，要么是networks中一级key指定的网络。
volumes：用于指定Docker将test-vol卷挂载到容器内的/code。test-vol卷不存在的话，会用volumes下面一级key定义的。
https://juejin.cn/post/7200679596122783805
```



# CIG容器监控系统

### 1、下载安装

CAdvisor + InfluxDB + Granfana

```yaml
version: "3.9"
volumes:
  grafana_data: {}

services:
  influxdb:
    image: tutum/influxdb:0.13
    restart: always
    environment:
      - PRE_CREATE_DB=cadvisor
    ports:
      - "8083:8083"
      - "8086:8086"
    volumes:
      - ./data/influxdb:/data

  cadvisor:
    image: google/cadvisor:v0.33.0
    links:
      - influxdb:influxsrv
    command:
    - storage_driver=influxdb
    - storage_driver_db=cadvisor
    - storage_driver_host=influxsrv:8086
    restart: always
    ports:
      - "8080:8080"
    volumes:
      - /:/rootfs:ro
      - /var/run:/var/run:rw
      - /sys:/sys:ro
      - /var/lib/docker:/var/lib/docker:ro

  grafana:
    user: "104"
    image: grafana/grafana:8.5.16
    restart: always
    links:
      - influxdb:influxsrv
    ports:
      - "3000:3000"
    volumes:
      - grafana_data:/var/lib/grafana
    environment:
      - HTTP_USER=admin
      - HTTP_PASS=admin
      - INFLUXDB_HOST=influxsrv
      - INFLUXDB_PORT=8088
      - INFLUXDB_NAME=cadvisor
      - INFLUXDB_USER=root
      - INFLUXDB_PASS=root
```

### 2、测试启动

```shell
浏览cAdvisor收集服务http://ip:8080/
浏览influxdb储存服务http://ip:8083/
浏览grafana展现服务http://ip:3000/
```

# postgresql数据库安装

### 一、镜像安装

1、安装镜像

```undefined
docker pull postgres:11
```

2、新建目录

```bash
mkdir -p /home/apps/postgres/{postgresql,data}
```

3、创建并启动

```diff
docker run -d --name postgres -p 5432:5432 \
-v /home/apps/postgres/postgresql:/var/lib/postgresql \
-v /home/apps/postgres/data:/var/lib/postgresql/data \
-v /etc/localtime:/etc/localtime:ro \
-e POSTGRES_USER=root \
-e POSTGRES_PASSWORD=123456 \
-e POSTGRES_DB=postgres \
-e TZ=Asia/Shanghai \
--restart always \
--privileged=true \
postgres:11
```

4、postgres基本操作

```python
# 进入docker容器
docker exec -it postgres /bin/bash

# 用户登录(sonar)
psql -U username

# 创建新用户
create user admin with password '123456';

# 创建数据库，指定用户
create database testDB with owner admin;

# 退出
\q

# 查看用户
\du

# 列出数据库
\l

# 删除用户
drop user admin;

# 删除数据库
drop database dbtest;
```

# Harbor仓库搭建及其使用

### 二、安装docker-compose

1、下载docker-compose文件

官网地址：https://get.daocloud.io/

```ruby
curl -L "https://get.daocloud.io/docker/compose/releases/download/v2.15.1/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
```

```bash
#这个命令是一键安装docker以及dockercompose
curl -L https://get.daocloud.io/docker/compose/releases/download/v2.0.1/docker-compose-`uname -s`-`uname -m` > /usr/local/bin/docker-compose

修改daemon.json文件

解决docker login 时，提示https问题
添加insecure-registries

#1.创建一个目录
sudo mkdir -p /etc/docker

#2.编写配置文件
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["http://hub-mirror.c.163.com"],
  "insecure-registries":["192.168.222.145:6007"]
}
EOF

#3.重启服务
sudo systemctl daemon-reload
sudo systemctl restart docker
```

2、为docker-compose文件赋予可执行权限

```bash
chmod +x /usr/local/bin/docker-compose
```

- ### 三、安装harbor

  1、下载地址

  - 官网地址：https://github.com/goharbor/harbor/releases
  - 版本：harbor-offline-installer-v2.5.0.tgz

  ```ruby
  wget https://github.com/goharbor/harbor/releases/download/v2.5.0/harbor-offline-installer-v2.5.0.tgz
  ```

  2、解压文件

  ```bash
  tar -zxvf harbor-offline-installer-v2.5.0.tgz -C /usr/local
  
  # 复制配置文件
  cd /usr/local/harbor
  cp harbor.yml.tmpl harbor.yml
  ```

  3、修改harbor.yml配置

  ```yaml
  vim harbor.yml
  
  # 修改如下配置
  hostname: 192.168.222.145
  http:
    port: 6007
  harbor_admin_password: Harbor12345
  ```

  - hostname 这里设置本机的ip
  - harbor_admin_password 页面的密码
  - 注释掉https部分

![img](https://img2020.cnblogs.com/blog/364454/202111/364454-20211102174633061-170913744.png)

4、配置数据库

- 如果你需要将数据存储在数据库中，请按以下配置（非必选）
- harbor默认会将数据存储在本地，为了提高数据的安全性，可将数据存储在数据库中
- 目前Harbor仅支持PostgraSQL数据库，需要手动创建harbor、notary_signer、notary_servers三个数据库，Harbor启动时会自动在对应数据库下生成表
- postgresql数据库安装：https://www.cnblogs.com/lvlinguang/p/16583405.html

```yaml
vim harbor.yml

#注释掉本地数据库
#database:
  # The password for the root user of Harbor DB. Change this before any production use.
 # password: root123
  # The maximum number of connections in the idle connection pool. If it <=0, no idle connections are retained.
  #max_idle_conns: 50
  # The maximum number of open connections to the database. If it <= 0, then there is no limit on the number of open connections.
  # Note: the default number of connections is 1024 for postgres of harbor.
  #max_open_conns: 1000

#修改挂载目录
data_volume: /data/harbor

# 配置数据库
external_database:
  harbor:
    host: 127.0.0.1
    port: 5432
    db_name: harbor
    username: root
    password: 123456
    ssl_mode: disable
    max_idle_conns: 50
    max_open_conns: 100
  notary_signer:
    host: 127.0.0.1
    port: 5432
    db_name: notary_signer
    username: root
    password: 123456
    ssl_mode: disable
  notary_server:
    host: 127.0.0.1
    port: 5432
    db_name: notary_server
    username: root
    password: 123456
    ssl_mode: disable
    
external_redis:
  host: 127.0.0.1:6379
  password: 123456
  registry_db_index: 1
  jobservice_db_index: 2
  chartmuseum_db_index: 3
  trivy_db_index: 5
  idle_timeout_seconds: 30
```

6、运行/安装

```bash
./prepare
./install.sh
```

7、访问页面
[http://192.168.3.12:6007](http://192.168.3.12:6007/)

- 帐号admin，密码Harbor12345
