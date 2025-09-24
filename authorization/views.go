package authorization

import (
	"net/http"
	"site_backend/response"

	"github.com/gin-gonic/gin"
)

// @Description authorize user with login and password
// @Tags authorization
// @Param body body UserLogin false "authorization struct"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /auth/token [post]
func GetToken(c *gin.Context) {
	defer Recoverd(c, "GetToken: ")
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
