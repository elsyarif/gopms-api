package entities

type Authentication struct {
	Token string `db:"token"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type AuthRequest struct {
	Identity string `json:"identity,omitempty" form:"identity" binding:"required"`
	Password string `json:"password,omitempty" form:"password" binding:"required"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token,omitempty" from:"refresh_token" binding:"required"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token,omitempty"`
}
