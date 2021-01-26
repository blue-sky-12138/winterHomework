package main

import(
	"WinterHomework/database"
	"WinterHomework/router"
)

func main(){
	database.OpenMysql()
	defer database.DB.Close()

	database.LogInit()

	router.RoutersEntrance()
}