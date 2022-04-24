package user_service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"

	"github.com/KapitanD/cyan-project/pkg/model"
)

var (
	// Define a secure key string used
	// as a salt when hashing our tokens.
	// Please make your own way more secure than this,
	// use a randomly generated md5 hash or something.
	key = []byte("mySuperSecretKey")
)

// CustomClaims is our custom metadata, which will be hashed
// and sent as the second segment in our JWT
type CustomClaims struct {
	User *model.User
	jwt.StandardClaims
}

type TokenService struct {
}

// Decode a token string into a token object
func (srv *TokenService) Decode(token string) (*CustomClaims, error) {

	// Parse the token
	tokenType, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		return nil, errors.Wrap(err, "error while parsing auth token")
	}

	// Validate the token and return the custom claims
	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

// Encode a claim into a JWT
func (srv *TokenService) TokenPair(user *model.User) (map[string]string, error) {
	// Create the Claims
	claims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(60 * time.Minute).Unix(),
			Issuer:    "user.service",
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	at, err := token.SignedString(key)
	if err != nil {
		return nil, err
	}

	refreshClaims := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "user.service",
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	rt, err := refreshToken.SignedString(key)
	if err != nil {
		return nil, err
	}

	// Sign token and return
	return map[string]string{
		"access_token":  at,
		"refresh_token": rt,
	}, nil
}

func (srv *TokenService) Refresh(refreshToken string) (map[string]string, error) {
	tokenType, err := jwt.ParseWithClaims(refreshToken, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if claims, ok := tokenType.Claims.(*CustomClaims); ok && tokenType.Valid {
		// Get the user record from database or
		// run through your business logic to verify if the user can log in
		if claims.Subject == "user.service" {

			newTokenPair, err := srv.TokenPair(claims.User)
			if err != nil {
				return nil, err
			}

			return newTokenPair, nil
		}

		return nil, errors.New("unauthorized")
	}

	return nil, err
}
