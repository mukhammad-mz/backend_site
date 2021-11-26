package authorization

import (
	"site_backend/db"
	"site_backend/helper"

	log "github.com/sirupsen/logrus"
)

func getUserToken(user *UserLogin) (*UserToken, error) {
	result := &UserToken{}
	getDB := db.GetDB()
	getDB = getDB.Table("users").Where("login = ? and password = ?", user.Login, helper.GenerateMd5(user.Password)).Scan(&result)
	if getDB.Error != nil {
		log.Error("getUserToken: ", getDB.Error)
		return nil, getDB.Error
	}
	return result, nil
}

func ChekToken(t string) bool {
	count := 0
	con := db.GetDB()
	con = con.Table("users").Where("token = ?", t).
		Updates(map[string]interface{}{"update_at": helper.GetDateTime(), "last_visit": helper.GetDateTime()}).Count(&count)
	if con.Error != nil {
		log.Error("chekToken: ", con.Error)
		return false
	} else if count == 0 {
		return false
	}
	return true
}
