package controller

import (
	"cdnFiber/app/admin/api"
	"github.com/gofiber/fiber/v2"
)

var Auth = cAuth{}

type cAuth struct {
}

// @Tags 用户接口
// @Summary 登录接口
// @Accept  json
// @Produce  json
// @Param   req     body    api.AuthReq     true   "登录参数"
// @Success 200  {object} api.AuthRes
// @Router /testapi/get-string-by-int/{some_id} [get]
func (c *cAuth) Auth(ctx *fiber.Ctx) (err error) {
	req := new(api.AuthReq)
	err = BindReq(ctx, req)
	if err != nil {
		return err
	}

	//if err := admin.BindReq(ctx, req); err != nil {
	//	return err
	//}
	//return SuccessRes(ctx, "")
	//
	//u := gen.CdnDap.Admin
	//user, err := u.WithContext(ctx.Context()).Where(u.Nick.Eq("")).First()
	//if err != nil {
	//	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
	//		"message": err.Error(),
	//	})
	//}
	//claims := jwt.MapClaims{
	//	"name":  "",
	//	"admin": true,
	//	"exp":   time.Now().Add(time.Hour * 72).Unix(),
	//}
	return nil
}

func (c *cAuth) LogOut(ctx *fiber.Ctx) error {
	return nil
}

func (c *cAuth) RefreshToken(ctx *fiber.Ctx) error {
	return nil
}
