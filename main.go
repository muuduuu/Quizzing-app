package main

import (
	"amrita-quiz-backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routes.SetupRoutes(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Failed to start", err)

	}

}
