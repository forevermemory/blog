package controllers

import (
	"bee/blog/models"
	"bee/blog/utils"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"log"
	"strconv"
	"strings"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

type userRegister struct {
	Name     string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

// @Title  检查用户名是否存在
// @Description 检查用户名是否存在
// @Param	username query string true "用户名"
// @Success 200 {object} models.User
// @Failure 400 register failed
// @router /registe [get]
func (this *UserController) CheckUsernameExist() {

	o := orm.NewOrm()
	username := this.GetString("username")

	var user models.User
	if err := o.QueryTable("user").Filter("username__exact", username).One(&user); err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": "username  not exist"}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"code": "2", "msg": "username exist"}
	this.ServeJSON()
	return

}

// @Title  用户列表
// @Description 用户列表
// @Success 200 {object} models.User
// @Failure 400 register failed
// @router /all [get]
func (this *UserController) AllUser() {

	o := orm.NewOrm()

	var users []models.User
	if _, err := o.QueryTable("user").All(&users); err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": err.Error()}
		this.ServeJSON()
		return
	}
	this.Data["json"] = &users
	this.ServeJSON()
	return

}

// @Title  用户列表2
// @Description 用户列表 1,2,3
// @Param	ids query string true	 "1,2"
// @Success 200 {object} models.User
// @Failure 400 register failed
// @router /getList [get]
func (this *UserController) UserList() {

	o := orm.NewOrm()

	var users []models.User
	ids := this.GetString("ids")
	idsArr := strings.Split(ids, ",")

	userTmp := models.User{}
	for _, v := range idsArr {
		// 转成 int
		intv, strconvVErr := strconv.Atoi(v)
		if strconvVErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "2", "msg": strconvVErr.Error()}
			this.ServeJSON()
			return
		}
		userTmp.Id = intv
		if readUserTmpErr := o.Read(&userTmp); readUserTmpErr != nil {
			this.Data["json"] = map[string]interface{}{"code": "3", "msg": readUserTmpErr.Error()}
			this.ServeJSON()
			return
		}
		users = append(users, userTmp)

	}

	this.Data["json"] = &users
	this.ServeJSON()
	return

}

// @Title  注册
// @Description 注册
// @Param	username formDate string true	 "用户名"
// @Param	password formDate string true	 "密码"
// @Param	email formDate string true "邮箱"
// @Success 200 {object} models.User
// @Failure 400 register failed
// @router /registe [post]
func (this *UserController) Register() {

	// this.Ctx.WriteString("ok")

	o := orm.NewOrm()

	//参数长度由前端控制。。。
	userReg := userRegister{}
	if err := this.ParseForm(&userReg); err != nil {
		this.Data["json"] = map[string]interface{}{"code": "1", "msg": "param error"}
		this.ServeJSON()
		return
	}

	//参数校验
	usernameIsValid := utils.CheckNameIsFull(userReg.Name)
	if !usernameIsValid {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": "username too short"}
		this.ServeJSON()
		return
	}

	passwordIsValid := utils.CheckPasswordIsFull(userReg.Password)
	if !passwordIsValid {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": "password too short"}
		this.ServeJSON()
		return
	}

	emailIsValid := utils.CheckEmailIsValid(userReg.Email)
	if !emailIsValid {
		this.Data["json"] = map[string]interface{}{"code": "4", "msg": "email is invalid"}
		this.ServeJSON()
		return
	}

	//加密密码 和生成token
	var porfile models.Profile
	user := models.User{
		Username:      userReg.Name,
		Password:      userReg.Password,
		Email:         userReg.Email,
		Token:         utils.GenerateToken(),
		CryptPassword: utils.PassWordEncrypted(userReg.Password),
		Profile:       &porfile,
	}
	log.Println(user)

	_, insertErr := o.Insert(&user)
	if insertErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "5", "msg": "add fail"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &user
	this.ServeJSON()
	return

}

// @Title  登录
// @Description 登录
// @Param	username formDate string true "用户名"
// @Param	password formDate string true "密码"
// @Success 200 {object} models.User
// @Failure 400 login failed
// @router /login [post]
func (this *UserController) Login() {

	o := orm.NewOrm()
	username := this.GetString("username")
	password := this.GetString("password")
	log.Println(username)
	log.Println(password)

	user := models.User{Username: username}
	o.QueryTable("user").Filter("username__exact", username).One(&user)

	isValid := user.PasswordCheck(password)
	if !isValid {
		// 密码错误
		this.Data["json"] = map[string]interface{}{"code": "1", "message": "password is invaild"}
		this.ServeJSON()
		return

	}

	// 将token 存入缓存
	pool := utils.GetDefaultRedisPool()
	conn := pool.Get()

	defer conn.Close()

	if _, doErr := conn.Do("set", user.Token, user.Token); doErr != nil {
		this.Data["json"] = map[string]interface{}{"code": "2", "msg": doErr.Error()}
		this.ServeJSON()
		return
	}

	if _, doErr2 := conn.Do("EXPIRE", user.Token, 60*60); doErr2 != nil {
		this.Data["json"] = map[string]interface{}{"code": "3", "msg": doErr2.Error()}
		this.ServeJSON()
		return
	}

	//设置过期时间

	this.Data["json"] = &user
	this.ServeJSON()
	return
}

// @Title 更新密码
// @Description 更新密码
// @Param id formData int true "username"
// @Param password1 formData string true "原密码"
// @Param password2 formData string true "新密码"
// @Success 200 {object} models.User
// @Failure 400 register failed
// @router /update [post]
func (this *UserController) UpdatePassword() {
	o := orm.NewOrm()
	//接收参数
	id, _ := this.GetInt("id")
	password1 := this.GetString("password1")
	password2 := this.GetString("password2")

	log.Println(password1)
	log.Println(password2)

	// 验证原密码是否正确
	user := models.User{Id: id}
	o.Read(&user)
	isValid := user.PasswordCheck(password1)
	if !isValid {
		// 密码错误
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "old password is invaild"}
		this.ServeJSON()
		return
	}

	//加密新密码
	newCryptPassword := utils.PassWordEncrypted(password2)
	user.CryptPassword = newCryptPassword
	if _, err := o.Update(&user); err != nil {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "update fail"}
		this.ServeJSON()
		return
	}

	this.Data["json"] = &user
	this.ServeJSON()
	return

}
