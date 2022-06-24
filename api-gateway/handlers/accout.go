package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sy-yoon/krealtors/api-gateway/middleware"
	userpb "github.com/sy-yoon/krealtors/protos/g1/user"
)

const (
	InvalidPassword = 1
	InvalidUser
)

type LoginForm struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AccountHandler struct {
	UserClient userpb.UserServiceClient
}

func (me *AccountHandler) SignIn(c *gin.Context) {
	var loginForm LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err})
		log.Fatal(err)
		return
	}

	req := &userpb.GetUserRequest{
		Email: loginForm.Email,
	}

	user, err := me.UserClient.GetUserByEmail(context.Background(), req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		log.Fatal(err)
		return
	}

	if user == nil {
		log.Println("SignIn", "auth", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid User", "errorCode": InvalidUser})
		return
	}

	if user.Password != loginForm.Password {
		log.Println("SignIn", "session", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password", "errorCode": InvalidPassword})
		return
	}

	tokenDetails, err := middleware.NewToken(user)
	if err != nil {
		log.Println("SignIn", "session", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password", "errorCode": InvalidPassword})
		return
	}

	if err = middleware.StoreSession(user.Id, tokenDetails); err != nil {
		log.Println("SignIn", "session", err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid password", "errorCode": InvalidPassword})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userName": user.Name})
}

func (me *AccountHandler) SignOut(c *gin.Context) {

}
