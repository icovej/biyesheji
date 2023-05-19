package controller

import (
	"PlatformBackEnd/data"
	"PlatformBackEnd/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func GetK8SPod(c *gin.Context) {
	var mpod data.Monitor
	err := c.ShouldBindJSON(&mpod)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    data.API_PARAMETER_ERROR,
			"message": fmt.Sprintf("Method MonitorK8SResource gets invalid request payload, err is %v", err.Error()),
		})
		glog.Errorf("Method MonitorK8SResource gets invalid request payload, the error is %v", err.Error())
		return
	}
	glog.Info("Succeed to get request to get cluster resource info")

	podList, err := tools.GetAllPod(mpod.Namespace)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    data.API_PARAMETER_ERROR,
			"message": fmt.Sprintf("Failed to get pod info, the error is %v", err),
		})
		glog.Errorf("Failed to get pod info, the error is %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    data.SUCCESS,
		"message": "Succeed to get pod info",
		"data":    podList,
	})
	glog.Info("Succeed to get cluster pods info")
}

func GetK8SNamespace(c *gin.Context) {
	nsList, err := tools.GetAllNamespace()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    data.API_PARAMETER_ERROR,
			"message": fmt.Sprintf("Failed to get all namespace, the error is %v", err),
		})
		glog.Errorf("Failed to get all namespace, the error is %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    data.SUCCESS,
		"message": "Succeed to get namespace info",
		"data":    nsList,
	})
	glog.Info("Succeed to get cluster all namespace")
}

func GetK8SNodeGPU(c *gin.Context) {
	glog.Info("Succeed to get request to get cluster node and gpu info")
	result, err := tools.GetGPUCount()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    data.OPERATION_FAILURE,
			"message": fmt.Sprintf("Failed to get node and gp info, the error is %v", err),
		})
		glog.Errorf("Failed to get node and gp info, the error is %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    data.SUCCESS,
		"message": "Succeed to get node and gpu info",
		"data":    result,
	})
	glog.Info("Succeed to get node and gpu info")
}

func GetClusterNodeData(c *gin.Context) {
	result, err := tools.GetClusterNodeData()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    data.OPERATION_FAILURE,
			"message": fmt.Sprintf("Failed to get node data, the error is %v", err),
		})
		glog.Errorf("Failed to get node data, the error is %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    data.SUCCESS,
		"message": "Succeed to get node data",
		"data":    result,
	})
	glog.Info("Succeed to get node data")
}
