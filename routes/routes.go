package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func ServerRoutes(r *gin.Engine) {
	r.GET("/test", controllers.TestEndPoint)
	r.GET("/quotes", controllers.GetQuotes)
	r.GET("/quotes/:id", controllers.GetQuoteById)
	r.GET("/quotes/search", controllers.SearchQuote)
	r.GET("/quotes/random", controllers.GetRandomQuote)
	r.POST("/quotes", controllers.AddQuote)
	r.PUT("/quotes/:id", controllers.UpdateQuote)
	r.DELETE("/quotes/:id", controllers.DeleteQuote)

	//comments
	r.GET("/quotes/:id/comments", controllers.GetCommentsByQuoteId)
	r.POST("/quotes/:id/comments", controllers.AddCommentsById)
}
