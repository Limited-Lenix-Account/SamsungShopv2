package jwt

import "github.com/dgrijalva/jwt-go"

type DefaultHeaders struct {
	Alg  string `json:"alg"`
	Type string `json:"type"`
}

type JwtPayload struct {
	AppVersion       string `json:"app_version,omitempty"`       // 2.0.35059
	AccessToken      string `json:"access_token,omitempty"`      // from login
	UniqueUserID     string `json:"unique_user_id,omitempty"`    // in android "cream" by default?
	UserID           string `json:"user_id,omitempty"`           // from login
	UniqueAppId      string `json:"unique_app_id,omitempty"`     // uuid for app
	StoreArbitration bool   `json:"store_arbitration,omitempty"` // true
	ApiUrl           string `json:"api_server_url,omitempty"`    //optional maybe
	SignedAt         int64  `json:"signed_at,omitempty"`         // UNIX time
	KeyVersion       string `json:"key_version,omitempty"`       // v1
	jwt.StandardClaims
}

type AccessResp struct {
	AccessToken           string `json:"access_token"`
	AccessSecret          string `json:"access_secret"`
	TokenType             string `json:"token_type"`
	AccessTokenExpiresIn  string `json:"access_token_expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshSecret         string `json:"refresh_secret"`
	RefreshTokenExpiresIn string `json:"refresh_token_expires_in"`
	UUID                  string `json:"uuid"`
	UserID                string `json:"userId"`
	ClientID              string `json:"client_id"`
	InputEmailID          string `json:"inputEmailID"`
	APIServerURL          string `json:"api_server_url"`
	AuthServerURL         string `json:"auth_server_url"`
	Close                 bool   `json:"close"`
	ClosedAction          string `json:"closedAction"`
}
