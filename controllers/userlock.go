package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/favorspin/lockdb/models"
)

func GetAllUserLocks(c *gin.Context) {
    var userlocks []models.UserLock
    models.DB.Find(&userlocks)

    c.JSON(http.StatusOK, gin.H{"data": userlocks})
}

func GetOneUserLock(c *gin.Context) {
    var userlock models.UserLock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&userlock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": userlock})
}

func AddLockToUser(c *gin.Context) {
    var input models.CreateUserLockInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userlock := models.UserLock{LockID: input.LockID, OwnerID: input.OwnerID}
    models.DB.Create(&userlock)

    c.JSON(http.StatusOK, gin.H{"data": userlock})
}

func UpdateUserLock(c *gin.Context) {
    var userlock models.UserLock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&userlock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    var input models.UpdateUserLockInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&userlock).Updates(input)

    c.JSON(http.StatusOK, gin.H{"data": userlock})
}

func RemoveLockFromUser(c *gin.Context) {
    var userlock models.UserLock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&userlock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    models.DB.Delete(&userlock)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
