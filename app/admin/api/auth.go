package api

type AuthReq struct {
	Nick     string `json:"nick"  v:"required|length:5,32 #账户名为必填项|账户长度必须为5到32"`
	Password string `json:"password" v:"required|length:5,32 #密码为必填|密码长度必须为5到32"`
}

type AuthRes struct {
	Id        string `json:"id"`
	Nick      string `json:"nick"`
	Role      int    `json:"role"`
	Email     string `json:"email"`
	State     int    `json:"state"`
	LastTime  int64  `json:"lastTime"`
	LastIP    string `json:"lastIP"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
	Token     string `json:"token"`
}
