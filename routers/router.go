package routers

import (
	"bin_blog/models"
	"bin_blog/pkg/e"
	"bin_blog/pkg/setting"
	v1 "bin_blog/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
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
				"message": "test",
			})
		})
		apiv1.GET("/tq", GetTags)
		//查
		apiv1.GET("/tags", v1.GetTags)
		//增
		apiv1.POST("/tags", v1.AddTags)
		//改
		apiv1.PUT("/tags/:id", v1.UpdateTags)
		//删
		apiv1.DELETE("/tags/:id", v1.DeleteTags)
	}
	return r
}

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	//var state int =1
	//if arg :=c.Query("state"); arg !=""{
	//	state =com.StrTo(arg).MustInt()
	//	maps["state"] =state
	//}

	code := e.SUCCESS

	data["list"] = models.Gett()

	//	data["list"] =models.GetTags(2,setting.PageSize,maps)
	//data["count"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
