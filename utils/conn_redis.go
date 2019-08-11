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

func GetDefaultRedisPool() *redis.Pool {

	// conn, err := redis.Dial("tcp", "129.211.19.13:6379")
	// // conn, err := redis.Dial("tcp", "129.211.19.13:6379", redis.DialPassword(`"$129.211.19.13$"`))
	// return conn, err

	pool := &redis.Pool{
		// Other pool configuration not shown in this example.
		MaxActive: 1024,
		MaxIdle:   16,
		Wait:      true,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", beego.AppConfig.String("redisPool"))
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", beego.AppConfig.String("redisPoolPassword")); err != nil {
				c.Close()
				return nil, err
			}
			if _, err := c.Do("SELECT", 0); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
	}

	return pool

}

// 新增文章 评论 需要登录后才可以  必须在query传递 token=xxxxxxx
var GetRequestTokenFilter = func(ctx *context.Context) {
	pool := GetDefaultRedisPool()

	// 获取token
	token := ctx.Input.Query("token")

	// 到redis中查找token是否存在   0 不存在  1存在

	conn := pool.Get()
	defer conn.Close()
	res, _ := redis.String(conn.Do("get", token))
	if len(res) <= 0 {
		log.Println("token is not exist......")

		ctx.WriteString("token is invalid")
		// ctx.Output.JSON(`{"code":"0","msg":"token is null"}`, false, false)
	}
	log.Println("is running......")
}
