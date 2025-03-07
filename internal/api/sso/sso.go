package sso

import (
	"common-sync/pkg/third_sso"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Oauth(c *gin.Context) {
	redirectURI := c.Query("redirect_uri")
	state := c.Query("state")
	c.SetCookie("redirect_uri", redirectURI, 3600, "/", "", false, true)
	c.SetCookie("state", state, 3600, "/", "", false, true)
	authorizationURL := third_sso.ThirdSSOInstance.GetAuthorizationURL(redirectURI, state)
	c.Redirect(http.StatusFound, authorizationURL)
}
func Code(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}
	state := c.Query("state")
	redirectURI, err := c.Cookie("redirect_uri")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "redirect_uri cookie not found"})
		return
	}

	redirectURIModule := redirectURI + "&code=" + code + "&state=" + state
	c.Redirect(http.StatusFound, redirectURIModule)
}
func Token(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "code is required"})
		return
	}
	accessToken, err := third_sso.ThirdSSOInstance.ExchangeCodeForToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to exchange code for token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
func UserInfo(c *gin.Context) {
	accessToken := c.Query("access_token")
	unionId, err := third_sso.ThirdSSOInstance.GetUserInfo(accessToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to get user info"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"union_id": unionId})
}
