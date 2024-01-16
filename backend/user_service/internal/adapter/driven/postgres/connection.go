package postgres

import (
	"fmt"
	"user/configs"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection(config configs.PostgresDBConfig) (*gorm.DB, error) {

	fmt.Println(config)

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.Host, config.User, config.Password, config.DatabaseName, config.Port)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
