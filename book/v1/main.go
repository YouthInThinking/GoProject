package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Title string `json:"title"`
}

func main() {
	// 设置 Gin 模式为发布模式
	// gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	// Book Restful API

	//list of books
	server.GET("/api/books", func(c *gin.Context) {
		//api/books?page_number=1&page_size=10
		//c.Query("page_number")
		//c.Query("page_size")
	})

	//create new book
	// 对于post请求，我们一般都会选择将数据放在body正文中。
	server.POST("/api/books", func(c *gin.Context) {
		// //读取请求体中的数据。
		// payload, err := io.ReadAll(c.Request.Body)
		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"code":    http.StatusBadRequest,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }
		// //关闭请求体。
		// defer c.Request.Body.Close()

		//将进行反序列化解析。

		//通过json.Unmarshal函数将payload中的数据反序列化为bookInstance。
		bookInstance := &Book{}
		// if err := json.Unmarshal(payload, bookInstance); err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"code":    http.StatusBadRequest,
		// 		"message": err.Error(),
		// 	})
		// 	return
		// }

		// 上面的读取body数据以及将数据反序列化的逻辑其实就是gin框架中已经为我们封装好了的方法：
		if err := c.BindJSON(&bookInstance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		//返回响应
		c.JSON(http.StatusCreated, gin.H{
			"message": http.StatusCreated,
			"data":    bookInstance,
		})

	})

	//get book by book number
	server.GET("/api/books/:id", func(c *gin.Context) {

		//获取请求参数中的id。
		id := c.Param("id")

		//将id转换为int64类型。
		strconv.ParseInt(id, 10, 64)
		c.JSON(200, gin.H{
			"message": "Book with id: " + id,
		})
	})

	// update book
	server.PUT("/api/books/:id", func(c *gin.Context) {
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

	})

	//delete book
	server.DELETE("/api/books/:id", func(c *gin.Context) {
		//获取请求参数中的id。
		id := c.Param("id")

		//将id转换为int64类型。
		strconv.ParseInt(id, 10, 64)
		c.JSON(200, gin.H{
			"message": "Book with id: " + id,
		})
	})

	if err := server.Run(":8080"); err != nil {
		fmt.Println("Failed to start server:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("Server started on port 8080")

	}
}
