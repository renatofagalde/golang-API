package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang-basic/config/database/mongodb"
	"golang-basic/config/logger"
	"golang-basic/controller/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error ao conectar no no banco, error=%s", err.Error())
		return
	} else {
		fmt.Println("conexao com sucesso")
	}

	userController := initDependencies(database)

	//init dependencies
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
	logger.Info("Iniciando")
}
