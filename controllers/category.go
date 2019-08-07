package controllers

import (
	"bee/blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// Operations about Category
type CategoryController struct {
	beego.Controller
}

// @Title  编辑分类后保存
// @Description 修改分类
// @Param id formDate int true "分类的id"
// @Param name formDate string true "新的分类名称"
// @Param url formDate string true "新的分类名称"
// @Success 200 {object} models.Category
// @Failure 400 register failed
// @router /update [post]
func (this *CategoryController) Update() {

	o := orm.NewOrm()
	//接收参数 验证参数合法性

	name := this.GetString("name")
	url := this.GetString("url")
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"msg": "param error"}
		this.ServeJSON()
		return
	}

	category := models.Category{
		Id: id,
	}

	readErr := o.Read(&category)
	if readErr != nil {
		this.Data["json"] = map[string]interface{}{"msg": " not sxist"}
		this.ServeJSON()
		return
	}

	category.Name = name
	category.Url = url
	_, UpdateErr := o.Update(&category)
	if UpdateErr != nil {
		this.Data["json"] = map[string]interface{}{"msg": "update fail"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &category
	this.ServeJSON()
	return
}

// @Title  查询一个
// @Description 根据id查询分类
// @Param id path string true "分类的id"
// @Success 200 {object} models.Category
// @Failure 400 register failed
// @router /get/:id [get]
func (this *CategoryController) Get() {

	o := orm.NewOrm()
	//接收参数 验证参数合法性
	id := this.Ctx.Input.Param(":id")
	// log.Printf("id的类型为%T", id) //string

	newId, err := strconv.Atoi(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	category := models.Category{}

	readErr := o.QueryTable("category").RelatedSel().Filter("id__exact", newId).One(&category)
	if readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &category
	this.ServeJSON()
	return
}

// @Title  查询所有分类
// @Description 查询所有分类
// @Success 200 {object} models.Category
// @Failure 400 register failed
// @router /getall [get]
func (this *CategoryController) GetAll() {

	o := orm.NewOrm()

	var categories []models.Category
	_, err := o.QueryTable("category").RelatedSel().All(&categories)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &categories
	this.ServeJSON()
	return

}

// @Title  新增分类
// @Description 新增分类
// @Param name formDate string true "新增分类名称"
// @Param url formDate string true "新增分类名称"
// @Success 200 {object} models.Category
// @Failure 400 add failed
// @router /add [post]
func (this *CategoryController) Add() {

	o := orm.NewOrm()
	name := this.GetString("name")
	url := this.GetString("url")
	category := models.Category{
		Name: name,
		Url:  url,
	}
	_, err := o.Insert(&category)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &category
	this.ServeJSON()
	return

}

// @Title  删除分类
// @Description 删除分类
// @Param id path string true "分类的id"
// @Success 200 {object} models.Category
// @Failure 400 add failed
// @router /delete/:id [get]
func (this *CategoryController) Delete() {

	o := orm.NewOrm()
	//接收参数 验证参数合法性
	id := this.Ctx.Input.Param(":id")
	// log.Printf("id的类型为%T", id) //string

	newId, err := strconv.Atoi(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	category := models.Category{
		Id: newId,
	}
	_, Delerr := o.Delete(&category)

	if Delerr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": Delerr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": "3", "msg": "ok"}
	this.ServeJSON()
	return

}
