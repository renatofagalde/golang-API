package mysqldb

import (
	"context"
	"fmt"
	"golang-API/model/repository/entity"
	"gorm.io/driver/mysql"
	"os"

	"golang-API/config/logger"
	"gorm.io/gorm"
)

var (
	MYSQL_DB_URL      = "MYSQL_DB_URL"
	MYSQL_DB_PORT     = "MYSQL_DB_PORT"
	MYSQL_DB_SCHEMA   = "MYSQL_DB_SCHEMA"
	MYSQL_DB_USER     = "MYSQL_DB_USER"
	MYSQL_DB_PASSWORD = "MYSQL_DB_PASSWORD"
)

func NewMySQLGORMConnection(ctx context.Context) (*gorm.DB, error) {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv(MYSQL_DB_USER), os.Getenv(MYSQL_DB_PASSWORD), os.Getenv(MYSQL_DB_URL),
		os.Getenv(MYSQL_DB_PORT), os.Getenv(MYSQL_DB_SCHEMA))

	logger.Info(connectionString)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	db.AutoMigrate(&entity.UserEntity{})

	if err != nil {
		return nil, err
	}
	return db, nil

	return nil, nil
}
