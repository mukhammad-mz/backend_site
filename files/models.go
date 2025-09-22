package files

import (
	"site_backend/db"

	log "github.com/sirupsen/logrus"
)

//musicsTable ....
var musicsTable = "musics"

func (id *musicID) musicID() bool {
	db := db.GetDB()
	//db = db.Raw("SELECT  m.`id` FROM `musics` m ORDER BY m.`id` DESC LIMIT 1").Scan(&id)
	db = db.Table(musicsTable).Select("id").Last(&id)
	
	if db.Error != nil {
		log.Error("maxIDSelect", db.Error)
		return false
	}
	return true
}

func (arr *MusicsArray) selectFile(id int) bool {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id < ?", id).Order("id desc").Limit(10).Scan(&arr)
	if db.Error != nil {
		log.Error("selectNewFile", db.Error)
		return false
	}
	return true
}

func (arr *MusicsArray) userFiles(id string) bool {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id_user = ?", id).Scan(&arr)
	if db.Error != nil {
		log.Error("selectNewFile: ", db.Error)
		return false
	}
	return true
}

func (file *Musics) File(id int) bool {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id = ?", id).Scan(&file)
	if db.Error != nil {
		log.Error("selectNewFile: ", db.Error)
		return false
	}
	return true
}

func (uploadDate *Musics) insertMusic() bool {
	db := db.GetDB()
	db = db.Table(musicsTable).Save(&uploadDate)
	if db.Error != nil {
		log.Error("insertMusic DB", db.Error)
		return false
	}
	return true
}

func deletFile(id int) bool {
	db := db.GetDB()
	db = db.Table(musicsTable).Where("id =? ", id).Delete(id)
	if db.Error != nil {
		log.Error("deletFile DB", db.Error)
		return false
	}
	return true
}

func checkFile(name string) bool {
	db := db.GetDB()
	var count int64 = 0
	err := db.Table(musicsTable).Select("name_orig").Where("name_orig = ?", name).Count(&count)
	if err.Error != nil {
		log.Error("Сheck File", err.Error)
		return false
	} else if count != 0 {
		return false
	}
	return true
}

func (direct *direct) direct() bool{
	db := db.GetDB()
	err := db.Table("direct").Scan(&direct)
	if err.Error != nil {
		log.Error("direct", err.Error)
		return false
	} 
	return true
}