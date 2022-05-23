package service

import (
	"github.com/gofiber/fiber/v2"
	"go_web_example/app/admin/global"
	"go_web_example/app/admin/service/model"
	"go_web_example/gen/dal/query"
	"go_web_example/utility/errors"
)

var Admin = sAdmin{}

type sAdmin struct {
}

func Admins() sAdmin {
	return Admin
}

func (s *sAdmin) Auth(ctx *fiber.Ctx) {

}

func (s *sAdmin) Query(ctx *fiber.Ctx, r *model.QueryAdminInput) (res *model.QueryAdminOutput, err error) {
	adminModel := query.Use(global.DB).Admin
	adminDao := adminModel.WithContext(ctx.Context())
	if r.KeyWord != "" {
		adminDao = adminDao.Where(adminModel.Nick.Like(r.KeyWord))
	}

	if len(r.Ids) != 0 {
		adminDao = adminDao.Where(adminModel.Id.In(r.Ids...))
	}

	res.Total, err = adminDao.Count()
	if err != nil {
		return nil, errors.Wrapf(err, "统计Admin数量发生错误")
	}
	if r.Limit != 0 && r.Offset != 0 {
		adminDao = adminDao.Limit(r.Limit).Offset(r.Offset).Order()
	}
	res.List, err = adminDao.Find()
	if err != nil {
		return nil, errors.Wrapf(err, "查询Admin数据发生错误")
	}
	return res, nil
}
