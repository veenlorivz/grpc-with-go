package handler

import (
	"context"
	"learn-grpc/server/blog"
	"learn-grpc/server/models"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

type BlogHandler struct {
	blogUseCase blog.BlogUseCase
}

func CreateBlogHandler(gr *grpc.Server, blogUseCase blog.BlogUseCase) {
	userHandler := &BlogHandler{blogUseCase}

	models.RegisterBlogsServer(gr, userHandler)
}

func (b *BlogHandler) GetBlogs(ctx context.Context, e *empty.Empty) (*models.BlogList, error) {
	blogs, err := b.blogUseCase.GetBlogs()

	if err != nil {
		return nil, err
	}

	return blogs, nil
}

func (b *BlogHandler) GetBlogById(ctx context.Context, id *models.BlogId) (*models.Blog, error) {
	blog, err := b.blogUseCase.GetBlogById(id)

	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (b *BlogHandler) CreateBlog(ctx context.Context, blog *models.Blog) (*empty.Empty, error) {
	_, err := b.blogUseCase.CreateBlog(blog)

	if err != nil {
		return new(empty.Empty), err
	}

	return new(empty.Empty), nil
}
