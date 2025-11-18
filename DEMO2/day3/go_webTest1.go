package main

// 第一行是固定代码：package main
import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// 创建一个路由对象
	router := gin.Default()
	// 创建get请求
	router.POST("/", func(content *gin.Context) {
		content.JSON(200, gin.H{
			"userName": "张三",
			"age":      "20",
		})

	})

	router.LoadHTMLGlob("D:/测试代码/DEMO2/day3/templates/*")

	router.GET("/html", func(ctx *gin.Context) {

		fmt.Println("不为空1")
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"title": "前台首页",
		})
		fmt.Println("不为空2")
	})
	//  指定端口
	router.Run()

}
