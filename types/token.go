package types

import (
	"time"
)

type AuthToken struct {
	AccessToken    string `json:"accessToken,omitempty"`
	ExpiresIn      int    `json:"expiresIn,omitempty"`
	RefreshToken   string `json:"refreshToken,omitempty"`
	TokenType      string `json:"tokenType,omitempty"`
	expirationTime time.Time
}

// NewToken creates a new struct from a refresh token
func NewToken(refreshToken string) *AuthToken {
	tk := &AuthToken{
		RefreshToken: refreshToken,
		ExpiresIn:    0,
	}
	tk.SetExpirationTime()
	return tk
}

// SetExpirationTime converts the ExpiresIn value to a time.Time
func (m *AuthToken) SetExpirationTime() {
	m.expirationTime = time.Now().UTC().Add(time.Duration(m.ExpiresIn) * time.Second)
}

// IsExpired returns whether the access token has expired
func (m *AuthToken) IsExpired() bool {
	return m.expirationTime.Unix() < time.Now().UTC().Unix()
}

// SetAccessToken updates the struct and sets a new access token
func (m *AuthToken) SetAccessToken(tk *AuthToken) {
	m.AccessToken = tk.AccessToken
	m.ExpiresIn = tk.ExpiresIn
	m.SetExpirationTime()
}
