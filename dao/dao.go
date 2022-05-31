package dao

import (
	"blog/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	Register(user *models.User)
	Login(username string) models.User

	//博客操作
	AddPost(post *models.Post)
	GetAllPost() []models.Post
	GetPost(pid int) models.Post
}

type manager struct {
	db *gorm.DB
}

var Mgr Manager

func init() {
	dsn := "root:Exosshowtime123@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
}

func (mgr *manager) Login(username string) models.User {
	var user models.User
	mgr.db.Where("username=?", username).First(&user)
	return user
}

func (mgr *manager) Register(user *models.User) {
	mgr.db.Create(user)
}

func (mgr *manager) AddPost(post *models.Post) {
	mgr.db.Create(post)
}

func (mgr *manager) GetAllPost() []models.Post {
	var posts = make([]models.Post, 10)
	mgr.db.Find(&posts)
	return posts
}

func (mgr *manager) GetPost(pid int) models.Post {
	var post models.Post
	mgr.db.First(&post, pid)
	return post
}
