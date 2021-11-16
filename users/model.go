package users

import (
	"fmt"
	"site_backend/db"

	log "github.com/sirupsen/logrus"
)

func CheckPermission(userID, loc string) bool {
	fmt.Println(userID,len(userID),len(loc), loc)
	db := db.GetDB()
	count := 0
	db = db.Table("users").Where("id_role = ? and uid = ?", 1, userID).Count(&count)
	if db.Error != nil {
		log.Error("Permissions: ", db.Error)
		return false
	} else if count > 0 {
		return true
	}

	// sql := fmt.Sprintf("SELECT COUNT(0) FROM users u " +
	// "JOIN accses a ON a.role_id = u.id_role" +
	// "JOIN handlers h	ON h.id = a.id_handler"+
	// "WHERE u.uid = %s AND h.name = %s",userID, loc )
	// c := Users{}
	// db = db.Raw(sql).Scan(&c)

	result := Users{}

	db = db.Table("users").
		Joins("inner JOIN accses ON accses.role_id = users.id_role").
		Joins("inner JOIN handlers ON handlers.id = accses.id_handler").
		Where("users.uid = ? AND handlers.name = ?", userID, loc).Find(&result)

	fmt.Println(result)

	if db.Error != nil {
		log.Error("Permissions: ", db.Error)
		return false
	}

	return true
}
