Gin是一个非常受欢迎的Golang Web框架，在GitHub上已经有47k的星星，它和Golang的语法一样简洁明了，使得初学者得以迅速入门。

# 初始Gin

## gin hello world

使用gin编写一个接口也是非常简单

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()
    router.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })
    router.Run(":8000") 
}
```

1. `router:=gin.Default()`：这是默认的服务器。使用gin的`Default`方法创建一个路由`Handler`；
2. 然后通过Http方法绑定路由规则和路由函数。不同于`net/http`库的路由函数，gin进行了封装，把`request`和`response`都封装到了`gin.Context`的上下文环境中。
3. 最后启动路由的Run方法监听端口。还可以用`http.ListenAndServe(":8080", router)`，或者自定义Http服务器配置。

## 两种启动方式

```go
// 启动方式一
router.Run(":8000")
// 启动方式二
http.ListenAndServe(":8000", router)CopyErrorOK!
```

## 修改ip为内网ip

```go
router.Run("0.0.0.0:8000")
```

```go
package main

import (
  "github.com/gin-gonic/gin"
  "net/http"
)

func Index(context *gin.Context) {
  context.String(200, "Hello 枫枫!")
}
func main() {

  // 创建一个默认的路由
  router := gin.Default()

  // 绑定路由规则和路由函数，访问/index的路由，将由对应的函数去处理
  router.GET("/index", Index)

  // 启动监听，gin会把web服务运行在本机的0.0.0.0:8080端口上
  router.Run("0.0.0.0:8080")
  // 用原生http服务的方式， router.Run本质就是http.ListenAndServe的进一步封装
  http.ListenAndServe(":8080", router)
}
```

------

# 响应

## 状态码

200 表示正常响应 `http.StatusOK`

## 返回字符串

```go
router.GET("/txt", func(c *gin.Context) {
  c.String(http.StatusOK, "返回text")
})
```

## 返回json

```go
router.GET("/json", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
})
// 结构体转json
router.GET("/moreJSON", func(c *gin.Context) {
  // You also can use a struct
  type Msg struct {
    Name    string `json:"user"`
    Message string
    Number  int
  }
  msg := Msg{"fengfeng", "hey", 21}
  // 注意 msg.Name 变成了 "user" 字段
  // 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
  c.JSON(http.StatusOK, msg)
})
```

## 返回xml

```go
router.GET("/xml", func(c *gin.Context) {
  c.XML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
})
```

## 返回yaml

```go
router.GET("/yaml", func(c *gin.Context) {
  c.YAML(http.StatusOK, gin.H{"user": "hanru", "message": "hey", "status": http.StatusOK})
})
```

## 返回html

先要使用 `LoadHTMLGlob()`或者`LoadHTMLFiles()`方法来加载模板文件

```go
//加载模板
router.LoadHTMLGlob("gin框架/templates/*")
//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
//定义路由
router.GET("/tem", func(c *gin.Context) {
  //根据完整文件名渲染模板，并传递参数
  c.HTML(http.StatusOK, "index.html", gin.H{
    "title": "Main website",
  })
})
```

在模板中使用这个title，需要使用`{{  .title  }}`

不同文件夹下模板名字可以相同，此时需要 LoadHTMLGlob() 加载两层模板路径。

```go
router.LoadHTMLGlob("templates/**/*")
router.GET("/posts/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
        "title": "Posts",
    })
    c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
        "title": "Users",
    })
})
```

## 文件响应

```go
// 在golang总，没有相对文件的路径，它只有相对项目的路径
// 网页请求这个静态目录的前缀， 第二个参数是一个目录，注意，前缀不要重复
router.StaticFS("/static", http.Dir("static/static"))
// 配置单个文件， 网页请求的路由，文件的路径
router.StaticFile("/titian.png", "static/titian.png")
```

## 重定向

```go
router.GET("/redirect", func(c *gin.Context) {
    //支持内部和外部的重定向
    c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
})
```

### 301 Moved Permanently

被请求的资源已永久移动到新位置，并且将来任何对此资源的引用都应该使用本响应返回的若干个 URI 之一。如果可能，拥有链接编辑功能的客户端应当自动把请求的地址修改为从服务器反馈回来的地址。除非额外指定，否则这个响应也是可缓存的。

### 302 Found

请求的资源现在临时从不同的 URI 响应请求。由于这样的重定向是临时的，客户端应当继续向原有地址发送以后的请求。只有在Cache-Control或Expires中进行了指定的情况下，这个响应才是可缓存的。

------

# 请求参数

## 查询参数 Query

```go
func _query(c *gin.Context) {
  fmt.Println(c.Query("user"))
  fmt.Println(c.GetQuery("user"))
  fmt.Println(c.QueryArray("user")) // 拿到多个相同的查询参数
  fmt.Println(c.DefaultQuery("addr", "四川省"))
}
```

## 动态参数 Param

```go
func _param(c *gin.Context) {
  fmt.Println(c.Param("user_id"))
  fmt.Println(c.Param("book_id"))
}


router.GET("/param/:user_id/", _param)
router.GET("/param/:user_id/:book_id", _param)

// ?param/12
// ?param/12/123
```

## 表单 PostForm

可以接收 `multipart/form-data; `和`application/x-www-form-urlencoded`

```go
func _form(c *gin.Context) {
  fmt.Println(c.PostForm("name"))
  fmt.Println(c.PostFormArray("name"))
  fmt.Println(c.DefaultPostForm("addr", "四川省")) // 如果用户没传，就使用默认值
  forms, err := c.MultipartForm()               // 接收所有的form参数，包括文件
  fmt.Println(forms, err)
}
```

## 原始参数 GetRawData

#### form-data

```go
----------------------------638149124879484626406689
Content-Disposition: form-data; name="name"

枫枫
----------------------------638149124879484626406689
Content-Disposition: form-data; name="name"

zhangsan
----------------------------638149124879484626406689
Content-Disposition: form-data; name="addr"

长沙市
----------------------------638149124879484626406689--
```

#### x-www-form-urlencoded

```go
name=abc&age=23
```

#### json

```go
{
    "name": "枫枫",
    "age": 21
}
CopyErrorOK!
func _raw(c *gin.Context) {
  body, _ := c.GetRawData()
  contentType := c.GetHeader("Content-Type")
  switch contentType {
  case "application/json":
  
    // json解析到结构体
    type User struct {
      Name string `json:"name"`
      Age  int    `json:"age"`
    }
    var user User
    err := json.Unmarshal(body, &user)
    if err != nil {
      fmt.Println(err.Error())
    }
    fmt.Println(user)
  }
}
```

#### 封装一个解析json到结构体上的函数

```go
func bindJson(c *gin.Context, obj any) (err error) {
  body, _ := c.GetRawData()
  contentType := c.GetHeader("Content-Type")
  switch contentType {
  case "application/json":
    err = json.Unmarshal(body, &obj)
    if err != nil {
      fmt.Println(err.Error())
      return err
    }
  }
  return nil
}
```

## 四大请求方式

```
GET` `POST` `PUT` `DELETE
```

Restful风格指的是网络应用中就是资源定位和资源操作的风格。不是标准也不是协议。

GET：从服务器取出资源（一项或多项）

POST：在服务器新建一个资源

PUT：在服务器更新资源（客户端提供完整资源数据）

PATCH：在服务器更新资源（客户端提供需要修改的资源数据）

DELETE：从服务器删除资源

```go
// 以文字资源为例

// GET    /articles          文章列表
// GET    /articles/:id      文章详情
// POST   /articles          添加文章
// PUT    /articles/:id      修改某一篇文章
// DELETE /articles/:id      删除某一篇文章
```

```go
package main

import (
  "encoding/json"
  "fmt"
  "github.com/gin-gonic/gin"
)

type ArticleModel struct {
  Title   string `json:"title"`
  Content string `json:"content"`
}

type Response struct {
  Code int    `json:"code"`
  Data any    `json:"data"`
  Msg  string `json:"msg"`
}

func _bindJson(c *gin.Context, obj any) (err error) {
  body, _ := c.GetRawData()
  contentType := c.GetHeader("Content-Type")
  switch contentType {
  case "application/json":
    err = json.Unmarshal(body, &obj)
    if err != nil {
      fmt.Println(err.Error())
      return err
    }
  }
  return nil
}

// _getList 文章列表页面
func _getList(c *gin.Context) {
  // 包含搜索，分页
  articleList := []ArticleModel{
    {"Go语言入门", "这篇文章是《Go语言入门》"},
    {"python语言入门", "这篇文章是《python语言入门》"},
    {"JavaScript语言入门", "这篇文章是《JavaScript语言入门》"},
  }
  c.JSON(200, Response{0, articleList, "成功"})
}

// _getDetail 文章详情
func _getDetail(c *gin.Context) {
  // 获取param中的id
  fmt.Println(c.Param("id"))
  article := ArticleModel{
    "Go语言入门", "这篇文章是《Go语言入门》",
  }
  c.JSON(200, Response{0, article, "成功"})
}

// _create 创建文章
func _create(c *gin.Context) {
  // 接收前端传递来的json数据
  var article ArticleModel

  err := _bindJson(c, &article)
  if err != nil {
    fmt.Println(err)
    return
  }

  c.JSON(200, Response{0, article, "添加成功"})
}

// _update 编辑文章
func _update(c *gin.Context) {
  fmt.Println(c.Param("id"))
  var article ArticleModel
  err := _bindJson(c, &article)
  if err != nil {
    fmt.Println(err)
    return
  }
  c.JSON(200, Response{0, article, "修改成功"})
}

// _delete 删除文章
func _delete(c *gin.Context) {
  fmt.Println(c.Param("id"))
  c.JSON(200, Response{0, map[string]string{}, "删除成功"})
}

func main() {
  router := gin.Default()
  router.GET("/articles", _getList)       // 文章列表
  router.GET("/articles/:id", _getDetail) // 文章详情
  router.POST("/articles", _create)       // 添加文章
  router.PUT("/articles/:id", _update)    // 编辑文章
  router.DELETE("/articles/:id", _delete) // 删除文章
  router.Run(":80")
}
```

## 请求头相关

### 请求头参数获取

`GetHeader`，可以大小写不分，且返回切片中的第一个数据

```go
router.GET("/", func(c *gin.Context) {
  // 首字母大小写不区分  单词与单词之间用 - 连接
  // 用于获取一个请求头
  fmt.Println(c.GetHeader("User-Agent"))
  //fmt.Println(c.GetHeader("user-agent"))
  //fmt.Println(c.GetHeader("user-Agent"))
  //fmt.Println(c.GetHeader("user-AGent"))

  // Header 是一个普通的 map[string][]string
  fmt.Println(c.Request.Header)
  // 如果是使用 Get方法或者是 .GetHeader,那么可以不用区分大小写，并且返回第一个value
  fmt.Println(c.Request.Header.Get("User-Agent"))
  fmt.Println(c.Request.Header["User-Agent"])
  // 如果是用map的取值方式，请注意大小写问题
  fmt.Println(c.Request.Header["user-agent"])

  // 自定义的请求头，用Get方法也是免大小写
  fmt.Println(c.Request.Header.Get("Token"))
  fmt.Println(c.Request.Header.Get("token"))
  c.JSON(200, gin.H{"msg": "成功"})
})
```

## 响应头相关

### 设置响应头

```go
// 设置响应头
router.GET("/res", func(c *gin.Context) {
  c.Header("Token", "jhgeu%hsg845jUIF83jh")
  c.Header("Content-Type", "application/text; charset=utf-8")
  c.JSON(0, gin.H{"data": "看看响应头"})
})
```

------

# Bind绑定器

gin中的bind可以很方便的将 前端传递 来的数据与 `结构体` 进行 `参数绑定` ，以及参数校验

在使用这个功能的时候，需要给结构体加上Tag `json` `form` `uri` `xml` `yaml`

## Must Bind

不用，校验失败会改状态码

## Should Bind

可以绑定json，query，param，yaml，xml

如果校验不通过会返回错误

### ShouldBindJSON

```go
package main

import "github.com/gin-gonic/gin"

type UserInfo struct {
  Name string `json:"name"`
  Age  int    `json:"age"`
  Sex  string `json:"sex"`
}

func main() {
  router := gin.Default()
  router.POST("/", func(c *gin.Context) {

    var userInfo UserInfo
    err := c.ShouldBindJSON(&userInfo)
    if err != nil {
      c.JSON(200, gin.H{"msg": "你错了"})
      return
    }
    c.JSON(200, userInfo)

  })
  router.Run(":80")
}
```

### ShouldBindQuery

绑定查询参数,tag对应为form

```go
// ?name=枫枫&age=21&sex=男
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type UserInfo struct {
  Name string `json:"name" form:"name"`
  Age  int    `json:"age" form:"age"`
  Sex  string `json:"sex" form:"sex"`
}

func main() {
  router := gin.Default()

  router.POST("/query", func(c *gin.Context) {

    var userInfo UserInfo
    err := c.ShouldBindQuery(&userInfo)
    if err != nil {
      fmt.Println(err)
      c.JSON(200, gin.H{"msg": "你错了"})
      return
    }
    c.JSON(200, userInfo)

  })
  router.Run(":80")
}
```

### ShouldBindUri

绑定动态参数,tag对应为uri

```go
// /uri/fengfeng/21/男

package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type UserInfo struct {
  Name string `json:"name" form:"name" uri:"name"`
  Age  int    `json:"age" form:"age" uri:"age"`
  Sex  string `json:"sex" form:"sex" uri:"sex"`
}

func main() {
  router := gin.Default()

  router.POST("/uri/:name/:age/:sex", func(c *gin.Context) {

    var userInfo UserInfo
    err := c.ShouldBindUri(&userInfo)
    if err != nil {
      fmt.Println(err)
      c.JSON(200, gin.H{"msg": "你错了"})
      return
    }
    c.JSON(200, userInfo)

  })

  router.Run(":80")
}
```

### ShouldBind

会根据请求头中的content-type去自动绑定,form-data的参数也用这个，tag用form,默认的tag就是form

#### 绑定form-data、x-www-form-urlencode

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type UserInfo struct {
  Name string `form:"name"`
  Age  int    `form:"age"`
  Sex  string `form:"sex"`
}

func main() {
  router := gin.Default()
  
  router.POST("/form", func(c *gin.Context) {
    var userInfo UserInfo
    err := c.ShouldBind(&userInfo)
    if err != nil {
      fmt.Println(err)
      c.JSON(200, gin.H{"msg": "你错了"})
      return
    }
    c.JSON(200, userInfo)
  })

  router.Run(":80")
}
```

## bind绑定器

需要使用参数验证功能，需要加binding tag

### 常用验证器

```go
// 不能为空，并且不能没有这个字段
required： 必填字段，如：binding:"required"  

// 针对字符串的长度
min 最小长度，如：binding:"min=5"
max 最大长度，如：binding:"max=10"
len 长度，如：binding:"len=6"

// 针对数字的大小
eq 等于，如：binding:"eq=3"
ne 不等于，如：binding:"ne=12"
gt 大于，如：binding:"gt=10"
gte 大于等于，如：binding:"gte=10"
lt 小于，如：binding:"lt=10"
lte 小于等于，如：binding:"lte=10"

// 针对同级字段的
eqfield 等于其他字段的值，如：PassWord string `binding:"eqfield=Password"`
nefield 不等于其他字段的值


- 忽略字段，如：binding:"-"
```

### gin内置验证器

```go
// 枚举  只能是red 或green
oneof=red green 

// 字符串  
contains=fengfeng  // 包含fengfeng的字符串
excludes // 不包含
startswith  // 字符串前缀
endswith  // 字符串后缀

// 数组
dive  // dive后面的验证就是针对数组中的每一个元素

// 网络验证
ip
ipv4
ipv6
uri
url
// uri 在于I(Identifier)是统一资源标示符，可以唯一标识一个资源。
// url 在于Locater，是统一资源定位符，提供找到该资源的确切路径

// 日期验证  1月2号下午3点4分5秒在2006年
datetime=2006-01-02
```

### 自定义验证的错误信息

当验证不通过时，会给出错误的信息，但是原始的错误信息不太友好，不利于用户查看

只需要给结构体加一个msg 的tag

```go
type UserInfo struct {
  Username string `json:"username" binding:"required" msg:"用户名不能为空"`
  Password string `json:"password" binding:"min=3,max=6" msg:"密码长度不能小于3大于6"`
  Email    string `json:"email" binding:"email" msg:"邮箱地址格式不正确"`
}
```

当出现错误时，就可以来获取出错字段上的msg。

- `err`：这个参数为`ShouldBindJSON`返回的错误信息
- `obj`：这个参数为绑定的结构体
- **还有一点要注意的是，validator这个包要引用v10这个版本的，否则会出错**

```go
// GetValidMsg 返回结构体中的msg参数
func GetValidMsg(err error, obj any) string {
  // 使用的时候，需要传obj的指针
  getObj := reflect.TypeOf(obj)
  // 将err接口断言为具体类型
  if errs, ok := err.(validator.ValidationErrors); ok {
    // 断言成功
    for _, e := range errs {
      // 循环每一个错误信息
      // 根据报错字段名，获取结构体的具体字段
      if f, exits := getObj.Elem().FieldByName(e.Field()); exits {
        msg := f.Tag.Get("msg")
        return msg
      }
    }
  }
  return err.Error()
}
```

### 自定义验证器

#### 1、注册验证器函数

```go
// github.com/go-playground/validator/v10
// 注意这个版本得是v10的

if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
   v.RegisterValidation("sign", signValid)
}
```

#### 2、编写函数

```go
// 如果用户名不等于fengfeng就校验失败
func signValid(fl validator.FieldLevel) bool {
  name := fl.Field().Interface().(string)
  if name != "fengfeng" {
    return false
  }
  return true
}
```

#### 3、使用

```go
type UserInfo struct {
  Name string `json:"name" binding:"sign" msg:"用户名错误"`
  Age  int    `json:"age" binding:""`
}
```

```go
package main

import (
  "github.com/gin-gonic/gin"
  "github.com/gin-gonic/gin/binding"
  "github.com/go-playground/validator/v10"
  "reflect"
)

func GetValidMsg(err error, obj interface{}) string {
  // obj为结构体指针
  getObj := reflect.TypeOf(obj)
  // 断言为具体的类型，err是一个接口
  if errs, ok := err.(validator.ValidationErrors); ok {
    for _, e := range errs {
      if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
        return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
      }
    }
  }
  return err.Error()
}
// 如果用户名不等于fengfeng就校验失败
func signValid(fl validator.FieldLevel) bool {
  name := fl.Field().Interface().(string)
  if name != "fengfeng" {
    return false
  }
  return true
}


func main() {
  router := gin.Default()
  router.POST("/", func(c *gin.Context) {
    type UserInfo struct {
      Name string `json:"name" binding:"sign" msg:"用户名错误"`
      Age  int    `json:"age" binding:""`
    }
    var user UserInfo
    err := c.ShouldBindJSON(&user)
    if err != nil {
      // 显示自定义的错误信息
      msg := GetValidMsg(err, &user)
      c.JSON(200, gin.H{"msg": msg})
      return
    }
    c.JSON(200, user)
  })
  // 注册
  if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
    v.RegisterValidation("sign", signValid)
  }
  router.Run(":80")
}
```

------

# 文件上传和下载

## 文件上传

### 单文件

```go
func main() {
  router := gin.Default()
  // 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
  // 单位是字节， << 是左移预算符号，等价于 8 * 2^20
  // gin对文件上传大小的默认值是32MB
  router.MaxMultipartMemory = 8 << 20  // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // 单文件
    file, _ := c.FormFile("file")
    log.Println(file.Filename)

    dst := "./" + file.Filename
    // 上传文件至指定的完整文件路径
    c.SaveUploadedFile(file, dst)

    c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
  })
  router.Run(":8080")
}
```

## 服务端保存文件的几种方式

### SaveUploadedFile

```go
c.SaveUploadedFile(file, dst)  // 文件对象  文件路径，注意要从项目根路径开始写
```

### Create+Copy

file.Open的第一个返回值就是我们讲文件对象中的那个文件（只读的），我们可以使用这个去直接读取文件内容

```go
file, _ := c.FormFile("file")
log.Println(file.Filename)
// 读取文件中的数据，返回文件对象
fileRead, _ := file.Open()
dst := "./" + file.Filename
// 创建一个文件
out, err := os.Create(dst)
if err != nil {
  fmt.Println(err)
}
defer out.Close()
// 拷贝文件对象到out中
io.Copy(out, fileRead)
```

### 读取上传的文件

```go
file, _ := c.FormFile("file")
// 读取文件中的数据，返回文件对象
fileRead, _ := file.Open()
data, _ := io.ReadAll(fileRead)
fmt.Println(string(data))
```

这里的玩法就很多了,例如我们可以基于文件中的内容，判断是否需要保存到服务器中

## 多文件上传

```go
func main() {
  router := gin.Default()
  // 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
  router.MaxMultipartMemory = 8 << 20 // 8 MiB
  router.POST("/upload", func(c *gin.Context) {
    // Multipart form
    form, _ := c.MultipartForm()
    files := form.File["upload[]"]  // 注意这里名字不要对不上了

    for _, file := range files {
      log.Println(file.Filename)
      // 上传文件至指定目录
      c.SaveUploadedFile(file, "./"+file.Filename)
    }
    c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
  })
  router.Run(":8080")
}
```

![img](http://python.fengfengzhidao.com/pic/20221210202205.png)

## 文件下载

### 直接响应一个路径下的文件

```go
c.File("uploads/12.png")
```

有些响应，比如图片，浏览器就会显示这个图片，而不是下载，所以我们需要使浏览器唤起下载行为

```go
c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
c.Header("Content-Disposition", "attachment; filename="+"牛逼.png") // 用来指定下载下来的文件名
c.Header("Content-Transfer-Encoding", "binary")                   // 表示传输过程中的编码形式，乱码问题可能就是因为它
c.File("uploads/12.png")
```

注意，文件下载浏览器可能会有缓存，这个要注意一下

解决办法就是加查询参数

### 前后端模式下的文件下载

如果是前后端模式下，后端就只需要响应一个文件数据

文件名和其他信息就写在请求头中

```go
c.Header("fileName", "xxx.png")
c.Header("msg", "文件下载成功")
c.File("uploads/12.png")
```

### 前端写法

```go
async downloadFile(row) {
   this.$http({
      method: 'post',
      url: 'file/upload',
      data:postData,
      responseType: "blob"
   }).then(res => {
      const _res = res.data
      let blob = new Blob([_res], {
            type: 'application/png'
          });
      let downloadElement = document.createElement("a");
      let href = window.URL.createObjectURL(blob); //创建下载的链接
      downloadElement.href = href;
      downloadElement.download = res.headers["fileName"]; //下载后文件名
      document.body.appendChild(downloadElement);
      downloadElement.click(); //点击下载
      document.body.removeChild(downloadElement); //下载完成移除元素
      window.URL.revokeObjectURL(href); //释放掉blob对象
    })}
```

------

# 中间件和路由

Gin框架允许开发者在处理请求的过程中，加入用户自己的钩子（Hook）函数。这个钩子函数就叫中间件，中间件适合处理一些公共的业务逻辑，比如登录认证、权限校验、数据分页、记录日志、耗时统计等 即比如，如果访问一个网页的话，不管访问什么路径都需要进行登录，此时就需要为所有路径的处理函数进行统一一个中间件

Gin中的中间件必须是一个gin.HandlerFunc类型

## 单独注册中间件

```go
import (
  "fmt"
  "github.com/gin-gonic/gin"
  "net/http"
)
func indexHandler(c *gin.Context) {
  fmt.Println("index.....")
  c.JSON(http.StatusOK, gin.H{
    "msg": "index",
  })
}

//定义一个中间件
func m1(c *gin.Context) {
  fmt.Println("m1 in.........")
}
func main() {
  r := gin.Default()
  //m1处于indexHandler函数的前面,请求来之后,先走m1,再走index
  r.GET("/index", m1, indexHandler)

  _ = r.Run()
}
```

## 多个中间件

router.GET，后面可以跟很多HandlerFunc方法，这些方法其实都可以叫中间件

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...")
    c.JSON(200, gin.H{"msg": "响应数据"})
  }, m2)

  router.Run(":8080")
}

/*
m1  ...in
index ...
m2  ...in
*/
```

## 中间件拦截响应

c.Abort()拦截，后续的HandlerFunc就不会执行了

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.JSON(200, gin.H{"msg": "第一个中间件拦截了"})
  c.Abort()
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...")
    c.JSON(200, gin.H{"msg": "响应数据"})
  }, m2)

  router.Run(":8080")
}
```

## 中间件放行

c.Next()，Next前后形成了其他语言中的请求中间件和响应中间件

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Next()
  fmt.Println("m1 ...out")
}
func m2(c *gin.Context) {
  fmt.Println("m2 ...in")
  c.Next()
  fmt.Println("m2 ...out")
}

func main() {
  router := gin.Default()

  router.GET("/", m1, func(c *gin.Context) {
    fmt.Println("index ...in")
    c.JSON(200, gin.H{"msg": "响应数据"})
    c.Next()
    fmt.Println("index ...out")
  }, m2)

  router.Run(":8080")
}

/*
m1 ...in
index ...in
m2 ...in   
m2 ...out  
index ...out
m1 ...out
*/
```

![img](http://python.fengfengzhidao.com/pic/20221210220434.png)

如果其中一个中间件响应了c.Abort()，后续中间件将不再执行，直接按照顺序走完所有的响应中间件

## 全局注册中间件

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Next()
  fmt.Println("m1 ...out")
}

func main() {
  router := gin.Default()

  router.Use(m10)
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    c.JSON(200, gin.H{"msg": "index"})
    c.Next()
    fmt.Println("index ...out")
  })

  router.Run(":8080")

}
```

使用Use去注册全局中间件，Use接收的参数也是多个HandlerFunc

### 中间件传递数据

使用Set设置一个key-value,

在后续中间件中使用Get接收数据

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Set("name", "fengfeng")
}

func main() {
  router := gin.Default()

  router.Use(m10)
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    name, _ := c.Get("name")
    fmt.Println(name)
    
    c.JSON(200, gin.H{"msg": "index"})
  })

  router.Run(":8080")

}
```

value的类型是any类型，所有我们可以用它传任意类型，在接收的时候做好断言即可

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type User struct {
  Name string
  Age  int
}

func m10(c *gin.Context) {
  fmt.Println("m1 ...in")
  c.Set("name", User{"枫枫", 21})
  c.Next()
  fmt.Println("m1 ...out")
}

func main() {
  router := gin.Default()

  router.Use(m10)
  router.GET("/", func(c *gin.Context) {
    fmt.Println("index ...in")
    name, _ := c.Get("name")
    user := name.(User)
    fmt.Println(user.Name, user.Age)
    c.JSON(200, gin.H{"msg": "index"})
  })

  router.Run(":8080")
}
```

## 路由分组

将一系列的路由放到一个组下，统一管理

例如，以下的路由前面统一加上api的前缀

```go
package main

import "github.com/gin-gonic/gin"

func main() {
  router := gin.Default()

  r := router.Group("/api")
  r.GET("/index", func(c *gin.Context) {
    c.String(200, "index")
  })
  r.GET("/home", func(c *gin.Context) {
    c.String(200, "home")
  })

  router.Run(":8080")
}
CopyErrorOK!
```

### 路由分组注册中间件

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func middle(c *gin.Context) {
  fmt.Println("middle ...in")
}

func main() {
  router := gin.Default()

  r := router.Group("/api").Use(middle)  // 可以链式，也可以直接r.Use(middle)
  r.GET("/index", func(c *gin.Context) {
    c.String(200, "index")
  })
  r.GET("/home", func(c *gin.Context) {
    c.String(200, "home")
  })

  router.Run(":8080")
}
```

这样写我们就可以指定哪一些分组下可以使用中间件了

当然，中间件还有一种写法，就是使用函数加括号的形式

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func middle(c *gin.Context) {
  fmt.Println("middle ...in")
}
func middle1() gin.HandlerFunc {
  // 这里的代码是程序一开始就会执行
  return func(c *gin.Context) {
    // 这里是请求来了才会执行
    fmt.Println("middle1 ...inin")
  }
}

func main() {
  router := gin.Default()

  r := router.Group("/api").Use(middle, middle1())
  r.GET("/index", func(c *gin.Context) {
    c.String(200, "index")
  })
  r.GET("/home", func(c *gin.Context) {
    c.String(200, "home")
  })

  router.Run(":8080")
}
```

## gin.Default

```go
func Default() *Engine {
  debugPrintWARNINGDefault()
  engine := New()
  engine.Use(Logger(), Recovery())
  return engine
}
CopyErrorOK!
```

gin.Default()默认使用了Logger和Recovery中间件，其中：

Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。 Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。 如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

使用gin.New，如果不指定日志，那么在控制台中就不会有日志显示

## 中间件案例

### 权限验证

以前后端最流行的jwt为例，如果用户登录了，前端发来的每一次请求都会在请求头上携带上token

后台拿到这个token进行校验，验证是否过期，是否非法

如果通过就说明这个用户是登录过的

不通过就说明用户没有登录

```go
package main

import (
  "github.com/gin-gonic/gin"
)

func JwtTokenMiddleware(c *gin.Context) {
  // 获取请求头的token
  token := c.GetHeader("token")
  // 调用jwt的验证函数
  if token == "1234" {
    // 验证通过
    c.Next()
    return
  }
  // 验证不通过
  c.JSON(200, gin.H{"msg": "权限验证失败"})
  c.Abort()
}

func main() {
  router := gin.Default()

  api := router.Group("/api")

  apiUser := api.Group("")
  {
    apiUser.POST("login", func(c *gin.Context) {
      c.JSON(200, gin.H{"msg": "登录成功"})
    })
  }
  apiHome := api.Group("system").Use(JwtTokenMiddleware)
  {
    apiHome.GET("/index", func(c *gin.Context) {
      c.String(200, "index")
    })
    apiHome.GET("/home", func(c *gin.Context) {
      c.String(200, "home")
    })
  }

  router.Run(":8080")
}
```

### 耗时统计

统计每一个视图函数的执行时间

```go
func TimeMiddleware(c *gin.Context) {
  startTime := time.Now()
  c.Next()
  since := time.Since(startTime)
  // 获取当前请求所对应的函数
  f := c.HandlerName()
  fmt.Printf("函数 %s 耗时 %d\n", f, since)
}
```

------

# 日志

## 为什么要使用日志

1. 记录用户操作，猜测用户行为
2. 记录bug

## gin自带日志系统

### 输出日志到log文件

```go
package main

import (
  "github.com/gin-gonic/gin"
  "io"
  "os"
)

func main() {
  // 输出到文件
  f, _ := os.Create("gin.log")
  //gin.DefaultWriter = io.MultiWriter(f)
  // 如果需要同时将日志写入文件和控制台，请使用以下代码。
  gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
  router := gin.Default()
  router.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{"msg": "/"})
  })
  router.Run()
}
```

### 定义路由格式

启动gin，它会显示所有的路由，默认格式如下

```go
[GIN-debug] POST   /foo    --> main.main.func1 (3 handlers)
[GIN-debug] GET    /bar    --> main.main.func2 (3 handlers)
[GIN-debug] GET    /status --> main.main.func3 (3 handlers)
```

```go
gin.DebugPrintRouteFunc = func(
  httpMethod,
  absolutePath,
  handlerName string,
  nuHandlers int) {
  log.Printf(
    "[ feng ] %v %v %v %v\n",
    httpMethod,
    absolutePath,
    handlerName,
    nuHandlers,
  )
}
/*  输出如下
2022/12/11 14:10:28 [ feng ] GET / main.main.func3 3
2022/12/11 14:10:28 [ feng ] POST /index main.main.func4 3
2022/12/11 14:10:28 [ feng ] PUT /haha main.main.func5 3
2022/12/11 14:10:28 [ feng ] DELETE /home main.main.func6 3
*/
```

### 查看路由

```go
router.Routes()  // 它会返回已注册的路由列表
```

### 环境切换

![img](http://python.fengfengzhidao.com/pic/20221211142056.png)

如果不想看到这些debug日志，那么我们可以改为release模式

```go
gin.SetMode(gin.ReleaseMode)
router := gin.Default()
```

### 修改log的显示

默认的是这样的

```go
[GIN] 2022/12/11 - 14:22:00 | 200 |  0s |  127.0.0.1 | GET  "/"
```

如果觉得不好看，我们可以自定义

```go
package main

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

func LoggerWithFormatter(params gin.LogFormatterParams) string {

  return fmt.Sprintf(
    "[ feng ] %s  | %d | \t %s | %s | %s \t  %s\n",
    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
    params.StatusCode,  // 状态码
    params.ClientIP,  // 客户端ip
    params.Latency,  // 请求耗时
    params.Method,  // 请求方法
    params.Path,  // 路径
  )
}
func main() {
  router := gin.New()
  router.Use(gin.LoggerWithFormatter(LoggerWithFormatter))
  router.Run()

}
```

也可以这样

```go
func LoggerWithFormatter(params gin.LogFormatterParams) string {
  return fmt.Sprintf(
    "[ feng ] %s  | %d | \t %s | %s | %s \t  %s\n",
    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
    params.StatusCode,
    params.ClientIP,
    params.Latency,
    params.Method,
    params.Path,
  )
}
func main() {
  router := gin.New()
  router.Use(
    gin.LoggerWithConfig(
      gin.LoggerConfig{Formatter: LoggerWithFormatter},
    ),
  )
  router.Run()

}
```

但是你会发现自己这样输出之后，没有颜色了，不太好看，我们可以输出有颜色的log

```go
func LoggerWithFormatter(params gin.LogFormatterParams) string {
  var statusColor, methodColor, resetColor string
  statusColor = params.StatusCodeColor()
  methodColor = params.MethodColor()
  resetColor = params.ResetColor()
  return fmt.Sprintf(
    "[ feng ] %s  | %s %d  %s | \t %s | %s | %s %-7s %s \t  %s\n",
    params.TimeStamp.Format("2006/01/02 - 15:04:05"),
    statusColor, params.StatusCode, resetColor,
    params.ClientIP,
    params.Latency,
    methodColor, params.Method, resetColor,
    params.Path,
  )
}
```

------

