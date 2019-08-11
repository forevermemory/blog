package models

import (
	// "github.com/astaxie/beego/orm"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func init() {

}

type User struct {
	Id            int
	Username      string
	Password      string
	Email         string
	CryptPassword string
	Token         string
	Superuser     int       `orm:"default(0)"`
	Date          time.Time `orm:"auto_now_add;type(datetime)"`

	Profile    *Profile      `orm:"reverse(one)"`
	Message    []*Message    `orm:"reverse(many)"`
	Article    []*Article    `orm:"reverse(many)"`
	Comment    []*Comment    `orm:"reverse(many)"`
	SubComment []*SubComment `orm:"reverse(many)"`
}

// 一对一关系
type Profile struct {
	Id          int
	Gender      int    `orm:"null;default(0)"`
	Age         int    `orm:"null;default(0)"`
	Address     string `orm:"null"`
	Telephone   string `orm:"null"`
	Avator      string `orm:"null"`
	Personality string `orm:"null"`
	User        *User  `orm:"null;rel(one);on_delete(set_null)"`
}

func (this *User) PasswordCheck(password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(this.CryptPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
