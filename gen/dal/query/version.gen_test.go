// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"fmt"
	"testing"

	"go_web_example/gen/dal/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := db.AutoMigrate(&model.Version{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Version{}) fail: %s", err)
	}
}

func Test_versionQuery(t *testing.T) {
	version := newVersion(db)
	version = *version.As(version.TableName())
	do := version.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(version.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <version> fail:", err)
		return
	}

	_, ok := version.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from version success")
	}

	err = do.Create(&model.Version{})
	if err != nil {
		t.Error("create item in table <version> fail:", err)
	}

	err = do.Save(&model.Version{})
	if err != nil {
		t.Error("create item in table <version> fail:", err)
	}

	err = do.CreateInBatches([]*model.Version{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <version> fail:", err)
	}

	_, err = do.Select(version.ALL).Take()
	if err != nil {
		t.Error("Take() on table <version> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <version> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <version> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <version> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Version{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <version> fail:", err)
	}

	_, err = do.Select(version.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <version> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <version> fail:", err)
	}

	_, err = do.Select(version.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <version> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <version> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <version> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <version> fail:", err)
	}

	_, err = do.ScanByPage(&model.Version{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <version> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <version> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <version> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <version> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <version> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <version> fail:", err)
	}
}
