package controllers

import (
	"bee/blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
)

// Operations about SubCategory
type SubCategoryController struct {
	beego.Controller
}

// @Title  编辑分类后保存二级分类
// @Description 编辑分类后保存二级分类
// @Param id formDate int true "分类的id"
// @Param sub_id formDate int true "二级分类的id"
// @Param name formDate string true "新二级分类名称"
// @Success 200 {object} models.SubCategory
// @Failure 400 register failed
// @router /update [post]
func (this *SubCategoryController) Update() {

	o := orm.NewOrm()
	//接收参数 验证参数合法性

	name := this.GetString("name")
	id, err1 := this.GetInt("id")
	if err1 != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err1.Error()}
		this.ServeJSON()
		return
	}

	sub_id, err2 := this.GetInt("sub_id")
	if err2 != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": err2.Error()}
		this.ServeJSON()
		return
	}

	category := models.Category{
		Id: id,
	}

	readErr := o.Read(&category)
	if readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}

	log.Printf("category的类型  --%T", category)
	log.Printf("category的类型  --%T", category)
	log.Printf("category的类型  --%T", category)

	sub_category := models.SubCategory{
		Id:       sub_id,
		Name:     name,
		Category: &category,
	}
	_, UpdateErr := o.Update(&sub_category)
	if UpdateErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": UpdateErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &sub_category
	this.ServeJSON()
	return
}

// @Title  查询一个
// @Description 根据id查询分类
// @Param id path string true "二级分类的id"
// @Success 200 {object} models.SubCategory
// @Failure 400 register failed
// @router /get/:id [get]
func (this *SubCategoryController) Get() {

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

	sub_category := models.SubCategory{}

	readErr := o.QueryTable("sub_category").RelatedSel().Filter("id__exact", newId).One(&sub_category)
	if readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &sub_category
	this.ServeJSON()
	return
}

// @Title  查询所有二级分类
// @Description 查询所有二级分类
// @Success 200 {object} models.SubCategory
// @Failure 400 register failed
// @router /getall [get]
func (this *SubCategoryController) GetAll() {

	o := orm.NewOrm()

	var sub_categories []models.SubCategory
	_, err := o.QueryTable("sub_category").RelatedSel().All(&sub_categories)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &sub_categories
	this.ServeJSON()
	return

}

// @Title  新增二级分类
// @Description 新增二级分类
// @Param id formDate int true "对应一级分类的id"
// @Param name formDate string true "新增分类名称"
// @Success 200 {object} models.SubCategory
// @Failure 400 add failed
// @router /add [post]
func (this *SubCategoryController) Add() {

	o := orm.NewOrm()
	name := this.GetString("name")
	id, err := this.GetInt("id")
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	category := models.Category{
		Id: id,
	}
	if readErr := o.Read(&category); readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}

	sub_category := models.SubCategory{
		Name:     name,
		Category: &category,
	}
	_, insertErr := o.Insert(&sub_category)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": insertErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &sub_category
	this.ServeJSON()
	return

}

// @Title  删除二级分类
// @Description 删除二级分类
// @Param id path string true "二级分类的id"
// @Success 200 {object} models.SubCategory
// @Failure 400 add failed
// @router /delete/:id [get]
func (this *SubCategoryController) Delete() {

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

	sub_category := models.SubCategory{
		Id: newId,
	}
	_, Delerr := o.Delete(&sub_category)

	if Delerr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": Delerr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": "3", "msg": "ok"}
	this.ServeJSON()
	return

}
