package routes

import (
	"amrita-quiz-backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Home route
	r.GET("/", controllers.Home)

	// Future routes can be added here, e.g., user or quiz routes
	// r.POST("/quiz", controllers.CreateQuiz)
	// r.GET("/quiz/:id", controllers.GetQuiz)
}
