package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"api-gin/controllers"
	"api-gin/middlewares"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// set db to gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	r.GET("/phone", controllers.GetAllPhone)
	r.GET("/:id", controllers.GetMovieById)

	phonesMiddlewareRoute := r.Group("/phones")
	phonesMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	phonesMiddlewareRoute.POST("/", controllers.CreatePhones)
	phonesMiddlewareRoute.PATCH("/:id", controllers.UpdatePhones)
	phonesMiddlewareRoute.DELETE("/:id", controllers.DeletePhones)

	r.GET("/review-rating-categories", controllers.GetAllRating)
	r.GET("/review-rating-categories/:id", controllers.GetRatingById)
	r.GET("/review-rating-categories/:id/phones", controllers.GetPhonesByRatingId)

	ratingMiddlewareRoute := r.Group("/review-rating-categories")
	ratingMiddlewareRoute.Use(middlewares.JwtAuthMiddleware())
	ratingMiddlewareRoute.POST("/", controllers.CreateRating)
	ratingMiddlewareRoute.PATCH("/:id", controllers.UpdateRating)
	ratingMiddlewareRoute.DELETE("/:id", controllers.DeleteRating)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
