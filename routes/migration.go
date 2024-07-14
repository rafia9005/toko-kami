package routes

import "toko-kami/database"

func RunMigrate(dataModel interface{}) {
	database.DB.AutoMigrate(dataModel)
}
