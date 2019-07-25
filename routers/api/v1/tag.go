package v1

import "C"
import (
	"bin_blog/models"
	"bin_blog/pkg/e"
	"bin_blog/pkg/setting"
	"bin_blog/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
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
	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}
	code := e.SUCCESS
	data["list"] = models.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["count"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签
func AddTags(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createBy := c.Query("created_by")
	//验证
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name, 9, "name").Message("名称最长为9字符")
	valid.Required(createBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createBy, 9, "created_by").Message("创建人最长不超过9字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTags(name, state, createBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//修改文章标签
func UpdateTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modifiedBy ")
	vaild := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state := com.StrTo(arg).MustInt()
		vaild.Range(state, 0, 1, "state").Message("状态只能0或者1")
	}
	vaild.Required(id, "id").Message("Id不能为空")
	vaild.Required(modifiedBy, "modifiedBy").Message("修改人不能为空")
	vaild.MaxSize(modifiedBy, 10, "modifiedBy").Message("长度最长为10")
	vaild.MaxSize(name, 5, "name").Message("长度最长为5")

	code := e.INVALID_PARAMS
	if !vaild.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			data := make(map[string]interface{})
			data["modified_by"] = modifiedBy
			if name != "" {
				data["name"] = name
			}
			if state != -1 {
				data["state"] = state
			}
			models.EditTag(id, data)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})

}

//删除文章标签
func DeleteTags(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	vaild := validation.Validation{}
	vaild.Min(id, 1, "id").Message("iD必须大于0")

	code := e.INVALID_PARAMS

	if !vaild.HasErrors() {
		code = e.SUCCESS
		if models.ExistTagByID(id) {
			models.DeleteTag(id)
		} else {
			code = e.ERROR_NOT_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
