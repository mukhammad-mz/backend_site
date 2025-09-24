package authorization

import (
	"net/http"
	"site_backend/db"
	"site_backend/helper"
	"site_backend/response"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func getUserToken(user *UserLogin) (*UserToken, error) {
	result := &UserToken{}
	getDB := db.GetDB()
	getDB = getDB.Table("users").
		Where("login = ? and password = ? and is_delete = ?",
			user.Login, helper.GenerateMd5(user.Password), 1).
		//Update("token", helper.GenerateToken()).
		Scan(&result)

	if getDB.Error != nil {
		log.Error("getUserToken: ", getDB.Error)
		return nil, getDB.Error
	}
	return result, nil
}

func ChekToken(t string) (string, error) {
	userToken := &UserToken{}
	con := db.GetDB()
	con = con.Table("users").Where("token = ?", t).
		Updates(map[string]interface{}{"update_at": helper.GetDateTime(),
			"last_visit": helper.GetDateTime()}).
		Scan(userToken)
	if con.Error != nil {
		log.Error("chekToken: ", con.Error)
		return "", con.Error
	}
	return userToken.UID, con.Error
}

func Recoverd(c *gin.Context, funcName string) {
	if err := recover(); err != nil {
		log.Error(funcName, err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response.ServerError())
	}
}
