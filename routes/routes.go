package routes

import (
	"brm_api/controllers/authcontroller"
	"brm_api/controllers/packagecontroller"
	"brm_api/controllers/pincontroller"
	"brm_api/controllers/usercontroller"
	"brm_api/models"
	"brm_api/services/userservice"
	"brm_api/utils/apperror"
	"brm_api/utils/response"
	"flag"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

var (
	prefork = flag.Bool("prefork", false, "Enable Prefork") // go run . -prefork=true
)

// New create an instance of BRM app routes
func New() *fiber.App {
	// file, err := os.OpenFile("./123.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	// if err != nil {
	// 	log.Fatalf("error opening file: %v", err)
	// }
	// defer file.Close()

	app := fiber.New(fiber.Config{
		Prefork: *prefork,
	})

	app.Use(cors.New())
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		// For more options, see the Config section
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
		// Output: file,
	}))
	app.Use(recover.New())
	app.Static("/", "./public")
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v1.Post("/login", authcontroller.Login)
	v1.Post("/register", authcontroller.Register)
	v1.Post("/forgot-password", authcontroller.ForgotPassword)
	v1.Post("/verify-otp", authcontroller.VerifyOTP)

	// JWT Middleware for v1 api
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_ACCESS_SECRET")),
		ErrorHandler: func(c *fiber.Ctx, e error) error {
			if e != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
					Success: false,
					Message: e.Error(),
					Data:    nil,
				})
			}
			return nil
		},
		SuccessHandler: func(c *fiber.Ctx) error {
			user := c.Locals("user").(*jwt.Token)
			token := user.Raw
			claims := user.Claims.(jwt.MapClaims)
			userID := uint(claims["user_id"].(float64))
			loggedIdUser, err := userservice.GetUserByID(userID)
			if ae, ok := err.(*apperror.AppError); ok && ae.Status != 0 {
				return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
					Success: false,
					Message: ae.Msg,
					Data:    nil,
				})
			}
			if err != nil && err.Error() != "" {
				return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
					Success: false,
					Message: err.Error(),
					Data:    nil,
				})
			}
			if loggedIdUser.AccessToken != token {
				return c.Status(fiber.StatusUnauthorized).JSON(response.HTTPResponse{
					Success: false,
					Message: "Invalid Token",
					Data:    nil,
				})
			}
			c.Locals(models.UserKey, &loggedIdUser)
			return c.Next()
		},
	}))

	user := v1.Group("/user")
	user.Get("/roles", usercontroller.Roles)
	user.Get("/packages", usercontroller.Packages)
	user.Get("/", usercontroller.List)
	user.Get("/:id", usercontroller.First)
	user.Post("/", usercontroller.Create)
	user.Put("/:id", usercontroller.Update)

	pin := v1.Group("/pin")
	pin.Get("/", pincontroller.List)
	pin.Post("/", pincontroller.Create)
	pin.Post("/transfer", pincontroller.TransferPins)

	packages := v1.Group("/package")
	packages.Get("/achievers", packagecontroller.ListAchievers)
	packages.Get("/classic", packagecontroller.ListClassicUsers)

	// family := v1.Group("/families")
	// family.Get("/", familycontroller.List)
	// family.Get("/:id", familycontroller.First)
	// family.Post("/", familycontroller.Create)
	// family.Put("/:id", familycontroller.Update)
	// family.Delete("/:id", familycontroller.Delete)

	return app
}
