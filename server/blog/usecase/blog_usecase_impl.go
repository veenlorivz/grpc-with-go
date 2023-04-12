package usecase

import (
	"learn-grpc/server/blog"
	"learn-grpc/server/models"
)

type BlogUseCaseImpl struct {
	blogRepo blog.BlogRepo
}

func CreateBlogUseCaseImpl(blogRepo blog.BlogRepo) blog.BlogUseCase {
	return &BlogUseCaseImpl{blogRepo}
}

func (b *BlogUseCaseImpl) GetBlogs() (*models.BlogList, error) {
	return b.blogRepo.GetBlogs()
}

func (b *BlogUseCaseImpl) GetBlogById(id *models.BlogId) (*models.Blog, error) {
	return b.blogRepo.GetBlogById(id)
}

func (b *BlogUseCaseImpl) CreateBlog(blog *models.Blog) (*models.Blog, error) {
	return b.blogRepo.CreateBlog(blog)
}
