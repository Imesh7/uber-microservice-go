package api

import (
	"log"
	"net/http"
	"time"
	"uber-microservice-auth/config"
	"uber-microservice-auth/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Driver struct {
	Id        int32     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	VehicleNo   string    `json:"vehicle_no"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

func CreateDriver(c *gin.Context, userData *config.CreateDriverRequest) {
	var driver Driver
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	driver = Driver{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(encryptedPassword),
	}

	db.ConnectedDataBase.DB.Save(&driver)
	c.JSON(200, gin.H{
		"message": "driver created",
	})

}

func LoginDriver(c *gin.Context, loginUser *config.LoginUserRequest) {
	var driver Driver

	err := db.ConnectedDataBase.DB.Where("email = ?", loginUser.Email).First(&driver).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
			"err":     err,
		})
		return
	}
	errs := bcrypt.CompareHashAndPassword([]byte(driver.Password), []byte(loginUser.Password))
	if errs != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Password",
			"errr":    err,
		})
		return
	}

	token, err := CreateToken(10*time.Minute)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Generate Error",
			"errr":    err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "user Found",
		"user":    driver,
		"token":   token,
	})
}

