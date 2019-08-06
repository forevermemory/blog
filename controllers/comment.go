package controllers

import (
	"bee/blog/models"
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"strconv"
)

type CommentController struct {
	beego.Controller
}

// @Title  添加评论
// @Description 添加评论
// @Param	articleId formDate string true "文章id"
// @Param	userId formDate string	true "评论人id"
// @Param	content	formDate string	true	"评论内容"
// @Success 200 {object} models.Category
// @router /add [post]
func (this *CommentController) Add() {

	//接受参数
	content := this.GetString("content")
	if !(utils.TrimString(content) > 0) {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": "content is null"}
		this.ServeJSON()
		return
	}

	userId, getInterr := this.GetInt("userId")
	if getInterr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": getInterr.Error()}
		this.ServeJSON()
		return
	}

	articleId, articleIdErr := this.GetInt("articleId")
	if articleIdErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": articleIdErr.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	// 查出数据

	user := models.User{
		Id: userId,
	}
	if readUserErr := o.Read(&user); readUserErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": readUserErr.Error()}
		this.ServeJSON()
		return
	}

	article := models.Article{
		Id: articleId,
	}
	if readErr2 := o.Read(&article); readErr2 != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": readErr2.Error()}
		this.ServeJSON()
		return
	}

	comment := models.Comment{
		Content: content,
		User:    &user,
		Article: &article,
	}
	_, insertErr := o.Insert(&comment)
	if insertErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "6", "msg": insertErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &comment
	this.ServeJSON()
	return

}

// @Title  删除评论
// @Description 删除评论
// @Param	commentId query string true "评论id"
// @Success 200 {object} models.Category
// @router /delete/:id [get]
func (this *CommentController) DeleteComment() {
	//获取当前 用户的id   与文章作者的 id 想比较 相等显示出删除按钮    或者管理员显示删除按钮
	strCommentId := this.Ctx.Input.Param(":id")

	commentId, atoiErr := strconv.Atoi(strCommentId)
	if atoiErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": atoiErr.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	comment := models.Comment{Id: commentId}

	if _, deleteErr := o.Delete(&comment); deleteErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": deleteErr.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"code": "3", "msg": "delete ok"}
	this.ServeJSON()
	return
}
