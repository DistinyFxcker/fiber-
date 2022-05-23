package biz

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go_web_example/gen/dal/model"
	"go_web_example/gen/dal/query"
	"go_web_example/utility/errors"
	"go_web_example/utility/uuid"
	"gorm.io/gorm"
)

type GetAdminInput struct {
	Id       string
	State    int
	Nick     string
	Password string
	Email    string
}

func Get(ctx context.Context, db *gorm.DB, input *GetAdminInput) (output *model.Admin, err error) {
	adminModel := query.Use(db).Admin
	adminDao := adminModel.WithContext(ctx)
	if input.Id != "" {
		adminDao = adminDao.Where(adminModel.Id.Eq(input.Id))
	}
	if input.State != 0 {
		adminDao = adminDao.Where(adminModel.State.Eq(input.State))
	}
	if input.Nick != "" {
		adminDao = adminDao.Where(adminModel.Nick.Eq(input.Nick))
	}
	if input.Password != "" {
		adminDao = adminDao.Where(adminModel.Password.Eq(input.Password))
	}
	if input.Email != "" {
		adminDao = adminDao.Where(adminModel.Email.Eq(input.Email))
	}
	output, err = adminDao.First()
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.NotFound.Wrapf(err, "")
		default:
			return nil, errors.Wrapf(err, "查询管理员Admin发生错误")
		}
	}
	return
}

type ListAdminInput struct {
	Ids   []string
	State int
	Nick  []string
	Email []string
}

func List(ctx context.Context, db *gorm.DB, input *ListAdminInput) (output []*model.Admin, err error) {
	adminModel := query.Use(db).Admin
	adminDao := adminModel.WithContext(ctx)
	if len(input.Ids) != 0 {
		adminDao = adminDao.Where(adminModel.Id.In(input.Ids...))
	}
	if len(input.Nick) != 0 {
		adminDao = adminDao.Where(adminModel.Nick.In(input.Nick...))
	}
	if input.State != 0 {
		adminDao = adminDao.Where(adminModel.State.Eq(input.State))
	}

	output, err = adminDao.Find()
	if err != nil {
		return nil, errors.Wrapf(err, "查询多条Admin数据发生错误")
	}
	return
}

type AddAdminInput struct {
	Nick     string
	Password string
	Email    string
	State    int
	Region   string
}

func Add(ctx *fiber.Ctx, db *gorm.DB, input *AddAdminInput) (err error) {
	adminModel := query.Use(db).Admin
	adminDao := adminModel.WithContext(ctx.Context())
	adminInfo := model.Admin{
		Id:        uuid.New(),
		Nick:      input.Nick,
		Password:  input.Password,
		Email:     input.Email,
		State:     input.State,
		Region:    input.Region,
		CreatedIp: ctx.IP(),
	}
	err = adminDao.Create(&adminInfo)
	if err != nil {
		return errors.Wrapf(err, "创建Admin出现错误")
	}
	return
}

type UpdateAdminInput struct {
	Id       string
	Nick     string
	Password string
	Email    string
	State    int
	LastTime int64
	LastIp   string
	Region   string
}
