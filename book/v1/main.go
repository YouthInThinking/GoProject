package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BookSet struct {
	//总共有多少个
	Total int64 `json:"total"`

	//
	Items []*Book `json:"items"`
}

type Book struct {
	ID uint `json:"id" gorm:"primaryKey;column:id"`

	//type字段 如果要使用gorm来自动创建和更新表的时候才需要被定义
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

// 初始化数据库
func SetupDatebase() *gorm.DB {
	dsn := "root:123456@tcp(172.16.160.12:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{}) // 自动迁移表结构
	fmt.Println("数据库连接成功")
	return db.Debug()
}

var db = SetupDatebase()

type BookApiHandler struct {
	Books []Book
}

var h = &BookApiHandler{}

func (h *BookApiHandler) ListBook(c *gin.Context) {
	//c.JSON(http.StatusOK, h.Books)

	set := &BookSet{}

	//查询书籍大小
	pageSize := c.Query("page_size")
	ps, err := strconv.ParseInt(pageSize, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
		})
	}

	//查询书籍id
	pageNumber := c.Query("page_number")
	pn, err := strconv.ParseInt(pageNumber, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "参数错误",
		})
	}

	query := db

	//查询带有关键字的书籍名称
	kws := c.Query("keywords")

	if kws != "" {
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	//设置分页参数
	offset := (pn - 1) * ps

	//定义一个空的bookList切片，用于存储查询结果。
	//bookList := []Book{}

	//db.Find接口用于查询数据库中所有的book数据。
	//通过控制offset limit实现分页
	if err := query.Count(&set.Total).Offset(int(offset)).Limit(int(ps)).Find(&set.Items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	//获取总数，总共有多少个, 总共多少页

	//返回查询结果。
	c.JSON(http.StatusOK, set)
}

func (h *BookApiHandler) CreateBook(c *gin.Context) {
	bookInstance := &Book{}

	if err := c.BindJSON(&bookInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	//数据入库,补充自增ID的值
	if err := db.Save(bookInstance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	//返回创建成功的消息和数据。
	c.JSON(http.StatusCreated, gin.H{
		"message": http.StatusCreated,
		"data":    bookInstance,
	})

}

func (h *BookApiHandler) GetBook(c *gin.Context) {

	//获取请求参数中的id。
	id := c.Param("id")

	//将id转换为int64类型。
	strconv.ParseInt(id, 10, 64)
	c.JSON(200, gin.H{
		"message": "Book with id: " + id,
	})
}
func (h *BookApiHandler) UpdateBook(c *gin.Context) {

	//获取请求参数中的id。
	id := c.Param("id")

	//将id转换为int64类型。
	strconv.ParseInt(id, 10, 64)
	c.JSON(200, gin.H{
		"message": "Book with id: " + id,
	})

	bookInstance := &Book{}
	if err := c.BindJSON(&bookInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
}

func (h *BookApiHandler) DeleteBook(c *gin.Context) {
	//获取请求参数中的id。
	id := c.Param("id")

	//将id转换为int64类型。
	strconv.ParseInt(id, 10, 64)
	c.JSON(200, gin.H{
		"message": "Book with id: " + id,
	})
}

func main() {
	// 设置 Gin 模式为发布模式
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// Book Restful API

	//list of books
	server.GET("/api/books", h.ListBook)

	//create new book
	// 对于post请求，我们一般都会选择将数据放在body正文中。
	server.POST("/api/books", h.CreateBook)

	//get book by book number
	server.GET("/api/books/:id", h.GetBook)

	// update book
	server.PUT("/api/books/:id", h.UpdateBook)

	//delete book
	server.DELETE("/api/books/:id", h.DeleteBook)

	if err := server.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Server started on port 8080")

	}
}
