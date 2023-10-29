package router

import (
    "TaskBTPN/controllers"
    "TaskBTPN/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

   
    r.POST("/users/register", controllers.Register)
    r.POST("/users/login", controllers.Login)

    
    authGroup := r.Group("/api")
    authGroup.Use(middlewares.AuthMiddleware())

    
    authGroup.POST("/photos", controllers.CreatePhoto)
    authGroup.GET("/photos", controllers.GetPhotos)
    authGroup.DELETE("/photos/:photoId", controllers.DeletePhoto)

    return r
}
