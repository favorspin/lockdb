package controllers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/favorspin/lockdb/models"
)

func GetAllLocks(c *gin.Context) {
    var locks []models.Lock
    models.DB.Find(&locks)

    c.JSON(http.StatusOK, gin.H{"data": locks})
}

func GetOneLock(c *gin.Context) {
    var lock models.Lock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&lock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": lock})
}

func CreateLock(c *gin.Context) {
    var input models.CreateLockInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    lock := models.Lock{Brand: input.Brand, ModelName: input.ModelName, CoreType: input.CoreType}
    models.DB.Create(&lock)

    c.JSON(http.StatusOK, gin.H{"data": lock})
}

func UpdateLock(c *gin.Context) {
    var lock models.Lock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&lock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    var input models.UpdateLockInput
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Model(&lock).Updates(input)

    c.JSON(http.StatusOK, gin.H{"data": lock})
}

func DeleteLock(c *gin.Context) {
    var lock models.Lock
    if err := models.DB.Where("id = ?", c.Param("id")).First(&lock).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
        return
    }

    models.DB.Delete(&lock)

    c.JSON(http.StatusOK, gin.H{"data": true})
}
