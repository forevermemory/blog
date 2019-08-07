package models

import (
	"github.com/astaxie/beego/orm"
	"log"
	"time"
)

func init() {

	log.Println("init registerModels ...")

	regisertModels()
	log.Println("finish   init registerModels ...")

}

func regisertModels() {
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Profile))

	orm.RegisterModel(new(Category))
	orm.RegisterModel(new(SubCategory))
	orm.RegisterModel(new(Article))
	orm.RegisterModel(new(Tag))
	orm.RegisterModel(new(Comment))
	orm.RegisterModel(new(SubComment))
	orm.RegisterModel(new(Message))

}

//留言板
type Message struct {
	Id      int
	Date    time.Time `orm:"auto_now_add;type(datetime)"`
	Content string    `orm:""  description:"xss  cross-site scriptong 跨站脚本漏洞"`
}

type Category struct {
	Id          int
	Name        string
	Url         string
	SubCategory []*SubCategory `orm:"reverse(many)"`
	Article     []*Article     `orm:"reverse(many)"`
}

//   多  分类下面对应子分类
type SubCategory struct {
	Id       int
	Name     string
	Url      string
	Category *Category  `orm:"rel(fk)"`
	Article  []*Article `orm:"reverse(many)"`
}

type Article struct {
	Id         int
	Title      string
	Content    string    `orm:" type(text)"`
	State      int       `orm:"null;default(0)"  description:"这是状态字段 0 已发表    1 已删除"`
	Image      string    `orm:"null"`
	View       int       `orm:"null;default(0)"  description:"阅读数量"`
	Zan        int       `orm: "null;default(0)"  description:"点赞数量"`
	CreateDate time.Time `orm:"auto_now_add;type(datetime)"`
	UpdateDate time.Time `orm:"auto_now;type(datetime)"`

	Author      *User        `orm:"rel(fk)"`
	Category    *Category    `orm:"rel(fk)"`
	SubCategory *SubCategory `orm:"rel(fk)"`
	Tag         []*Tag       `orm:"rel(m2m)"`

	Comment    []*Comment    `orm:"reverse(many)"`
	SubComment []*SubComment `orm:"reverse(many)"`
}

// 标签   与article 多对多
type Tag struct {
	Id      int
	Name    string
	Article []*Article `orm:"reverse(many)"`
}

//评论   与article   一对多
type Comment struct {
	Id      int
	Content string    `orm:"size(2048)"`
	Zan     int       `orm: description:"点赞数量"`
	Date    time.Time `orm:"auto_now_add;type(datetime)"`

	SubComment []*SubComment `orm:"reverse(many)"`
	// UserId       *User
	User    *User    `orm:"rel(fk)"`
	Article *Article `orm:"rel(fk)"`
}

//二级评论表  与一级评论  一对多
type SubComment struct {
	Id      int
	Content string    `orm:"size(2048)"`
	Zan     int       `orm: description:"点赞数量"`
	Date    time.Time `orm:"auto_now_add;type(datetime)"`

	Comment *Comment `orm:"rel(fk)"`
	// FromUser *User
	// ToUser   *User
	User    *User    `orm:"rel(fk)"`
	Article *Article `orm:"rel(fk)"`
}
