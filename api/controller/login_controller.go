package controller

import (
	"fmt"
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/internal/authutil"
	// "stontactics/internal/tokenutil"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// ConnectVK	godoc
// @Summary	    Вход по социальной сети
// @Tags        Login
// @Router      /auth/{provider} [get]
// @Param       provider	path	string	true	"vk/google"
// @Param       token	    query	string	false	"jwt token"
// @Produce		json
// @Security 	Bearer
func (lc *LoginController) BeginLogin(c *gin.Context) {
	provider := c.Param("provider")
	oauthState := authutil.GenerateStateOauthCookie(c.Writer)

	var url string
	switch provider {
	case "google":
		url = authutil.GetConfigGoogle(lc.Env).AuthCodeURL(oauthState)
	case "vk":
		url = authutil.GetConfigVK(lc.Env).AuthCodeURL(oauthState)
	default:
		c.String(http.StatusInternalServerError, "Invalid provider")
		return
	}

	token := c.Query("token")
	if token != "" {
		var expiration = time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "token", Value: token, Expires: expiration}
		http.SetCookie(c.Writer, &cookie)
	}
	
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (lc *LoginController) Callback(c *gin.Context) {
	provider := c.Param("provider")

	// Проверка валидности провайдера
	if provider != "google" && provider != "vk" {
		c.String(http.StatusBadRequest, "Invalid provider")
		return
	}

	oauthState, err := c.Cookie("oauthstate")
	if err != nil {
		c.String(http.StatusInternalServerError, "Invalid oauth google state")
		return
	}

	if c.Query("state") != oauthState {
		c.String(http.StatusInternalServerError, "Invalid oauth state")
		return
	}

	// Проверка валидности параметра "code"
	if c.Query("code") == "" {
		c.String(http.StatusBadRequest, "Code parameter is required")
		return
	}

	var user domain.User
	var userEntry domain.User

	switch provider {
	case "google":
		// Получение данных от Google
		data, err := authutil.GetUserDataFromGoogle(lc.Env, c.Query("code"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		// Валидация данных пользователя от Google
		if data.ID == "" || !data.VerifiedEmail {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Incomplete user data received from Google"})
			return
		}

		now := time.Now()
		userEntry = domain.User{
			ID:        "",
			Name:      data.Name,
			AvatarURL: data.Picture,
			Pro: domain.UserPro{
				Active: false,
				Until:  nil,
			},
			Auth: domain.UserAuth{
				Email: domain.EmailAuth{},
				Google: domain.SocialAuth{
					ID: data.ID,
				},
				VK: domain.SocialAuth{},
			},
			CreatedAt: &now,
		}

		dbUser, err := lc.LoginUsecase.GetUserByGoogleID(c, userEntry.Auth.Google.ID)
		if err != nil {
			user = dbUser
		}
	case "vk":
		// Получение данных от VK
		data, err := authutil.GetUserDataFromVK(lc.Env, c.Query("code"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		// Валидация данных пользователя от VK
		if data.ID == 0 {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Incomplete user data received from VK"})
			return
		}

		now := time.Now()
		userEntry = domain.User{
			ID:        "",
			Name:      data.FirstName + " " + data.LastName,
			AvatarURL: data.Photo200,
			Pro: domain.UserPro{
				Active: false,
				Until:  nil,
			},
			Auth: domain.UserAuth{
				Email:  domain.EmailAuth{},
				Google: domain.SocialAuth{},
				VK: domain.SocialAuth{
					ID: strconv.Itoa(data.ID),
				},
			},
			CreatedAt: &now,
		}

		dbUser, err := lc.LoginUsecase.GetUserByVKID(c, userEntry.Auth.VK.ID)
		if err != nil {
			user = dbUser
		}
	default:
		c.String(http.StatusInternalServerError, "Invalid provider")
		return
	}

	/* token, err := c.Cookie("token")
	if err == nil {
		userID, err := tokenutil.ExtractIDFromToken(token, lc.Env.AccessTokenSecret)
		if err == nil {
			tokenUser, err := lc.LoginUsecase.GetUserByID(c, userID)
		}
	}
 	*/

	var accessToken, refreshToken string
	if user.ID != "" {
		// Обновление данных пользователя, если они изменились
		if userEntry.Name != user.Name || userEntry.AvatarURL != user.AvatarURL {
			lc.LoginUsecase.UpdateUser(c, user.ID, user.Name, user.AvatarURL)
		}
	} else {
		// Создание нового пользователя, если его нет в базе данных
		user = userEntry
		user.ID, err = lc.LoginUsecase.Create(c, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
	}

	// Создание токенов доступа и обновления
	accessToken, err = lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err = lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Перенаправление пользователя на фронтенд с токенами
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/callback?accessToken=%s&refreshToken=%s", lc.Env.FrontendAddress, accessToken, refreshToken))
}

// LoginEmail	godoc
// @Summary		Вход по почте и паролю
// @Tags        Login
// @Router      /login [post]
// @Success		200		{object}	domain.RefreshTokenResponse
// @Failure		400		{object}	domain.ErrorResponse
// @Param       loginRequest	body	domain.LoginRequest	true	"login request"
// @Produce		json
func (lc *LoginController) LoginEmail(c *gin.Context) {
	var loginRequest domain.LoginRequest
	err := c.ShouldBindBodyWith(&loginRequest, binding.JSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	email := strings.ToLower(loginRequest.Email)
	user, err := lc.LoginUsecase.GetUserByEmail(c, email)
	if err != nil || bcrypt.CompareHashAndPassword(user.Auth.Email.Password, []byte(loginRequest.Password)) != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "User not found"})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
}
