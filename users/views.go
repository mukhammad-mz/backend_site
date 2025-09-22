package users

import (
	"net/http"
	"site_backend/authorization"
	"site_backend/helper"
	"site_backend/response"

	"github.com/gin-gonic/gin"
)

// @Security ApiKey
// @Tags user
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user [get]
func GetUser(c *gin.Context) {
	defer authorization.Recoverd(c, "GetUser: ")
	userUID, ext := c.Get("userUID")
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	user := userInfo{}
	ext = user.userInfo(userUID.(string))
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	c.JSON(http.StatusOK, response.CorrectWithData(user))
}

// @Security ApiKey
// @Tags user
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/users [get]
func GetUsers(c *gin.Context) {
	defer authorization.Recoverd(c, "GetUsers: ")
	userUID, ext := c.Get("userUID")
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	users := usersInfo{}

	ext = users.usersInfo(userUID.(string))
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	c.JSON(http.StatusOK, response.CorrectWithData(users))
}

// @Security ApiKey
// @Tags user
// @Param body body userInfo false "user struct"
// @Produce json
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user [post]
func PostUser(c *gin.Context) {
	defer authorization.Recoverd(c, "PostUser: ")
	user := &Users{}
	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}

	user.Token = helper.GenerateToken()
	if user.Token == "" {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	user.Password = helper.GenerateMd5(user.Password)
	user.UID = helper.UUID()
	user.DataRegist = helper.GetDateTime()
	user.CreateAt = helper.GetDateTime()
	user.UpdateAt = helper.GetDateTime()
	ext := user.userInsert()
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	c.JSON(http.StatusCreated, response.Createdd())
}

// @Security ApiKey
// @Tags user
// @Produce json
// @Param uid path string true "User Uid"
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/{uid} [delete]
func DeleteUser(c *gin.Context) {
	defer authorization.Recoverd(c, "DelUser: ")
	uid := c.Param("uid")
	if len(uid) < 30 {
		c.JSON(http.StatusBadRequest, response.ErrorFromString(http.StatusText(400)))
		return
	}
	ext := userDel(uid)
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}
	c.JSON(http.StatusOK, response.Correct())
}

// @Security ApiKey
// @Tags user
// @Param body body userInfo false "user struct"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user [put]
func PutUser(c *gin.Context) {
	defer authorization.Recoverd(c, "PutUser: ")
	uid, _ := c.Get("userUID")
	user := &Users{}
	if err := c.BindJSON(user); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}
	user.UpdateAt = helper.GetDateTime()
	ext := user.userUpdate(uid.(string))
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}
	c.JSON(http.StatusOK, response.Correct())
}

// @Security ApiKey
// @Tags user
// @Param body body password false "user struct"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/password [put]
func UserChangePassword(c *gin.Context) {
	defer authorization.Recoverd(c, "UserChangePassword: ")
	passw := &password{}
	if err := c.BindJSON(passw); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}
	uid, _ := c.Get("userUID")
	user := Users{}

	ext := user.userinfo(uid.(string))
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	if user.Password != helper.GenerateMd5(passw.OldPassword) {
		c.JSON(http.StatusBadRequest, response.ErrorFromString("invalid old password"))
		return
	}

	user.Password = helper.GenerateMd5(passw.NewPassword)
	ext = changePassword(uid.(string), user.Password)
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	c.JSON(http.StatusOK, response.Correct())
}

// @Security ApiKey
// @Tags user
// @Param body body login false "login struct"
// @Produce json
// @Success 200 {object} response.Response
// @Success 204 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /user/login  [get]
func chenckLogin(c *gin.Context) {
	defer authorization.Recoverd(c, "chenckLogin: ")

	login := &login{}
	if err := c.BindJSON(login); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}

	ext, count := login.chenckLogin()
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	} else if count == 0 {
		c.JSON(http.StatusNoContent, "")
		return
	}

	c.JSON(http.StatusOK, response.Correct())
}

func permission(c *gin.Context) {
	defer authorization.Recoverd(c, "permission: ")

	uid, _ := c.Get("userUID")
	user := &Users{}
	ext := user.userinfo(uid.(string))
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	perm := &permissions{}
	ext = perm.perms(user.IDRole)
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
		return
	}

	c.JSON(http.StatusOK, response.CorrectWithData(perm))
}
