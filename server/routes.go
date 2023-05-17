package server

import (
	_ "main/docs"
	"main/server/handler"

	"main/server/gateway"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {
	server.engine.GET("/hello",func(ctx *gin.Context) {
		ctx.JSON(200,gin.H{
			"message":"Hello World",
		})
	})

	server.engine.Use(gateway.CORSMiddleware())
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	//create notice 


	server.engine.POST("/create-notice",handler.AddNoticeHandler)

	server.engine.GET("/get-notice",handler.GetNoticeHandler)

	server.engine.DELETE("/delete-notice/*any",handler.DeleteNoticeHandler)

	server.engine.PUT("/update-notice",handler.UpdateNoticeHandler)

	server.engine.GET("/get-all-notices",handler.GetNoticesHandler)

	server.engine.DELETE("/delete-all-notices",handler.DeleteAllNoticesHandler)



	//TESTING
	server.engine.POST("/create-token",handler.CreateTOKEN)
	server.engine.POST("/decode-token",handler.DecodeTOKEN)	

}
