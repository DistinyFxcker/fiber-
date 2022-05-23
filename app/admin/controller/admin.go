package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"go_web_example/app/admin/api"
	"go_web_example/app/admin/service"
	"go_web_example/app/admin/service/model"
	db "go_web_example/gen/dal/model"
	"log"
)

var Admin = cAdmin{}

type cAdmin struct {
}

func (c *cAdmin) GetAdmin(ctx *gin.Context) (err error) {
	Req := model.GetAdminInput{}
	err = BindReq(ctx, &Req)
	if Req.Id == "" {
		return ErrorRes(ctx, api.Admin_Not_found, nil)
	}
	var AdminInfo *db.Admin
	admins := service.Admins()
	AdminInfo, err = admins.Get(ctx, &model.GetAdminInput{
		Id: Id,
	})
	if err != nil {
		return ErrorRes(ctx, api.Admin_Not_found, err)
	}
	res := api.AuthRes{
		Id:        AdminInfo.Id,
		Nick:      AdminInfo.Nick,
		Email:     AdminInfo.Email,
		State:     AdminInfo.State,
		LastTime:  AdminInfo.LastTime,
		LastIP:    AdminInfo.LastIP,
		CreatedAt: AdminInfo.CreatedAt,
		UpdatedAt: AdminInfo.UpdatedAt,
	}
	SuccessRes(ctx, res)
	return
}

func (c *cAdmin) GetAdminList(ctx *fiber.Ctx) (err error) {
	Req := api.GetAdminReq{}
	//err = BindReq(ctx , &Req)
	log.Fatalln(Req)
	if err != nil {
		//return ErrorRes(ctx , api.Admin_Not_found , err)
	}
	return nil
}

//func (c *cAdmin) AddAdmin(ctx *fiber.Ctx, req *api.AddAdminReq) error {
//	return nil
//}

//func (c *cAdmin) DeleteAdmin(ctx *fiber.Ctx , req *api.) error {
//	return nil
//}
//
//func (c *cAdmin) EditAdmin(ctx *fiber.Ctx) error {
//	return nil
//}

const ()
