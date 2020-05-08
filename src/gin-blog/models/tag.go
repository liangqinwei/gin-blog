package models

import (
	"ginBlog/src/gin-blog/package/exception"
	"github.com/jinzhu/gorm"
	"math/big"
	"time"
)

type Tag struct {
	Model
	Name         string    `json:"name"`
	Created_by   string    `json:"create_by"`
	Modified_by  string    `json:"modified_by"`
	Summary      string    `json:"summary"`
	State        int       `json:"state"`
	Deleted_time *time.Time `json:deleted_time`
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	//cmp前大于参数，返回1，等于返回0，小于返回-1
	//if tag.Id.Cmp(big.NewInt(int64(0)))==1{
	if big.NewInt( int64(tag.Id)).Cmp(big.NewInt(int64(0)))==1{
		return true
	}
	return false
}

//创建前更新create时间
func (t *Tag) BeforeCreate(scope *gorm.Scope) error{
	scope.SetColumn("CreateTime",time.Now())
	scope.SetColumn("UpdateTime",time.Now())
	return nil
}

//创建前更新create时间
func (t *Tag) BeforeUpdate(scope *gorm.Scope) error{
	scope.SetColumn("UpdateTime",time.Now())
	return nil
}

func DeleteTag(id int) bool {
	db.Where("id=?",id).Delete(&Tag{})
	return true
}

func EditTag(id int,data interface{}) bool {
	db.Model(&Tag{}).Where("id=?",id).Update(data)
	return true
}

func (Tag)TableName()  string{
	return "blog_tag"
}

func InsertTag(tag Tag) {

	if err := db.Create(&tag).Error; err != nil {
		panic(exception.SqlError(err.Error()))
	}

}

