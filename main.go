package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang-basic/config/database/mongodb"
	"golang-basic/config/logger"
	"golang-basic/controller/routes"
	controller "golang-basic/controller/user"
	"golang-basic/model/service"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongodb.NewMongoDBConnection(context.Background())

	//init dependencies
	service := service.NewUserDomainService()
	userCotnroller := controller.NewUserControllerInterface(service)
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userCotnroller)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(os.Getenv("TESTE"))
	logger.Info("Iniciando")
}
