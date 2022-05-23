package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"go_web_example/app/admin/api"
	"go_web_example/app/admin/global"
	"go_web_example/app/admin/model"
	"go_web_example/app/admin/service"
	db "go_web_example/gen/dal/model"
	"time"
)

var Auth = cAuth{}

type cAuth struct {
}

//@Tags 用户接口
//@Summary 登录接口
//@Accept  json
//@Produce  json
//@Param   req     body    api.AuthReq     true   "登录参数"
//@Success 200  {object} api.AuthRes
//@Router /api/v1/auth [post]
func (c *cAuth) Auth(ctx *gin.Context) {
	var err error
	req := new(api.AuthReq)
	err = BindReq(ctx, req)
	if err != nil {
		ErrorRes(ctx, "登录失败,请稍后重试", err)
		return
	}

	AdminService := service.Admins()
	var adminInfo *db.Admin
	adminInfo, err = AdminService.Get(ctx, &model.GetAdminInput{
		Id:       req.Nick,
		Password: req.Password,
	})
	if err != nil {
		ErrorRes(ctx, "登录失败,请稍后重试", err)
		return
	}

	claims := jwt.MapClaims{
		"nick":         adminInfo.Nick,
		"email":        adminInfo.Email,
		"state":        adminInfo.State,
		"adminService": true,
		"exp":          time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	t, err := token.SignedString([]byte(global.Viper.GetString("cdnAdmin.SecretKey")))
	res := api.AuthRes{
		Id:        adminInfo.Id,
		Nick:      adminInfo.Nick,
		Email:     adminInfo.Email,
		State:     adminInfo.State,
		LastTime:  adminInfo.LastTime,
		LastIP:    adminInfo.LastIP,
		CreatedAt: adminInfo.CreatedAt,
		UpdatedAt: adminInfo.UpdatedAt,
		Token:     t,
	}
	SuccessRes(ctx, res)
	return
}

func (c *cAuth) LogOut(ctx *gin.Context) {
	return
}

func (c *cAuth) RefreshToken(ctx *gin.Context) {
	return
}

const ()
