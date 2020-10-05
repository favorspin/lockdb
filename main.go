package main

import (
    "github.com/gin-gonic/gin"
    "github.com/favorspin/lockdb/models"
    "github.com/favorspin/lockdb/controllers"
)

func main() {
    r := gin.Default()

    // Connect to database
    models.ConnectDatabase()

    // Routes
    v1 := r.Group("/api/v1")
    {
        // user routes
        v1.POST("/user", controllers.CreateUser)
        v1.GET("/user", controllers.GetAllUsers)
        v1.GET("/user/:id", controllers.GetOneUser)
        v1.PATCH("/user/:id", controllers.UpdateUser)
        v1.DELETE("/user/:id", controllers.DeleteUser)

        // lock routes
        // v1.POST("/lock", controllers.CerateLock)
        // v1.GET("/lock", controllers.GetAllLocks)
        // v1.GET("/lock/:id", controllers.GetOneLock)
        // v1.PATCH("/lock/:id", controllers.UpdateLock)
        // v1.DELETE("/lock/:id", controllers.DeleteLock)

        // user-lock routes
        // v1.POST("/userLock", controllers.AddLockToUser)
        // v1.GET("/userLock", controllers.GetAllUserLocks)
        // v1.GET("/userLock/:id", controllers.GetOneUserLock)
        // v1.PATCH("/userLock/:id", controllers.UpdateUserLock)
        // v1.DELETE("/userLock/:id", controllers.RemoveLockFromUser)
    }

    // Run the server
    r.Run()
}
