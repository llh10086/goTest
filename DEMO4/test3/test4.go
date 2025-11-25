package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey"`               // 主键ID
	Username  string         `gorm:"size:50;not null;unique"`  // 用户名，非空且唯一
	Email     string         `gorm:"size:100;not null;unique"` // 邮箱，非空且唯一
	Password  string         `gorm:"size:100;not null"`        // 密码，非空
	CreatedAt time.Time      //创建时间
	UpdatedAt time.Time      //更新时间
	PostCount int            `gorm:"default:0"` // 文章数量，默认值为0
	DeletedAt gorm.DeletedAt `gorm:"index"`     //软删除

	// 一对多关系：用户拥有多篇文章
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	Id            uint           `gorm:"primaryKey"`         // 主键ID
	Title         string         `gorm:"size:200;not null"`  // 标题，非空
	Content       string         `gorm:"type:text;not null"` // 内容，非空
	UserID        uint           `gorm:"not null;index"`     // 外键，关联用户ID
	CreatedAt     time.Time      //创建时间
	UpdatedAt     time.Time      //更新时间
	DeletedAt     gorm.DeletedAt `gorm:"index"`                 //软删除
	CommentStatus string         `gorm:"size:20;default:'无评论'"` // 新增字段：评论状态

	//一对多反关联，一个文章只有一个用户
	User     User      `gorm:"foreignkey:UserId"`
	Comments []Comment `gorm:"foreignKey:PostID"`
}

type Comment struct {
	Id        uint           `gorm:"primaryKey"`         //主键id
	Content   string         `gorm:"type:text;not null"` // 评论内容，非空
	PostID    uint           `gorm:"not null"`           // 外键，关联 Post 表的 ID
	UserID    uint           `gorm:"not null"`           // 外键，关联 User 表的 ID（记录评论者）
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除

	// 关联关系
	Post Post `gorm:"foreignKey:PostID"` // 多对一关系：评论属于一篇文章
	User User `gorm:"foreignKey:UserID"` // 多对一关系：评论属于一个用户
}

// 定义TableName方法自定义表名
func (User) TableName() string {
	return "users"
}

func (Post) TableName() string {
	return "posts"
}

func (Comment) TableName() string {
	return "comments"
}

var db *gorm.DB

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	log.Println("BeforeCreate钩子被调用：在创建文章之前执行一些操作")

	// 文章创建前，更新用户的文章数量
	result := tx.Model(&User{}).Where("id = ?", p.UserID).UpdateColumn("post_count", gorm.Expr("post_count + ?", 1))
	if result.Error != nil {
		log.Printf("更新用户文章数量失败: %v", result.Error)
	}
	log.Println("用户:%s文章数量更新成功", p.User.Username)
	return nil
}

func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	log.Println("AfterDelete钩子被调用：在删除文章之后执行一些操作")
	// 计算该文章的评论数量
	var count int64
	result := tx.Model(&Comment{}).Where("post_id = ?", c.PostID).Count(&count)
	if result.Error != nil {
		log.Printf("统计评论数异常", result.Error)
	}

	log.Printf("文章ID:%d删除评论后，当前评论数量为:%d", c.PostID, count)

	// 如果评论数量为0，更新文章的评论状态为 "无评论"
	if count == 0 {
		result = tx.Model(&Post{}).Where("id = ?", c.PostID).Update("comment_status", "无评论")
		if result.Error != nil {
			log.Printf("更新文章评论状态失败: %v", result.Error)
		} else {
			log.Printf("文章ID:%d评论数量为0，评论状态已更新为'无评论'", c.PostID)
		}
	}

	return nil
}

func main() {
	// 1. 数据库连接配置
	dsn := "root:Llh..123@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	sqlDb, err := db.DB() // 获取通用数据库对象
	if err != nil {
		log.Fatalf("获取数据库对象失败: %v", err)
	}
	// 关闭数据库连接
	defer sqlDb.Close()

	// 查询某个用户的所有文章及其评论
	var users User
	userId := 1
	result := db.Preload("Posts.Comments").First(&users, userId)
	if result.Error != nil {
		log.Fatalf("查询用户及其文章和评论失败: %v", result.Error)
	}

	// 输出查询结果
	log.Printf("用户: %v", users.Username)
	for _, v := range users.Posts {
		log.Printf("文章：%s", v.Title)
		for _, v := range v.Comments {
			log.Printf("评论：%s", v.Content)
		}

	}
	//给用户添加文章
	newPost := Post{
		Title:   "GORM钩子详解",
		Content: "本文介绍了GORM中的钩子函数及其应用场景。",
		UserID:  1,
	}
	result = db.Create(&newPost)
	if result.Error != nil {
		log.Fatalf("创建文章失败: %v", result.Error)
	} else {
		log.Printf("文章创建成功，文章ID: %d", newPost.Id)
	}

	//删除文章评论
	commentId := 1
	var comment Comment
	result = db.First(&comment, commentId)
	if result.Error != nil {
		log.Fatalf("查询评论失败: %v", result.Error)
	}

	result = db.Delete(&comment)
	if result.Error != nil {
		log.Fatalf("删除评论失败: %v", result.Error)
	} else {
		log.Printf("评论删除成功，评论ID: %d", commentId)
	}

}
