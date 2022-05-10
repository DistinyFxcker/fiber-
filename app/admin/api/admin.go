package api

type AddAdminReq struct {
	Nick     string `json:"nick" v:"required|length:5,32 #账户名为必填项|账户长度必须为5到32"`    //昵称
	Password string `json:"password" v:"required|length:5,32 #密码为必填项|密码长度必须为5到32"` //密码，md5加密
	Role     int    `json:"role" v:"required|in:1,2 #创建角色为必填项|角色类型错误"`             //角色，1：管理员，2：普通会员
	Email    string `json:"email" v:"required-if:role,2|email #创建普通用户邮箱为必填项"`      //邮箱
	State    int    `json:"state" v:"required|in:1,2 #状态为必填项|状态类型错误"`              //状态
}

type GetAdminReq struct {
}
