package routes

import (
	"github.com/gin-gonic/gin"
	"n8n-workflow/curl"
)

func NewRouter() *gin.Engine {
	ginRouter := gin.Default()
	v1 := ginRouter.Group("/api/v1")
	{

		v1.GET("ping", func(c *gin.Context) {
			c.JSON(200, "success")
		})
		workflow := v1.Group("/workflows")
		{
			// 创建工作流
			workflow.POST("", curl.HttpRequest)
			//查看单个工作流
			workflow.GET("/:id", curl.HttpRequest)
			//查询全部工作流
			workflow.GET("", curl.HttpRequest)
			//删除单个工作流
			workflow.DELETE("/:id", curl.HttpRequest)
			//更新单个工作流
			workflow.PUT("/:id", curl.HttpRequest)
			//激活工作流
			workflow.POST("/:id/activate", curl.HttpRequest)
			//关闭工作流
			//workflow.POST("/execute", curl.HttpRequest)

		}
		webhook := v1.Group("webhook")
		{
			//执行工作流
			webhook.POST("/execute", curl.HttpRequest)
		}

	}

	return ginRouter
}
