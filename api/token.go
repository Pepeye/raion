package api

// AuthToken struct
type AuthToken struct {
	Token        string `json:"token"`
	IDToken      `json:"idToken"`
	RefreshToken `json:"refreshToken"`
	AccessToken  `json:"accessToken"`
}

// Token struct
type Token struct {
	JwtToken string `json:"jwtToken"`
}

// IDToken Struct
type IDToken Token

// RefreshToken Struct
type RefreshToken Token

// AccessToken Struct
type AccessToken Token
