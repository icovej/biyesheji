package main

import (
	"flag"
	"platform_back_end/controller"
	"platform_back_end/data"
	"platform_back_end/tools"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	// parse cmdline
	flag.Parse()

	var srcfilepath = flag.String("srcfilepath", "", "the original dockerfile path")
	data.Srcfilepath = *srcfilepath

	var logdir = flag.String("logdir", "", "The path to save glog")
	flag.Lookup("log_dir").Value.Set(*logdir)

	defer glog.Flush()

	// 初始化Gin框架
	router := gin.Default()
	router.Use(tools.Core())

	// 获取api操作信息
	router.GET("/operation", controller.OperationInfo)

	// 账号注册
	router.POST("/register", controller.RegisterHandler)

	// 账号登陆
	router.POST("/login", controller.Login)

	// 查询目录容量
	router.GET("/search_dir", controller.GetDirInfo)

	// 创建镜像
	router.POST("/image", controller.CreateImage)

	// 创建Pod
	router.POST("/pod", controller.CreatePod)

	// 获取训练模型的损失值和正确率
	router.POST("/data", controller.GetData)

	router.Run(":8080")
}
