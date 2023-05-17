package middlewares

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/ssssshel/fiber-service-template/src/shared/config"
	"github.com/ssssshel/restponses-go"
)

func AccessTokenChecker(c *fiber.Ctx) error {
	accessToken := c.Request().Header.Peek("authorization")

	if len(accessToken) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(restponses.Response4xxClientError(restponses.Status401Unauthorized, &restponses.BaseClientErrorInput{
			ErrorName:        "Unauthorized",
			ErrorDescription: "Access token is missing",
		}))
	}

	token := []byte(strings.TrimPrefix(string(accessToken), "Bearer "))
	authChan := make(chan jwt.Token)
	errChan := make(chan error)

	go func() {
		defer close(authChan)
		defer close(errChan)

		verificationToken := config.ATKey()

		verifiedToken, err := jwt.Parse(token, jwt.WithVerify(jwa.HS256, []byte(verificationToken)))
		if err != nil {
			errChan <- err
			return
		}

		tokenExpiration := verifiedToken.Expiration().Unix()
		currentTime := time.Now().Unix()

		if tokenExpiration < currentTime {
			errChan <- errors.New("Token expired")
			return
		}

		roles := map[string]bool{"admin": true, "user": true, "superadmin": true}

		rol, ok := verifiedToken.Get("rol")

		if !ok || !roles[rol.(string)] {
			errChan <- errors.New("Forbidden")
			return
		}

		authChan <- verifiedToken
	}()

	select {
	case verifiedToken := <-authChan:
		fmt.Printf("Verified token: %v", verifiedToken)
		return c.Next()

	case err := <-errChan:
		if err.Error() == "Forbidden" {
			return c.Status(fiber.StatusForbidden).JSON(restponses.Response4xxClientError(restponses.Status403Forbidden, &restponses.BaseClientErrorInput{
				ErrorName:        "Forbidden",
				ErrorDescription: "Access denied",
				Detail:           err.Error(),
			}))
		}
		return c.Status(fiber.StatusUnauthorized).JSON(restponses.Response4xxClientError(restponses.Status401Unauthorized, &restponses.BaseClientErrorInput{
			ErrorName:        "Unauthorized",
			ErrorDescription: "Access token is invalid",
			Detail:           err.Error(),
		}))
	}
}
