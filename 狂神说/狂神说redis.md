### Linux下载

1、下载安装包 `redis-7.0.5.tar.gz`

2、解压`redis`的安装包

```shell
#一般文件放在 /opt 目录下
tar -zxvf redis-7.0.5.tar.gz
```

3、进入`redis`文件目录

```shell
yum install -y gcc-c++
gcc -v
make 
make install
```

4、redis默认安装路径在 `/usr/local/bin`

5、将 `redis`配置文件，复制到我们当前目录下,以后使用这个配置文件

```shell
cd /usr/local/bin
mkdir config
cp /opt/redis-7.0.5/redis.conf ./config
```

6、redis默认不是后台启动的，修改配置文件

```shell
vim redis.conf
#使用/快速查找
daemonize no ---> yes
bind 0.0.0.0
requirepass 123456
```

7、启动`redis`服务

```shell
#回到/usr/local/bin目录
cd ..
redis-server ./config/redis.conf
redis-cli启动客户端
使用ping测试是否联通
```

8、查看`redis`进程

```shell
ps -ef | grep redis
```

9、关闭`shutdown`

```shell
redis-cli启动客户端里面：shutdown
```

10、压测工具`redis-benchmark`

```shell
-h <hostname> 指定服务器主机名 (默认 127.0.0.1)
 -p <port> 指定服务器端口 (默认 6379)
 -s <socket> 指定服务器 socket
 -a <password> Redis 认证密码
 -c <clients> 指定并发连接数 (默认 50)
 -n <requests> 指定请求数 (默认 100000)
 -d <size> 以字节的形式指定 SET/GET 值的数据大小 (默认 2)
 --dbnum <db> 选择指定的数据库号 (默认 0)
 -k <boolean> 1=keep alive 0=reconnect (默认 1)
 -r <keyspacelen> SET/GET/INCR 使用随机 key, SADD 使用随机值
 -P <numreq> 通过管道传输 <numreq> 请求 (no pipeline)
 -q 退出，仅显示 query/sec 值
 --csv 以 CSV 格式输出
 -l 生成循环，永久执行测试
 -t <tests> 仅运行以逗号分隔的测试命令列表
 -I Idle 模式，仅打开 N 个 idle 连接并等待
 
 #对127.0.0.1使用20个并行客户端，总共10万个请求
 在`/usr/local/bin`目录下
 ./redis-benchmark -h 127.0.0.1 -p 6379 -n 100000 -c 20 -q
```

![img](https://img-blog.csdnimg.cn/ed4c4b48d8324cd6a6507b42623b963e.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAwrfmooXoirHljYHkuIk=,size_20,color_FFFFFF,t_70,g_se,x_16)

```shell
./redis-benchmark -h localhost -p 6369 -c 100 -n 100000 
```

![img](https://img-blog.csdnimg.cn/7b3a9b117b694e8a988dc9d6d7a995bf.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAwrfmooXoirHljYHkuIk=,size_20,color_FFFFFF,t_70,g_se,x_16)

![img](https://img-blog.csdnimg.cn/6e41f545e90843a69fef3e95cba54cd0.png?x-oss-process=image/watermark,type_ZHJvaWRzYW5zZmFsbGJhY2s,shadow_50,text_Q1NETiBAwrfmooXoirHljYHkuIk=,size_20,color_FFFFFF,t_70,g_se,x_16)

### Redis基本命令

```shell
#127.0.0.1:6379> keys * 查看数据库所有key
#127.0.0.1:6379> select 0 切换数据库
#127.0.0.1:6379> DBSIZE 查看DB大小
#127.0.0.1:6379> flushdb 清除当前数据库
#127.0.0.1:6379> FLUSHALL 清除全部数据库
redis是基于内存的，CPU不是瓶颈，是根据机器的内存和网络带宽，所有就使用单线程，官方提供的数据为100000+QPS
```

### 五大数据类型

命令参考：http://redisdoc.com/index.html

官网命令：https://redis.io/commands/

###### Redis-Key

```shell
#127.0.0.1:6379>set key value 设置
#127.0.0.1:6379>get key
#127.0.0.1:6379>exists key 判断是否存在
#127.0.0.1:6379>move key 1 状态设置为1，移除key
#127.0.0.1:6379>expire key time 设置有效期,以秒为单位
#127.0.0.1:6379>ttl key 查看有效期
#127.0.0.1:6379>type key 查看当前key的一个类型
```

###### String字符串

```shell
#127.0.0.1:6379>append key value 追加字符串
#127.0.0.1:6379>strlen key 获取字符串长度
#127.0.0.1:6379>incr key  自增
#127.0.0.1:6379>decr key  自减
#127.0.0.1:6379>incrby key +数字 步长
#127.0.0.1:6379>decrby key -数字
#127.0.0.1:6379>getrange key start end 切片
#127.0.0.1:6379>mset key1 value1 key2 value2...
#127.0.0.1:6379>mget key1 key2...
#127.0.0.1:6379>setrange key 索引 value 替换索引对应的值
#127.0.0.1:6379>setex key time value 设置过期时间
#127.0.0.1:6379>setnx key value key不存在创建，在创建失败,在分布式锁中会常常使用,是一个原子性操作，要么一起成功，要么一起失败
#127.0.0.1:6379>set user:1 {"id":1,"name":"tom"}
#127.0.0.1:6379>met user:1:name jerry user:1:age 18
#127.0.0.1:6379>getset db redis 先get后set,如果不存在值，则返回nil，存在值就更改原来的值
```

###### List数组

```shell
#127.0.0.1:6379>左边进左边拿就是栈，左边进右边拿就是队列
#127.0.0.1:6379>是一个链表，左右两边都可以插入值，key不存在创建新的链表，存在就新增内容
#127.0.0.1:6379>lset key value
#127.0.0.1:6379>lget key
#127.0.0.1:6379>lpush key value 往左边添加值，头部
#127.0.0.1:6379>lrange key start end 倒序拿数据
#127.0.0.1:6379>rpush key value 往右边添加值，尾部
#127.0.0.1:6379>lpop key 移除左边第一个元素，也就是第一个元素
#127.0.0.1:6379>rpop key 移除右边第一个元素，也就是最后一个元素
#127.0.0.1:6379>lrem key count value 移除几个元素
#127.0.0.1:6379>trim key start end 修剪,通过下标截取指定的长度
#127.0.0.1:6379>rpoppush key value 移除列表的最后一个元素，将它移动到新的列表中
#127.0.0.1:6379>lset key index value 替换下标的值
#127.0.0.1:6379>linsert key before/after value newvalue在什么值前面或后面插入值
```

###### Set无序集合

```shell
#127.0.0.1:6379>set中的值不能重复
#127.0.0.1:6379>sadd key value 添加值
#127.0.0.1:6379>smembers key 查看
#127.0.0.1:6379>sismebere key value 判断值是否在集合中存在
#127.0.0.1:6379>scard key 获取元素个数
#127.0.0.1:6379>srem key value 移除
#127.0.0.1:6379>srandmember key count 随机读取某几个值,count可以不填
#127.0.0.1:6379>spop key count 随机删除值,,count可以不填
#127.0.0.1:6379>smove key1 key2 value 把集合中的指定值移动到另一个集合里
#127.0.0.1:6379>sunion key1 key2 并集，合并两个集合
#127.0.0.1:6379>sinter key1 key2 交集，共同的值
#127.0.0.1:6379>sdiff key1 key2 差集,不同的值
```

###### Hash键值对

```shell
#127.0.0.1:6379>hash容易变更数据，常存储用户信息类，经常变动的信息
#127.0.0.1:6379>hset key field value
#127.0.0.1:6379>hset user:1 name tom 举例
#127.0.0.1:6379>hget key field
#127.0.0.1:6379>hmset key field1 value1 field2 value2
#127.0.0.1:6379>hmget key field1 field2
#127.0.0.1:6379>hgetall key 获取所有的值
#127.0.0.1:6379>hdel key field 删除指定的key，也删除可以对应的value
#127.0.0.1:6379>hexists key field 判断指定的字段是否存在
#127.0.0.1:6379>hkeys key 只获取所有的field
#127.0.0.1:6379>hvals key 只获取所有的值
#127.0.0.1:6379>hincrby key field count
#127.0.0.1:6379>hdecrby key field count
#127.0.0.1:6379>hsetnx key field value 如果不存在就可以设置
```

###### Zset有序集合

```shell
#127.0.0.1:6379>排序
#127.0.0.1:6379>zadd key number value
#127.0.0.1:6379>zrangebyscore key min max 排序(-inf,+inf正负无穷)
#127.0.0.1:6379>ZRANGEBYSCORE key -inf +inf WITHSCORES # 显示整个有序集及成员的 score值
#127.0.0.1:6379>zrem key value 移除
#127.0.0.1:6379>ZRANGE key 1 2 # 显示有序集下标区间 1 至 2 的值的个数
#127.0.0.1:6379>ZCARD key 自增
#127.0.0.1:6379>EXISTS non_exists_key   # 对不存在的 key 进行 ZCARD 操作
#127.0.0.1:6379>ZRANGE key min max # 递增排列
#127.0.0.1:6379>ZREVRANGE key min max # 递减排列
#127.0.0.1:6379>ZCOUNT key min max #计算在min max之间的数量
```

### 三种特殊数据类型

###### geospatial地理空间

```shell
geoadd key value(纬度、经度、名字)
有效的经度介于 -180 度至 180 度之间。
有效的纬度介于 -85.05112878 度至 85.05112878 度之间。
#127.0.0.1:6379>geoadd china:city 13.361389 38.115556 beijing  15.087269 37.502669 tianjing 添加数据

#127.0.0.1:6379>geodist china:city beijing tianjing km 返回两个给定位置之间的距离
m:米 km:千米 mi:英里 ft:英尺 如果两个位置之间的其中一个不存在，那么命令返回空值
#127.0.0.1:6379>geopos china:city beijing tianjing 从key里返回所有给定元素的位置

#127.0.0.1:6379>georadius key 13 38 半径(1000) km 获取当前经纬度为中心的1000公里的数据
WITHDIST ： 在返回位置元素的同时， 将位置元素与中心之间的距离也一并返回。
WITHCOORD ： 将位置元素的经度和维度也一并返回。
ASC ： 根据中心的位置， 按照从近到远的方式返回位置元素。
DESC ： 根据中心的位置， 按照从远到近的方式返回位置元素。
#127.0.0.1:6379>GEORADIUS Sicily 15 37 200 km WITHDIST/WITHCOORD count number(可忽略)

#127.0.0.1:6379>GEORADIUSBYMEMBER key value 100 km 找出指定key范围内的元素

#127.0.0.1:6379>geohash key value1 value2 返回一个或多个位置元素的hash,该命令会返回11个字符的geohash字符，将二维转换为一维，如果两个字符串越接近，那么距离越近
redis> GEOHASH Sicily Palermo Catania
1) "sqc8b49rny0"
2) "sqdtr74hyu0"
GEO底层的实现原理其实都是zset，可以使用zset命令来操作geo
zrem key value
zrange key 0 -1
```

###### hyperloglog基数统计

基数(不重复的数),hyperloglog基数统计的算法

```shell
#127.0.0.1:6379>PFADD key value value1... 添加元素
#127.0.0.1:6379>PFADD databases "Redis" "MongoDB" "MySQL"
#127.0.0.1:6379>PFADD databases "Redis" # Redis 已经存在，不必对估计数量进行更新

#127.0.0.1:6379>PFCOUNT key 统计元素个数
#127.0.0.1:6379>PFCOUNT databases

#127.0.0.1:6379>PFADD nosql "Redis" "MongoDB" "Memcached"
#127.0.0.1:6379>PFADD RDBMS "MySQL" "MSSQL" "PostgreSQL"
#127.0.0.1:6379>PFMERGE databases nosql RDBMS 合并为一个新的列表
#127.0.0.1:6379>PFCOUNT databases
```

###### bitmaps位图

数据结构，都是操作二进制位来进行记录，就只有0和1 两个状态

```shell
#127.0.0.1:6379>SETBIT key offset value 设置
#127.0.0.1:6379>SETBIT bit 10086 1/0
#127.0.0.1:6379>GETBIT bit 10086 获取
#127.0.0.1:6379>GETBIT bit 100   # bit 默认被初始化为 0
#127.0.0.1:6379>BITCOUNT bits 统计
```

### 事务

事务本质:一组命令的集合，一个事务中的所有命令都会被序列化，在事务执行过程中，会按照顺序执行,单条命令保存原子性，事务不保存，一次性，顺序性，排他性，执行一些列的命令

```shell
redis> MULTI            # 标记事务开始
OK
redis> INCR user_id     # 多条命令按顺序入队
QUEUED
redis> INCR user_id     
QUEUED
redis> PING
QUEUED
redis> EXEC             # 执行
1) (integer) 1
2) (integer) 2
3) PONG
```

放弃/取消事务

```shell
redis> MULTI
OK
redis> SET greeting "hello"
QUEUED
redis> DISCARD
OK
```

### Redis乐观锁

- 悲观锁：很悲观，认为什么时候都会出问题，无论做什么都会加锁
- 乐观锁：很乐观，认为什么时候都不会出问题，所以不会上锁，更新数据的时候去判断一下，在此期间是否有人修改过这个数据，获取version，更新的时候比较version

```shell
127.0.0.1:6379> set money 100
OK
127.0.0.1:6379> watch money #监视 money 对象
OK
127.0.0.1:6379> multi #事务正常结束，数据期间没有发生变动，这个时候就正常执行成功
OK
127.0.0.1:6379> incrby money 20
QUEUED
127.0.0.1:6379> exec
1) (integer) 120
127.0.0.1:6379>
```

测试多线程修改值，使用watch可以当做redis的乐观锁操作

执行之前，另外一个线程，修改了我们的值，这个时候，就是导致事务执行失败

```shell
redis> WATCH key1 key2 ...
OK
redis> UNWATCH #发现失败执行失败，就先解锁
OK
https://www.cnblogs.com/jasonZh/p/9522772.html
```

### Redis.conf 配置文件

```shell
bind 0.0.0.0 #监听地址，可以用空格隔开后多个监听IP

protected-mode yes #redis3.2 之后加入的新特性，在没有设置bind IP和密码的时候,redis只允许访问

127.0.0.1:6379，可以远程连接，但当访问将提示警告信息并拒绝远程访问

port 6379 #监听端口

tcp-backlog 511 #三次握手的时候server端收到client ack确认号之后的队列值，即全队列长度

timeout 0 #客户端和Redis服务端的连接超时时间，默认是0，表示永不超时

tcp-keepalive 300 #tcp 会话保持时间300s

daemonize no #默认redis-server不作为守护进程运行的，而前台运行，如果想在后台运行，就把它改成 yes,当redis作为守护进程运行的时候，它会写一个 pid 到 /var/run/redis.pid 文件

supervised no #和OS相关参数，可设置通过upstart和systemd管理Redis守护进程，centos7后都使用systemd

pidfile /var/run/redis_6379.pid #pid文件路径

loglevel notice #日志级别

logfile "/path/redis.log" #日志路径

databases 16 #设置数据库数量，默认：0-15，共16个库

always-show-logo yes #在启动redis 时是否显示redis的logo

save 900 1 #在900秒内有一个键内容发生更改就出就快照机制
save 300 10
save 60 10000  #60秒内如果有10000个健以上的变化，就自动快照备份

stop-writes-on-bgsave-error yes #yes时因空间满等原因快照无法保存出错时，禁止redis写入操作，建议为no

rdbcompression yes #持久化到RDB文件时，是否压缩，"yes"为压缩，"no"则反之

rdbchecksum yes #是否对备份文件开启RC64校验，默认是开启

dbfilename dump.rdb #快照文件名

dir ./ #快照文件保存路径，示例：dir "/apps/redis/data"

# replicaof <masterip> <masterport>  #指定复制的master主机地址和端口，5.0版之前的指令为slaveof 
# masterauth <master-password> #指定复制的master主机的密码

replica-serve-stale-data yes #当从库同主库失去连接或者复制正在进行，从机库有两种运行方式：
1、设置为yes(默认设置)，从库会继续响应客户端的读请求，此为建议值
2、设置为no，除去指定的命令之外的任何请求都会返回一个错误"SYNC with master in progress"。

replica-read-only yes #是否设置从库只读，建议值为yes,否则主库同步从库时可能会覆盖数据，造成数据丢失

repl-diskless-sync no #是否使用socket方式复制数据(无盘同步)，新slave连接连接时候需要做数据的全量同步，redis server就要从内存dump出新的RDB文件，然后从master传到slave，有两种方式把RDB文件传输给客户端：
1、基于硬盘（disk-backed）：为no时，master创建一个新进程dump生成RDB磁盘文件，RDB完成之后由父进程（即主进程）将RDB文件发送给slaves，此为推荐值
2、基于socket（diskless）：master创建一个新进程直接dump RDB至slave的网络socket，不经过主进程和硬盘

基于硬盘（为no），RDB文件创建后，一旦创建完毕，可以同时服务更多的slave，但是基于socket(为yes)， 新slave连接到master之后得逐个同步数据。当磁盘I/O较慢且网络较快时，可用diskless(yes),否则使用磁盘(no)

repl-diskless-sync-delay 5 #diskless时复制的服务器等待的延迟时间，设置0为关闭，在延迟时间内到达的客户端，会一起通过diskless方式同步数据，但是一旦复制开始，master节点不会再接收新slave的复制请求，直到下一次同步开始才再接收新请求。即无法为延迟时间后到达的新副本提供服务，新副本将排队等待下一次RDB传输，因此服务器会等待一段时间才能让更多副本到达。推荐值：30-60

repl-ping-replica-period 10 #slave根据master指定的时间进行周期性的PING master 监测master状态

repl-timeout 60 #复制连接的超时时间，需要大于repl-ping-slave-period，否则会经常报超时

repl-disable-tcp-nodelay no #是否在slave套接字发送SYNC之后禁用 TCP_NODELAY，如果选择"yes"，Redis将合并多个报文为一个大的报文，从而使用更少数量的包向slaves发送数据，但是将使数据传输到slave上有延迟，Linux内核的默认配置会达到40毫秒，如果 "no" ，数据传输到slave的延迟将会减少，但要使用更多的带宽

repl-backlog-size 512mb #复制缓冲区内存大小，当slave断开连接一段时间后，该缓冲区会累积复制副本数据，因此当slave 重新连接时，通常不需要完全重新同步，只需传递在副本中的断开连接后没有同步的部分数据即可。只有在至少有一个slave连接之后才分配此内存空间。

repl-backlog-ttl 3600 #多长时间内master没有slave连接，就清空backlog缓冲区

replica-priority 100 #当master不可用，Sentinel会根据slave的优先级选举一个master，此值最低的slave会当选master，而配置成0，永远不会被选举，一般多个slave都设为一样的值，让其自动选择

#min-replicas-to-write 3  #至少有3个可连接的slave，mater才接受写操作
#min-replicas-max-lag 10  #和上面至少3个slave的ping延迟不能超过10秒，否则master也将停止写操作

requirepass foobared #设置redis 连接密码，如果有特殊符号，用" "引起来

rename-command #重命名一些高危命令，示例：rename-command FLUSHALL "" 禁用命令

maxclients 10000 #Redis最大连接客户端

maxmemory #redis使用的最大内存，单位为bytes字节，0为不限制，建议设为物理内存一半，8G内存的计算方式8(G)*1024(MB)1024(KB)*1024(Kbyte)，需要注意的是缓冲区是不计算在maxmemory内。

appendonly no #是否开启AOF日志记录，默认redis使用的是rdb方式持久化，这种方式在许多应用中已经足够用了，但是redis如果中途宕机，会导致可能有几分钟的数据丢失(取决于dumpd数据的间隔时间)，根据save来策略进行持久化，Append Only File是另一种持久化方式，可以提供更好的持久化特性，Redis会把每次写入的数据在接收后都写入 appendonly.aof 文件，每次启动时Redis都会先把这个文件的数据读入内存里，先忽略RDB文件。默认不启用此功能,改成yes表示开启

appendfilename "appendonly.aof" #AOF文件名，是文本文件，存放在dir指令指定的目录中
appendfsync everysec #aof持久化策略的配置,no表示不执行fsync，由操作系统保证数据同步到磁盘,always表示每次写入都执行fsync，以保证数据同步到磁盘,everysec表示每秒执行一次fsync，可能会导致丢失这1s数据。

no-appendfsync-on-rewrite no #在aof rewrite期间,是否对aof新记录的append暂缓使用文件同步策略,主要考虑磁盘IO开支和请求阻塞时间。默认为no,表示"不暂缓",新的aof记录仍然会被立即同步，Linux的默认fsync策略是30秒，如果为yes 可能丢失30秒数据，但由于yes性能较好而且会避免出现阻塞因此比较推荐。

auto-aof-rewrite-percentage 100 # 当Aof log增长超过指定百分比例时，重写AOF文件， 设置为0表示不自动重写Aof 日志，重写是为了使aof体积保持最小，但是还可以确保保存最完整的数据

auto-aof-rewrite-min-size 64mb #触发aof rewrite的最小文件大小

aof-load-truncated yes #是否加载由于其他原因导致的末尾异常的AOF文件(主进程被kill/断电等)，建议yes

aof-use-rdb-preamble yes #redis4.0新增RDB-AOF混合持久化格式，在开启了这个功能之后，AOF重写产生的文件将同时包含RDB格式的内容和AOF格式的内容，其中RDB格式的内容用于记录已有的数据，而AOF格式的内存则用于记录最近发生了变化的数据，这样Redis就可以同时兼有RDB持久化和AOF持久化的优点（既能够快速地生成重写文件，也能够在出现问题时，快速地载入数据）。

lua-time-limit 5000 #lua脚本的最大执行时间，单位为毫秒

cluster-enabled yes #是否开启集群模式，默认是单机模式

cluster-config-file nodes-6379.conf #由node节点自动生成的集群配置文件名称

cluster-node-timeout 15000 #集群中node节点连接超时时间，超过此时间，会踢出集群

cluster-replica-validity-factor 10 #在执行故障转移的时候可能有些节点和master断开一段时间数据比较旧，这些节点就不适用于选举为master，超过这个时间的就不会被进行故障转移，计算公式：(node-timeout * replica-validity-factor) + repl-ping-replica-period 

cluster-migration-barrier 1 #集群迁移屏障，一个主节点至少拥有一个正常工作的从节点，即如果主节点的slave节点故障后会将多余的从节点分配到当前主节点成为其新的从节点。

cluster-require-full-coverage yes #集群请求槽位全部覆盖，如果一个主库宕机且没有备库就会出现集群槽位不全，那么yes情况下redis集群槽位验证不全就不再对外提供服务，而no则可以继续使用但是会出现查询数据查不到的情况(因为有数据丢失)。建议为no

cluster-replica-no-failover no #如果为yes,此选项阻止在主服务器发生故障时尝试对其主服务器进行故障转移。 但是，主服务器仍然可以执行手动强制故障转移，一般为no

#Slow log 是 Redis 用来记录超过指定执行时间的日志系统， 执行时间不包括与客户端交谈，发送回复等I/O操作，而是实际执行命令所需的时间（在该阶段线程被阻塞并且不能同时为其它请求提供服务）slow log 保存在内存里面，读写速度非常快，因此可放心地使用，不必担心因为开启 slow log 而影响 Redis 的速度

slowlog-log-slower-than 10000 #以微秒为单位的慢日志记录，为负数会禁用慢日志，为0会记录每个命令操作。

slowlog-max-len 128 #最多记录多少条慢日志的保存队列长度，达到此长度后，记录新命令会将最旧的命令从命令队列中删除，以此滚动删除

127.0.0.1:6379> SLOWLOG len  #查看慢日志的记录条数
(integer) 14
127.0.0.1:6379> slowlog get  #查看慢日志的记录
1) 1) (integer) 14
2) (integer) 1544690617
3) (integer) 4
4) 1) "slowlog"
127.0.0.1:6379> SLOWLOG reset #清空慢日志
OK
```

### Redis 持久化RDB

在指定的时间间隔内将内存中的数据集快照写入磁盘，也就是行话讲的Snapshot快照，它恢复时将快照文件直接读到内存里，redis会单独创建(fork)一个子进程来进行持久化，会先将数据写入到一个临时文件中，待持久化过程都结束了，再用这个临时文件替换上次持久化的文件，整个过程中，主进程是不进行任何IO操作的，这就确保了极高的性能，如果需要进行大规模数据的恢复，且对于数据恢复的完整性不是非常敏感，那RDB方式要比AOF方式更加的高效，RDB的缺点是最后一次持久化后的数据可能丢失，我们默认的就是RDB，一般情况下不需要修改这个配置

```shell
#RDB保存的文件是dump.rdb，都是在我们配置文件中快照中进行配置的
#dbfilename dump.rdb
如何恢复rdb文件，只需要将rdb文件放在我们的redis启动目录下就可以了
127.0.0.1:6379> config get dir
"dir"
"/usr/local/bin"
几乎就使用它的默认配置就行了
```

### Redis 持久化AOF

以日志的形式来记录每个操作，将Redis执行过的所有指令记录下来(读操作不记录)，只许追加文件但不可以改写文件，redis启动之初会读取该文件重新构建数据，换言之，redis重启的话就根据日志文件的内容将写指令从前到后执行一次以完成数据的恢复工作

```shell
# Aof保存的是appendonly.aof文件
开启 yes 后，如果这个aof文件有错位，这时候redis启动不起来，我们可以使用命令行 `redis-check-aof --fix appendonly.aof` 然后输入y(yes)
```

### Redis发布订阅

Redis 发布订阅 (pub/sub) 是一种消息通信模式：发送者 (pub) 发送消息，订阅者 (sub) 接收消息。

Redis 客户端可以订阅任意数量的频道。

下图展示了频道 channel1 ， 以及订阅这个频道的三个客户端 —— client2 、 client5 和 client1 之间的关系：

![img](https://www.runoob.com/wp-content/uploads/2014/11/pubsub1.png)

当有新消息通过 PUBLISH 命令发送给频道 channel1 时， 这个消息就会被发送给订阅它的三个客户端：

![img](https://www.runoob.com/wp-content/uploads/2014/11/pubsub2.png)

```shell
在我们实例中我们创建了订阅频道名为 runoobChat:
第一个 redis-cli 客户端
#redis 127.0.0.1:6379> SUBSCRIBE runoobChat

第二个 redis-cli 客户端
#redis 127.0.0.1:6379> PUBLISH runoobChat "Redis PUBLISH test"
这个时候在第一个 redis-cli 客户端就会看到由第二个 redis-cli 客户端发送的测试消息。
```

##### Redis 发布订阅命令

| 序号 | 命令及描述                                                   |
| :--- | :----------------------------------------------------------- |
| 1    | [PSUBSCRIBE pattern [pattern ...\]](https://www.runoob.com/redis/pub-sub-psubscribe.html) 订阅一个或多个符合给定模式的频道。 |
| 2    | [PUBSUB subcommand [argument [argument ...\]]](https://www.runoob.com/redis/pub-sub-pubsub.html) 查看订阅与发布系统状态。 |
| 3    | [PUBLISH channel message](https://www.runoob.com/redis/pub-sub-publish.html) 将信息发送到指定的频道。 |
| 4    | [PUNSUBSCRIBE [pattern [pattern ...\]]](https://www.runoob.com/redis/pub-sub-punsubscribe.html) 退订所有给定模式的频道。 |
| 5    | [SUBSCRIBE channel [channel ...\]](https://www.runoob.com/redis/pub-sub-subscribe.html) 订阅给定的一个或多个频道的信息。 |
| 6    | [UNSUBSCRIBE [channel [channel ...\]]](https://www.runoob.com/redis/pub-sub-unsubscribe.html) 指退订给定的频道。 |

### Redis集群环境搭建

只配置从库，不配置主库，默认自己就是主库

```shell
#redis 127.0.0.1:6379>info replication #查看当前库的信息
role:master #角色
connected_slaves:0 #从库数量
测试搭建，在`/usr/local/bin/config/`下复制三个配置文件
cp redis.conf redis1.conf
cp redis.conf redis2.conf
cp redis.conf redis3.conf
```

更改三个复制文件的内容

```shell
vim redis1.conf
port 6379
pidfile /var/run/redis_6379.pid
logfile "redis1.log"
dbfilename dumpredis1.rdb
```

```shell
vim redis2.conf
port 6380
pidfile /var/run/redis_6380.pid
logfile "redis2.log"
dbfilename dumpredis2.rdb
```

```shell
vim redis3.conf
port 6381
pidfile /var/run/redis_6381.pid
logfile "redis3.log"
dbfilename dumpredis3.rdb
```

启动三个服务

```shell
redis-server ./config/redis1.conf
redis-server ./config/redis2.conf
redis-server ./config/redis3.conf
```

一主二从

```shell
redis1.conf(主) redis2.conf/redis3.conf(从)
#redis 127.0.0.1:6380>SLAVEOF 127.0.0.1(ip) 6379(port)
#redis 127.0.0.1:6381>SLAVEOF 127.0.0.1 6379
```

真实的主从配置应该在配置文件中配置，上面的命令是临时的

```shell
#从库的配置文件
replicaof <masterip> <masterport>
#如果主机有密码，配置即可
masterauth <master-password>
```

主机可以写，从机只能读

假设redis1是redis2的主机，redis2是redis3 的主机，redis1挂掉以后，使用命令代替redis1成为主节点

```shell
SLAVEOF no one
```

### Redis哨兵模式

Sentinel(哨兵)是用于监控Redis集群中Master状态的工具，是Redis高可用解决方案，哨兵可以监视一个或者多个redis master服务，以及这些master服务的所有从服务。 某个master服务宕机后，会把这个master下的某个从服务升级为master来替代已宕机的master继续工作。

![2.png](http://ypimg.muzhuangnet.com/Collect/csharp/upload/image/138/862/882/1620703914529826.png)

```shell
在`/uer/local/bin/config`文件下新建哨兵文件
vim sentinel.conf

# 哨兵sentinel实例运行的端口 默认26379
port 26379
# 保护模式关闭，这样其他服务起就可以访问此台redis
protected-mode no
# 哨兵模式是否后台启动，默认no，改为yes
daemonize yes
pidfile /var/run/redis-sentinel.pid
# log日志保存位置
logfile /usr/local/redis/sentinel/redis-sentinel.log
# 工作目录
dir /usr/local/redis/sentinel
# 第三个参数：哨兵名字，可自行修改。（若修改了，那后面涉及到的都得同步） 
# 第四个参数：master主机ip地址
# 第五个参数：redis端口号
# 第六个参数：哨兵的数量。比如2表示，当至少有2个哨兵发现master的redis挂了，
#               那么就将此master标记为宕机节点。
#               这个时候就会进行故障的转移，将其中的一个从节点变为master
sentinel monitor mymaster 192.168.217.151 6379 2
# master中redis的密码
sentinel auth-pass mymaster 123456
# 哨兵从master节点宕机后，等待多少时间（毫秒），认定master不可用。
# 默认30s，这里为了测试，改成10s
sentinel down-after-milliseconds mymaster 10000
# 当替换主节点后，剩余从节点重新和新master做同步的并行数量，默认为 1
sentinel parallel-syncs mymaster 1
# 主备切换的时间，若在3分钟内没有切换成功，换另一个从节点切换
sentinel failover-timeout mymaster 180000


因为配置文件中设定了自己的log存储位置，所以要把相应的文件创建出来，在redis1、redis2和redis3中都需要执行
mkdir /usr/local/bin/config/sentinel -p

`启动哨兵`,在`/usr/local/bin`目录下,分别在redis1、redis2和redis3中执行下面命令，启动哨兵
redis-sentinel ./config/sentinel.conf
```

从测试结果可以看到，即使后来重启之前的master，也不会替换，而是作为slave

优点：

- 哨兵集群，基于主从复制模式，所有的主从配置优点，他全有

- 主从可以切换，故障可以转移，系统的可用性就会更好
- 哨兵模式就是主从复制的升级，自动版

缺点：

- redis不好在线扩容，集群数量一旦达到上线，在线扩容麻烦
- 实现哨兵模式的配置很麻烦，里面有很多选择

###### 哨兵模式的全部配置sentinel.conf

```shell
# Example sentinel.conf
 
# 哨兵sentinel实例运行的端口 默认26379
port 26379
 
# 哨兵sentinel的工作目录
dir /tmp
 
# 哨兵sentinel监控的redis主节点的 ip port 
# master-name  可以自己命名的主节点名字 只能由字母A-z、数字0-9 、这三个字符".-_"组成。
# quorum 当这些quorum个数sentinel哨兵认为master主节点失联 那么这时 客观上认为主节点失联了
# sentinel monitor <master-name> <ip> <redis-port> <quorum>
sentinel monitor mymaster 127.0.0.1 6379 1
 
# 当在Redis实例中开启了requirepass foobared 授权密码 这样所有连接Redis实例的客户端都要提供密码
# 设置哨兵sentinel 连接主从的密码 注意必须为主从设置一样的验证密码
# sentinel auth-pass <master-name> <password>
sentinel auth-pass mymaster MySUPER--secret-0123passw0rd
 
 
# 指定多少毫秒之后 主节点没有应答哨兵sentinel 此时 哨兵主观上认为主节点下线 默认30秒
# sentinel down-after-milliseconds <master-name> <milliseconds>
sentinel down-after-milliseconds mymaster 30000
 
# 这个配置项指定了在发生failover主备切换时最多可以有多少个slave同时对新的master进行 同步，
这个数字越小，完成failover所需的时间就越长，
但是如果这个数字越大，就意味着越 多的slave因为replication而不可用。
可以通过将这个值设为 1 来保证每次只有一个slave 处于不能处理命令请求的状态。
# sentinel parallel-syncs <master-name> <numslaves>
sentinel parallel-syncs mymaster 1
 
# 故障转移的超时时间 failover-timeout 可以用在以下这些方面： 
#1. 同一个sentinel对同一个master两次failover之间的间隔时间。
#2. 当一个slave从一个错误的master那里同步数据开始计算时间。直到slave被纠正为向正确的master那里同步数据时。
#3.当想要取消一个正在进行的failover所需要的时间。  
#4.当进行failover时，配置所有slaves指向新的master所需的最大时间。不过，即使过了这个超时，slaves依然会被正确配置为指向master，但是就不按parallel-syncs所配置的规则来了
# 默认三分钟
# sentinel failover-timeout <master-name> <milliseconds>
sentinel failover-timeout mymaster 180000
 
# SCRIPTS EXECUTION
 
#配置当某一事件发生时所需要执行的脚本，可以通过脚本来通知管理员，例如当系统运行不正常时发邮件通知相关人员。
#对于脚本的运行结果有以下规则：
#若脚本执行后返回1，那么该脚本稍后将会被再次执行，重复次数目前默认为10
#若脚本执行后返回2，或者比2更高的一个返回值，脚本将不会重复执行。
#如果脚本在执行过程中由于收到系统中断信号被终止了，则同返回值为1时的行为相同。
#一个脚本的最大执行时间为60s，如果超过这个时间，脚本将会被一个SIGKILL信号终止，之后重新执行。
 
#通知型脚本:当sentinel有任何警告级别的事件发生时（比如说redis实例的主观失效和客观失效等等），将会去调用这个脚本，
#这时这个脚本应该通过邮件，SMS等方式去通知系统管理员关于系统不正常运行的信息。调用该脚本时，将传给脚本两个参数，
#一个是事件的类型，
#一个是事件的描述。
#如果sentinel.conf配置文件中配置了这个脚本路径，那么必须保证这个脚本存在于这个路径，并且是可执行的，否则sentinel无法正常启动成功。
#通知脚本
# sentinel notification-script <master-name> <script-path>
  sentinel notification-script mymaster /var/redis/notify.sh
 
# 客户端重新配置主节点参数脚本
# 当一个master由于failover而发生改变时，这个脚本将会被调用，通知相关的客户端关于master地址已经发生改变的信息。
# 以下参数将会在调用脚本时传给脚本:
# <master-name> <role> <state> <from-ip> <from-port> <to-ip> <to-port>
# 目前<state>总是“failover”,
# <role>是“leader”或者“observer”中的一个。 
# 参数 from-ip, from-port, to-ip, to-port是用来和旧的master和新的master(即旧的slave)通信的
# 这个脚本应该是通用的，能被多次调用，不是针对性的。
# sentinel client-reconfig-script <master-name> <script-path>
sentinel client-reconfig-script mymaster /var/redis/reconfig.sh
```

### Redis缓存穿透和雪崩

### Redis集群

#### 前置条件

```shell
mkdir -p /app/redis/data


#官网下载的默认地址为/root/redis.conf
wget http://download.redis.io/redis-stable/redis.conf

将一个默认出厂的原始redis.conf文件拷贝近该目录下
cp /redis.conf /app/redis

vi redis.conf

bind 127.0.0.1        #注释掉这部分，使redis可以外部访问
daemonize no          #用守护线程的方式启动
requirepass 123456    #给redis设置密码
appendonly yes        #redis持久化　　默认是no
tcp-keepalive 300     #防止出现远程主机强迫关闭了一个现有的连接的错误 默认是300
```

#### 启动redis

```shell
docker run -d -p 6379:6379 --privileged=true \
-v /app/redis/redis.conf:/etc/redis/redis.conf \
-v /app/redis/data:/data \
--name redis7 redis:7.0.8

docker exec -it redis7 /bin/bash

#连接客户端
redis-cli
```

#### 单机版三主三从

```shell
master1     master2    master3
6381        6382       6383

slave1      slave2     slave3
6384        6385       6386

--net host #使用宿主机的IP和端口，默认
--cluster-enabled yes #开启redis集群

#哈希槽默认分成三段
0-5460 5461-10922 10923-16383
```

```shell
[root@localhost:]mkdir -p /data/redis/share

[root@localhost:]docker run -d --name redis-node-1 --net host --privileged=true -v /data/redis/share/redis-node-1:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6381
 
[root@localhost:]docker run -d --name redis-node-2 --net host --privileged=true -v /data/redis/share/redis-node-2:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6382
 
[root@localhost:]docker run -d --name redis-node-3 --net host --privileged=true -v /data/redis/share/redis-node-3:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6383
 
[root@localhost:]docker run -d --name redis-node-4 --net host --privileged=true -v /data/redis/share/redis-node-4:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6384
 
[root@localhost:]docker run -d --name redis-node-5 --net host --privileged=true -v /data/redis/share/redis-node-5:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6385
 
[root@localhost:]docker run -d --name redis-node-6 --net host --privileged=true -v /data/redis/share/redis-node-6:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6386
```

#### 构建集群关系

```shell
#链接进入6381作为切入点，查看集群状态

[root@localhost:]docker exec -it redis-node-1 /bin/bash

#注意，进入docker容器后才能执行一下命令，且注意自己的真实IP地址
[root@localhost:/data]redis-cli --cluster create 192.168.222.15:6381 192.168.222.15:6382 192.168.222.15:6383 192.168.222.15:6384 192.168.222.15:6385 192.168.222.15:6386 --cluster-replicas 1

#连接
[root@localhost:/data]redis-cli -p 6381

#查看详情
127.0.0.1:6381>cluster info

#查看节点
127.0.0.1:6381>cluster nodes
```

#### 主从容错切换迁移案例

```shell
#数据读写存储
#防止路由失效加参数-c并新增两个key
[root@localhost:/data]redis-cli -p 6381 -c #这里是集群，不是单机，
127.0.0.1:6381>set k1 v1
127.0.0.1:6381>get k1

[root@localhost:/data]redis-cli -p 6382 -c #这里是集群，不是单机，
127.0.0.1:6382>set k2 v2
127.0.0.1:6382>get k2

[root@localhost:/data]redis-cli -p 6383 -c #这里是集群，不是单机，
127.0.0.1:6383>set k3 v3
127.0.0.1:6383>get k3

#查看集群信息
[root@localhost:/data]redis-cli --cluster check 192.168.222.15:6381
```

#### 主从扩容案例

```shell
#新建6387、6388两个节点+新建后启动+查看是否8节点

[root@localhost~]docker run -d --name redis-node-7 --net host --privileged=true -v /data/redis/share/redis-node-7:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6387

[root@localhost~]docker run -d --name redis-node-8 --net host --privileged=true -v /data/redis/share/redis-node-8:/data redis:7.0.8 --cluster-enabled yes --appendonly yes --port 6388

[root@localhost~]docker ps


#进入6387容器实例内部
docker exec -it redis-node-7 /bin/bash

#将新增的6387节点(空槽号)作为master节点加入原集群
redis-cli --cluster add-node 192.168.222.15:6387 192.168.222.15:6381

#检查集群情况第1次
redis-cli --cluster check 192.168.222.15:6381

#重新分派槽号,参考尚硅谷的docker.mmap文件，有图说明
redis-cli --cluster reshard 192.168.222.15:6381

#为主节点6387分配从节点6388
redis-cli --cluster add-node 192.168.222.15:6388 192.168.222.15:6387 --cluster-slave --cluster-master-id a26a4dde9ec5f5ae6b50c96aea00147d65dec232

#检查集群情况第3次
redis-cli --cluster check 192.168.222.15:6382
```

#### 主从缩容案例

```shell
#检查集群情况1获得6388的节点ID
redis-cli --cluster check 192.168.222.15:6382

#从集群中将4号从节点6388删除
redis-cli --cluster del-node 192.168.222.15:6388 e9d5362dcfaf4145411daa29b66d9989d4a4f594

#将6387的槽号清空，重新分配，本例将清出来的槽号都给6381
redis-cli --cluster reshard 192.168.222.15:6381

#将6387删除
redis-cli --cluster del-node 192.168.222.15:6387 a26a4dde9ec5f5ae6b50c96aea00147d65dec232
```

