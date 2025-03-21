package main

import (
	"fmt"
	"log"
	"myGo/bcrypt"
	jwt "myGo/jwt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID       uint   `gorm:"primary_key"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func main() {
	//读取配置文件，连接数据库
	viper.SetConfigFile("config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file:%v", err)
	}
	dsn := viper.GetString("database.dsn")
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database:%v", err)
	}
	if err = db.AutoMigrate(&User{}); err != nil {
		log.Fatalf("迁移表失败:%v", err)
	}
	fmt.Println("数据库连接成功")
	//创建路由
	r := gin.Default()

	r.Use(CORSMiddleware())

	r.POST("/register", register)
	r.POST("/login", login)
	r.POST("/checkToken", checkToken)

	r.Run(":8081")
}

func register(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	var count int64
	db.Model(&User{}).Where("username=?", user.Username).Count(&count)
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "username already exists"})
		return
	}
	user.Password, _ = bcrypt.HashPassword(user.Password)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "register failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "register success"})
}

func login(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "message": "invalid request"})
		return
	}
	var storedUser User
	if err := db.Where("username=?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "wrong username"})
		return
	}
	if !bcrypt.CheckPassword(storedUser.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"code": 1, "message": "wrong username or password"})
		return
	}
	token, err := jwt.GenToken(int64(user.ID), user.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "message": "login failed"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "message": "login success", "token": token})
}

func checkToken(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": 2003,
			"msg":  "请求头中的auth为空",
		})
		fmt.Println("请求头中的auth为空")
		c.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"code": 2004,
			"msg":  "请求头中的auth格式错误",
		})
		fmt.Println("请求头中的auth格式错误")
		c.Abort()
		return
	}

	mc, err := jwt.ParseToken(parts[1])
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2005,
			"msg":  "invalid Token",
		})
		fmt.Println("invalid Token")
		c.Abort()
		return
	}
	fmt.Println(mc)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "Valid Token",
	})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
