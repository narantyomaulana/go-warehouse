package controller

import (
	"micro-warehouse/user-service/controller/request"
	"micro-warehouse/user-service/controller/response"
	"micro-warehouse/user-service/pkg/conv"
	"micro-warehouse/user-service/pkg/validator"
	"micro-warehouse/user-service/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

type AuthControllerInterface interface {
	Login(c *fiber.Ctx) error
}

type AuthController struct {
	AuthService usecase.UserUsecaseInterface
}

// Login implements AuthControllerInterface.
func (a *AuthController) Login(c *fiber.Ctx) error {
	ctx := c.Context()
	var loginRequest request.LoginRequest
	if err := c.BodyParser(&loginRequest); err != nil {
		log.Errorf("[AuthController.Login] Login - 1: %v", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	if err := validator.Validate(loginRequest); err != nil {
		log.Errorf("[AuthController.Login] Login - 2: %v", err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	user, err := a.AuthService.GetUserByEmail(ctx, loginRequest.Email)
	if err != nil {
		log.Errorf("[AuthController.Login] Login - 3: %v", err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if user == nil {
		log.Errorf("[AuthController.Login] Login - 4: %v", err.Error())
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	isSame := conv.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isSame {
		log.Errorf("[AuthController.Login] Login - 5: %v", err.Error())
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid email or password",
		})
	}

	var roles []string
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}

	loginResp := response.LoginResponse{
		UserID: user.ID,
		Email:  user.Email,
		Role:   roles,
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Login successful",
		"data":    loginResp,
	})

}

func NewAuthController(authService usecase.UserUsecaseInterface) AuthControllerInterface {
	return &AuthController{
		AuthService: authService,
	}
}
