package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 操作接口信息
type Operation struct {
	Api    string `json:"api"`
	Params string `json:"params"`
	Remark string `json:"remark"`
}

// 操作接口信息
func OperationInfo(c *gin.Context) {
	list := []Operation{
		Operation{
			"/regoster",
			"username, password, permission, workpath",
			"注册账号",
		},
		Operation{
			"/login",
			"username, password",
			"登录账号",
		},
		Operation{
			"/search_dir",
			"dir_name, depth",
			"查询目录存储",
		},
		Operation{
			"/image",
			"dstpath, osversion, pyversion, imagearray, imagename",
			"创建镜像",
		},
		Operation{
			"/pod",
			"podname, container_name, memeory, cpu, etc",
			"创建容器",
		},
		Operation{
			"/data",
			"logdir",
			"模型训练的损失值和正确率",
		},
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "list": list})
	return
}