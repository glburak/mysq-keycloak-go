package controller

import (
	"github.com/Nerzal/gocloak/v11"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"context"
	"log"
	"fmt"
)


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
	token, err := client.Login(ctx, "login", "7MSrOXewwhx4guCuV0aBGLYaSSSxAWeH", "password-manager", username, password)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": "Başarısız"})

		panic(err.Error())
	} else {
		session := sessions.Default(c)
		session.Set("token", token.AccessToken )

		log.Println(token)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"message": token})
		

	}
}
func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()
	c.JSON(http.StatusOK, gin.H{
		"message": "User Sign out successfully",
	})
}

func IsLogin(token string) (bool){
	client := gocloak.NewClient("http://localhost:8090/")
	ctx := context.Background()
	
	rptResult, err := client.RetrospectToken(ctx, token, "login", "7MSrOXewwhx4guCuV0aBGLYaSSSxAWeH", "password-manager")
	if err != nil {
		panic("Inspection failed:" + err.Error())
	}
	var active bool
	active = *rptResult.Active

	if !active {
		
		return false
	}

	return true
}

func User(c *gin.Context){
	client := gocloak.NewClient("http://localhost:8090/")
	ctx := context.Background()
	session := sessions.Default(c)
	token := fmt.Sprint(session.Get("token"))
	fmt.Printf("var1 = %T\n", token)
	userinfo,_ :=client.GetUserInfo(ctx,token,"password-manager")


	if IsLogin(token){
		c.JSON(http.StatusOK, gin.H{"message": userinfo.Sub})
	}else{
		c.JSON(http.StatusForbidden, gin.H{"message":"forbidden" })
	}
	
}