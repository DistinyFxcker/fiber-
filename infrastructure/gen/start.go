package gen

import (
	"cdnFiber/infrastructure/gen/cdn"
	"cdnFiber/infrastructure/gen/cdn/dal/query"
	"gorm.io/gen"
	"gorm.io/gorm"
	"reflect"
)

var (
	CdnDap *query.Query
)

func New(db *gorm.DB) {
	g := gen.NewGenerator(
		gen.Config{
			OutPath: "infrastructure/gen/cdn/dal/query",
		})
	g.UseDB(db)

	InitCdnQuery(g, db)

	g.Execute()

}

//生成cdn查询
func InitCdnQuery(generator *gen.Generator, db *gorm.DB) {

	cdnDao, err := cdn.New(db)
	if err != nil {
		panic(err)
	}
	t := reflect.TypeOf(*cdnDao)
	v := reflect.ValueOf(*cdnDao)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i).Interface()
		generator.ApplyBasic(value)
	}
	CdnDap = query.Use(db)
	//generator.ApplyInterface()

	//g.ApplyInterface(func(method model.))
}
