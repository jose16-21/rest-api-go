package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

func fetchPaginatedData[T any](c *gin.Context, model *gorm.DB, result *[]T) (int64, error) {
	page, limit, err := paginate(c)
	if err != nil {
		return 0, err
	}

	offset := (page - 1) * limit
	var total int64

	model.Count(&total)
	model.Limit(limit).Offset(offset).Find(result)

	return total, nil
}

func paginate(c *gin.Context) (int, int, error) {
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	pageInt, err := strconv.Atoi(page)
	if err != nil || pageInt < 1 {
		return 0, 0, fmt.Errorf("invalid page parameter")
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil || limitInt < 1 {
		return 0, 0, fmt.Errorf("invalid limit parameter")
	}

	return pageInt, limitInt, nil
}

func getUsers(c *gin.Context) {
	var users []User
	total, err := fetchPaginatedData(c, db.Model(&User{}), &users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  users,
		"total": total,
		"page":  c.DefaultQuery("page", "1"),
		"limit": c.DefaultQuery("limit", "10"),
	})
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&user)
	c.JSON(http.StatusCreated, user)
}

func main() {
	initDB()
	r := gin.Default()
	r.GET("/users", getUsers)
	r.POST("/users", createUser)

	r.Run(":8080")
}
