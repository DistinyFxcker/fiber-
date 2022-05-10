package service

import (
	"cdnFiber/app/admin/model"
	"context"
)

var Admin = sAdmin{}

type sAdmin struct {
}

func (s *sAdmin) Get(ctx context.Context, r *model.GetAdminInput) (res *model.GetAdminOutPut, err error) {
	return nil, nil
}

func (s *sAdmin) Auth(ctx context.Context) {
}
