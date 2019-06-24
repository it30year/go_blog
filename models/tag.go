package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func GetTags(pagNum int, pagesize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Limit(pagesize).Find(&tags)
	return
}
func Gett() (tags []Tag) {
	db.Find(&tags)
	return
}
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}
