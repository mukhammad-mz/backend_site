package files

//MusicsType for select DB
type Musics struct {
	ID         int    `json:"id"`
	DateUpload string `json:"date_upload"`
	Atist      string `json:"atist"`
	Name       string `json:"name"`
	NameOrig   string `json:"name_orig"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	IDUser     string `json:"id_user"`
}

//MusicsArray Array Type
type MusicsArray []Musics

//MaxID ..... : )
type musicID struct {
	ID int `json:"id"`
}

type direct struct {
	Direct string `json:"direct"`
}
