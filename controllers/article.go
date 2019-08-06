package controllers

import (
	"bee/blog/models"
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
)

// Operations about Article
type ArticleController struct {
	beego.Controller
}

type addArticle struct {
	UserId      int    `form:"user_id"`
	Title       string `form:"title"`
	Content     string `form:"content"`
	Category    int    `form:"category"`
	SubCategory int    `form:"sub_category"`
}

// @Title  新增文章
// @Description 新增文章
// @Param	user_id	formDate  int true	"当前用户id"
// @Param	title	formDate	string true	 "文章title"
// @Param	content	formDate string	true	"文章内容"
// @Param	category	formDate int	true	"一级分类"
// @Param	sub_category formDate int	true	"二级分类"
// @Success 200 {strings} "ok"
// @router /add [post]
func (this *ArticleController) Add() {

	o := orm.NewOrm()

	var a addArticle
	if errParseForm := this.ParseForm(&a); errParseForm != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": errParseForm.Error()}
		this.ServeJSON()
		return
	}

	if a.Title == "" || len(strings.TrimSpace(a.Title)) == 0 {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": "title is null"}
		this.ServeJSON()
		return
	}

	if a.Content == "" || len(strings.TrimSpace(a.Content)) == 0 {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": "content is null"}
		this.ServeJSON()
		return
	}

	//关联关系处理
	// author
	author := models.User{
		Id: a.UserId,
	}
	if errauthor := o.Read(&author); errauthor != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": errauthor.Error()}
		this.ServeJSON()
		return
	}

	//category
	category := models.Category{
		Id: a.Category,
	}
	if errcategory := o.Read(&category); errcategory != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": errcategory.Error()}
		this.ServeJSON()
		return
	}

	//sub_category
	sub_category := models.SubCategory{
		Id: a.SubCategory,
	}
	if errsub_category := o.Read(&sub_category); errsub_category != nil {
		this.Data["json"] = map[string]interface{}{"code": "6", "msg": errsub_category.Error()}
		this.ServeJSON()
		return
	}
	// 标签？？？

	article := models.Article{
		Title:       a.Title,
		Content:     a.Content,
		Author:      &author,
		Category:    &category,
		SubCategory: &sub_category,
	}

	conn, _ := utils.GetDefaultRedisConn()
	defer conn.Close()

	//插入数据库
	_, insertErr := o.Insert(&article)
	if insertErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "7", "msg": insertErr.Error()}
		this.ServeJSON()
		return
	}

	// 初始化view 和 zan 数量
	if _, doErr := conn.Do("hmset", "article_"+strconv.Itoa(article.Id), "view", "0", "zan", "0"); doErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "8", "msg": doErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &article
	this.ServeJSON()
	return
}

// @Title  查询一篇文章
// @Description 查询一篇文章
// @Param	id	path 	string true	 "根据id查询详情，点击进入文章详情"
// @Success 200 {object} models.Article
// @router /get/:id [get]
func (this *ArticleController) Get() {
	conn, cacheError := utils.GetDefaultRedisConn()
	defer conn.Close()

	if cacheError != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": cacheError.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	//接收参数 验证参数合法性
	id := this.Ctx.Input.Param(":id")
	// log.Printf("id的类型为%T", id) //string

	newId, err := strconv.Atoi(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	// 查询文章基本信息
	article := models.Article{Id: newId}
	if readErr := o.Read(&article); readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}

	// 查询文章的一级评论
	if _, loadErr := o.LoadRelated(&article, "SubComment"); loadErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": loadErr.Error()}
		this.ServeJSON()
		return
	}

	// 查询文章的二级评论
	if _, loadcommentErr := o.LoadRelated(&article, "comment"); loadcommentErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": loadcommentErr.Error()}
		this.ServeJSON()
		return
	}

	// // 设置阅读数量自增 1   没有自动创建
	// if incrErr := conn.Incr("article_" + id); incrErr != nil {
	// 	this.Data["json"] = map[string]interface{}{"code": "6", "msg": incrErr.Error()}
	// 	this.ServeJSON()
	// 	return
	// }

	// 设置 view 自增 1
	if _, incrbyErr := conn.Do("hincrby", "article_"+id, "view", 1); incrbyErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "7", "msg": incrbyErr.Error()}
		this.ServeJSON()
		return
	}

	// 获取 view 的最新值

	view, viewErr := redis.Int(conn.Do("hget", "article_"+id, "view"))
	if viewErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "8", "msg": viewErr.Error()}
		this.ServeJSON()
		return
	}

	// log.Printf("v的类型为 %T", v) //v的类型为 []uint8
	// intvalue, _ := strconv.Atoi(string(v.([]byte))) //int
	// log.Printf("value:%T", intvalue)
	//HINCRBY key field 1

	// 设置阅读数量
	article.View = view

	this.Data["json"] = &article
	this.ServeJSON()
	return
}

// @Title  条件查询所有分类/作者/标签
// @Description 条件查询所有
// @Param	page query string false "页码"
// @Param	title query string false "标题"
// @Param	tags query string false "标签"
// @Param	authorId query string false "作者"
// @Param	categoryId query string false "分类"
// @Param	subCategoryId query string false "二级分类"
// @Success 200 {object} models.Article
// @router /getall [get]
func (this *ArticleController) GetList() {

	var (
		page int = 1
	)

	// 查询
	o := orm.NewOrm()
	qs := o.QueryTable("article")

	// title   标题
	title := this.GetString("title")
	if utils.TrimString(title) > 0 {
		qs = qs.Filter("title__icontains", title)
	}

	// 作者 id
	strAuthorId := this.GetString("authorId")
	if utils.TrimString(strAuthorId) > 0 {
		authorId, authorErr := strconv.Atoi(strAuthorId)
		if authorErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "1", "msg": authorErr.Error()}
			this.ServeJSON()
			return
		}

		qs = qs.Filter("author__id__exact", authorId)
	}

	// 分类 id
	strCategoryId := this.GetString("categoryId")
	if utils.TrimString(strCategoryId) > 0 {
		categoryId, categoryErr := strconv.Atoi(strCategoryId)
		if categoryErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "2", "msg": categoryErr.Error()}
			this.ServeJSON()
			return
		}
		qs = qs.Filter("category__id__exact", categoryId)

	}

	// 二级分类 id

	strSubCategoryId := this.GetString("subCategoryId")
	if utils.TrimString(strSubCategoryId) > 0 {
		subCategoryId, subCategoryErr := strconv.Atoi(strSubCategoryId)
		if subCategoryErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "3", "msg": subCategoryErr.Error()}
			this.ServeJSON()
			return
		}
		qs = qs.Filter("subCategory__id__exact", subCategoryId)

	}

	// page
	strPage := this.GetString("page", "1")
	if utils.TrimString(strPage) > 0 {
		var pageErr error
		page, pageErr = strconv.Atoi(strPage)
		if pageErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "4", "msg": pageErr.Error()}
			this.ServeJSON()
			return
		}
		if page <= 0 {
			page = 1
		}

		pagesize, _ := beego.AppConfig.Int("pagesize")
		qs = qs.Limit(pagesize).Offset((page - 1) * pagesize)
	}

	// 标签组

	// TODO

	var articles []models.Article
	_, allErr := qs.All(&articles)
	if allErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": allErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &articles
	this.ServeJSON()
}

// @Title  删除一篇文章，
// @Description 删除一篇文章，
// @Param	id path string true "删除一篇文章，用于后台"
// @Success 200 {object} models.Article
// @router /delete/:id [get]
func (this *ArticleController) DeleteById() {
	conn, cacheError := utils.GetRedisConn()

	if cacheError != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": cacheError.Error()}
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	//接收参数 验证参数合法性
	id := this.Ctx.Input.Param(":id")
	// log.Printf("id的类型为%T", id) //string

	newId, err := strconv.Atoi(id)

	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	article := models.Article{
		Id: newId,
	}

	// 删除文章
	if _, deleteErr := o.Delete(&article); deleteErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": deleteErr.Error()}
		this.ServeJSON()
	}

	// 删除redis中的key
	if deleteErr := conn.Delete("article_" + id); deleteErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": deleteErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": "4", "msg": "ok"}
	this.ServeJSON()
	return
}

// @Title  修改文章
// @Description 修改文章
// @Param	article_id	formDate  int true	"当前文章id"
// @Param	title formDate string true "文章title"
// @Param	content formDate string true "文章内容"
// @Success 200 {object} models.Article
// @router /edit [post]
func (this *ArticleController) Edit() {

	o := orm.NewOrm()

	//获取参数

	article_id, getArticleErr := this.GetInt("article_id")
	if getArticleErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": getArticleErr.Error()}
		this.ServeJSON()
		return
	}

	title := this.GetString("title")
	if title == "" || len(strings.TrimSpace(title)) == 0 {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": "title is null"}
		this.ServeJSON()
		return
	}

	content := this.GetString("content")
	if content == "" || len(strings.TrimSpace(content)) == 0 {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": "content is null"}
		this.ServeJSON()
		return
	}

	// 标签  tags TODO

	article := models.Article{
		Id: article_id,
	}

	if readErr := o.Read(&article); readErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": readErr.Error()}
		this.ServeJSON()
		return
	}
	//插入数据库

	article.Content = content
	article.Title = title

	_, updateErr := o.Update(&article)
	if updateErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": updateErr.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &article
	this.ServeJSON()
	return
}
