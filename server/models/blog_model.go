package models

type BlogDB struct {
	Id    int `gorm:"primaryKey"`
	Title string
	Body  string
}

func (b *BlogDB) TableName() string {
	return "blogs"
}
