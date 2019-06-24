package v1

import (
	"bin_blog/pkg/e"
	"github.com/gin-gonic/gin"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}
	//var state int =-1
	//if arg :=c.Query("state"); arg !=""{
	//	state =com.StrTo(arg).MustInt()
	//	maps["state"] =state
	//}

	code := e.SUCCESS

	//data["list"] =models.GetTags(2,setting.PageSize,maps)
	//data["count"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTags(c *gin.Context) {

}

//修改文章标签
func UpdateTags(c *gin.Context) {

}

//删除文章标签
func DeleteTags(c *gin.Context) {

}
