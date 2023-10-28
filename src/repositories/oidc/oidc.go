package oidc

import (
	"encoding/json"
	"fmt"
	"goth/src/config"
	"goth/src/utils"
	"net/url"
	"time"
)

type oidcWellKnown struct {
	Issuer                string `json:"issuer"`
	AuthorizationEndpoint string `json:"authorization_endpoint"`
	TokenEndpoint         string `json:"token_endpoint"`
	UserInfoEndpoint      string `json:"userinfo_endpoint"`
}

type oidc struct {
	wellKnown    oidcWellKnown
	clientID     string
	clientSecret string
	redirectURI  string
}

func NewOIDC(cfg config.ServerConfigs) (OIDC, error) {
	baseClient := utils.HTTPClient{BaseUrl: cfg.OIDCIssuer}
	wellKnownResponse, err := baseClient.Request("GET", "/.well-known/openid-configuration", nil)
	if err != nil {
		return nil, err
	}

	var wellKnown oidcWellKnown
	err = json.Unmarshal(wellKnownResponse.Body, &wellKnown)
	if err != nil {
		return nil, err
	}

	return &oidc{
		wellKnown:    wellKnown,
		clientID:     cfg.OIDCClientID,
		clientSecret: cfg.OIDCClientSec,
		redirectURI:  cfg.OIDCRedirectURI,
	}, nil
}

func (o *oidc) GetAuthorizeURL(scope, state string) string {
	query := url.Values{}
	query.Set("client_id", o.clientID)
	query.Set("response_type", "code")
	query.Set("redirect_uri", o.redirectURI)
	query.Set("scope", scope)
	query.Set("state", state)

	return fmt.Sprintf("%s/oauth/v2/authorize?%s", o.wellKnown.Issuer, query.Encode())
}

func (o *oidc) getTokens(code string) (string, int, error) {
	headers := utils.Headers{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	httpClient := utils.HTTPClient{
		Headers: headers,
		BaseUrl: o.wellKnown.TokenEndpoint,
	}

	body := url.Values{}
	body.Set("grant_type", "authorization_code")
	body.Set("client_id", o.clientID)
	body.Set("client_secret", o.clientSecret)
	body.Set("code", code)
	body.Set("redirect_uri", o.redirectURI)

	res, err := httpClient.Request("POST", "", []byte(body.Encode()))
	if err != nil {
		return "", 0, err
	}

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		// IdToken     string `json:"id_token"`
		// TokenType   string `json:"token_type"`
	}

	var tokenResponse TokenResponse
	err = json.Unmarshal(res.Body, &tokenResponse)
	if err != nil {
		return "", 0, err
	}

	return tokenResponse.AccessToken, tokenResponse.ExpiresIn, nil
}

func (o *oidc) getUserInfo(accessToken string) (UserInfoResponse, error) {
	headers := utils.Headers{
		"Authorization": fmt.Sprintf("Bearer %s", accessToken),
	}

	httpClient := utils.HTTPClient{
		Headers: headers,
		BaseUrl: o.wellKnown.UserInfoEndpoint,
	}

	res, err := httpClient.Request("GET", "", nil)
	if err != nil {
		return UserInfoResponse{}, err
	}

	var userInfoResponse UserInfoResponse
	err = json.Unmarshal(res.Body, &userInfoResponse)
	if err != nil {
		return UserInfoResponse{}, err
	}

	return userInfoResponse, nil
}

func (o *oidc) GetUser(code string) (UserInfoResponse, time.Time, error) {
	accessToken, expiresIn, err := o.getTokens(code)
	if err != nil {
		return UserInfoResponse{}, time.Time{}, err
	}

	userInfoResponse, err := o.getUserInfo(accessToken)
	if err != nil {
		return UserInfoResponse{}, time.Time{}, err
	}

	return userInfoResponse, time.Now().Add(time.Duration(expiresIn) * time.Second), nil
}
