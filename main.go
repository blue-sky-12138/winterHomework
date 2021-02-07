package main

import(
	"WinterHomework/database"
	"WinterHomework/control"
	"WinterHomework/utilities"
)

func main(){
	database.OpenMysql()
	defer database.DB.Close()

	utilities.LogInit()

	control.RoutersEntrance()
}