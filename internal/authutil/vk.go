package authutil

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/vk"
	"io/ioutil"
	"net/http"
	"stontactics/bootstrap"
	"stontactics/domain"
)

const (
	oauthVKUrlAPI   = "https://api.vk.com/method/users.get?fields=photo_200,email&access_token="
	vkAPIVersion    = "5.131"
	vkUserDataField = "response"
)

func GetConfigVK(env *bootstrap.Env) *oauth2.Config {
	return &oauth2.Config{
		RedirectURL:  fmt.Sprintf("%s/auth/vk/callback", env.ServerAddress),
		ClientID:     env.VKClientID,
		ClientSecret: env.VKClientSecret,
		Scopes:       []string{"email"},
		Endpoint:     vk.Endpoint,
	}
}

func GetUserDataFromVK(env *bootstrap.Env, code string) (*domain.VKUserResponse, error) {
	// Use code to get token and get user info from VK.
	token, err := GetConfigVK(env).Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	client := &http.Client{}
	request, err := http.NewRequest("GET", oauthVKUrlAPI+token.AccessToken+"&v="+vkAPIVersion, nil)
	if err != nil {
		return nil, fmt.Errorf("failed creating request: %s", err.Error())
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	var vkResponse struct {
		Response []domain.VKUserResponse `json:"response"`
	}
	if err := json.Unmarshal(contents, &vkResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal user data: %s", err.Error())
	}

	if len(vkResponse.Response) == 0 {
		return nil, fmt.Errorf("no user data found in VK response")
	}

	return &vkResponse.Response[0], nil
}
