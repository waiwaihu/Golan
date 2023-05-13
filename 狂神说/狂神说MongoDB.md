### 一、mongodb安装

1、安装镜像

```undefined
docker pull mongo:5.0.2
```

2、新建目录

```bash
mkdir -p /home/apps/mongo/data
```

3、创建并启动

```haskell
docker run \
-d \
--name mongo \
--restart=always \
--privileged=true \
-p 27017:27017 \
-v /home/apps/mongo/data:/data/db \
mongo:6.0.4 --auth
```

```haskell
docker run \
-d \
--name mongo \
--restart=always \
--privileged=true \
-p 27017:27017 \
-v /home/apps/mongo/data:/data/db \
-e MONGO_INITDB_ROOT_USERNAME=admin \
-e MONGO_INITDB_ROOT_PASSWORD=123456 \
mongo:6.0.4 --auth
```



4、创建登录用户

```lua
-- 进入docker并登录mongo
docker exec -it mongo mongo admin
或
-- 进入容器
docker exec -it mongo /bin/bash
-- 登录mobodb
mongo admin

-- 创建一个名为 root，密码为 123456 的用户。
db.createUser({ user:'root',pwd:'123456',roles:[ { role:'userAdminAnyDatabase', db: 'admin'},"readWriteAnyDatabase"]});

-- 尝试使用上面创建的用户信息进行连接。
db.auth('root', '123456')
```

5、创建用户

```lua
-- 使用admin
use admin
-- 验证权限
db.auth('root','123456')

-- use db01 可以指定数据库创建用户
-- 创建新用户
db.createUser(
    {
      user: "user01",
      pwd: "123456A",
      roles: ["readWrite"]
    }
 )
 
--查询所有用户
db.getUsers()
```