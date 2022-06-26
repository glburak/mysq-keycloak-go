package keycloak

import (
	"context"
	"net/http"
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-gonic/gin"


)

func CreateUser(c *gin.Context) {
	client := gocloak.NewClient("http://localhost:8090/")
	ctx := context.Background()
	token, err := client.LoginAdmin(ctx, "burak", "123123", "password-manager")
	if err != nil {
		panic(err.Error())
		panic("Something wrong with the credentials or url")
	}

	user := gocloak.User{
		FirstName: gocloak.StringP("Bob"),
		LastName:  gocloak.StringP("Uncle"),
		Email:     gocloak.StringP("something@really.wrong"),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP("CoolGuy"),
	}
	_, err = client.CreateUser(ctx, token.AccessToken, "password-manager", user)
	if err != nil {
		panic(err.Error())

		panic("Oh no!, failed to create user :(")
	}
}

func Login(c *gin.Context) {
	client := gocloak.NewClient("http://localhost:8090/")
	ctx := context.Background()

	type Login struct {
		Username string
		Password string
	}
	var json Login
	c.BindJSON(&json)
	username := json.Username
	password := json.Password
	_, err := client.Login(ctx, "login", "7MSrOXewwhx4guCuV0aBGLYaSSSxAWeH", "password-manager", username, password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Unauthorized"})

		panic(err.Error())
	} else {
		
		c.JSON(http.StatusOK, gin.H{"message": "Successfuly"})
	}
}
func IsLogin(c *gin.Context) {
	client := gocloak.NewClient("http://localhost:8090/")
	ctx := context.Background()
	token, _ := c.Params.Get("id")
	rptResult, err := client.RetrospectToken(ctx, token, "login", "7MSrOXewwhx4guCuV0aBGLYaSSSxAWeH", "password-manager")
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}
	var active bool
	active = *rptResult.Active

	if !active {
		
		c.JSON(http.StatusForbidden, gin.H{"message": "Başarısız"})
		
	}else {
		
		c.JSON(http.StatusOK,  gin.H{"message": "Başarılı"})
	}

	

}
