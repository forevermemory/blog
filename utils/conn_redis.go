package utils

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/context"
	"github.com/gomodule/redigo/redis"
	"log"
)

func GetRedisConn() (cache.Cache, error) {

	conn, cacheError := cache.NewCache("redis", beego.AppConfig.String("redis"))
	return conn, cacheError
}

func GetDefaultRedisConn() (redis.Conn, error) {

	conn, err := redis.Dial("tcp", "129.211.19.13:6379")
	return conn, err

}

// 新增文章 评论 需要登录后才可以  必须在query传递 token=xxxxxxx
var GetRequestTokenFilter = func(ctx *context.Context) {
	conn, _ := GetDefaultRedisConn()
	defer conn.Close()
	// 获取token
	token := ctx.Input.Query("token")

	// 到redis中查找token是否存在   0 不存在  1存在

	res, _ := redis.String(conn.Do("get", token))
	if len(res) <= 0 {
		log.Println("token is not exist......")

		ctx.WriteString("token is invalid")
		ctx.Output.JSON(`{"code":"0","msg":"token is null"}`, false, false)
	}
	log.Println("is running......")
}
