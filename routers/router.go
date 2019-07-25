package routers

import (
	"bin_blog/pkg/setting"
	v1 "bin_blog/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "tes123123t",
			})
		})
		//获取文章列表
		apiv1.GET("/tag", v1.GetTags)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddTags)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.UpdateTags)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteTags)
	}
	return r
}
