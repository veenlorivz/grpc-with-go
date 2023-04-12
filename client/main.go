package main

import (
	"learn-grpc/client/blog/handler"
	"learn-grpc/client/models"
	"log"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	servicePort := ":9000"
	targetPort := ":8000"

	conn, err := grpc.Dial(targetPort, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err.Error())
	}

	blogsClient := models.NewBlogsClient(conn)

	router := gin.Default()

	handler.CreateBlogHandler(router, blogsClient)

	err = router.Run(servicePort)

	if err != nil {
		log.Fatal(err.Error())
	}
}
