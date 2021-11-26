package authorization

import (
	"site_backend/db"
	"site_backend/users"
	"time"

	log "github.com/sirupsen/logrus"
)

func getUserToken(user *UserLogin) (*UserToken, error) {
	result := &UserToken{}
	token := users.NewToken()
	println(user.Password)
	getDB := db.GetDB()
	getDB = getDB.Table("users").Where("login = ? and password = md5(?)", user.Login, user.Password).Updates(map[string]interface{}{"token": token, "update_at": time.Now()}).Scan(&result)
	if getDB.Error != nil {
		log.Error("getUserToken: ", getDB.Error)
		return nil, getDB.Error
	}
	return result, nil
}

func ChekToken(t string) bool {
	count := 0
	con := db.GetDB()
	con = con.Table("users").Where("token = ?", t).Count(&count)
	if con.Error != nil {
		log.Error("chekToken: ", con.Error)
		return false
	} else if count == 0 {
		return false
	}
	return true
}
