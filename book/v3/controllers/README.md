# 控制器

业务处理

![](./biz.png)

## 单元测试 (TDD)

```go
func TestGetBook(t *testing.T) {
	book, err := controllers.Book.GetBook(context.Background(), controllers.NewGetBookRequest(3))
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)
}

func TestCreateBook(t *testing.T) {
	book, err := controllers.Book.CreateBook(context.Background(), &models.BookSpec{
		Title:  "unit test for go controller obj",
		Author: "will",
		Price:  99.99,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(book)
}

func init() {
	// 执行配置的加载
	err := config.LoadConfigFromYaml(fmt.Sprintf("%s/book/v3/application.yaml", os.Getenv("workspaceFolder")))
	if err != nil {
		panic(err)
	}
}
```