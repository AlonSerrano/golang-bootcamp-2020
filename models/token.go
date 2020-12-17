package models

// TokenInfoResp this structure is intended to return a user authentication
type TokenInfoResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
