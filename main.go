package main

import (
	"hacktiv8-chapter-three-session-two/database"
	"hacktiv8-chapter-three-session-two/routers"

	"github.com/gin-gonic/gin"
)

const PORT = ":8080"

func main() {
	router := gin.Default()

	database.StartDB()
	db := database.GetDB()

	routers.SetupAuthRoute(router, db)
	routers.SetupProductRoute(router, db)

	router.Run(PORT)
}
