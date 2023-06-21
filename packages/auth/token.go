package auth

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	errorpkg "github.com/example-golang-projects/clean-architecture/packages/error"
	"github.com/google/uuid"

	// "github.com/example-golang-projects/clean-architecture/packages/id"

	"github.com/dgrijalva/jwt-go"
)

const (
	AuthorizationHeader         = "Authorization"
	AuthorizationScheme         = "Bearer"
	TokenExpiresDuration        = 60
	RefreshTokenExpiresDuration = 24 * 7
)

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AtExpires    int64
	RtExpires    int64
}

type SessionInfo struct {
	UserID uuid.UUID
}

type JWTCustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}

func GenerateToken(userID uuid.UUID) (*TokenDetails, error) {
	var err error
	//Creating Access Token
	atExpires := time.Now().Add(time.Minute * TokenExpiresDuration).Unix()
	atClaims := &JWTCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: atExpires,
		},
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	accessToken, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtExpires := time.Now().Add(time.Minute * RefreshTokenExpiresDuration).Unix()
	rtClaims := &JWTCustomClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: rtExpires,
		},
	}
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	refreshToken, err := rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return &TokenDetails{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		AtExpires:    atExpires,
		RtExpires:    rtExpires,
	}, nil
}

// Get token string from Authorization Header (Ex: Bearer bncdyfcg812h3ndsya8dg68sd12...)
func getTokenStrFromHeader(header http.Header) (string, error) {
	s := header.Get(AuthorizationHeader)
	if s == "" {

		return "", errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Missing authorization string.")))
	}
	splits := strings.SplitN(s, " ", 2)
	if len(splits) < 2 {
		return "", errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Bad authorization string.")))
	}
	if splits[0] != AuthorizationScheme {
		return "", errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Request unauthenticated with %v", AuthorizationScheme)))
	}
	return splits[1], nil
}

func getTokenFromRequest(r *http.Request) (*jwt.Token, error) {
	tokenStr, err := getTokenStrFromHeader(r.Header)
	if err != nil {
		return nil, err
	}
	atClaims := &JWTCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, atClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Validation Error Malformed")))
				return nil, errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Validation Error Malformed")))
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired

				return nil, errorpkg.ErrAccessTokenExpired(errors.New(fmt.Sprintf("Token have already expried")))

			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errorpkg.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Validation Error Not Valid Yet")))
			} else {
				return nil, errorpkg.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))
			}
		}
	}
	if time.Now().Unix() > atClaims.ExpiresAt {
		return nil, errorpkg.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))
	}
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetCustomClaimsFromRequest(r *http.Request) (*JWTCustomClaims, error) {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return nil, errorpkg.ErrInvalidAccessToken(errors.New(fmt.Sprintf("Token Invalid")))

	}

	claim, ok := token.Claims.(*JWTCustomClaims)
	if ok && token.Valid {
		return claim, nil
	}
	return nil, nil
}

func RefreshToken(refreshToken string) (*TokenDetails, error) {
	// Verify the token
	var customClaims JWTCustomClaims
	token, err := jwt.ParseWithClaims(refreshToken, customClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})

	// If there is an error, the token must have expired
	if err != nil {
		return nil, errorpkg.ErrRefreshTokenExpired(errors.New(fmt.Sprintf("Refresh token expired")))
	}

	// Check token valid
	if _, ok := token.Claims.(JWTCustomClaims); !ok || !token.Valid {
		return nil, errorpkg.ErrInvalidRefreshToken(errors.New(fmt.Sprintf("Refresh token does not valid")))
	}

	// Generate new tokens
	customClaims, _ = token.Claims.(JWTCustomClaims)
	userID := customClaims.UserID
	if userID.String() == "" {
		return nil, errorpkg.ErrAuthFailure(errors.New(fmt.Sprintf("Unauthorized")))
	}

	//Create new pairs of refresh and access tokens
	ts, err := GenerateToken(userID)
	if err != nil {
		return nil, errorpkg.ErrForbidden(err)

	}
	return ts, nil
}

func ValidateToken(r *http.Request) error {
	token, err := getTokenFromRequest(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(JWTCustomClaims); !ok || !token.Valid {
		return err
	}
	return nil
}
