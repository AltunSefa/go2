package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/cruxia/go/initializers"
	"github.com/cruxia/go/models"
	"github.com/cruxia/go/utils"
	"gorm.io/gorm"
)

func DeserializeUser(c *fiber.Ctx) error {
	var access_token string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		access_token = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		access_token = c.Cookies("token")
	}

	if access_token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "You are not logged in"})
	}

	config, _ := initializers.LoadConfig(".")

	tokenClaims, err := utils.ValidateToken(access_token, config.AccessTokenPublicKey)
	if err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	var token_details utils.TokenDetails
	err = initializers.DB.First(&token_details, "token_uuid = ?", tokenClaims.TokenUuid).Error
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "could not find token"})
	}


	var user models.User
	err = initializers.DB.First(&user, "id = ?", token_details.UserID).Error

	if err == gorm.ErrRecordNotFound {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "fail", "message": "the user belonging to this token no logger exists"})
	}

	c.Locals("user", models.FilterUserRecord(&user))
	c.Locals("access_token_uuid", tokenClaims.TokenUuid)

	return c.Next()
}

