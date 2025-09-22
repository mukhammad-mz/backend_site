package users

import (
	"site_backend/db"
	"time"

	log "github.com/sirupsen/logrus"
)

func CheckPermission(userID, handlerName string) bool {
	db := db.GetDB()
	var count int64 = 0
	errdb := db.Table("users").Where("id_role = ? and uid = ?", 1, userID).Count(&count)
	if errdb.Error != nil {
		log.Error("Permissions: ", errdb.Error)
		return false
	} else if count > 0 {
		return true
	}

	errdb = db.Table("users").
		Joins("JOIN accses ON accses.role_id = users.id_role and users.uid=?", userID).
		Joins("JOIN handlers ON handlers.id = accses.id_handler and handlers.name=?", handlerName).
		Joins("join role on role.id=users.id_role").Count(&count)

	if errdb.Error != nil {
		log.Error("Permissions: ", errdb.Error)
		return false
	}
	return count > 0
}

func (user *userInfo) userInfo(userUID string) bool {
	db := db.GetDB()
	res := db.Table("users").Where("uid=?", userUID).Scan(user)
	if res.Error != nil {
		log.Error("userInfo: ", res.Error)
		return false
	}
	return true
}

func (users *usersInfo) usersInfo(uid string) bool {
	db := db.GetDB()
	res := db.Table("users").Where("uid != ?", uid).Scan(users)
	if res.Error != nil {
		log.Error("usersInfo: ", res.Error)
		return false
	}
	return true
}

func (user *Users) userInsert() bool {
	db := db.GetDB()
	err := db.Table("users").Save(&user)
	if err.Error != nil {
		log.Error("user Insert: ", err.Error)
		return false
	}
	return true
}

func userDel(uid string) bool {
	db := db.GetDB()
	db = db.Table("users").Where("uid = ?", uid).
		Updates(map[string]interface{}{"update_at": time.Now(), "is_delete": 0})
	if db.Error != nil {
		log.Error("user Delet: ", db.Error, " User Id", uid)
		return false
	}
	return true
}

func (user *Users) userUpdate(uid string) bool {
	db := db.GetDB()
	// Use Updates to update from a struct or map in GORM
	res := db.Table("users").Where("uid = ?", uid).Updates(user)
	if res.Error != nil {
		log.Error("user Update: ", res.Error)
		return false
	}
	return true
}

func (user *Users) userinfo(uid string) bool {
	db := db.GetDB()
	res := db.Table("users").Where("uid=?", uid).Scan(user)
	if res.Error != nil {
		log.Error("user info update password: ", res.Error)
		return false
	}
	return true
}

func changePassword(uid, pass string) bool {
	db := db.GetDB()
	err := db.Table("users").Where("uid=?", uid).Update("password", pass).Error
	if err != nil {
		log.Error("user change Password ", err.Error())
		return false
	}
	return true
}

func (login *login) chenckLogin() (bool, int64) {
	db := db.GetDB()
	var count int64 = 0
	err := db.Table("users").Where("login LIKE ?", login.Login).Count(&count)
	if err.Error != nil {
		log.Error("chenck Login ", err.Error)
		return false, 0
	}
	return true, count
}

func (perm *permissions) perms(role int) bool {
	db := db.GetDB()
	err := db.Table("accses").Where("role_id = ?", role).Scan(perm)
	if err.Error != nil {
		log.Error("chenck Login ", err.Error)
		return false
	}
	return true
}

