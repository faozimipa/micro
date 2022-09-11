package keyloack

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/faozimipa/micro/shared/config"
)


type TokenResponse struct {
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
}

type CredentialForm struct {
    Type        string  `json:"type"`
    Value       string  `json:"value"`
    Temporary   bool    `json:"temporary"`
}
type UserRegisterForm struct {
    FirstName       string           `json:"firstName"`
    LastName        string           `json:"lastName"`
    Email           string           `json:"email"`
    Username        string           `json:"username"`
    Enabled         string           `json:"enabled"`
    Credential      []CredentialForm `json:"credentials"`
}

func GetToken () (TokenResponse, error ){

    var result TokenResponse
    apiUrl := config.getKeyloackUrl()
    resource := "/token"
    data := url.Values{}
    data.Set("client_id", config.AppConfig.KeyloackClientId)
    data.Set("client_secret", config.AppConfig.KeyloackClientSecret)
    data.Set("grant_type", config.AppConfig.KeyloackGrantType)
    data.Set("username", config.AppConfig.KeyloackUsername)
    data.Set("password", config.AppConfig.KeyloackPassword)

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := u.String()

    client := &http.Client{}
    r, _ := http.NewRequest(http.MethodPost, urlStr, strings.NewReader(data.Encode())) // URL-encoded 
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    resp, err := client.Do(r)
	fmt.Println(resp.Status)
    

    if err != nil {
        return result, err
    }
    defer resp.Body.Close()

    err = json.NewDecoder(resp.Body).Decode(&result)
    if err != nil {
        return result, err
    }

    return result, nil
    
	
}

func RegisterUser(user UserRegisterForm) (string, error) {
    return "", nil
}