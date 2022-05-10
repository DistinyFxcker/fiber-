package valid

import (
	"context"
	"github.com/gogf/gf/v2/util/gvalid"
)

var (
	Validator *gvalid.Validator
)

func Init() {
	Validator = gvalid.New()
}

func Validate(v interface{}) (err error) {
	err = gvalid.New().Data(v).Run(context.Background())
	if err != nil {
		return err
	}
	return nil
}
