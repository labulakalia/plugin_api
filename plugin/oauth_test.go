package plugin

import (
	"fmt"
	"testing"
	"time"
)

func TestOAuth(t *testing.T) {
	oauthConfig := OauthConfig{
		ClientId:     "your_client_id",
		ClientSecret: "your_client_secret",
		RedirectUri:  "http://localhost:8080/callback",
		Scopes:       []string{"openid", "profile"},
		AuthUrl:      "http://auth_url",
		TokenUrl:     "http://token_url",
		TokenReqType: "json",
	}
	t.Log(oauthConfig.GetAuthAddr(fmt.Sprint(time.Now().Unix()), map[string]string{"param": "param1"}))
}
