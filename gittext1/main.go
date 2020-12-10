package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"
)
//User 该结构体创造个人信息
type User struct {
	gorm.Model
	Name      string `gorm:"type:varchar(20);not null"`
	Telephone string `gorm:"varchar(11);not null;unique"`
	Password  string `gorm:"size:255;not null"`
}

func main() {
	db := InitDB()
	db.Close()


	r := gin.Default()
	r.POST("/api/auth/register", func(c *gin.Context) {

		//获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")

		//数据验证
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": "手机号输入错误！至少为11位哦！"})
			return
		}

		if len(password) < 6 {
			c.JSON(422, gin.H{
				"code": "密码不能小于六位",
			})
			return
		}

		if len(name) == 0 {
			name = Randomstring(10)
			return
		}

		log.Println(name, telephone, password)

		//判断手机号是否纯在
		if isTelephoneExist(db,telephone) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"code": "用户已经被注册"})
			return
		}

		//创建用户

		newUser := User{
			Name: name,
			Telephone: telephone,
			Password: password,
		}
		db.Create(&newUser)

		//返回结果
		c.JSON(200, gin.H{
			"message": "注册成功",
		})
	})
	r.Run(":9090") // listen and serve on 0.0.0.0:8080
}

// Randomstring :随机数生成id
func Randomstring(n int) string {
	var letters = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM")
	reuslt := make([]byte, n)

	rand.Seed(time.Now().Unix())
	for i := range reuslt {
		reuslt[i] = letters[rand.Intn(len(letters))]
	}

	return string(reuslt)
}

//InitDB :数据库的链接操作
func InitDB() *gorm.DB {
	driverName :="mysql"
	host :="localhost"
	port :="3306"
	database :="main"
	username :="root"
	password :="root"
	charset :="utf8"
	args := fmt.Sprintf("%s%s@tcp(%s:%s)/%s?charset=%sparstTime=true",
	username,
	password,
	host,
	port,
	database,
	charset)
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	db , err := gorm.Open(driverName,args)
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	if err!=nil {
		panic("db error"+err.Error())
	}

	db.AutoMigrate(&User{})
	return db
}


//isTelephoneExist 判断电话是否存在 是返回true 无则返回false
func isTelephoneExist(db *gorm.DB , telephone string)bool{
	var user User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}

	return false
}
