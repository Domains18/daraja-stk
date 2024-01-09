package main

import "auth/database"

func main(){
	database.Connect("root:root@tcp(localhost:3306)/jwt_app?parseTime=true")
	database.Migrate()
}
