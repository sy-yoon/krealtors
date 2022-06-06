package middleware

// // import (
// // 	"fmt"
// // 	"net/http"
// // 	"os"
// // 	"time"

// // 	jwt "github.com/dgrijalva/jwt-go"
// // 	userpb "github.com/sy-yoon/krealtors/protos/g1/user"
// // )

// var ACCESS_SECRET = []byte(os.Getenv("ACCESS_SECRET"))
// var REFRESH_SECRET = []byte(os.Getenv("REFRESH_SECRET"))

// type AuthTokenClaims struct {
// 	UserID             string   `json:"id"`   // 유저 ID
// 	Name               string   `json:"name"` // 유저 이름
// 	Email              string   `json:"mail"` // 유저 메일
// 	Role               []string `json:"role"` // 유저 역할
// 	jwt.StandardClaims          // 표준 토큰 Claims
// }

// type authHeader struct {
// 	IDToken string `header:"Authorization"`
// }

/*
func NewToken(user *userpb.User) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = user.Id
	atClaims["name"] = user.Name
	atClaims["email"] = user.Email
	//atClaims["role"] = user.Role
	atClaims["exp"] = time.Now().Add(time.Minute * 15)

	atoken := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	atoken.SignedString(ACCESS_SECRET)

	rtClaims := jwt.MapClaims{}
	rtClaims["id"] = user.Id
	rtClaims["name"] = user.Name
	rtClaims["email"] = user.Email
	//atClaims["role"] = user.Role
	rtClaims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 180) // 6개월
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	rt.SignedString(REFRESH_SECRET)

}

func IsAuthorized(c *gin.Context) {

	header := authHeader{}
	if err := c.ShouldBindHeader(&header); err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	token, err := jwt.Parse(header.IDToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(("invalid Signing Method"))
		}

		// aud := "billing.jwtgo.io"
		// checkAudience := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
		// if !checkAudience {
		// 	return nil, fmt.Errorf(("invalid aud"))
		// }
		// // verify iss claim
		// iss := "jwtgo.io"
		// checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
		// if !checkIss {
		// 	return nil, fmt.Errorf(("invalid iss"))
		// }
		return ACCESS_SECRET, nil
	})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
	}

	if token.Valid {
		c.Next()
	}


}
*/
