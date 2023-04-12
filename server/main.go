package main

import (
	"log"
	"net"

	"learn-grpc/server/blog/handler"
	"learn-grpc/server/blog/repo"
	"learn-grpc/server/blog/usecase"
	"learn-grpc/server/models"

	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	_ "gorm.io/driver/mysql"
)

func main() {
	servicePort := ":8000"

	db, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/learn_grpc?parseTime=true")

	if err != nil {
		log.Fatal(err.Error())
	}

	db.Debug().AutoMigrate(&models.BlogDB{})

	blogRepo := repo.CreateBlogRepoImpl(db)
	blogUseCase := usecase.CreateBlogUseCaseImpl(blogRepo)

	server := grpc.NewServer()

	handler.CreateBlogHandler(server, blogUseCase)

	conn, err := net.Listen("tcp", servicePort)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Server starting at Port", servicePort)
	log.Println(server.Serve(conn))
}
