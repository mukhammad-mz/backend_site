package authorization

import (
	"net/http"
	"site_backend/response"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	userLogin := &UserLogin{}
	if err := c.BindJSON(userLogin); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}
	userToken, err := getUserToken(userLogin)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromString("invalid login/password"))
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(userToken))
}
