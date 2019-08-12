package utils

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 生成token
func GenerateToken() string {

	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}

//加密密码
func PassWordEncrypted(value string) string {
	passWord := []byte(string(value))

	hashedPassword, err := bcrypt.GenerateFromPassword(passWord, bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%s", hashedPassword)

}

// 检查用户名是否合理 >3
func CheckNameIsFull(value string) bool {
	if value == "" {
		return false
	}
	trimValue := strings.TrimSpace(value)
	if len(trimValue) < 2 {
		return false
	}
	return true
}

// 检查用户名是否合理 >3
func CheckPasswordIsFull(value string) bool {
	if value == "" {
		return false
	}
	trimValue := strings.TrimSpace(value)
	if len(trimValue) < 6 {
		return false
	}
	return true
}

//邮箱
//^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$
func CheckEmailIsValid(value string) bool {
	if value == "" {
		return false
	}

	reg := `^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`
	retrunrst := regexp.MustCompile(reg).MatchString(value)
	return retrunrst
}

// 手机号码
func CheckPhone(value string) bool {
	if value == "" {
		return false
	}
	trimValue := strings.TrimSpace(value)
	if len(trimValue) < 11 {
		return false
	}

	reg := `^[1]([3-9])[0-9]{9}$`
	retrunrst := regexp.MustCompile(reg).MatchString(value)
	// s := []string{"18505921256", "18330823069", "19159029321"}
	// for _, v := range s {
	//    fmt.Println(rgx.MatchString(v))
	// }
	return retrunrst
}
