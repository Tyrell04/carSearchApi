package middleware

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/marcleonschulz/carSearchApi/config"
	"github.com/marcleonschulz/carSearchApi/pkg/models"
)

func AuthenticateRoles(role string, config config.Config) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Jwt.Secret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			roles := claims["roles"].([]interface{})
			for _, roleInterface := range roles {
				roleMap := roleInterface.(map[string]interface{})
				if roleMap["role"] == role {
					return ctx.Next()
				}
			}

			return ctx.
				Status(fiber.StatusUnauthorized).
				JSON(models.GeneralResponse{
					Code:    401,
					Message: "Unauthorized",
					Data:    "Invalid Role",
				})
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(models.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(models.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    "Invalid or expired JWT",
					})
			}
		},
	})
}

func AuthenticateUser(config config.Config) func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: []byte(config.Jwt.Secret),
		SuccessHandler: func(ctx *fiber.Ctx) error {
			user := ctx.Locals("user").(*jwt.Token)
			claims := user.Claims.(jwt.MapClaims)
			email := claims["email"].(string)
			ctx.Locals("email", email)
			return ctx.Next()
		},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			if err.Error() == "Missing or malformed JWT" {
				return c.
					Status(fiber.StatusBadRequest).
					JSON(models.GeneralResponse{
						Code:    400,
						Message: "Bad Request",
						Data:    "Missing or malformed JWT",
					})
			} else {
				return c.
					Status(fiber.StatusUnauthorized).
					JSON(models.GeneralResponse{
						Code:    401,
						Message: "Unauthorized",
						Data:    err.Error(),
					})
			}
		},
	})
}
