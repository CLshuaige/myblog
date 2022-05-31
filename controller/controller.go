package controller

import (
	"blog/dao"
	"blog/models"
	"fmt"
	"html/template"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
	_ "github.com/russross/blackfriday/v2"
)

func RegisterUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := models.User{
		Username: username,
		Password: password,
	}

	dao.Mgr.Register(&user)

	c.Redirect(301, "/")
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username)

	u := dao.Mgr.Login(username)

	if u.Username == "" {
		c.HTML(200, "login.html", "does not exist")
		fmt.Println("does not exist")
	} else {
		if u.Password != password {
			fmt.Println("wrong password")
			c.HTML(200, "login.html", "wrong password")
		} else {
			fmt.Println("login success")
			c.Redirect(301, "/")
		}
	}
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)
}

//操作博客

//博客列表
func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

//添加博客
func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := models.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}

	dao.Mgr.AddPost(&post)

	c.Redirect(302, "/post_index")
}

//跳转到添加博客
func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)
}

func PostDetail(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)

	content := blackfriday.Run([]byte(p.Content))

	c.HTML(200, "detail.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})
}
