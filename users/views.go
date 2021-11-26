package users

import (
	"net/http"
	"site_backend/helper"
	"site_backend/response"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
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

func GetUsers(c *gin.Context) {
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
func PostUser(c *gin.Context) {
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
func DeleteUser(c *gin.Context) {
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
func PutUser(c *gin.Context) {
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
func UserChangePassword(c *gin.Context) {
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

func chenckLogin(c *gin.Context) {
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
