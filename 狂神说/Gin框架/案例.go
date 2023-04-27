// go get github.com/gin-gonic/gin
//官网https://gin-gonic.com/zh-cn/docs/

package main
import (
    "github.com/gin-gonic/gin"
    "github.com/thinkerou/favicon"
    "log"
    "net/http"
    "encoding/json"
)

func main() {
    //创建一个服务
    ginServer := gin.Default()

    //给网页添加一个logo,favicon.ico展示网站个性的缩略logo标志
    ginServer.Use(favicon.New("./favicon.ico"))

    // (1)举例：访问地址，处理我们的请求
    ginServer.GET("/hello",func (context *gin.Context{
        context.JSON(200,gin.H{"msg":"Hello World"})
    })

    // (2) 加载前端静态页面
    ginServer.LoadHTMLGlob("templates/*")

    //加载静态资源文件
    ginServer.Static("/static","./static")

    //(3) Restful API,响应一个页面给前端
    ginServer.GET("/index",func (c *gin.Context{
        c.HTML(http.StatusOK,"login.html",gin.H{"msg":"这是Go后台传递来的数据"})
    })

    /* ginServer.POST("/index")
    ginServer.PUT("/index")
    ginServer.DELETE("/index") */

    // (4) 接收前端传参: user/info?userid=xxx&username=kuangshen
    ginServer.GET("/user/info",func (context *gin.Context{
        userid := context.Query("userid")
        username := context.Query("username")
        context.JSON(http.StatusOK,gin.H{"userid":userid,"username":username})
    })
    //或者 user/info/1/kuangshen
    ginServer.GET("user/info/:userid/:username",func (context *gin.Context{
        //param:单个，params多个
        userid := context.Param("userid")
        username := context.Param("username")
        context.JSON(http.StatusOK,gin.H{"userid":userid,"username":username})
    })

    // (5) 前端给后端传递一个json，序列化
    ginServer.POST("/json",func (context *gin.Context{
        //获取前端传递的内容
        data,_ := context.GetRawData()
        //定义一个接口
        var m map[string] interface{}
        _ = json.Unmarshal(data,&m)
        context.JSON(http.StatusOK,m)
    })

    // (6) 前端给后端传递表单数据
    ginServer.POST("/user/add",func (context *gin.Context{
        username := context.PostForm("username")
        password := context.PostForm("password")
        context.JSON(http.StatusOK,gin.H{"username":username,"password":password})
    })

    // (7) 路由
    ginServer.GET("/test",func (context *gin.Context{
        //重定向,301
        context.Redirect(http.StatusMovedPermanently,"https://kuangstudy.com")
    })
    //没有路由 404
    ginServer.NoRoute(func (context *gin.Context{
        context.HTML(http.StatusNotFound,"404.html",nil)
    })

    // (8)路由组,管理路由接口
    userGroup := ginServer.Group("/user")
    {
        userGroup.GET("/login")
        userGroup.GET("/add")
        userGroup.GET("/logout")
    }

    // (9) 拦截器 中间件(举例说明：登录授权，验证，分页),自定义一个中间件,接口没有指定MyHandler()，就默认全局使用MyHandler()
    func MyHandller() (gin.HandlerFunc){
        f := func (context *gin.Context{
            //通过自定义的中间件，设置的值，在后续处理只要调用了这个中间件的都可以拿到这里的参数
            context.Set("usession","dj4rys45ecv4e9v42tz405nos15l29v3sr7")
            if xxx,xxx{
                context.Next() //放行
            } else {
                context.Abort() //阻断,不放行
            }
        }
        return f
    }
    ginServer.GET("/handler",MyHandller(),func (context *gin.context){
        //取出中间件的值
        usersession := context.MustGet("usession").(string)
        log.Println("取出中间件的值",usersession)
    })

    //服务器端口
    ginServer.Run(":8083")

    //修改ip为内网ip
    //ginServer.Run("0.0.0.0:8083")
}