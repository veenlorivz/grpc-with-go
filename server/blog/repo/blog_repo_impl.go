package repo

import (
	"learn-grpc/server/blog"
	"learn-grpc/server/models"

	"github.com/jinzhu/gorm"
)

type BlogRepoImpl struct {
	db *gorm.DB
}

func CreateBlogRepoImpl(db *gorm.DB) blog.BlogRepo {
	return &BlogRepoImpl{db}
}

func (b *BlogRepoImpl) GetBlogs() (*models.BlogList, error) {
	var blogs []*models.Blog

	if err := b.db.Table("blogs").Find(&blogs).Error; err != nil {
		return nil, err
	}

	blogList := models.BlogList{
		Blogs: blogs,
	}

	return &blogList, nil
}

func (b *BlogRepoImpl) GetBlogById(id *models.BlogId) (*models.Blog, error) {
	var blog models.Blog

	if err := b.db.First(&blog, "id = ?", id.Id).Error; err != nil {
		return nil, err
	}

	return &blog, nil
}

func (b *BlogRepoImpl) CreateBlog(blog *models.Blog) (*models.Blog, error) {
	var newBlog = &models.BlogDB{
		Title: blog.Title,
		Body:  blog.Body,
	}

	if err := b.db.Create(&newBlog).Error; err != nil {
		return nil, err
	}

	return blog, nil
}
