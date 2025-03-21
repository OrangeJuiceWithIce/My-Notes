package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

// 创建一个存储器，可以用.Get()方法来获取特定名字的Session,如果没有会自动创建
var store = sessions.NewCookieStore([]byte("Your-Secret-Key"))

func main() {
	r := gin.Default()
	//通过修改store.Options的MaxAge属性，可以设置Session的过期时间
	store.Options.MaxAge = 30
	//登录时，我们在存储中创建或查询特定的会话，并修改键值对
	r.GET("/login", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "Your-Session-Name")
		session.Values["authenticated"] = true
		session.Values["username"] = "Username from your frontend"

		session.Save(c.Request, c.Writer)
		c.String(http.StatusOK, "Login success")
	})
	//当访问主页面时，需要确认Session中键值对是否符合预期
	r.GET("/home", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "Your-Session-Name")
		if session.Values["authenticated"] == true {
			c.String(http.StatusOK, "Welcome "+session.Values["username"].(string))
		} else {
			c.String(http.StatusUnauthorized, "Please login first")
		}
	})
	//我们可以通过删除Session来实现注销
	r.GET("/logout", func(c *gin.Context) {
		session, _ := store.Get(c.Request, "Your-Session-Name")
		session.Options.MaxAge = -1 //设置MaxAge为负数，表示立即销毁Session
		session.Save(c.Request, c.Writer)
		c.String(http.StatusOK, "Logout success")
	})
	r.Run(":8000")
}
