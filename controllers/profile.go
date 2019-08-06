package controllers

import (
	"bee/blog/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

// Operations about profile
type ProfileController struct {
	beego.Controller
}

type profil struct {
	UserId      int    `form:"userId"`
	Gender      int    `form:"gender"`
	Age         int    `form:"age"`
	Telephone   string `form:"telephone"`
	Address     string `form:"address"`
	Avator      string `form:"avator"`
	Personality string `form:"personality"`
}

// @Title  完善用户信息
// @Description 完善用户信息
// @Param	userId formDate int	true	"用户id"
// @Param	gender formDate int	true	"性别"
// @Param	age formDate int	true	"年龄"
// @Param	address formDate string	true	"地址"
// @Param	telephone formDate string	true	"手机"
// @Param	avator formDate string true "头像的路径"
// @Param	personality formDate string true "个人简介"
// @Success 200 {strings} "ok"
// @router /profile/add [post]
func (this *ProfileController) Add() {
	o := orm.NewOrm()

	// 解析到结构体
	p := profil{}
	if parseFormErr := this.ParseForm(&p); parseFormErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": parseFormErr.Error()}
		this.ServeJSON()
		return
	}

	user := models.User{
		Id: p.UserId,
	}

	if readUserErr := o.Read(&user); readUserErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readUserErr.Error()}
		this.ServeJSON()
		return
	}

	f, h, getFileErr := this.GetFile("avator")
	// 没有传文件 发生异常
	if getFileErr != nil {
		//invalid memory address or nil pointer dereference
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": "请上传头像"}
		this.ServeJSON()
		return
	}
	defer f.Close()

	path := "static/avator/" + user.Username + "_" + h.Filename
	if saveFileErr := this.SaveToFile("avator", path); saveFileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": getFileErr.Error()}
		this.ServeJSON()
		return
	}
	// 保存位置在 static/upload, 没有文件夹要先创建

	profile := models.Profile{
		Gender:      p.Gender,
		Age:         p.Age,
		Address:     p.Address,
		Telephone:   p.Telephone,
		Avator:      path,
		Personality: p.Personality,
		User:        &user,
	}

	if _, insertProfileErr := o.Insert(&profile); insertProfileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": insertProfileErr.Error()}
		this.ServeJSON()
		return
	}

	// 成功之后返回数据

	this.Data["json"] = map[string]interface{}{"code": "6", "msg": "ok"}
	this.ServeJSON()
	return

}

// @Title  查询一个
// @Description 查询一个
// @Param	userId path int true "用户id"
// @Success 200 {strings} "ok"
// @router /profile/get/:userId [get]
func (this *ProfileController) Get() {
	o := orm.NewOrm()

	strId := this.Ctx.Input.Param(":userId")
	id, atoiErr := strconv.Atoi(strId)
	if atoiErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": atoiErr.Error()}
		this.ServeJSON()
		return
	}

	user := models.User{Id: id}
	if readUserErr := o.Read(&user); readUserErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": readUserErr.Error()}
		this.ServeJSON()
		return
	}

	profile := models.Profile{}

	if queryTableErr := o.QueryTable("profile").Filter("user__exact", user).One(&profile); queryTableErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": queryTableErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &profile
	this.ServeJSON()
	return

}

// @Title  更新用户profile信息
// @Description 更新用户profile信息
// @Param	profileId formDate int	true	"用户 的 profileId"
// @Param	gender formDate int	true	"性别"
// @Param	age formDate int	true	"年龄"
// @Param	address formDate string	true	"地址"
// @Param	telephone formDate string	true	"手机"
// @Param	personality formDate string true "个人简介"
// @Success 200 {strings} "ok"
// @router /profile/update [post]
func (this *ProfileController) Update() {
	o := orm.NewOrm()

	strId := this.GetString("profileId")
	profileId, err := strconv.Atoi(strId)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	// 解析到结构体
	p := profil{}
	if parseFormErr := this.ParseForm(&p); parseFormErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": parseFormErr.Error()}
		this.ServeJSON()
		return
	}

	profile := models.Profile{Id: profileId}
	if readProfileRawErr := o.Read(&profile); readProfileRawErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": readProfileRawErr.Error()}
		this.ServeJSON()
		return
	}

	// 设置新的值
	profile.Gender = p.Gender
	profile.Age = p.Age
	profile.Address = p.Address
	profile.Telephone = p.Telephone
	profile.Personality = p.Personality

	if _, updateProfileErr := o.Update(&profile); updateProfileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": updateProfileErr.Error()}
		this.ServeJSON()
		return
	}

	// 成功之后返回数据

	this.Data["json"] = &profile
	this.ServeJSON()
	return

}

// @Title  更新用户头像
// @Description 更新用户头像
// @Param	profileId formDate string	true	"用户 的 profileId"
// @Param	userId formDate string	true	"用户 的 userId"
// @Param	avator formDate string true "头像"
// @Success 200 {strings} "ok"
// @router /profile/update/avator [post]
func (this *ProfileController) UpdateAvator() {
	o := orm.NewOrm()

	strProfileId := this.GetString("profileId")
	profileId, err := strconv.Atoi(strProfileId)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	strUserId := this.GetString("userId")
	userId, strconverr := strconv.Atoi(strUserId)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": strconverr.Error()}
		this.ServeJSON()
		return
	}
	user := models.User{Id: userId}
	if readUserErr := o.Read(&user); readUserErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": readUserErr.Error()}
		this.ServeJSON()
		return
	}

	f, h, getFileErr := this.GetFile("avator")
	if getFileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": "请上传头像"}
		this.ServeJSON()
		return
	}
	defer f.Close()

	profile := models.Profile{
		Id: profileId,
	}
	if readProfileErr := o.Read(&profile); readProfileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": getFileErr.Error()}
		this.ServeJSON()
		return
	}

	path := "static/avator/" + user.Username + "_" + h.Filename
	if saveFileErr := this.SaveToFile("avator", path); saveFileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "6", "msg": getFileErr.Error()}
		this.ServeJSON()
		return
	}

	profile.Avator = path

	if _, updateProfileErr := o.Update(&profile); updateProfileErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "7", "msg": updateProfileErr.Error()}
		this.ServeJSON()
		return
	}

	// 成功之后返回数据

	this.Data["json"] = &profile
	this.ServeJSON()
	return

}
