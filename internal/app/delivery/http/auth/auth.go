package auth

import (
	"github.com/dgrijalva/jwt-go"
)

const _secret = "1f635616d380bc3bcda2f255ea796575"

type claimsWithUserID struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// Encode generates JWT token from claimsWithUserID
func (c *claimsWithUserID) Encode() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	signed, err := token.SignedString([]byte(_secret))
	if err != nil {
		return "", err
	}

	return signed, nil
}

// NewClaims creates claimsWithUserID object
func NewClaims(userID int, expiresAt int64) *claimsWithUserID {
	return &claimsWithUserID{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}
}

// Decode decodes provided token to claimsWithUserID object
func DecodeClaims(token string) (*claimsWithUserID, error) {
	claims := &claimsWithUserID{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(_secret), nil
	})
	if err != nil {
		return nil, err
	}

	return claims, nil
}
