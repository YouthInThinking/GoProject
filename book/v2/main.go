package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/YouthInThinking/GoProject/book/v2/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Book 结构体定义
type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
}

// 创建通过调用配置文件字段的方式初始化数据库连接函数逻辑
func setupDatabase() *gorm.DB {

	mc := config.C().MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mc.Username,
		mc.Password,
		mc.Host,
		mc.Port,
		mc.DB,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")

	}
	// 自动迁移表结构
	db.AutoMigrate(&Book{})
	return db.Debug()
}

func main() {
	//首先 加载默认配置文件
	path := os.Getenv("CONFIG_PATH")
	if path == "" {
		path = "application.yaml"
	}

	config.LoadConfigFromYaml(path)
	fmt.Println(config.C())
	// 初始化数据库连接
	db := setupDatabase()

	//初始化gin框架
	r := gin.Default()

	//创建书籍
	r.POST("/api/books", func(c *gin.Context) {
		//将结构体参数初始化变量
		var book Book
		//绑定请求体到结构体中
		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(
				http.StatusBadRequest, gin.H{
					"error": err.Error(),
					"code":  http.StatusBadRequest,
				},
			)
			return //将错误信息以json格式返回
		}
		//如果body数据成功绑定到结构体中，调用数据库创建方法创建书籍并返回成功信息
		db.Create(&book)
		c.JSON(http.StatusCreated, gin.H{"message": "Book created successfully", "data": book})
	})

	//获取所有书籍
	r.GET("/api/books/:id", func(c *gin.Context) {
		//将结构体定义为切边类型变量，方便获取顺序表数据
		var book []Book
		id := c.Param("id")
		//如果查找失败，返回错误信息
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found", "code": http.StatusNotFound})
			return
		}

		//查找成功返回书籍信息

		c.JSON(http.StatusOK, gin.H{"message": "Book retrieved successfully", "data": book})

	})

	//根据ID更新书籍
	//更新的逻辑就是先根据ID查找书籍信息，再根据查到的书籍信息更新数据库中的书籍信息
	r.PUT("/api/books/:id", func(c *gin.Context) {
		var book Book
		id := c.Param("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "code": http.StatusNotFound})
			return
		}

		if err := c.ShouldBindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "code": http.StatusBadRequest})
			return
		}
		//如果查找成功，将传入的body数据绑定到结构体中，并调用数据库更新方法更新书籍信息
		db.Save(&book)

		c.JSON(http.StatusOK, gin.H{"message": "Book updated successfully", "data": book})
	})

	//删除书籍
	r.DELETE("/api/books/:id", func(c *gin.Context) {
		var book Book

		id := c.Param("id")
		if err := db.First(&book, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found", "code": http.StatusNotFound})
			return
		}
		db.Delete(&book, id)

	})

	ac := config.C().Application
	r.Run(fmt.Sprintf("%s:%d", ac.Host, ac.Port)) // 启动服务

}
