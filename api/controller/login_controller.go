package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
	"stontactics/internal/authutil"
	"strconv"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (sc *LoginController) BeginLogin(c *gin.Context) {
	provider := c.Param("provider")
	oauthState := authutil.GenerateStateOauthCookie(c.Writer)

	var url string
	switch provider {
	case "google":
		url = authutil.GetConfigGoogle(sc.Env).AuthCodeURL(oauthState)
	case "vk":
		url = authutil.GetConfigVK(sc.Env).AuthCodeURL(oauthState)
	default:
		c.String(http.StatusInternalServerError, "Invalid provider")
		return
	}
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func (sc *LoginController) Login(c *gin.Context) {
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
	switch provider {
	case "google":
		// Получение данных от Google
		data, err := authutil.GetUserDataFromGoogle(sc.Env, c.Query("code"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		// Валидация данных пользователя от Google
		if data.ID == "" || !data.VerifiedEmail {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Incomplete user data received from Google"})
			return
		}

		user = domain.User{
			ID:        data.ID,
			Name:      data.Name,
			Email:     data.Email,
			AvatarURL: data.Picture,
			Pro: domain.UserPro{
				Active: true,
				Until:  nil,
			}, // TODO: set false in prod
		}
	case "vk":
		// Получение данных от VK
		data, err := authutil.GetUserDataFromVK(sc.Env, c.Query("code"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}

		// Валидация данных пользователя от VK
		if data.ID == 0 {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: "Incomplete user data received from VK"})
			return
		}

		user = domain.User{
			ID:        strconv.Itoa(data.ID),
			Name:      data.FirstName + " " + data.LastName,
			Email:     data.Email,
			AvatarURL: data.Photo200,
		}
	default:
		c.String(http.StatusInternalServerError, "Invalid provider")
		return
	}

	// Получение пользователя из базы данных
	dbUser, err := sc.LoginUsecase.GetUserByID(c, user.ID)
	var accessToken, refreshToken string
	if err == nil {
		// Обновление данных пользователя, если они изменились
		if dbUser.Name != user.Name || dbUser.AvatarURL != user.AvatarURL {
			sc.LoginUsecase.UpdateUser(c, user.ID, user.Name, user.AvatarURL)
		}
		user = dbUser
	} else {
		// Создание нового пользователя, если его нет в базе данных
		err = sc.LoginUsecase.Create(c, &user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
			return
		}
	}

	// Создание токенов доступа и обновления
	accessToken, err = sc.LoginUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err = sc.LoginUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	// Перенаправление пользователя на фронтенд с токенами
	c.Redirect(http.StatusFound, fmt.Sprintf("%s/callback?accessToken=%s&refreshToken=%s", sc.Env.FrontendAddress, accessToken, refreshToken))
}
