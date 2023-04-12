package handler

import (
	"context"
	"learn-grpc/client/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/ptypes/empty"
)

type BlogHandler struct {
	blogsClient models.BlogsClient
}

func CreateBlogHandler(r *gin.Engine, blogsClient models.BlogsClient) {
	blogHandler := &BlogHandler{blogsClient}

	r.GET("/blogs", blogHandler.GetBlogs)
	r.GET("/blogs/:id", blogHandler.GetBlogById)
	r.POST("/blogs", blogHandler.CreateBlog)
}

func (b *BlogHandler) CreateBlog(c *gin.Context) {
	var blog *models.Blog

	if err := c.Bind(&blog); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	if blog.Body == "" || blog.Title == "" {
		c.JSON(http.StatusBadRequest, "Fields are required")
		return
	}

	_, err := b.blogsClient.CreateBlog(context.Background(), blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, "Blog has been created")
}

func (b *BlogHandler) GetBlogs(c *gin.Context) {
	blogList, err := b.blogsClient.GetBlogs(context.Background(), new(empty.Empty))

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, blogList)
}

func (b *BlogHandler) GetBlogById(c *gin.Context) {
	id := c.Param("id")
	BlogId := &models.BlogId{
		Id: id,
	}
	blog, err := b.blogsClient.GetBlogById(context.Background(), BlogId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, blog)
}
