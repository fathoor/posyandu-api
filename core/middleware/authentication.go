package middleware

import (
	"github.com/fathoor/posyandu-api/core/config"
	"github.com/fathoor/posyandu-api/core/exception"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Authenticate(role string) func(*fiber.Ctx) error {
	cfg := config.ProvideConfig()
	jwtSecret := cfg.Get("JWT_SECRET")

	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwt.SigningMethodHS256.Alg(),
			Key:    []byte(jwtSecret),
		},

		SuccessHandler: func(c *fiber.Ctx) error {
			claims := c.Locals("user").(*jwt.Token).Claims.(jwt.MapClaims)
			user := claims["role"].(string)

			switch role {
			case "public":
				return c.Next()
			case "kader":
				if user == "kader" || user == "bidan" || user == "admin" {
					return c.Next()
				} else {
					panic(exception.ForbiddenError{
						Message: "Restricted access!",
					})
				}
			case "bidan":
				if user == "bidan" || user == "admin" {
					return c.Next()
				} else {
					panic(exception.ForbiddenError{
						Message: "Restricted access!",
					})
				}
			default:
				if user == role || user == "admin" {
					return c.Next()
				} else {
					panic(exception.ForbiddenError{
						Message: "Restricted access!",
					})
				}
			}
		},

		ErrorHandler: func(c *fiber.Ctx, e error) error {
			if e.Error() == "Missing or malformed JWT" {
				panic(exception.BadRequestError{
					Message: "Missing or malformed JWT",
				})
			} else {
				panic(exception.UnauthorizedError{
					Message: "Invalid or expired JWT",
				})
			}
		},
	})
}
