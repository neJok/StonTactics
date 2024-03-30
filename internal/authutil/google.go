package authutil

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
)

func GetConfigGoogle(env *bootstrap.Env) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/auth/google/callback", env.ServerAddress),
		ClientID:     env.GoogleClientID,
		ClientSecret: env.GoogleClientSecret,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile", "openid"},
		Endpoint:     google.Endpoint,
	}
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func GetUserDataFromGoogle(env *bootstrap.Env, code string) (*domain.GoogleUserResponse, error) {
	// Use code to get token and get user info from Google.

	token, err := GetConfigGoogle(env).Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	var user domain.GoogleUserResponse
	if err := json.Unmarshal(contents, &user); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %s", err.Error())
	}

	return &user, nil
}
