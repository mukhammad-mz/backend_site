package files

import (
	"fmt"
	"net/http"
	"os"
	"site_backend/helper"
	"site_backend/response"
	st "strconv"
	s "strings"

	"github.com/gin-gonic/gin"
)

//MaxID ..... : )
type MaxID struct {
	ID int `json:"id"`
}

//GetInfoNewFiles get new files info in Datebase
func GetListMusic(c *gin.Context) {
	result := MusicsArrayList{}
	id, _ := st.Atoi(c.Param("idFile"))
	if id == 0 {
		maxID := MaxID{}
		maxID.maxIDSelect()
		id = maxID.ID
		id++
	}
	if err := result.selectFile(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "Select music", err))
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(result))
}

func GetMisicInfo(c *gin.Context) {

}

func GetUserFile(c *gin.Context) {
	id, _ := c.Get("userID")
	result := MusicsArrayList{}
	if err := result.selectUserFile(id.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "select User File", err))
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(result))
}

func GetFile(c *gin.Context) {
	result := MusicsArrayList{}
	if err := result.selectUserFile(""); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "select File", err))
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(result))
}

type direcr struct {
	Direct string `json:"direct"`
}

//PostFile Upload File function
func PostFile(c *gin.Context) {
	addFile := MusicsType{}

	// TODO: delete Coment ↓
	// addFile.Atist = DecodeB64(c.GetHeader("atist"))
	// addFile.Name = DecodeB64(c.GetHeader("name"))
	// addFile.Duration = DecodeB64(c.GetHeader("duration"))
	// addFile.Id_direct = c.GetHeader("id_direct")

	addFile.Atist = c.GetHeader("atist")
	addFile.Name = c.GetHeader("name")
	addFile.Duration = c.GetHeader("duration")
	addFile.IDDirect = c.GetHeader("id_direct")
	str := addFile.Atist + " - " + addFile.Name + ".mp3"
	addFile.NameOrig = s.Replace(str, " ", "_", -1)
	if check := checkFile(addFile.NameOrig); !check {
		c.JSON(http.StatusConflict, response.ErrorFromString(http.StatusText(http.StatusConflict)))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, response.ErrorFrom("Bad Request", "Form File", err))
		return
	}

	dir := direcr{}
	id, _ := st.Atoi(addFile.IDDirect)
	if err = dir.selectDir(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFromString(http.StatusText(500)))
	}

	if err := c.SaveUploadedFile(file, dir.Direct+"/"+addFile.NameOrig); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "Save Uploaded File", err))
		return
	}

	sz := float64(file.Size)
	size := fmt.Sprintf("%.2f", float64(sz/1024/1024)) + " mb"
	addFile.Size = size
	addFile.DateUpload = helper.GetDateTime()
	userID, _ := c.Get("userID")
	addFile.IDUser = userID.(string)
	if err := addFile.insertMusic(); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "Save Uploaded File", err))
		return
	}
	c.JSON(http.StatusCreated, response.Createdd())
}

func PutFile(c *gin.Context) {

}

//DeleteFile ...
func DeleteFile(c *gin.Context) {
	id, _ := st.Atoi(c.Param("idFile"))
	fileInfo := URLType{}
	err := fileInfo.creatURL(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "deletFile creat name", err))
		return
	}
	path := fileInfo.Direct + "/" + fileInfo.NameOrig

	if err = os.Remove(path); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "deletFile in memory", err))
		return
	}
	if err := deletFile(id); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "deletFile in DB", err))
		return
	}
	c.JSON(200, response.Correct())
}
