package authorization

type UserLogin struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserToken struct {
	UID    string `json:"uid"`
	Token  string `json:"token"`
	IDRole string `json:"id_role"`
}
