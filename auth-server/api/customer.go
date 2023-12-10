package api

import (
	"context"
	"log"
	"net/http"
	"strings"
	"time"
	"uber-microservice-auth/db"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	customerPb "uber-microservice-auth/gen/customer/proto"
)

type Customer struct {
	Id        int32     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	// UpdatedAt time.Time `gorm:"autoCreateTime"`
}

type CustomerAuthServiceServer struct {
	customerPb.UnimplementedCustomerAuthServiceServer
}

func (au *CustomerAuthServiceServer) VerifyToken(ctx context.Context, customerRequest *customerPb.VerifyTokenRequestProto) (*customerPb.VerifyTokenResponseProto, error) {
	return &customerPb.VerifyTokenResponseProto{},nil
}

func (au *CustomerAuthServiceServer) Login(ctx context.Context, customerRequest *customerPb.LoginRequestProto) (*customerPb.LoginResponseProto, error) {
	var customer Customer

	err := db.ConnectedDataBase.DB.Where("email = ?", customerRequest.Email).First(&customer).Error
	if err != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
			"err":     err,
		}) */
		return &customerPb.LoginResponseProto{
			StatusCode: 500,
			Message:    "User does not exists",
			Token:      "",
		}, err
	}
	errs := bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(customerRequest.Password))
	if errs != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Password",
			"errr":    err.Error(),
		}) */
		return &customerPb.LoginResponseProto{
			StatusCode: 401,
			Message:    "Wrong Password",
			Token:      "",
		}, err
	}

	token, err := CreateToken(10 * time.Minute)
	if err != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Generate Error",
			"errr":    err.Error(),
		}) */
		return &customerPb.LoginResponseProto{
			StatusCode: 400,
			Message:    "Token Generate Error",
			Token:      "",
		}, err
	}

	/* c.JSON(200, gin.H{
		"message": "user Found",
		"user":    customer,
		"token":   token,
	}) */
	return &customerPb.LoginResponseProto{
		StatusCode: 200,
		Message:    "User Found",
		Token:      token,
	}, nil
}

func (au *CustomerAuthServiceServer) Create(ctx context.Context, userData *customerPb.CreateCustomerRequestProto) (*customerPb.CreateCustomerResponseProto, error) {
	var customer Customer
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), 2)
	if err != nil {
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		}) */
		return &customerPb.CreateCustomerResponseProto{
			StatusCode: 400,
			Message:    "Error",
		}, nil
	}

	customer = Customer{
		Name:     userData.Name,
		Email:    userData.Email,
		Password: string(encryptedPassword),
	}

	db.ConnectedDataBase.DB.Save(&customer)
	/* c.JSON(200, gin.H{
	"message": "user created",
	}) */
	return &customerPb.CreateCustomerResponseProto{
		StatusCode: 201,
		Message:    "user created",
	}, nil

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
