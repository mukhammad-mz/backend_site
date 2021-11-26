package files

import (
	"site_backend/db"

	log "github.com/sirupsen/logrus"
)

//MusicsArray ....


//MusicsArray Array Type
type MusicsArray []MusicsType

//musicsTable ....
var musicsTable = "musics"

func (id *MaxID) maxIDSelect() error {
	db := db.GetDB()
	//db = db.Raw("SELECT  m.`id` FROM `musics` m ORDER BY m.`id` DESC LIMIT 1").Scan(&id)
	db = db.Table(musicsTable).Select("id").Last(&id)
	if db.Error != nil {
		log.Error("maxIDSelect", db.Error)
	}
	return db.Error
}

func (arr *MusicsArray) selectFile(id int) error {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id < ?", id).Order("id desc").Limit(10).Scan(&arr)
	if db.Error != nil {
		log.Error("selectNewFile", db.Error)
	}
	return db.Error
}
func (arr *MusicsArray) selectUserFile(id string) error {
	db := db.GetDB()
	if id != "" {
		db = db.Table(musicsTable).Where("id_user = ?", id).Scan(&arr)
		if db.Error != nil {
			log.Error("selectNewFile: ", db.Error)
		}
		return db.Error
	}
	db = db.Table(musicsTable).Where("id_user = ?", id).Scan(&arr)
	if db.Error != nil {
		log.Error("selectNewFile: ", db.Error)
	}
	return db.Error
}

func (downType *MusicsType) insertMusic() error {
	db := db.GetDB()
	db = db.Table(musicsTable).Save(&downType)
	if db.Error != nil {
		log.Error("insertMusic DB", db.Error)
	}
	return db.Error
}

func deletFile(id int) error {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id =? ", id).Delete(id)
	if db.Error != nil {
		log.Error("deletFile DB", db.Error)
	}
	return db.Error
}

func checkFile(name string) bool {
	db := db.GetDB()
	count := 0
	err := db.Table(musicsTable).Select("name_orig").Where("name_orig = ?", name).Count(&count)
	if err.Error != nil {
		log.Error("Сheck File", err.Error)
		return false
	} else if count != 0 {
		return false
	}
	return true
}


func (url *URLType) creatURL(id int) error {
	db := db.GetDB()
	db = db.Raw(`SELECT d.direct, m.name_orig FROM musics m , directs d 
				WHERE m.id_direct = d.id 
				AND m.id = ?`, id).Scan(&url)
	if db.Error != nil {
		log.Error("creatURL", db.Error)
	}
	return db.Error
}
