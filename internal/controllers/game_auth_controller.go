package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	dtos "smaash-web/internal/DTOs"
	"smaash-web/internal/repository"
	"smaash-web/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type GameAuthController struct {
	authService    services.Authentication
	profileService services.PlayerProfileService
	userRepo       repository.UserRepository
}

func NewGameAuthController(
	authService services.Authentication,
	profileService services.PlayerProfileService,
	userRepo repository.UserRepository,
) *GameAuthController {
	return &GameAuthController{
		authService:    authService,
		profileService: profileService,
		userRepo:       userRepo,
	}
}

func (g GameAuthController) GameLogin(c *gin.Context) {
	var body dtos.UserLoginDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	user, err := g.userRepo.ReadByEmail(c.Request.Context(), body.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, dtos.NewErrResp(
				"User doesn't exist. Please sign up on the website first.",
				c.Request.URL.Path,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, dtos.NewErrResp("Incorrect password", c.Request.URL.Path))
		return
	}

	profile, err := g.profileService.GetProfileByUserID(c.Request.Context(), user.ID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, dtos.NewErrResp(
				"Player profile not found. Please create your profile on the website first.",
				c.Request.URL.Path,
			))
			return
		}
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	profile.LastLogin = time.Now()
	if err := g.profileService.UpdateProfile(c.Request.Context(), profile); err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"playerProfileId": profile.ID,
		"userId":          user.ID,
		"email":           user.Email,
		"displayName":     profile.DisplayName,
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.NewErrResp(err.Error(), c.Request.URL.Path))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token":   tokenString,
		"profile": dtos.PlayerProfileToDTO(*profile),
	})
}
