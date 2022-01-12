package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CreateAccessTokenRequest struct {
	Subject     string   `json:"subject,omitempty"`
	Authorizer  string   `json:"authorizer"`
	Requester   string   `json:"requester"`
	Identity    string   `json:"identity"`
	Service     string   `json:"service"`
	Credentials []string `json:"credentials"`
}

type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

func requestAccessToken(reqValues CreateAccessTokenRequest) (AccessTokenResponse, error) {
	jsonData, _ := json.Marshal(reqValues)
	resp, err := http.Post("http://localhost:2323/internal/auth/v1/request-access-token", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Error().Msg(err.Error())
		return AccessTokenResponse{}, errors.New("request access token failed")
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		res := make(map[string]interface{})
		json.NewDecoder(resp.Body).Decode(&res)
		log.Error().Msgf("%+v", res)
		return AccessTokenResponse{}, errors.New("request access token failed")
	}
	var res AccessTokenResponse
	json.NewDecoder(resp.Body).Decode(&res)
	log.Info().Msgf("%+v", res)
	return res, nil
}

func HandleAccessToken(c *gin.Context) {

	values := CreateAccessTokenRequest{
		Authorizer: "did:nuts:D1tuEqEua3SEntiA8wm48fBACzPaeJ4d2eAqPC4KQqob",
		Requester:  "did:nuts:Hr7YWKvuPZrLKBQZPSxHcqVcamBUDgmHgMdXSb79SzGi",
		Service:    "validated-query",
	}
	jwt, err := requestAccessToken(values)
	if err != nil {
		log.Error().Msgf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, jwt)
}
