package files

import (
	"fmt"
	"net/http"
	"os"
	"site_backend/authorization"
	"site_backend/helper"
	"site_backend/response"
	st "strconv"
	s "strings"

	"github.com/gin-gonic/gin"
)

//GetInfoNewFiles get new files
//@Tags Music
// @Summary return music list
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /music/list [get]
func GetListMusic(c *gin.Context) {
	defer authorization.Recoverd(c, "GetListMusic: ")
	result := MusicsArray{}
	id, _ := st.Atoi(c.Param("idFile"))
	if id == 0 {
		maxID := musicID{}
		maxID.musicID()
		id = maxID.ID
		id++
	}
	
	if err := result.selectFile(id); !err {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(map[string]MusicsArray{"lists": result}))
}

func GetMisicInfo(c *gin.Context) {
	defer authorization.Recoverd(c, "GetMisicInfo: ")

}

// @Security ApiKey
// @Tags files
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /file [get]
func GetUserFile(c *gin.Context) {
	defer authorization.Recoverd(c, "GetUserFile: ")
	id, _ := c.Get("userUID")
	result := MusicsArray{}
	if err := result.userFiles(id.(string)); !err {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	c.JSON(http.StatusOK, response.CorrectWithData(result))
}

func GetFile(c *gin.Context) {
	defer authorization.Recoverd(c, "GetFile: ")
	result := MusicsArray{}
	if err := result.userFiles(""); !err {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}
	c.JSON(http.StatusOK, response.CorrectWithData(result))
}

// @Security ApiKey
// @Tags files
// @Param body body Musics false "file struct"
// @Produce json
// @Success 200 {object} response.Response
// @Failure 409 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /file [post]
func PostFile(c *gin.Context) {
	defer authorization.Recoverd(c, "PostFile: ")
	addFile := Musics{}
	if err := c.BindJSON(&addFile); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFromError(err))
		return
	}
	str := addFile.Atist + " - " + addFile.Name + ".mp3"
	addFile.NameOrig = s.Replace(str, " ", "_", -1)
	if check := checkFile(addFile.NameOrig); !check {
		c.JSON(http.StatusConflict, response.ErrorFromString(http.StatusText(http.StatusConflict)))
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorFrom("Bad Request", "post file:  Form File", err))
		return
	}

	direct := direct{}
	ext := direct.direct()
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	if err := c.SaveUploadedFile(file, direct.Direct+addFile.NameOrig); err != nil {
		c.JSON(http.StatusInternalServerError, response.ErrorFrom(http.StatusText(500), "post file: Save Uploaded File", err))
		return
	}

	sz := float64(file.Size)
	size := fmt.Sprintf("%.2f", float64(sz/1024/1024)) + " mb"
	addFile.Size = size
	addFile.DateUpload = helper.GetDateTime()
	userID, _ := c.Get("userUID")
	addFile.IDUser = userID.(string)

	if err := addFile.insertMusic(); !err {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	c.JSON(http.StatusCreated, response.Createdd())
}

func PutFile(c *gin.Context) {
	defer authorization.Recoverd(c, "PutFile: ")

}

//DeleteFile ...
// @Security ApiKey
// @Tags files
// @Produce json
// @Success 200 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /file/{idFile} [delete]
func DeleteFile(c *gin.Context) {
	defer authorization.Recoverd(c, "DeleteFile: ")
	id, _ := st.Atoi(c.Param("idFile"))
	fileInfo := Musics{}
	ext := fileInfo.File(id)
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	direct := direct{}
	ext = direct.direct()
	if !ext {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	path := direct.Direct + "/" + fileInfo.NameOrig
	if err := os.Remove(path); err != nil {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}

	if ext = deletFile(id); !ext {
		c.JSON(http.StatusInternalServerError, response.ServerError())
		return
	}
	c.JSON(200, response.Correct())
}
