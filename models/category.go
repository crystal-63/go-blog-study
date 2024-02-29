package models

// Category 从数据库获取的
type Category struct {
	Cid      int
	Name     string
	CreateAt string
	UpdateAt string
}

// CategoryResponse 从数据库获取的
type CategoryResponse struct {
	*HomeResponse
	CategoryName string
}