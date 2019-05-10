package middlewares

import (
	"fmt"

	"github.com/thiepwong/resident-manager/models"

	"github.com/kataras/iris"
)

func Authorization(c iris.Context) {
	_token := c.GetHeader("Authorization")
	if _token == "" {
		return
	}
	var _auth = &models.Auth{
		Issuer:  "Nguyen Ai Quoc",
		IssueId: "2342342342",
	}
	if _auth != nil {
		fmt.Print("Da co authen")
		c.Values().Set("Auth", _auth)
		c.Next()
	} else {
		return
	}
}
