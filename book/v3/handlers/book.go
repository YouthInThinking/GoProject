package handlers

import (
	"net/http"
	"strconv"

	"github.com/YouthInThinking/GoProject/book/v3/config"
	"github.com/YouthInThinking/GoProject/book/v3/controllers"
	"github.com/YouthInThinking/GoProject/book/v3/models"
	"github.com/gin-gonic/gin"
)

type BookApiHandler struct {
	// 这里可以定义一些字段，用于存储一些需要在多个方法中使用的数据。
}

// BookApiHandler 实现了 BookApi 接口。
var Book = &BookApiHandler{}

// listBook 方法实现了 BookApi 接口的 listBook 方法。
func (h *BookApiHandler) listBook(c *gin.Context) {

	// 这里可以实现具体的业务逻辑，比如从数据库中查询数据，并返回给客户端。
	// 为了简化示例，这里直接返回一个固定的值。实际应用中，需要根据请求参数和业务逻辑来查询数据库。
	set := models.BookSet{}

	pn, ps := 1, 20

	pageNumber := c.Query("page_number")

	if pageNumber != "" {
		pnInt, err := strconv.ParseInt(pageNumber, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}
		// 如果查询参数有效，则使用查询参数的值。否则，使用默认值。
		pn = int(pnInt)
	}

	pageSize := c.Query("page_size")
	if pageSize != "" {
		psInt, err := strconv.ParseInt(pageSize, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
			return
		}

		ps = int(psInt)

	}

	//获取Book对象
	query := config.C().MySQL.GetDB().Model(&models.Book{})

	// 根据查询参数进行过滤。
	kws := c.Query("keywords")
	if kws != "" {
		query = query.Where("title LIKE ?", "%"+kws+"%")
	}

	//根据book对象进行分页查询，逻辑就是先查总数，再从总数上偏移当量，限制每次查询的记录数，在此之内获取所有复合条件的记录。
	//如果查询失败，就返回错误信息
	if err := query.Count(&set.Total).Offset((pn - 1) * ps).Limit(ps).Find(&set.Items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}

	//如果查询成功就返回书籍属性信息
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": set,
	})
}

func (h *BookApiHandler) createBook(c *gin.Context) {

	// 创建一个BookSpec属性对象实例。
	bookSpecInstences := &models.BookSpec{}

	// 获取BookSpec对象实例的body数据。如果获取失败，就返回错误信息。
	if err := c.BindJSON(bookSpecInstences); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}
	// 创建一个Book对象实例,将bookSpecInstences实例信息保存值Book对象中
	bookInstences := &models.Book{BookSpec: *bookSpecInstences}

	// 将Book对象实例保存到数据库中。如果保存失败，就返回错误信息。
	if err := config.DB().Save(&bookInstences).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err,
		})
		return
	}

	// 如果保存成功就返回创建的书籍属性信息。
	c.JSON(http.StatusCreated, gin.H{
		"code": http.StatusCreated,
		"data": bookInstences,
	})

}

func (h *BookApiHandler) getBook(c *gin.Context) {
	// 创建一个Book属性对象实例。
	// bookInstences := &models.Book{}
	// if err := config.DB().Where("id = ?", c.Param("id")).Take(&bookInstences).Error; err != nil {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"code":    http.StatusNotFound,
	// 		"message": "Book not found",
	// 	})
	// 	return
	// }
	book, err := controllers.Book.GetBooks(c, controllers.NewGetBookRequest(c.Param("id")))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, book)
}

func (h *BookApiHandler) updateBook(c *gin.Context) {

	isbnStr := c.Param("id")
	isbn, err := strconv.ParseInt(isbnStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "Invalid ISBN",
		})
		return
	}

	//读取body中的参数

	bookInstance := &models.Book{
		ID: int(isbn),
	}

	if err := c.BindJSON(&bookInstance.BookSpec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if err := config.DB().Where("id = ?", bookInstance.ID).Updates(bookInstance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, bookInstance)
}

func (h BookApiHandler) deleteBook(c *gin.Context) {
	if err := config.DB().Where("id = ?", c.Param("id")).Delete(&models.Book{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Failed to delete book",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": "Book deleted successfully",
	})
}

// 注册路由
func (h *BookApiHandler) Registry(r gin.IRouter) {
	r.GET("/api/books", h.listBook)
	r.POST("/api/books", h.createBook)
	r.GET("/api/books/:id", h.getBook)
	r.PUT("/api/books/:id", h.updateBook)
	r.DELETE("/api/books/:id", h.deleteBook)

}
