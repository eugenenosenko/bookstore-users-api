package app

import (
	"gihub.com/eugenenosenko/bookstore-users-api/controllers/ping"
	"gihub.com/eugenenosenko/bookstore-users-api/controllers/users"
)

func mapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUsers)
	//router.GET("/users/search", controllers.SearchUsers)
}
