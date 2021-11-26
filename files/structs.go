package files

//MusicsType for select DB
type MusicsType struct {
	ID         int    `json:"id"`
	DateUpload string `json:"date_upload"`
	Atist      string `json:"atist"`
	Name       string `json:"name"`
	NameOrig   string `json:"name_orig"`
	Duration   string `json:"duration"`
	Size       string `json:"size"`
	IDDirect   string `json:"id_direct"`
	IDUser     string `json:"id_user"`
}


//URLType struct
type URLType struct {
	Direct   string `json:"direct"`
	NameOrig string `json:"name_orig"`
}
