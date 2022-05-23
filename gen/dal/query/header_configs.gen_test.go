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
	err := db.AutoMigrate(&model.HeaderConfig{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.HeaderConfig{}) fail: %s", err)
	}
}

func Test_headerConfigQuery(t *testing.T) {
	headerConfig := newHeaderConfig(db)
	headerConfig = *headerConfig.As(headerConfig.TableName())
	do := headerConfig.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(headerConfig.TableName(), clause.PrimaryKey)
	_, err := do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <header_configs> fail:", err)
		return
	}

	_, ok := headerConfig.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from headerConfig success")
	}

	err = do.Create(&model.HeaderConfig{})
	if err != nil {
		t.Error("create item in table <header_configs> fail:", err)
	}

	err = do.Save(&model.HeaderConfig{})
	if err != nil {
		t.Error("create item in table <header_configs> fail:", err)
	}

	err = do.CreateInBatches([]*model.HeaderConfig{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <header_configs> fail:", err)
	}

	_, err = do.Select(headerConfig.ALL).Take()
	if err != nil {
		t.Error("Take() on table <header_configs> fail:", err)
	}

	_, err = do.First()
	if err != nil {
		t.Error("First() on table <header_configs> fail:", err)
	}

	_, err = do.Last()
	if err != nil {
		t.Error("First() on table <header_configs> fail:", err)
	}

	_, err = do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <header_configs> fail:", err)
	}

	err = do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.HeaderConfig{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <header_configs> fail:", err)
	}

	_, err = do.Select(headerConfig.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <header_configs> fail:", err)
	}

	_, err = do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <header_configs> fail:", err)
	}

	_, err = do.Select(headerConfig.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <header_configs> fail:", err)
	}

	_, err = do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <header_configs> fail:", err)
	}

	_, err = do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <header_configs> fail:", err)
	}

	_, _, err = do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <header_configs> fail:", err)
	}

	_, err = do.ScanByPage(&model.HeaderConfig{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <header_configs> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <header_configs> fail:", err)
	}

	_, err = do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <header_configs> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), clause.PrimaryKey)

	err = do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <header_configs> fail:", err)
	}

	err = do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <header_configs> fail:", err)
	}

	_, err = do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <header_configs> fail:", err)
	}
}
