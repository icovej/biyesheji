package controller

import (
	"PlatformBackEnd/data"
	"PlatformBackEnd/tools"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func MonitorK8SResource(c *gin.Context) {
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
		"code":       data.SUCCESS,
		"message":    "Succeed to get pod and namespace info",
		"pods":       podList,
		"namespaces": nsList,
	})
	glog.Info("Succeed to getcluster resource info")
}