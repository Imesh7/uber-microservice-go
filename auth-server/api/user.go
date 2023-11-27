package api

import (
	"log"
	"net/http"
	"strings"
	"time"
	"uber-microservice-auth/config"
	"uber-microservice-auth/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id        int32     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt time.Time `gorm:"autoCreateTime"`
}

func CreatUser(c *gin.Context, userData *config.CreateUserRequest) {
	var user User
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 2)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		})
		return
	}

	user = User{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(encryptedPassword),
	}

	db.ConnectedDataBase.DB.Save(&user)
	c.JSON(200, gin.H{
		"message": "user created",
	})

}

func LoginUser(c *gin.Context, loginUser *config.LoginUserRequest) {
	var user User

	err := db.ConnectedDataBase.DB.Where("email = ?", loginUser.Email).First(&user).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
			"err":     err,
		})
		return
	}
	errs := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password))
	if errs != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Password",
			"errr":    err,
		})
		return
	}

	token, err := CreateToken(user, 10*time.Minute)
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
		"user":    user,
		"token":   token,
	})
}

func ValidateTokenClaims(c *gin.Context) {
	if c.Request.Header["Authorization"][0] != "" {
		authToken := strings.Split(c.Request.Header["Authorization"][0], " ")[1]
		token, err := ValidateToken(authToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		if token.Valid {
			claimOutput, ok := token.Claims.(*CustomClaims)
			if !ok {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
				return
			}
			println(claimOutput.UserID)
			if claimOutput.UserID != 1 {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong claim"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"message": "Valid token",
			})
			return
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Auth Required"})
		return
	}
}
