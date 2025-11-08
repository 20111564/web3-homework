package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 题目一：模型定义============================================
type User struct {
	ID      uint `gorm:"primarykey"`
	Name    string
	PostNum int    `gorm:"default:0"`
	Posts   []Post `gorm:"foreignKey:UserID"`
}
type Post struct {
	ID           uint `gorm:"primarykey"`
	UserID       uint
	Title        string
	Content      string
	CommentState int       `gorm:"default:0"`
	Comments     []Comment `gorm:"foreignKey:PostID"`
}
type Comment struct {
	ID      uint `gorm:"primarykey"`
	PostID  uint
	Content string
}

var db *gorm.DB

func initDB() {
	fmt.Println("Connecting to the database...")
	initDb, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:123456@tcp(127.0.0.1:3308)/web3-test?charset=utf8&parseTime=True&loc=Local",
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect success!")

	err = initDb.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		fmt.Println("AutoMigrate failed:", err)
	}
	fmt.Println("数据库初始化成功")

	db = initDb
}

// 题目2：关联查询============================================
// 根据用户ID查询文章列表
func getPostListByUserId(userId uint) []Post {
	// db := initDB()
	var posts = []Post{}
	err := db.Model(&Post{}).Preload("Comments").Where("user_id = ?", userId).Find(&posts).Error
	if err != nil {
		fmt.Println("查询错误：", err)
	}
	return posts
}

// 查询评论数最多的文章
func getMostCommentsPost() Post {
	// db := initDB()
	var post Post
	err := db.Model(&Post{}).
		Select("posts.*, COUNT(1) as comment_count").
		Joins("left join comments on comments.post_id = posts.id").
		Group("posts.id").
		Order("comment_count DESC").
		Limit(1).
		Scan(&post).Error

	if err != nil {
		fmt.Println("查询错误：", err)
	}
	return post
}

func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println("AfterSave start")
	var user User
	err = tx.Model(&User{}).Where("id = ?", p.UserID).First(&user).Error
	if err != nil {
		return err
	}
	user.PostNum++
	err = tx.Model(&User{}).Where("id = ?", p.UserID).Update("post_num", user.PostNum).Error
	if err != nil {
		return err
	}
	fmt.Println("post hook success")
	return err
}

func (p *Comment) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println("AfterDelete start")
	var post Post
	err = tx.Model(&Post{}).Where("id = ?", p.PostID).First(&post).Error
	fmt.Println("post=", post)
	if err != nil {
		return err
	}
	var commentCount int64
	err = tx.Model(&Comment{}).Where("post_id = ?", p.PostID).Count(&commentCount).Error
	if err != nil {
		return err
	}
	if commentCount == 0 {
		post.CommentState = 0
		err = tx.Model(&Post{}).Where("id = ?", post.ID).Update("comment_state", post.CommentState).Error
		if err != nil {
			return err
		}
	}
	fmt.Println("comment hook success")
	return
}

func postSaveHookDemo() {
	post := Post{
		UserID:  1,
		Title:   "GORM Hooks in Go",
		Content: "This post demonstrates GORM hooks.",
	}

	err := db.Model(&Post{}).Save(&post).Error
	if err != nil {
		fmt.Println("创建错误：", err)
	}
	fmt.Println("创建文章成功=", post)
}

func commentDeleteHookDemo() {
	comment := Comment{
		ID:     1,
		PostID: 1,
	}
	err := db.Delete(&comment).Error
	if err != nil {
		fmt.Println("删除错误：", err)
	}
	fmt.Println("删除成功：", err)
}
func main() {
	//题目1：模型定义
	initDB()
	// //题目2：关联查询
	// fmt.Println(getPostListByUserId(1))
	// fmt.Println(getMostCommentsPost())
	//题目3：钩子函数
	// postSaveHookDemo()
	commentDeleteHookDemo()

}
