package model

type GetAdminInput struct {
	Id       string `orm:"id"`
	Nick     string `orm:"nick"`
	Password string `orm:"password"`
	State    int64  `orm:"state"`
	Email    string `orm:"email"`
	Role     int64  `orm:"role"`
}

type GetAdminOutPut struct {
	Id        string //id
	Nick      string //昵称
	Password  string //密码，md5加密
	Role      int64  //角色，1：管理员，2：操作员
	Email     string //邮箱
	State     int64  //1:正常  2:禁用
	LastTime  int64  //最新登录时间
	LastIP    string //最新登录IP
	Region    string //最新登录所属地区
	CreatedIp string //创建IP
	CreatedAt int64  //创建时间
	UpdatedAt int64  //修改时间，新记录保持和创建时间一致
}
