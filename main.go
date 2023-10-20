package main

import (
	"example/grpc/application/grpc"
	"example/grpc/infrastructure/db"
	"os"

	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func main() {

	println("Rodando o GRPC server")
	database := db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)

}
