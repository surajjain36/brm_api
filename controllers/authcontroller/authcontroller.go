package authcontroller

import (
	"brm_api/services/smsservice"
	"brm_api/services/userservice"
	"brm_api/utils/apperror"
	"brm_api/utils/common"
	"brm_api/utils/otpgen"
	"brm_api/utils/response"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Login(c *fiber.Ctx) error {
	var loginReq LoginReq
	json.Unmarshal(c.Body(), &loginReq)
	errors := response.ValidateStruct(loginReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}
	user, err := userservice.GetUserByMobileNumber(loginReq.MobileNumber)
	if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: ae.Msg,
			Data:    nil,
		})
	}
	if user.MobileNumber == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "login failed",
			Data:    nil,
		})
	}
	if user.Password == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "please actiate your account",
			Data:    nil,
		})
	}
	isValid := common.ComparePasswords(user.Password, loginReq.Password)
	if isValid {
		// Create Access token
		accessToken := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := accessToken.Claims.(jwt.MapClaims)
		claims["user_id"] = user.ID
		claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

		// Generate encoded accessToken and send it as response.
		aT, err := accessToken.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
		if err != nil {
			log.Printf("accessToken.SignedString: %v", err)
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		user.AccessToken = aT
		user.PartnerID = 0
		err = userservice.UpdateUser(user)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
				Success: false,
				Message: "login failed",
				Data:    nil,
			})
		}
		return c.JSON(response.HTTPResponse{
			Success: true,
			Message: "loggedin",
			Data: fiber.Map{
				"access_token": aT,
				"role":         user.Role,
			},
		})
	}
	return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
		Success: isValid,
		Message: "login Failed",
		Data:    nil,
	})
}

func Register(c *fiber.Ctx) error {
	var registerReq RegisterReq
	json.Unmarshal(c.Body(), &registerReq)
	errors := response.ValidateStruct(registerReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}
	user, err := userservice.GetUserByMobileNumber(registerReq.MobileNumber)
	if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
		if ae.Status != 404 {
			return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
				Success: false,
				Message: ae.Msg,
				Data:    nil,
			})
		}
	}
	if user.MobileNumber != "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "User already registered",
			Data:    nil,
		})
	}
	hashedPw := common.HashAndSalt(registerReq.Password)
	user.Password = hashedPw
	user.MobileNumber = registerReq.MobileNumber
	userservice.CreateUser(user)
	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "Registered",
		Data:    nil,
	})
}

func ForgotPassword(c *fiber.Ctx) error {
	var forgotPWReq ForgotPWReq
	json.Unmarshal(c.Body(), &forgotPWReq)
	errors := response.ValidateStruct(forgotPWReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}
	user, err := userservice.GetUserByMobileNumber(forgotPWReq.MobileNumber)
	if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: ae.Msg,
			Data:    nil,
		})
	}
	if user.MobileNumber == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "User not registered",
			Data:    nil,
		})
	}
	user.OTP = otpgen.GenerateOTP(6)
	user.OTPCreatedAt = time.Now()
	userservice.UpdateUser(user)

	// sending mobile otp
	res, err := smsservice.SendOTP([]string{user.MobileNumber}, []string{user.OTP})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: "OTP not sent",
			Data:    err.Error(),
		})
	}
	if res["return"] != true {
		return c.Status(fiber.StatusInternalServerError).JSON(response.HTTPResponse{
			Success: false,
			Message: "OTP not sent",
			Data:    res["message"],
		})
	}
	//Sending otp in email.
	// err = mailservice.SendMail(user.EmailID, os.Getenv("PROJECT_NAME")+" | OTP", user.OTP)
	// if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
	// 		Success: false,
	// 		Message: ae.Msg,
	// 		Data:    nil,
	// 	})
	// }

	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: fmt.Sprintf("otp(%s) sent", user.OTP),
		Data:    nil,
	})
}

func VerifyOTP(c *fiber.Ctx) error {
	var verifyOTPReq VerifyOTPReq
	json.Unmarshal(c.Body(), &verifyOTPReq)
	errors := response.ValidateStruct(verifyOTPReq)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(response.HTTPResponse{
			Success: false,
			Message: "Field validation failed",
			Data:    errors,
		})
	}
	user, err := userservice.GetUserByMobileNumber(verifyOTPReq.MobileNumber)
	if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: ae.Msg,
			Data:    nil,
		})
	}
	if user.MobileNumber == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "User not registered",
			Data:    nil,
		})
	}

	if verifyOTPReq.OTP != user.OTP {
		return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
			Success: false,
			Message: "verification failed",
			Data:    nil,
		})
	}

	user.OTP = ""
	user.OTPCreatedAt = time.Now()
	hashedPw := common.HashAndSalt(verifyOTPReq.Password)
	user.Password = hashedPw
	userservice.UpdateUser(user)
	//todo: send only if email is present.(after integrating mailenator)
	// err = mailservice.SendMail(user.EmailID, os.Getenv("PROJECT_NAME")+" | Password changed", "Password Successfully changed.")
	// if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
	// 		Success: false,
	// 		Message: ae.Msg,
	// 		Data:    nil,
	// 	})
	// }

	return c.JSON(response.HTTPResponse{
		Success: true,
		Message: "password updated",
		Data:    nil,
	})
}
