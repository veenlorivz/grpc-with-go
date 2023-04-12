package blog

import (
	"learn-grpc/server/models"
)

type BlogUseCase interface {
	GetBlogs() (*models.BlogList, error)
	GetBlogById(id *models.BlogId) (*models.Blog, error)
	CreateBlog(blog *models.Blog) (*models.Blog, error)
}
