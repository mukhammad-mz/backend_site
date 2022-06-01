package users

type Users struct {
	UID        string `json:"uid"`
	LastName   string `json:"last_name"`
	FristName  string `json:"frist_name"`
	DataRegist string `json:"data_regist"`
	LastVisit  string `json:"last_visit"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Token      string `json:"token"`
	IDRole     int    `json:"id_role"`
	CreateAt   string `json:"create_at"`
	UpdateAt   string `json:"update_at"`
}

type userInfo struct {
	LastName   string `json:"last_name"`
	FristName  string `json:"frist_name"`
	Login      string `json:"login"`
	DataRegist string `json:"data_regist"`
	CreatAt    string `json:"create_at"`
}

type usersInfo []Users

type password struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

type login struct {
	Login string `json:"login"`
}

type Permission struct {
	ID        int `json:"id"`
	RoleID    int `json:"role_id"`
	IDHandler int `json:"id_handler"`
}

type permissions []Permission
