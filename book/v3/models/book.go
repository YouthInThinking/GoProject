package models

// 定义书籍属性字段
type Book struct {
	ID int `json:"id"`
	BookSpec
}

type BookSpec struct {
	//定义书籍状态字段

	//因为要将属性存储到数据库中，因此这里需要使用gorm来进行映射
	//将其属性映射为数据库字段以及字段类型

	Title  string  `json:"title" gorm:"cloumn:title varchar(255)" validate:"required"`
	Author string  `json:"author" gorm:"column:author varchar(255)" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale boolean"`
	Price  float64 `json:"price" gorm:"column:price decimal(10,2)" validate:"required"`
}

// 定义书籍清单字段
type BookSet struct {
	// 总共有多少本书籍
	Total int64 `json:"total"`
	// 书籍列表
	Items []*Book `json:"items"`
}

//Book属性映射数据库的表名称

func (b *Book) TableName() string {
	return "books"
}
