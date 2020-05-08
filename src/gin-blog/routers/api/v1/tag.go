package v1

import (
	"ginBlog/src/gin-blog/dto"
	"ginBlog/src/gin-blog/models"
	"ginBlog/src/gin-blog/package/exception"
	"ginBlog/src/gin-blog/package/setting"
	"ginBlog/src/gin-blog/utils"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

//获取tag标签
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
	code := exception.SUCCESS
	data["lists"] = models.GetTags(utils.GetPage(c), setting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  exception.GetMsg(code),
		"data": data,
	})

}



//新增标签
func AddTag(c *gin.Context) {

	var json dto.AddTagEntity
	if err := c.ShouldBind(&json); err != nil {

		panic(exception.BodyError("请求参数有误",nil))
		return
	}
	if err:= utils.Validate.Struct(json);err!=nil{
		utils.NewValidatorError(err)
	}
	var tag models.Tag
	tag.Name=json.Name
	tag.State=json.State
	tag.Created_by=json.CreatedBy
	tag.Deleted_time=json.DeletedTime
	tag.Modified_by=json.ModifiedBy
	tag.Summary=json.Summary
	models.InsertTag(tag)
	c.JSON(http.StatusOK, gin.H{"code":"200","data":nil,"msg":"success"})
	return

}

//批量新增标签
func AddTags(c *gin.Context) {

	var jsons dto.AddTagEntities
	if err := c.ShouldBind(&jsons); err != nil {

		panic(exception.BodyError("请求参数有误",nil))
		return
	}
	if err:= utils.Validate.Struct(jsons);err!=nil{
		utils.NewValidatorError(err)
	}

	for _,json:=range jsons.Params{
		var tag models.Tag
		tag.Name=json.Name
		tag.State=json.State
		tag.Created_by=json.CreatedBy
		tag.Deleted_time=json.DeletedTime
		tag.Modified_by=json.ModifiedBy
		tag.Summary=json.Summary
		models.InsertTag(tag)
	}

	c.JSON(http.StatusOK, gin.H{"code":"200","data":nil,"msg":"success"})
	return

}



//修改标签
func EditTag(c *gin.Context) {

}

//删除标签
func DeleteTag(c *gin.Context) {

}

func TestTag(c *gin.Context) {
	panic(exception.ServerError())
}
