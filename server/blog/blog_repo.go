package blog

import (
	"learn-grpc/server/models"
)

type BlogRepo interface {
	GetBlogs() (*models.BlogList, error)
	GetBlogById(id *models.BlogId) (*models.Blog, error)
	CreateBlog(blog *models.Blog) (*models.Blog, error)
}
