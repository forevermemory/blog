package controllers

import (
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/gomodule/redigo/redis"
	"strconv"
)

type MessageController struct {
	beego.Controller
}

func init() {

}

// @Title  留言板添加信息
// @Description 留言板添加信息
// @Param	content	formDate string	true	"留言板添加信息"
// @Success 200 {strings} "ok"
// @router /add [post]
func (this *MessageController) Add() {
	conn, cacheError := utils.GetDefaultRedisConn()

	defer conn.Close()
	if cacheError != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": cacheError.Error()}
		this.ServeJSON()
		return
	}

	content := this.GetString("content")

	// _, doErr := conn.Do("", args)
	_, err := conn.Do("lpush", "message", content)
	if err != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": err.Error()}
		this.ServeJSON()
		return
	}

	this.Data["json"] = map[string]interface{}{"code": "3", "msg": "ok"}
	this.ServeJSON()
	return

}

// @Title  取留言板记录
// @Description 取留言板记录
// @Param	page	query string	false	"默认展示十条,"
// @Success 200 {strings} "ok"
// @router /get [get]
func (this *MessageController) GetMessage() {

	conn, cacheError := utils.GetDefaultRedisConn()

	defer conn.Close()
	if cacheError != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": cacheError.Error()}
		this.ServeJSON()
		return
	}

	strPage := this.GetString("page", "1")
	page, pageErr := strconv.Atoi(strPage)
	if pageErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": pageErr.Error()}
		this.ServeJSON()
		return
	}

	// 查询
	pagesize := beego.AppConfig.DefaultInt("pagesize", 10)
	// sendErr := conn.Send("lrange", "message", (page-1)*pagesize, pagesize*page)
	// if sendErr != nil {
	// 	this.Data["json"] = map[string]interface{}{"code": "3", "msg": sendErr.Error()}
	// 	this.ServeJSON()
	// 	return
	// }

	// conn.Flush()
	// v, _ := conn.Receive()

	// log.Println(v)      //[[101 101 101 101 101] [100 100 100 100] [99 99 99] [98 98 98] [97 97 97]]
	// log.Printf("%T", v) // []interface {}
	// v1 := []string{}
	// for _, i := range v.([]interface{}) {
	// 	v1 = append(v1, string(i.([]byte)))
	// }

	v1, _ := redis.Strings(conn.Do("lrange", "message", (page-1)*pagesize, pagesize*page-1))

	// log.Println(v1)

	// log.Println(string(val.([]byte)))

	this.Data["json"] = &v1
	this.ServeJSON()
	return

}

// @Title  管理员删除一条记录
// @Description 管理员删除一条记录
// @Param	content	query string	true	"需要删除的item"
// @Success 200 {strings} "ok"
// @router /delete [get]
func (this *MessageController) DeleteByContent() {

	conn, cacheError := utils.GetDefaultRedisConn()
	defer conn.Close()
	if cacheError != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": cacheError.Error()}
		this.ServeJSON()
		return
	}

	content := this.GetString("content")

	// LREM KEY_NAME COUNT VALUE

	_, lremErr := conn.Do("lrem", "message", 1, content)
	if lremErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": lremErr.Error()}
		this.ServeJSON()
		return
	}

	// log.Println(string(val.([]byte)))

	this.Data["json"] = map[string]interface{}{"code": "2", "msg": "ok"}
	this.ServeJSON()
	return

}
