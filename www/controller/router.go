package controller

import (
	"github.com/gin-gonic/gin"
	"yespider-go/www/controller/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/index", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "success!"})
	})

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())

	{
		apiv1.GET("/tags", api.GetTags)
		apiv1.POST("/tags", api.PostTags)
		apiv1.GET("/task", api.GetTaskInfo)
	}

	return r

}
