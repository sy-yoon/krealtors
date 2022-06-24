package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	userpb "github.com/sy-yoon/krealtors/protos/g1/user"
	"github.com/sy-yoon/krealtors/utils"
)

var ACCESS_SECRET = []byte(os.Getenv("ACCESS_SECRET"))
var REFRESH_SECRET = []byte(os.Getenv("REFRESH_SECRET"))

type authHeader struct {
	IDToken string `header:"Authorization"`
}

type TokenDetails struct {
	AccessToken        string
	RefreshToken       string
	AccessTokenUUID    string
	RefreshTokenUUID   string
	AccessTokenExpiry  int64
	RefreshTokenExpiry int64
}

func NewToken(user *userpb.User) (*TokenDetails, error) {
	var tokenDetails TokenDetails
	var err error

	if err = newAccessToken(user.Id, &tokenDetails); err != nil {
		return nil, err
	}

	if err = newRefreshToken(user.Id, &tokenDetails); err != nil {
		return nil, err
	}

	return &tokenDetails, nil
}

func newAccessToken(userId int64, tokenDetails *TokenDetails) error {
	accessTokenClaims := jwt.MapClaims{}
	accessTokenClaims["id"] = userId
	accessTokenClaims["access_uuid"] = uuid.New().String()
	//accessTokenClaims["name"] = user.Name
	//accessTokenClaims["email"] = user.Email
	//atClaims["role"] = user.Role
	accessTokenClaims["exp"] = utils.UtcNow().Add(time.Minute * 15).Unix() // 15 minutes

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	tokenString, err := accessToken.SignedString(ACCESS_SECRET)
	if err != nil {
		return err
	}

	tokenDetails.AccessToken = tokenString
	tokenDetails.AccessTokenExpiry = accessTokenClaims["exp"].(int64)
	tokenDetails.AccessTokenUUID = accessTokenClaims["access_uuid"].(string)
	return nil
}

func newRefreshToken(userId int64, tokenDetails *TokenDetails) error {
	refreshTokenClaims := jwt.MapClaims{}
	refreshTokenClaims["id"] = userId
	refreshTokenClaims["refresh_uuid"] = uuid.New().String()
	//accessTokenClaims["name"] = user.Name
	//accessTokenClaims["email"] = user.Email
	//atClaims["role"] = user.Role
	refreshTokenClaims["exp"] = utils.UtcNow().Add(time.Minute * 15).Unix() // 15 minutes

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	tokenString, err := accessToken.SignedString(REFRESH_SECRET)
	if err != nil {
		return err
	}

	tokenDetails.RefreshToken = tokenString
	tokenDetails.RefreshTokenExpiry = refreshTokenClaims["exp"].(int64)
	tokenDetails.RefreshTokenUUID = refreshTokenClaims["refresh_uuid"].(string)
	return nil
}

func IsAuthorized(c *gin.Context) {
	header := authHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	token, err := jwt.Parse(header.IDToken,
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf(("invalid Signing Method"))
			}
			return ACCESS_SECRET, nil
		})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err})
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok && token.Valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token"})
	}

	c.Next()
}

func StoreSession(userId int64, tokenDetails *TokenDetails) error {
	redisClient := utils.GetRedisClient()
	ctx := context.Background()
	now := utils.UtcNow()
	err := redisClient.Set(ctx, tokenDetails.AccessTokenUUID, userId, time.Unix(tokenDetails.AccessTokenExpiry, 0).Sub(now)).Err()
	if err != nil {
		return err
	}

	err = redisClient.Set(ctx, tokenDetails.RefreshTokenUUID, userId, time.Unix(tokenDetails.RefreshTokenExpiry, 0).Sub(now)).Err()
	if err != nil {
		return err
	}

	return nil
}


