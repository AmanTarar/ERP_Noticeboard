package server

import (
	_ "main/docs"
	"main/server/handler"
	"main/server/services/notice"

	"main/server/gateway"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {


	server.engine.Use(gateway.CORSMiddleware())
	server.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))


	 

	//Notice Routes
	server.engine.POST("/create-notice",handler.AddNoticeHandler)

	server.engine.GET("/get-notice",handler.GetNoticeHandler)

	server.engine.DELETE("/delete-notice/*any",handler.DeleteNoticeHandler)

	server.engine.PUT("/update-notice",handler.UpdateNoticeHandler)

	server.engine.GET("/get-all-notices",handler.GetNoticesHandler)

	// server.engine.DELETE("/delete-all-notices",handler.DeleteAllNoticesHandler)

	server.engine.GET("/get-notices-history",handler.GetNoticesHistoryHandler)



	//testing 


	server.engine.GET("/get-notices",notice.UserGetNotices)
	


	


}
