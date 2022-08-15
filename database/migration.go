package database

import (
	"fmt"

	"back_end_waysbucks/models"
	"back_end_waysbucks/pkg/mysql"
)

func RunMigration() {
	err := mysql.DB.AutoMigrate(
		&models.Topping{},
		&models.Product{})

	if err != nil {
		fmt.Println(err)
		panic("Migration failed")
	}

	fmt.Println("Migration success")
}
