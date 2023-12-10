package api

import (
	"context"
	"log"
	"time"
	"uber-microservice-auth/db"

	driverPb "uber-microservice-auth/gen/driver/proto"

	"golang.org/x/crypto/bcrypt"
)

type Driver struct {
	Id        int32     `gorm:"AUTO_INCREMENT;PRIMARY_KEY;not null" json:"id"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	VehicleNo string    `json:"vehicle_no"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}

type DriverAuthServiceServer struct {
	driverPb.UnimplementedDriverAuthServiceServer
}

func (s *DriverAuthServiceServer) Create(ctx context.Context, driverRequest *driverPb.CreateDriverRequestProto) (*driverPb.CreateDriverResponseProto, error) {
	var driver Driver
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(driverRequest.Password), 2)
	if err != nil {
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error",
		}) */
		return &driverPb.CreateDriverResponseProto{
			StatusCode: 400,
			Message:    "Error",
		}, err
	}

	driver = Driver{
		Name:     driverRequest.Name,
		Email:    driverRequest.Email,
		Password: string(encryptedPassword),
	}

	db.ConnectedDataBase.DB.Save(&driver)
	/* c.JSON(200, gin.H{
		"message": "driver created",
	}) */
	return &driverPb.CreateDriverResponseProto{
		StatusCode: 201,
		Message:    "driver created",
	}, nil
}

func (s *DriverAuthServiceServer) Login(ctx context.Context, driverRequest *driverPb.LoginRequestProto) (*driverPb.LoginResponseProto, error) {
	var driver Driver

	err := db.ConnectedDataBase.DB.Where("email = ?", driverRequest.Email).First(&driver).Error
	if err != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "User does not exists",
			"err":     err,
		}) */
		return &driverPb.LoginResponseProto{
			StatusCode: 400,
			Message:    "User does not exists",
		}, err
	}
	errs := bcrypt.CompareHashAndPassword([]byte(driver.Password), []byte(driverRequest.Password))
	if errs != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Wrong Password",
			"errr":    err,
		}) */
		return &driverPb.LoginResponseProto{
			StatusCode: 401,
			Message:    "Wrong Password",
		}, err
	}

	token, err := CreateToken(10 * time.Minute)
	if err != nil {
		log.Println(err)
		/* c.JSON(http.StatusBadRequest, gin.H{
			"message": "Token Generate Error",
			"errr":    err.Error(),
		}) */
		return &driverPb.LoginResponseProto{
			StatusCode: 400,
			Message:    "Token Generate Error",
		}, err
	}

	/* c.JSON(200, gin.H{
		"message": "user Found",
		"user":    driver,
		"token":   token,
	}) */
	return &driverPb.LoginResponseProto{
		StatusCode: 200,
		Message:    "user Found",
		Token:      token,
	}, nil
}

func (s *DriverAuthServiceServer) VerifyToken(ctx context.Context, driverRequest *driverPb.VerifyTokenRequestProto) (*driverPb.VerifyTokenResponseProto, error) {
	return &driverPb.VerifyTokenResponseProto{}, nil
}
