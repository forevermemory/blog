package controllers

import (
	"bee/blog/models"
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// Operations about SubCategory
type TagController struct {
	beego.Controller
}

// @Title  新增一个标签
// @Description 新增一个标签
// @Param name formDate string true "标签名称"
// @Success 200 {object} models.Tag
// @Failure 400 register failed
// @router /add [post]
func (this *TagController) Add() {

	name := this.GetString("name")

	if utils.TrimString(name) <= 0 {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": "name is null"}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	tag := models.Tag{
		Name: name,
	}
	if _, insertErr := o.Insert(&tag); insertErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": insertErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &tag
	this.ServeJSON()
	return

}

// @Title  根据标签查询出来对应类型的文章
// @Description 根据标签查询出来对应类型的文章
// @Param id query int true "标签id"
// @Success 200 {object} models.article
// @Failure 400 register failed
// @router /getArticle [get]
func (this *TagController) GetAllArticle() {

	id, getIntErr := this.GetInt("id")
	if getIntErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": getIntErr.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()

	tag := models.Tag{Id: id}
	if readTagErr := o.Read(&tag); readTagErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readTagErr.Error()}
		this.ServeJSON()
		return
	}

	// 查询关联
	if _, loadRelatedErr := o.LoadRelated(&tag, "article"); loadRelatedErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": loadRelatedErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &tag
	this.ServeJSON()
	return
}

// @Title  查询所有标签
// @Description 查询所有标签
// @Success 200 {object} models.Tag
// @Failure 400 register failed
// @router /getall [get]
func (this *TagController) GetAll() {

	o := orm.NewOrm()
	var tags []models.Tag

	if _, err := o.QueryTable("tag").All(&tags); err != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &tags
	this.ServeJSON()
	return

}
