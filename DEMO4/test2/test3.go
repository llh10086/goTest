package main

import (
	"fmt"
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
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除

	// 一对多关系：用户拥有多篇文章
	Posts []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	Id        uint           `gorm:"primaryKey"`         // 主键ID
	Title     string         `gorm:"size:200;not null"`  // 标题，非空
	Content   string         `gorm:"type:text;not null"` // 内容，非空
	UserID    uint           `gorm:"not null;index"`     // 外键，关联用户ID
	CreatedAt time.Time      //创建时间
	UpdatedAt time.Time      //更新时间
	DeletedAt gorm.DeletedAt `gorm:"index"` //软删除

	//一对多反关联，一个文章只有一个用户
	User User `gorm:"foreignkey:"UserId"`
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
	return "suers"
}

func (Post) TableName() string {
	return "posts"
}

func (Comment) TableName() string {
	return "comments"
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
	defer sqlDb.Close()

	// 2. 自动迁移，创建表
	err = db.AutoMigrate(&User{}, &Post{}, &Comment{})
	if err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}

	fmt.Println("数据库迁移成功！")
}
